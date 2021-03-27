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
	GameAdd
	GameDel
	GameUpdate
	GameQuery
	GamePages

	GameTypeAdd
	GameTypeDel
	GameTypeUpdate
	GameTypeQuery
	GameTypePages
	GameTypeFind
	GameTypeAll

	GameGroupAdd
	GameGroupDel
	GameGroupUpdate
	GameGroupQuery
	GameGroupPages
	GameGroupFind
	GameGroupAll

	CoinAdd
	CoinDel
	CoinUpdate
	CoinQuery
	CoinPages
	CoinFind
	CoinAll

	CoinChainAdd
	CoinChainDel
	CoinChainUpdate
	CoinChainQuery
	CoinChainPages
	CoinChainFind
	CoinChainAll

	GradeAdd
	GradeDel
	GradeUpdate
	GradeQuery
	GradePages
	GradeFind
	GradeAll

	LogOperateAdd
	LogOperateDel
	LogOperateUpdate
	LogOperateQuery
	LogOperatePages

	LogLoginAdd
	LogLoginDel
	LogLoginUpdate
	LogLoginQuery
	LogLoginPages

	LogUserLoginAdd
	LogUserLoginDel
	LogUserLoginUpdate
	LogUserLoginQuery
	LogUserLoginPages

	PowerAdd
	PowerDel
	PowerUpdate
	PowerQuery
	PowerPages

	SecurityQuestionsAdd
	SecurityQuestionsDel
	SecurityQuestionsUpdate
	SecurityQuestionsQuery
	SecurityQuestionsPages

	UserInfoAdd
	UserInfoDel
	UserInfoUpdate
	UserInfoQuery
	UserInfoPages

	UserPowerAdd
	UserPowerDel
	UserPowerUpdate
	UserPowerQuery
	UserPowerPages

	UserSecurityQuestionAdd
	UserSecurityQuestionDel
	UserSecurityQuestionUpdate
	UserSecurityQuestionQuery
	UserSecurityQuestionPages

	UserTokenAdd
	UserTokenDel
	UserTokenUpdate
	UserTokenQuery
	UserTokenPages

	UserWalletAdd
	UserWalletDel
	UserWalletUpdate
	UserWalletQuery
	UserWalletPages

	UserAdd
	UserDel
	UserUpdate
	UserQuery
	UserPages
)
