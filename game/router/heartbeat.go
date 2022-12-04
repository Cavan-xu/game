package router

import (
	"gameserver/game/service"
	"gameserver/public/rpc"
	"github.com/Cavan-xu/van/vnet"
	"github.com/golang/protobuf/proto"
)

type HeartBeatRouter struct {
	BaseRouter
}

func NewHeartBeatRouter() *HeartBeatRouter {
	return &HeartBeatRouter{
		BaseRouter{
			MsgId: rpc.MsgId_CToG_HeartBeat,
		},
	}
}

func (r *HeartBeatRouter) Handle(request vnet.IRequest) {
	req := &rpc.HeartBeatReq{}
	if err := proto.Unmarshal(request.GetMessage().GetData(), req); err != nil {
		request.GetServer().LogErr("protoId: %d proto Unmarshal data err: %v", request.GetMsgId(), err)
		return
	}

	p := service.GetPlayer(req.RoleId)
	if p == nil {
		return
	}

	p.CheckRoleMails()
}
