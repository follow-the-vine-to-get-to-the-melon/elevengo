package elevengo

import (
	"fmt"
	"github.com/deadblue/elevengo/core"
	"github.com/deadblue/elevengo/internal"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	cookieDomain = ".115.com"
	cookieUrl    = "https://115.com"

	apiUserInfo = "https://my.115.com/"
)

// Credentials contains required data to make remote server considers as
// you have signed in. You can get these from you browser cookies.
type Credentials struct {
	UID  string
	CID  string
	SEID string
}

// Import the credentials into client.
func (c *Client) ImportCredentials(cr *Credentials) (err error) {
	cks := []*http.Cookie{
		{Name: "UID", Value: cr.UID, Domain: cookieDomain, Path: "/", HttpOnly: true},
		{Name: "CID", Value: cr.CID, Domain: cookieDomain, Path: "/", HttpOnly: true},
		{Name: "SEID", Value: cr.SEID, Domain: cookieDomain, Path: "/", HttpOnly: true},
	}
	u, _ := url.Parse(cookieUrl)
	c.cj.SetCookies(u, cks)
	return c.getUserInfo()
}

// A new and graceful way to get user information.
func (c *Client) getUserInfo() (err error) {
	cb := fmt.Sprintf("jQuery%d_%d", rand.Uint64(), time.Now().Unix())
	qs := core.NewQueryString().
		WithString("ct", "ajax").
		WithString("ac", "nav").
		WithString("callback", cb).
		WithInt64("_", time.Now().Unix())
	result := &internal.UserInfoResult{}
	if err = c.hc.JsonpApi(apiUserInfo, qs, result); err != nil {
		return
	}
	if c.ui == nil {
		c.ui = new(internal.UserInfo)
	}
	c.ui.UserId = result.Data.UserId
	c.ui.UserName = result.Data.UserName
	return
}
