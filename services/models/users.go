package models

type Users struct {
	UUID		int 	`json:"uuid" form:"-"`
	Username	string 	`json:"username" form:"username"`
	Password	string 	`json:"password" form:"password"`
}