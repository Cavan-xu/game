package model

import (
	"fmt"
)

var (
	PlayerDataTableCount int32 = 100 // 分一百张表
)

type PlayerData struct {
	RoleId   int32  `xorm:"not null pk int"`
	Account  string `xorm:"not null varchar(50)"`
	Name     string `xorm:"not null varchar(30)"`
	Gender   int32  `xorm:"not null int(1)"`
	MailData []byte `xorm:"blob"`
}

func NewPlayerData(roleId, gender int32, account, name string) *PlayerData {
	return &PlayerData{
		RoleId:  roleId,
		Gender:  gender,
		Account: account,
		Name:    name,
	}
}

func (p *PlayerData) GetTableName() string {
	return fmt.Sprintf("player_data_%d", p.RoleId%PlayerDataTableCount)
}
