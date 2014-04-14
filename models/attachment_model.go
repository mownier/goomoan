// @filename attachment_model.go
// @author mownier

package models

import (
	"goomoan/goomoan"
)

type AttachmentModel struct {
	goomoan.Model
	File 		string	`json:"file"`
	Thumbnail	string	`json:"thumbnail"`
}
