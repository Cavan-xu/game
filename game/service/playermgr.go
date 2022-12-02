package service

import (
	"sync"
	"time"
)

const (
	PlayerMgrCount = 64
)

var (
	playerMgrArr = [PlayerMgrCount]*PlayerMgr{}
)

func init() {
	for i := 0; i < PlayerMgrCount; i++ {
		playerMgrArr[i] = &PlayerMgr{
			playerMap: make(map[int32]*Player),
		}
	}
}

func GetPlayer(roleId int32) *Player {
	mgr := playerMgrArr[roleId%PlayerMgrCount]
	return mgr.get(roleId)
}

func CreatePlayer(roleId int32) *Player {
	mgr := playerMgrArr[roleId%PlayerMgrCount]
	return mgr.create(roleId)
}

func CheckOffLine(tick int32) {
	mgr := playerMgrArr[tick%PlayerMgrCount]
	mgr.checkOffLine()
}

type PlayerMgr struct {
	sync.RWMutex
	playerMap map[int32]*Player
}

func (mgr *PlayerMgr) get(roleId int32) *Player {
	mgr.RLock()
	defer mgr.RUnlock()

	return mgr.playerMap[roleId]
}

func (mgr *PlayerMgr) create(roleId int32) *Player {
	mgr.Lock()
	defer mgr.Unlock()

	if p, ok := mgr.playerMap[roleId]; ok {
		return p
	}

	p := &Player{
		roleId: roleId,
	}
	p.Init()
	mgr.playerMap[roleId] = p
	return p
}

func (mgr *PlayerMgr) checkOffLine() {
	mgr.RLock()
	defer mgr.RUnlock()

	// 最大不活跃时间五分钟
	maxUnActiveTime := 300
	curTime := time.Now().Unix()
	for _, p := range mgr.playerMap {
		lastActiveTime := p.GetLastActiveTime().Unix()
		if int(curTime-lastActiveTime) >= maxUnActiveTime {
			// 关闭连接等动作
			delete(mgr.playerMap, p.GetRoleId())
		}
	}

}
