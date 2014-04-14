// @filename post.go
// @author mownier

package modules

import (
	"goomoan/goomoan"
)

type PostModule struct {
	*goomoan.Module
}

//@Method : POST
//@Path   : /post
func (m *PostModule) SendPost(params map[string]string) (int, interface{}) {
	return 200, "SendPost: Ok"
}

//@Method : GET
//@Path   : /post
func (m *PostModule) GetPosts(params map[string]string) (int, interface{}) {
	return 200, "GetPosts: Ok"
}

//@Method : DELETE
//@Path   : /post
func (m *PostModule) DeletePost(params map[string]string) (int, interface{}) {
	return 200, "DeletePost: Ok"
}

//@Method : GET
//@Path   : /post/search
func (m *PostModule) SearchPost(params map[string]string) (int, interface{}) {
	return 200, "SearchPost: Ok"
}

//@Method : GET
//@Path   : /post/:id
func (m *PostModule) GetPost(params map[string]string) (int, interface{}) {
	return 200, "GetPost: Ok"
}
