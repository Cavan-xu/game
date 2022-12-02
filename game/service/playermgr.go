package service

import "sync"

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
