// @filename post_model.go
// @author mownier

package models

import (
	"goomoan/goomoan"
)

type PostModel struct {
	goomoan.Model
	User		*UserModel			`json:"user"`
	Message 	string				`json:"message"`
	Attachments	[]*AttachmentModel	`json:"attachments"`
}

