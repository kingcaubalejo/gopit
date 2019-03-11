package models

type Users struct {
	UUID		int 	`json:"uuid" form:"uuid"`
	Username	string 	`json:"username" form:"username"`
	Password	string 	`json:"password" form:"password"`
}