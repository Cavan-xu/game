package model

type BaseRecord struct {
	isDirty   bool   // 是否存在脏数据
	fieldName string // 对应的数据库字段
}

func (record *BaseRecord) SetDirty(dirty bool) {
	record.isDirty = dirty
}

func (record *BaseRecord) IsDirty() bool {
	return record.isDirty
}

func (record *BaseRecord) GetFieldName() string {
	return record.fieldName
}
