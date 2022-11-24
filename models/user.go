package models

type User struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
}

type UserFull struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required" db:"password_hash"`
	IsAdmin  bool   `json:"is_admin" db:"is_admin"`
}

type SetAdmin struct {
	UserId  int  `json:"user_id"`
	IsAdmin bool `json:"is_admin"`
}
