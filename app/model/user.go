package model

type UserInfo struct {
	CID   string `json:"cid,omitempty" bson:"cid,omitempty"`
	CName string `json:"cname,omitempty" bson:"cname,omitempty"`
}
