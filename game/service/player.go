package service

import (
	"context"
	"time"

	"github.com/Cavan-xu/van/vnet"
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
}

func (p *Player) Init() {
	p.ctx, p.cancel = context.WithCancel(context.TODO())

}

func (p *Player) IsDataLoad() bool {
	return p.isDataLoad
}

func (p *Player) LoadData() (bool, error) {
	return true, nil
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
