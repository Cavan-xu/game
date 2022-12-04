package router

import (
	"gameserver/game/service"
	"gameserver/public/rpc"

	"github.com/Cavan-xu/van/vnet"
	"github.com/golang/protobuf/proto"
)

type LoginRouter struct {
	BaseRouter
}

func NewLoginRouter() *LoginRouter {
	return &LoginRouter{
		BaseRouter{
			MsgId: rpc.MsgId_CToG_RoleLogin,
		},
	}
}

func (r *LoginRouter) Handle(request vnet.IRequest) {
	req := &rpc.RoleLoginReq{}
	if err := proto.Unmarshal(request.GetMessage().GetData(), req); err != nil {
		request.GetServer().LogErr("protoId: %d proto Unmarshal data err: %v", request.GetMsgId(), err)
		return
	}
	if req.RoleId <= 0 {
		return
	}

	p := service.CreatePlayer(req.RoleId)
	if !p.IsDataLoad() {
		ok, err := p.LoadData()
		if err != nil {
			request.GetServer().LogErr("role: %d load data err: %v", req.RoleId, err)
			r.SendFailAck(request, rpc.ResultId_System, err.Error())
			return
		}
		if !ok {
			ack := &rpc.RoleLoginAck{
				ResultId:      rpc.ResultId_Success,
				HasCreateRole: false,
			}
			r.SendMsgToClient(request, ack)
			return
		}
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
