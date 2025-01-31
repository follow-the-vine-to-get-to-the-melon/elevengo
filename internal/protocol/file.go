package protocol

import (
	"encoding/json"

	"github.com/deadblue/elevengo/internal/util"
	"github.com/deadblue/elevengo/lowlevel/errors"
	"github.com/deadblue/elevengo/lowlevel/types"
)

const (
	FileListLimit = 32
)

//lint:ignore U1000 This type is used in generic.
type FileListResp struct {
	StandardResp

	AreaId     string         `json:"aid"`
	CategoryId util.IntNumber `json:"cid"`

	Count int `json:"count"`

	Order string `json:"order"`
	IsAsc int    `json:"is_asc"`

	Offset int `json:"offset"`
	Limit  int `json:"limit"`
}

func (r *FileListResp) Err() (err error) {
	// Handle special error
	if r.ErrorCode2 == errors.CodeFileOrderNotSupported {
		return &errors.ErrFileOrderNotSupported{
			Order: r.Order,
			Asc:   r.IsAsc,
		}
	}
	return r.StandardResp.Err()
}

func (r *FileListResp) Extract(v any) (err error) {
	ptr, ok := v.(*types.FileListResult)
	if !ok {
		return errors.ErrUnsupportedResult
	}
	ptr.Files = make([]*types.FileInfo, 0)
	if err = json.Unmarshal(r.Data, &ptr.Files); err != nil {
		return
	}
	ptr.DirId = r.CategoryId.String()
	ptr.Count = r.Count
	ptr.Order, ptr.Asc = r.Order, r.IsAsc
	return
}

//lint:ignore U1000 This type is used in generic.
type FileSearchResp struct {
	StandardResp

	Folder struct {
		CategoryId string `json:"cid"`
		ParentId   string `json:"pid"`
		Name       string `json:"name"`
	} `json:"folder"`

	Count       int `json:"count"`
	FileCount   int `json:"file_count"`
	FolderCount int `json:"folder_count"`

	Order string `json:"order"`
	IsAsc int    `json:"is_asc"`

	Offset int `json:"offset"`
	Limit  int `json:"page_size"`
}

func (r *FileSearchResp) Extract(v any) (err error) {
	ptr, ok := v.(*types.FileListResult)
	if !ok {
		return errors.ErrUnsupportedResult
	}
	ptr.Files = make([]*types.FileInfo, 0)
	if err = json.Unmarshal(r.Data, &ptr.Files); err != nil {
		return
	}
	ptr.DirId = r.Folder.CategoryId
	ptr.Count = r.Count
	ptr.Order, ptr.Asc = r.Order, r.IsAsc
	return
}

//lint:ignore U1000 This type is used in generic.
type FileGetDescResp struct {
	BasicResp

	Desc string `json:"desc"`
}

func (r *FileGetDescResp) Extract(v any) (err error) {
	if ptr, ok := v.(*string); !ok {
		err = errors.ErrUnsupportedResult
	} else {
		*ptr = r.Desc
	}
	return
}
