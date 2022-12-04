package service

import (
	"context"
	"errors"
	"gameserver/game/model"
	"gameserver/public/db"
	"time"

	"github.com/Cavan-xu/van/vnet"
	"gorm.io/gorm"
)

type Player struct {
	ctx            context.Context
	cancel         context.CancelFunc
	roleId         int32
	level          int32
	account        string
	name           string
	isOnline       bool
	isDataLoad     bool
	loginTime      time.Time
	lastActiveTime time.Time
	connection     vnet.IConnection

	MailRecord *model.MailRecord
}

func (p *Player) Init() {
	p.ctx, p.cancel = context.WithCancel(context.TODO())
}

func (p *Player) IsDataLoad() bool {
	return p.isDataLoad
}

func (p *Player) CreateData(roleId, gender int32, account, name string) error {
	playerData := model.NewPlayerData(roleId, gender, account, name)
	if err := db.GetGameDB().Table(playerData.GetTableName()).Create(p).Error; err != nil {
		return err
	}

	return nil
}

func (p *Player) LoadData() (bool, error) {
	playerData := &model.PlayerData{RoleId: p.roleId}
	err := db.GetGameDB().Table(playerData.GetTableName()).First(p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false, nil
	}
	if err != nil {
		return false, err
	}
	if err = p.LoadRecord(playerData); err != nil {
		return false, err
	}

	return true, nil
}

func (p *Player) LoadRecord(playerData *model.PlayerData) error {
	mailRecord := model.NewMailRecord()
	if err := mailRecord.Unmarshal(playerData.MailData); err != nil {
		return err
	}
	p.MailRecord = mailRecord

	return nil
}

func (p *Player) GetRoleId() int32 {
	return p.roleId
}

func (p *Player) GetAccount() string {
	return p.account
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetLevel() int32 {
	return p.level
}

func (p *Player) GetLastActiveTime() time.Time {
	return p.lastActiveTime
}

func (p *Player) Cancel() {
	p.cancel()
}
