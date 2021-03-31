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

	AnnouncementsAdd
	AnnouncementsDel
	AnnouncementsUpdate
	AnnouncementsQuery
	AnnouncementsPages

	AuthAdd
	AuthDel
	AuthUpdate
	AuthQuery
	AuthPages

	CoinAdd
	CoinDel
	CoinUpdate
	CoinQuery
	CoinPages

	CoinChainAdd
	CoinChainDel
	CoinChainUpdate
	CoinChainQuery
	CoinChainPages

	GameAdd
	GameDel
	GameUpdate
	GameQuery
	GamePages

	GameGroupAdd
	GameGroupDel
	GameGroupUpdate
	GameGroupQuery
	GameGroupPages

	GameResultAdd
	GameResultDel
	GameResultUpdate
	GameResultQuery
	GameResultPages

	GameTypeAdd
	GameTypeDel
	GameTypeUpdate
	GameTypeQuery
	GameTypePages

	GradeAdd
	GradeDel
	GradeUpdate
	GradeQuery
	GradePages

	IssueFactoryAdd
	IssueFactoryDel
	IssueFactoryUpdate
	IssueFactoryQuery
	IssueFactoryPages

	LogOperateAdd
	LogOperateDel
	LogOperateUpdate
	LogOperateQuery
	LogOperatePages

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

	TmpSeedAdd
	TmpSeedDel
	TmpSeedUpdate
	TmpSeedQuery
	TmpSeedPages

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

	UserSecurityQuestionsAdd
	UserSecurityQuestionsDel
	UserSecurityQuestionsUpdate
	UserSecurityQuestionsQuery
	UserSecurityQuestionsPages

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

	UsersAdd
	UsersDel
	UsersUpdate
	UsersQuery
	UsersPages
)
