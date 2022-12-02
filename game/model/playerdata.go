package model

type PlayerData struct {
	RoleId  int32  `xorm:"not null pk int"`
	Account string `xorm:"not null varchar(50)"`
	Name    string `xorm:"not null varchar(30)"`
}
