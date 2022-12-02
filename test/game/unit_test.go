package test

import (
	"gameserver/public/rpc"
	"io"
	"net"
	"sync"
	"testing"

	"github.com/Cavan-xu/van/vnet"
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func InitTcpClient() (*net.TCPConn, error) {
	rAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:12310")
	if err != nil {
		return nil, err
	}

	return net.DialTCP("tcp4", nil, rAddr)
}

func Test_RoleLogin(t *testing.T) {
	conn, err := InitTcpClient()
	assert.Nil(t, err)

	waitGroup := sync.WaitGroup{}
	go func() {
		for {
			select {
			default:
				head := make([]byte, 12) // 定长消息头长度
				_, err := io.ReadFull(conn, head)
				assert.Nil(t, err)
				pack := &vnet.DataPack{}
				message, err := pack.UnPack(head)
				var data []byte
				if message.GetDataLen() > 0 {
					data = make([]byte, message.GetDataLen())
					_, err := io.ReadFull(conn, data)
					assert.Nil(t, err)
				}
				ack := &rpc.RoleLoginAck{}
				err = proto.Unmarshal(data, ack)
				assert.Nil(t, err)
				t.Log(ack)
				waitGroup.Done()
				return
			}
		}
	}()

	waitGroup.Add(1)
	req := &rpc.RoleLoginReq{
		RoleId:  1,
		Account: "cavav.xu",
	}
	data, _ := proto.Marshal(req)
	message := &vnet.Message{
		MsgId:   uint32(rpc.MsgId_CToG_RoleLogin),
		ConnId:  1,
		DataLen: uint32(len(data)),
		Data:    data,
	}
	pack := vnet.NewDataPack()
	data = pack.Pack(message)
	_, err = conn.Write(data)
	assert.Nil(t, err)

	waitGroup.Wait()
}
