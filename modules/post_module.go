// @filename post.go
// @author mownier

package modules

import (
	"net/url"
	"goomoan/goomoan"
	"goomoan/models"
)

type PostModule struct {
	goomoan.Module
}

//@Method : POST
//@Path   : /post
func (m *PostModule) SendPost(params url.Values) (int, *models.PostModel) {
	u := new(models.UserModel)
	u.Id = "1023"
	u.Email = "email@email.com"
	u.FirstName = "Juan"
	u.LastName = "Dela Cruz"
	u.Username = "iamjuan"
	u.Avatar = "avatar.jpg"

	p := new(models.PostModel)
	p.Id = "100"
	p.User = u
	p.Message = "Hello Universe!"
	p.Attachments = make([]*models.AttachmentModel, 0)

	return 200, p
} 

//@Method : GET
//@Path   : /post
func (m *PostModule) GetPosts(params url.Values) (int, []models.PostModel) {
	return 200, make([]models.PostModel, 0)
}

//@Method : DELETE
//@Path   : /post
func (m *PostModule) DeletePost(params url.Values) (int, models.PostModel) {
	return 200, models.PostModel{}
}

//@Method : GET
//@Path   : /post/search
func (m *PostModule) SearchPost(params url.Values) (int, []models.PostModel) {
	return 200, make([]models.PostModel, 0)
}

//@Method : GET
//@Path   : /post/:id
func (m *PostModule) GetPost(params url.Values) (int, models.PostModel) {
	return 200, models.PostModel{}
}
