// @filename post.go
// @author mownier

package modules

import (
	"goomoan/core"
)

type PostModule struct {
	*core.Module
}

func NewPostModule() *core.Module {
	m := new(core.Module)
	m.Routes = map[string]*core.Route {
		"SendPost"   : {"POST"  , "/post"},
		"GetPosts"   : {"GET"   , "/posts"},
		"DeletePost" : {"DELETE", "/post"},
		"SearchPost" : {"GET"   , "/post/search"},
	}
	return m
}

func (m *PostModule) SendPost(params map[string]string) (int, interface{}) {
	return 200, "SendPost: Ok"
}

func (m *PostModule) GetPosts(params map[string]string) (int, interface{}) {
	return 200, "GetPosts: Ok"
}

func (m *PostModule) DeletePost(params map[string]string) (int, interface{}) {
	return 200, "DeletePost: Ok"
}

func (m *PostModule) SearchPost(params map[string]string) (int, interface{}) {
	return 200, "SearchPost: Ok"
}
