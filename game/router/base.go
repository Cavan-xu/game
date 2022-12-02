package router

import (
	"gameserver/public/Interceptor"
	"gameserver/public/rpc"

	"github.com/Cavan-xu/van/vnet"
	"github.com/golang/protobuf/proto"
)

type BaseRouter struct {
	MsgId rpc.MsgId
}

func (r *BaseRouter) GetMsgId() uint32 {
	return uint32(r.MsgId)
}

func (r *BaseRouter) PreHandle(request vnet.IRequest) {
	msgId := request.GetMessage().GetMsgId()
	if methodName, ok := rpc.MsgId_name[int32(msgId)]; ok {
		Interceptor.TcpRequestInterceptor(methodName)
	}
}

func (r *BaseRouter) Handle(request vnet.IRequest) {
	// implement me
}

func (r *BaseRouter) AfterHandle(request vnet.IRequest) {
	// implement me
}

func (r *BaseRouter) SendFailAck(request vnet.IRequest, resultId rpc.ResultId, errMsg string) {
	message := request.GetMessage()
	ack := &rpc.FailAck{
		ResultId: resultId,
		ErrMsg:   errMsg,
	}
	data, _ := proto.Marshal(ack)
	message.SetData(data)
	request.GetConnection().SendMsg(message)
}

func (r *BaseRouter) SendMsgToClient(request vnet.IRequest, ack proto.Message) {
	message := request.GetMessage()
	data, _ := proto.Marshal(ack)
	message.SetData(data)
	request.GetConnection().SendMsg(message)
}
