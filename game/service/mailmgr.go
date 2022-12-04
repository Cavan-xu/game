package service

import (
	"gameserver/public/rpc"
	"sync"
)

type MailMgr struct {
	sync.RWMutex
	mailArr []*rpc.SystemMail
}

var (
	mailMgr *MailMgr
)

func init() {
	mailMgr = &MailMgr{
		mailArr: make([]*rpc.SystemMail, 0),
	}
}

func GetSystemMails() []*rpc.SystemMail {
	return mailMgr.GetSystemMails()
}

func (mgr *MailMgr) LoadData() error {
	return nil
}

func (mgr *MailMgr) GetSystemMails() []*rpc.SystemMail {
	mgr.RLock()
	defer mgr.Unlock()

	return mgr.mailArr
}

func (mgr *MailMgr) SetSystemMails(mails []*rpc.SystemMail) {
	mgr.Lock()
	defer mgr.Unlock()

	mgr.mailArr = mails
}
