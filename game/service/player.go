package service

import "context"

type Player struct {
	ctx    context.Context
	cancel context.CancelFunc

	roleId  int32
	level   int32
	account string
	name    string

	isDataLoad bool
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

func (p *Player) Cancel() {
	p.cancel()
}
