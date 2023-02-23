package dal

func AddRecord(fromId, toId int64, content, time string) (*ChatRecord, error) {

	var rec = &ChatRecord{
		FromUserId: fromId,
		ToUserId:   toId,
		Content:    content,
		CreateTime: time,
	}

	res := DB.Create(&rec)

	return rec, res.Error
}
