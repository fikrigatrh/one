package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name      string `json:"name" validate:"max=55"`
	Username  string `column:"user_name" json:"user_name" validate:"required,max=100"`
	Password  string `json:"password" validate:"required"`
	CreatedBy int    `json:"created_by"`
	UpdatedBy int    `json:"updated_by"`
}

type Auth struct {
	gorm.Model
	AuthUUID string `json:"auth_uuid"`
	Username string `json:"username"`
	UserID   string `json:"user_id"`
	SessID   string `json:"sess_id"`
}
