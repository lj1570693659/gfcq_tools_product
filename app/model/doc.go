package model

type DocListInput struct {
	CID   string `json:"cid" bson:"cid"`
	CNAME string `json:"cname" bson:"cname"`
}

type DoctListInput struct {
	CID   string `v:"required#项目唯一标识不能为空" json:"cid" bson:"cid"`
	CNAME string `json:"cname" bson:"cname"`
}

type DocComputeInput struct {
	CID   string `v:"required#项目唯一标识不能为空" json:"cid" bson:"cid"`
	CNAME string `json:"cname" bson:"cname"`
}
