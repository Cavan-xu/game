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

func (record *MailRecord) HasReceiveSystemMail(mailId int32) bool {
	for _, id := range record.SystemMailId {
		if mailId == id {
			return true
		}
	}
	return false
}

func (record *MailRecord) Add(mail *rpc.MailInfo) {
	record.MailInfos = append(record.MailInfos, mail)
	record.SetDirty(true)
}

func (record *MailRecord) GetAll() []*rpc.MailInfo {
	return record.MailInfos
}

func (record *MailRecord) Get(mailId int64) *rpc.MailInfo {
	for _, mail := range record.MailInfos {
		if mail.MailId == mailId {
			return mail
		}
	}
	return nil
}

func (record *MailRecord) ReceiveSingleReward(mailId int64) ([]int64, []*rpc.ItemInfo) {
	mail := record.Get(mailId)
	if mail == nil || mail.Status == rpc.EMail_StatusReceiveReward {
		return nil, nil
	}

	mail.Status = rpc.EMail_StatusReceiveReward
	record.SetDirty(true)
	return []int64{mailId}, mail.ItemInfos
}

func (record *MailRecord) ReceiveAllReward() ([]int64, []*rpc.ItemInfo) {
	var (
		mailIds   []int64
		itemInfos []*rpc.ItemInfo
	)

	for _, mail := range record.MailInfos {
		if mail.Status == rpc.EMail_StatusReceiveReward {
			continue
		}
		mailIds = append(mailIds, mail.MailId)
		itemInfos = append(itemInfos, mail.ItemInfos...)
		mail.Status = rpc.EMail_StatusReceiveReward
	}
	if len(mailIds) > 0 {
		record.SetDirty(true)
	}

	return mailIds, itemInfos
}
