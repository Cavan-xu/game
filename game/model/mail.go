package model

import (
	"gameserver/public/rpc"

	"github.com/golang/protobuf/proto"
)

type MailRecord struct {
	BaseRecord
	*rpc.MailData
}

func NewMailRecord() *MailRecord {
	return &MailRecord{
		BaseRecord: BaseRecord{
			fieldName: "mail_data",
		},
		MailData: &rpc.MailData{},
	}
}

func (record *MailRecord) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, record.MailData)
}

func (record *MailRecord) Marshal() ([]byte, error) {
	return proto.Marshal(record.MailData)
}
