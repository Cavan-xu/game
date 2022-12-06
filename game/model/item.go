package model

import (
	"gameserver/public/rpc"
	"sort"

	"github.com/golang/protobuf/proto"
)

type ItemRecord struct {
	BaseRecord
	*rpc.ItemData
	itemMap map[int32]*rpc.ItemInfo
}

func NewItemRecord() *ItemRecord {
	return &ItemRecord{
		BaseRecord: BaseRecord{
			fieldName: "item_data",
		},
		ItemData: &rpc.ItemData{},
		itemMap:  make(map[int32]*rpc.ItemInfo),
	}
}

func (record *ItemRecord) Unmarshal(data []byte) error {
	return proto.Unmarshal(data, record.ItemData)
}

func (record *ItemRecord) Marshal() ([]byte, error) {
	record.GetAll()
	return proto.Marshal(record.ItemData)
}

func (record *ItemRecord) GetAll() []*rpc.ItemInfo {
	record.ItemInfos = record.ItemInfos[0:0]
	for _, itemInfo := range record.itemMap {
		record.ItemInfos = append(record.ItemInfos, itemInfo)
	}
	sort.Slice(record.ItemInfos, func(i, j int) bool {
		return record.ItemInfos[i].ItemId < record.ItemInfos[j].ItemId
	})
	return record.ItemInfos
}

func (record *ItemRecord) Get(itemId int32) *rpc.ItemInfo {
	return record.itemMap[itemId]
}

func (record *ItemRecord) GetItemCount(itemId int32) int32 {
	itemInfo, ok := record.itemMap[itemId]
	if !ok {
		return 0
	}
	return itemInfo.Count
}
