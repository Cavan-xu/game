package router

import (
	"gameserver/game/service"
	"gameserver/public/rpc"

	"github.com/Cavan-xu/van/vnet"
	"github.com/golang/protobuf/proto"
)

type RegisterRouter struct {
	BaseRouter
}

func NewRegisterRouter() *RegisterRouter {
	return &RegisterRouter{
		BaseRouter{
			MsgId: rpc.MsgId_CToG_RoleRegister,
		},
	}
}

func (r *RegisterRouter) Handle(request vnet.IRequest) {
	req := &rpc.RoleRegisterReq{}
	if err := proto.Unmarshal(request.GetMessage().GetData(), req); err != nil {
		request.GetServer().LogErr("protoId: %d proto Unmarshal data err: %v", request.GetMsgId(), err)
	}
	if req.RoleId <= 0 {
		r.SendFailAck(request, rpc.ResultId_InvalidParams, "roleId is empty")
		return
	}
	if req.Account == "" {
		r.SendFailAck(request, rpc.ResultId_InvalidParams, "account is empty")
		return
	}
	if req.Name == "" {
		r.SendFailAck(request, rpc.ResultId_InvalidParams, "name is empty")
		return
	}
	if req.Gender < 0 || req.Gender > 1 {
		req.Gender = 0
	}

	p := service.CreatePlayer(req.RoleId)
	if err := p.CreateData(req.RoleId, req.Gender, req.Account, req.Name); err != nil {
		r.SendFailAck(request, rpc.ResultId_System, err.Error())
		return
	}

	ack := &rpc.RoleLoginAck{
		ResultId:      rpc.ResultId_Success,
		RoleId:        p.GetRoleId(),
		Account:       p.GetAccount(),
		Name:          p.GetName(),
		Level:         p.GetLevel(),
		HasCreateRole: true,
		MailInfos:     p.MailRecord.GetAll(),
	}
	r.SendMsgToClient(request, ack)
}
