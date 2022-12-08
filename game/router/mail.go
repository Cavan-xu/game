package router

import (
	"gameserver/game/service"
	"gameserver/public/rpc"

	"github.com/Cavan-xu/van/vnet"
	"github.com/golang/protobuf/proto"
)

type MailRouter struct {
	BaseRouter
}

func NewMailRouter() *MailRouter {
	return &MailRouter{
		BaseRouter{
			MsgId: rpc.MsgId_CToG_ReceiveMailReward,
		},
	}
}

func (r *MailRouter) Handle(request vnet.IRequest) {
	req := &rpc.ReceiveMailRewardReq{}
	if err := proto.Unmarshal(request.GetMessage().GetData(), req); err != nil {
		request.GetServer().LogErr("protoId: %d proto Unmarshal data err: %v", request.GetMsgId(), err)
		return
	}

	p := service.GetPlayer(req.RoleId)
	if p == nil {
		return
	}

	var (
		mailIds   []int64
		itemInfos []*rpc.ItemInfo
	)
	switch req.ReceiveTYpe {
	case rpc.EMail_ReceiveTypeSingle:
		mailIds, itemInfos = p.MailRecord.ReceiveSingleReward(req.MailId)
	case rpc.EMail_ReceiveTypeAll:
		mailIds, itemInfos = p.MailRecord.ReceiveAllReward()
	}

	ack := &rpc.ReceiveMailRewardAck{
		ResultId:  rpc.ResultId_Success,
		ItemInfos: itemInfos,
	}
	r.SendMsgToClient(request, ack)

	if len(mailIds) > 0 {
		NotifyRoleMails(p, mailIds)
	}
}

type DeleteMailRouter struct {
	BaseRouter
}

func (r *DeleteMailRouter) Handle(request vnet.IRequest) {
	req := &rpc.CommonReq{}
	if err := proto.Unmarshal(request.GetMessage().GetData(), req); err != nil {
		request.GetServer().LogErr("protoId: %d proto Unmarshal data err: %v", request.GetMsgId(), err)
		return
	}

	p := service.GetPlayer(req.RoleId)
	if p == nil {
		return
	}

	mailIds := p.MailRecord.RemoveReadMails()
	ack := &rpc.CommonAck{
		ResultId: rpc.ResultId_Success,
	}
	r.SendMsgToClient(request, ack)

	NotifyRoleMails(p, mailIds)
}

func NotifyRoleMails(p *service.Player, mailIds []int64) {
	ntf := &rpc.RoleMailNtf{
		RoleId: p.GetRoleId(),
	}
	for _, mailId := range mailIds {
		mail := p.MailRecord.Get(mailId)
		if mail == nil {
			mail = &rpc.MailInfo{
				MailId: mailId,
				Status: rpc.EMail_StatusDeleted,
			}
		}
		ntf.MailInfos = append(ntf.MailInfos, mail)
	}
	p.SendMessageToClient(rpc.MsgId_GToC_RoleMailNtf, ntf)
}
