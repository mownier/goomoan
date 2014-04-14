// @filename post_model.go
// @author mownier

package models

import (
	"goomoan/goomoan"
)

type UserModel struct {
	goomoan.Model
	Email 		string	`json:"email"`
	Username 	string	`json:"username"`
	FirstName	string	`json:"first_name"`
	LastName 	string	`json:"last_name"`
	Password 	string	`json:"password"`
	Avatar		string	`json:"avatar"`
}
