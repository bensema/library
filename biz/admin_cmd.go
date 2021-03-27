package biz

type AdminRequest struct {
	Operator   string   `json:"operator"`
	OperatorId string   `json:"operator_id"`
	Token      string   `json:"token"`
	Auth       string   `json:"auth"`
	Ip         string   `json:"ip"`
	T          int64    `json:"t"`
	R          string   `json:"r"`
	Cmd        AdminCmd `json:"cmd"`
	Method     string   `json:"method"`
	Data       string   `json:"data"`
}

type AdminCmd int

const (
	_ AdminCmd = iota
	AdvertiseAdd
	AdvertiseDel
	AdvertiseUpdate
	AdvertiseQuery
	AdvertisePages
	AnnouncementAdd
	AnnouncementDel
	AnnouncementUpdate
	AnnouncementQuery
	AnnouncementPages
)
