package ecode

// main ecode interval is [0,990000]
var (
	UsernameIsEmpty = New(1000, " 账户名不能为空") // 账户名不能为空
	PasswordIsEmpty = New(1001, " 密码不能为空")  // 密码不能为空

	// archive
	ArchiveTypeForbidAdd = New(10001, " 该类型不支持投稿")                           // 该类型不支持投稿
	ArchiveExist         = New(10002, " 已经存在该稿件了")                           // 已经存在该稿件了
	ArchiveNotExist      = New(10003, " 不存在该稿件")                             // 不存在该稿件
	ArchiveAlreadyDel    = New(10004, " 稿件已经被删除")                            // 稿件已经被删除
	ArchiveDelayTimeErr  = New(10005, " 延迟变更的时间早于当前")                        // 延迟变更的时间早于当前
	ArchiveAlreadyDelay  = New(10006, " 延迟变更的时间已存在")                         // 延迟变更的时间已存在
	ArchiveOwnerErr      = New(10007, " 稿件不属于这个人")                           // 稿件不属于这个人
	VideoshotNotExist    = New(10008, " 稿件的缩略图不存在")                          // 稿件的缩略图不存在
	VideoAbnormalSubmit  = New(10009, " 异常视频提交")                             // 异常视频提交
	ArchiveBlocked       = New(10010, " 当前稿件已锁定，可能处于以下状态之一(审核锁定，网警锁定,用户删除)") // 当前稿件已锁定，可能处于以下状态之一(审核锁定，网警锁定,用户删除)

	// appeal
	AppealNotExist     = New(10101, " 不存在该申诉")          // 不存在该申诉
	AppealAlreadyClose = New(10102, " 申诉工单已经被关闭")       // 申诉工单已经被关闭
	AppealInterval     = New(10103, " 申诉间隔时间内，不能再发起申诉") // 申诉间隔时间内，不能再发起申诉
	AppealOpen         = New(10104, " 该稿件已处于申诉中")       // 该稿件已处于申诉中
	AppealOwner        = New(10105, " 只能申诉自己的稿件")       // 只能申诉自己的稿件
	AppealHasStar      = New(10106, " 该申诉已评分")          // 该申诉已评分
	AppealLimit        = New(10107, " 仅支持打回和锁定的稿件申诉")   // 仅支持打回和锁定的稿件申诉

	// favorite
	FavNameTooLong        = New(11001, " 收藏夹名称过长")            // 收藏夹名称过长
	FavMaxFolderCount     = New(11002, " 达到最大收藏夹数")           // 达到最大收藏夹数
	FavCanNotDelDefault   = New(11003, " 不能删除默认收藏夹")          // 不能删除默认收藏夹
	FavFolderNoPublic     = New(11004, " 收藏夹目录未公开")           // 收藏夹目录未公开
	FavMaxVideoCount      = New(11005, " 视频数达到目录最大收藏数")       // 视频数达到目录最大收藏数
	FavFolderExist        = New(11006, " 已经存在该收藏夹")           // 已经存在该收藏夹
	FavVideoExist         = New(11007, " 已经存在该视频了")           // 已经存在该视频了
	FavOnlyPublic         = New(11008, " 仅仅只能设置公开0 或 非公开1")   // 仅仅只能设置公开0 或 非公开1
	FavDefaultFolder      = New(11009, " 默认收藏夹")              // 默认收藏夹
	FavFolderNotExist     = New(11010, " 没有该收藏夹")             // 没有该收藏夹
	FavSearchReqErr       = New(11011, " 请求搜索出错")             // 请求搜索出错
	FavFloderAlreadyDel   = New(11012, " 已经删除该收藏夹了")          // 已经删除该收藏夹了
	FavVideoAlreadyDel    = New(11013, " 已经取消收藏该视频了")         // 已经取消收藏该视频了
	FavMaxOperNum         = New(11014, " 超出允许的最大操作数75")       // 超出允许的最大操作数75
	FavFolderSame         = New(11015, " 一样的收藏夹")             // 一样的收藏夹
	FavFolderMoveFailed   = New(11016, " 收藏夹视频移动失败")          // 收藏夹视频移动失败
	FavFolderSortErr      = New(11017, " 收藏夹列表信息错误")          // 收藏夹列表信息错误
	FavSearchWordIllegal  = New(11018, " 收藏夹视频搜索关键词非法")       // 收藏夹视频搜索关键词非法
	FavMaintenance        = New(11019, " 收藏服务维护中")            // 收藏服务维护中
	FavFolderBanned       = New(11020, " 收藏夹名称包含敏感词")         // 收藏夹名称包含敏感词
	FavCleaneInProgress   = New(11021, " 正在删除失效视频…请过段时间再来访问") // 正在删除失效视频…请过段时间再来访问
	FavCleanedLocked      = New(11022, " 清除操作锁定中")            // 清除操作锁定中
	FavResourceExist      = New(11201, " 已经收藏过了")             // 已经收藏过了
	FavResourceAlreadyDel = New(11202, " 已经取消收藏了")            // 已经取消收藏了
	FavResourceOverflow   = New(11203, " 达到收藏上限")             // 达到收藏上限
	FavDescTooLang        = New(11204, " 收藏夹描述过长")            // 收藏夹描述过长
	FavRetryLater         = New(11205, " 请稍后重试")              // 请稍后重试
	FavHitSensitive       = New(11206, " 收藏夹命中敏感词")           // 收藏夹命中敏感词

	// reply
	ReplySubjectExist      = New(12001, " 已经存在评论主题")             // 已经存在评论主题
	ReplyForbidReply       = New(12002, " 禁止评论")                 // 禁止评论
	ReplyForbidRootReply   = New(12003, " 禁止回复")                 // 禁止回复
	ReplyForbidAction      = New(12004, " 禁止操作： 赞或踩")            // 禁止操作： 赞或踩
	ReplyForbidReport      = New(12005, " 禁止举报")                 // 禁止举报
	ReplyNotExist          = New(12006, " 没有该评论")                // 没有该评论
	ReplyActioned          = New(12007, " 已经操作过了")               // 已经操作过了
	ReplyReported          = New(12008, " 已经举报过了")               // 已经举报过了
	ReplyIllegalSubType    = New(12009, " 评论主体的type不合法")         // 评论主体的type不合法
	ReplyIllegalRoot       = New(12010, " 不合法的父评论")              // 不合法的父评论
	ReplyIllegalAction     = New(12011, " 不合法的赞或踩")              // 不合法的赞或踩
	ReplyIllegalReport     = New(12012, " 不合法的举报")               // 不合法的举报
	ReplyTooManyAts        = New(12013, " at人数过多")               // at人数过多
	ReplyDeniedAsCD        = New(12014, " cd时间未到不能评论")           // cd时间未到不能评论
	ReplyDeniedAsCaptcha   = New(12015, " 验证码错误未到不能评论")          // 验证码错误未到不能评论
	ReplyDeniedByFilter    = New(12016, " 评论内容被 word filter 禁止") // 评论内容被 word filter 禁止
	ReplyMosaicByFilter    = New(12017, " 评论内容被过滤掉敏感词打***")      // 评论内容被过滤掉敏感词打***
	ReplyUpDeniedAsState   = New(12018, " 目前状态不能执行up主操作")        // 目前状态不能执行up主操作
	ReplyReportDeniedAsCD  = New(12019, " cd时间未到不能举报")           // cd时间未到不能举报
	ReplyReportNotExist    = New(12020, " 举报不存在")                // 举报不存在
	ReplyReportDealed      = New(12021, " 已经处理该举报了")             // 已经处理该举报了
	ReplyDeleted           = New(12022, " 已经被删除了")               // 已经被删除了
	ReplyRepeat            = New(12023, " 短时间内发重复评论")            // 短时间内发重复评论
	ReplyNoNeedCaptcha     = New(12024, " 用户不需要验证码")             // 用户不需要验证码
	ReplyContentOver       = New(12025, " 评论内容过少或过多")            // 评论内容过少或过多
	ReplyDeniedByLink      = New(12026, " 一级用户评论里含有链接")          // 一级用户评论里含有链接
	ReplyPending           = New(12027, " 评论进入待审")               // 评论进入待审
	ReplyDeniedAsGarbage   = New(12028, " 大数据垃圾评论")              // 大数据垃圾评论
	ReplyHaveTop           = New(12029, " 已经有置顶评论")              // 已经有置顶评论
	ReplyNotRootReply      = New(12030, " 不能置顶非一级评论")            // 不能置顶非一级评论
	ReplyIllegalSubState   = New(12031, " sub状态不合法")             // sub状态不合法
	ReplyEmojiOverMax      = New(12032, " vip表情超过上限")            // vip表情超过上限
	ReplyVipStatusIllegal  = New(12033, " vip状态异常无法使用评论特权")      // vip状态异常无法使用评论特权
	ReplyDelTopForbidden   = New(12034, " 置顶评论不允许删除")            // 置顶评论不允许删除
	ReplyBlacklistFilter   = New(12035, " 黑名单屏蔽不予评论")            // 黑名单屏蔽不予评论
	ReplyUpgrading         = New(12036, " 系统升级中")                // 系统升级中
	ReplyUserTelVerified   = New(12037, " 请绑定手机号后重试")            // 请绑定手机号后重试
	ReplyForbidList        = New(12038, " 评论已关闭")                // 评论已关闭
	ReplyHitBlacklist      = New(12039, " 评论含有违规内容")             // 评论含有违规内容
	ReplyOverRateLimit     = New(12040, " 一段时间内相似内容出现过多，请稍后再试")  // 一段时间内相似内容出现过多，请稍后再试
	ReplyContestNotPassed  = New(12041, " 该账号处于封禁中,点击申请答题")      // 该账号处于封禁中,点击申请答题
	ReplyNoticeConflict    = New(12042, " 时间冲突，发布公告失败")          // 时间冲突，发布公告失败
	ReplySubjectFrozen     = New(12043, " 该评论区已冻结")              // 该评论区已冻结
	ReplyEmojiExits        = New(12044, " 表情已存在")                // 表情已存在
	ReplyForbidReplyNotPay = New(12045, " 禁止评论，未付费")             // 禁止评论，未付费
	ReplyForbidFolded      = New(12048, " 该评论禁止被折叠")             // 该评论禁止被折叠
	ReplyFolded            = New(12049, " 该评论被折叠")               // 该评论被折叠

	// elec
	ElecUserForbid    = New(13001, " 用户被禁止参与充电计划")  // 用户被禁止参与充电计划
	ElecUserAudit     = New(13002, " 用户的是否充电正在审核中") // 用户的是否充电正在审核中
	ElecNotUpper      = New(13003, " 不是up主")        // 不是up主
	ElecArchiveForbid = New(13004, " 稿件被屏蔽")        // 稿件被屏蔽

	// stat
	ClickAesDecryptErr  = New(14001, " aes解密失败")      // aes解密失败
	ClickQueryFormatErr = New(14002, " 解密粗来的参数格式有问题") // 解密粗来的参数格式有问题
	ClickServerTimeout  = New(14003, " 服务端时间超时")      // 服务端时间超时
	ClickQuerySignErr   = New(14004, " 参数中的sign错误")   // 参数中的sign错误
	ClickHmacSignErr    = New(14005, " hmac算的签名错误")   // hmac算的签名错误

	// topic
	TopicNotExist      = New(15001, " 不存在该话题")     // 不存在该话题
	FavTopicExist      = New(15002, " 已经存在该话题了")   // 已经存在该话题了
	FavTopicAlreadyDel = New(15003, " 已经取消收藏该话题了") // 已经取消收藏该话题了

	// tag
	TagOperateFail         = New(16000, " 操作失败")                // 操作失败
	TagNotExist            = New(16001, " tag不存在")              // tag不存在
	TagAidNotExist         = New(16002, " 视频tag关联不存在")          // 视频tag关联不存在
	TagArcTagMaxNum        = New(16003, " 超过视频最大tag数")          // 超过视频最大tag数
	TagMaxSub              = New(16004, " 超过个人最大订阅数")           // 超过个人最大订阅数
	TagOpDenied            = New(16005, " tag操作权限不够")           // tag操作权限不够
	TagAidNoTags           = New(16006, " 该视频不存在tag")           // 该视频不存在tag
	TagRidNotExist         = New(16007, " rid分区不存在")            // rid分区不存在
	TagAvNotExist          = New(16008, " 该视频不存在")              // 该视频不存在
	TagArcTagExist         = New(16009, " 视频存在此tag")            // 视频存在此tag
	TagArcTagAddMaxFre     = New(16010, " 添加太多标签啦,休息休息~")       // 添加太多标签啦,休息休息~
	TagArcTagDelMaxFre     = New(16011, " 删除太多标签啦,休息休息~")       // 删除太多标签啦,休息休息~
	TagArcTagLikeMaxFre    = New(16012, " 顶/踩太多次啦,休息休息~")       // 顶/踩太多次啦,休息休息~
	TagArcTagRptMaxFre     = New(16013, " 举报太多次啦,休息休息~")        // 举报太多次啦,休息休息~
	TagArcCannotAddTag     = New(16014, " 管理员暂时不让添加标签哦~")       // 管理员暂时不让添加标签哦~
	TagArcCannotDelTag     = New(16015, " 管理员暂时不让删除标签哦~")       // 管理员暂时不让删除标签哦~
	TagArcAddLevelLower    = New(16016, " 升级到Lv3就能添加标签啦~")      // 升级到Lv3就能添加标签啦~
	TagUpAddNotAllowed     = New(16017, " 然而up主并不希望你添加~")       // 然而up主并不希望你添加~
	TagLikeNotDelAllowed   = New(16018, " 被大家顶起来的标签,不能随便删除哦~")  // 被大家顶起来的标签,不能随便删除哦~
	TagArcTagNotExist      = New(16020, " 视频不存在此tag")           // 视频不存在此tag
	TagArcTagLiked         = New(16021, " 已经顶过啦~")              // 已经顶过啦~
	TagArcTagHated         = New(16022, " 已经踩过啦~")              // 已经踩过啦~
	TagArcTagisLocked      = New(16023, " 已经是锁定状态")             // 已经是锁定状态
	TagArcTagnoLocked      = New(16024, " 已经是非锁定状态")            // 已经是非锁定状态
	TagIsSealing           = New(16025, " tag已经被封印了~")          // tag已经被封印了~
	TagArcDelLevelLower    = New(16026, " 升级到Lv4就能删除标签啦~")      // 升级到Lv4就能删除标签啦~
	TagArcRptLevelLower    = New(16027, " 升级到Lv1就能举报标签啦~")      // 升级到Lv1就能举报标签啦~
	TagArcLikeLevelLower   = New(16028, " 升级到Lv1就能喜欢标签啦~")      // 升级到Lv1就能喜欢标签啦~
	TagArcHateLevelLower   = New(16029, " 升级到Lv1就能讨厌标签啦~")      // 升级到Lv1就能讨厌标签啦~
	TagTagIsSubed          = New(16030, " tag 已经订阅啦~")          // tag 已经订阅啦~
	TagRankingTagNotExist  = New(16031, " 排行榜热门tag不存在")         // 排行榜热门tag不存在
	TagRankingBsNotExist   = New(16032, " 排行榜番剧不存在")            // 排行榜番剧不存在
	TagRankingPsNotExist   = New(16033, " 排行榜分区不存在")            // 排行榜分区不存在
	TagRankingSimNotExist  = New(16034, " 排行榜分区不存在")            // 排行榜分区不存在
	TagNotSub              = New(16035, " 未订阅")                 // 未订阅
	TagArcTagRpted         = New(16036, " 已经举报啦~")              // 已经举报啦~
	TagArcTagDelMore       = New(16037, " 超过单个视频日删除次数上线~")      // 超过单个视频日删除次数上线~
	TagRptNotRptPassed     = New(16038, " 已经被审核通过的,不能再继续举报哦~")  // 已经被审核通过的,不能再继续举报哦~
	TagArcTagLogNotExsit   = New(16039, " 视频tag操作日志不存在")        // 视频tag操作日志不存在
	TagDelNotRptPassed     = New(16040, " 该标签已经审核为正常标签,不能删除哦~") // 该标签已经审核为正常标签,不能删除哦~
	TagUpTagCannotDel      = New(16041, " 不能删除UP主添加的tag哦~")     // 不能删除UP主添加的tag哦~
	TagAddNotRptPassed     = New(16042, " 该标签已经认定为恶意标签,不能添加哦~") // 该标签已经认定为恶意标签,不能添加哦~
	TagAdminOpCanNotTpt    = New(16043, " 管理员操作的标签不能举报~")       // 管理员操作的标签不能举报~
	TagArcAccountBlocked   = New(16044, " 你被封禁，禁止操作~")          // 你被封禁，禁止操作~
	TagIsOfficailTag       = New(16045, " 此tag是官方活动tag,不可以操作~") // 此tag是官方活动tag,不可以操作~
	TagAlreadyExist        = New(16046, " Tag已存在，不能重复添加~")      // Tag已存在，不能重复添加~
	TagResTagMaxNum        = New(16047, " 资源绑定tag超过最大数量~")      // 资源绑定tag超过最大数量~
	TagResTagExist         = New(16048, " 资源绑定tag关系已存在~")       // 资源绑定tag关系已存在~
	TagResTagNotExist      = New(16049, " 资源绑定tag关系不存在~")       // 资源绑定tag关系不存在~
	TagServiceUpdate       = New(16050, " Tag服务升级中~")           // Tag服务升级中~
	TagUpdateLimitUserFail = New(16051, " Tag服务更新白名单用户失败")      // Tag服务更新白名单用户失败
	TagLimitUserExist      = New(16052, " Tag服务白名单用户已存在")       // Tag服务白名单用户已存在
	TagLimitUserFail       = New(16053, " Tag服务获取白名单用户信息失败")    // Tag服务获取白名单用户信息失败
	TagResTagFail          = New(16054, " 资源绑定tag失败")           // 资源绑定tag失败
	TagReportNotExist      = New(16055, " 该举报不存在")              // 该举报不存在
	TagReportIsDealed      = New(16056, " 该举报已经处理无法移动")         // 该举报已经处理无法移动
	TagPunishMsgFailed     = New(16057, " Tag惩罚通知用户失败")         // Tag惩罚通知用户失败
	TagReportLogAddFailed  = New(16058, " Tag举报日志增加失败")         // Tag举报日志增加失败
	TagReportUpdateFailed  = New(16059, " Tag举报更新失败")           // Tag举报更新失败
	TagHotUpdateFailed     = New(16060, " 热门Tag更新失败")           // 热门Tag更新失败
	TagAlreadyExamined     = New(16061, " tag已经审核过了")           // tag已经审核过了
	TagAlreadyShield       = New(16062, " tag已经被屏蔽")            // tag已经被屏蔽
	TagArcTagLockFail      = New(16063, " tag稿件锁定失败")           // tag稿件锁定失败
	TagUpdateFail          = New(16064, " tag更新失败")             // tag更新失败
	TagAddFail             = New(16065, " tag增加失败")             // tag增加失败
	TagChangeStateFail     = New(16066, " tag更改状态失败")           // tag更改状态失败
	TagArcTagUnLockFail    = New(16067, " tag稿件解除锁定失败")         // tag稿件解除锁定失败
	TagReportLogNotExist   = New(16068, " 举报日志不存在")             // 举报日志不存在
	TagAlreadyDelete       = New(16069, " tag已经被删除")            // tag已经被删除
	TagOnlyUpAdd           = New(16070, " 只有UP主可以添加频道")         // 只有UP主可以添加频道
	TagOnlyUpDel           = New(16071, " 只有UP主可以删除频道")         // 只有UP主可以删除频道
	ChannelAlreadyExist    = New(16101, " 频道已存在")               // 频道已存在
	ChannelNotExist        = New(16102, " 频道不存在")               // 频道不存在
	ChanRuleException      = New(16103, "频道规则有异常")              //频道规则有异常
	ChanTypeNotExist       = New(16104, "频道类型不存在")              //频道类型不存在
	ChanRuleCanotUse       = New(16105, "频道规则不符合标准")            //频道规则不符合标准
	ChannelNoLike          = New(16106, "频道不能点赞")               //频道不能点赞
	ChannelNoHate          = New(16107, "频道不能点踩")               //频道不能点踩
	ChannelNotAllowDel     = New(16108, "这个不能删除哦~")             //这个不能删除哦~
	ChannelAINoData        = New(16109, "该频道下没数据了哦")            //该频道下没数据了哦
	ChannelAITimeout       = New(16110, "AI超时了，技能冷却中")          //AI超时了，技能冷却中
	ChannelAleadyMigrated  = New(16111, "频道已经被迁移过了")            //频道已经被迁移过了
	ChannelCanNotDel       = New(16112, "频道分类下有数据，禁止操作")        //频道分类下有数据，禁止操作
	ChannelRecommended     = New(16113, "频道已经被推荐了，无须再推荐")       //频道已经被推荐了，无须再推荐
	ChannelTypeExist       = New(16114, "频道分类已存在")              //频道分类已存在
	ChannelTypeDeled       = New(16115, "频道分类已删除")              //频道分类已删除
	ChannelNoChange        = New(16116, "频道没变化")                //频道没变化
	CategoryNoChange       = New(16117, "频道分类没变化")              //频道分类没变化

	// zlimit
	ZlimitAllow     = New(17001, " 稿件无限制")    // 稿件无限制
	ZlimitForbidden = New(17002, " 稿件禁止查看")   // 稿件禁止查看
	ZlimitFormal    = New(17003, " 正式会员")     // 正式会员
	ZlimitPay       = New(17004, " 付费会员")     // 付费会员
	ZlimitShared    = New(17005, " 共享地址错误提示") // 共享地址错误提示

	// Feedback
	FeedbackSsnNotExist  = New(18001, " session不存在") // session不存在
	FeedbackBodyTooLarge = New(18002, " 上传的文件太大")    // 上传的文件太大
	FeedbackBodyNotExist = New(18003, " 上传的文件无内容")   // 上传的文件无内容
	FeedbackContentOver  = New(18004, "创建会话过多")      //创建会话过多

	// short utl
	ShortURLAlreadyExist = New(30001, " 短链已存在")        // 短链已存在
	ShortURLNotFound     = New(30002, " 短链不存在")        // 短链不存在
	ShortURLIllegalSrc   = New(30003, " 生成短链的长链来源不合法") // 生成短链的长链来源不合法

	// captcha
	CaptchaSignErr            = New(33001, " 签名错误")                // 签名错误
	CaptchaTSOverTolerant     = New(33002, " 时间戳超过规定时间范围")         // 时间戳超过规定时间范围
	CaptchaBusinessNotAllowed = New(33003, " 业务id不存在")             // 业务id不存在
	CaptchaCodeNotFound       = New(33004, " code不存在，一般是token已过期") // code不存在，一般是token已过期
	CaptchaTokenErr           = New(33005, " 验证时token错误")          // 验证时token错误
	CaptchaNotCreate          = New(33006, " 验证码未创建")              // 验证码未创建
	CaptchaTokenNotExist      = New(33007, " 验证码Token不存在")         // 验证码Token不存在
	CaptchaTokenExpired       = New(33008, " 验证码已过期")              // 验证码已过期

	// Coin
	CoinRatingErr       = New(34001, " 评级参数错误")     // 评级参数错误
	CoinCannotAddToSelf = New(34002, " up主不能自己投币")  // up主不能自己投币
	CoinIllegaMultiply  = New(34003, " 非法的投币数量")    // 非法的投币数量
	CoinDeniedAsCD      = New(34004, " 投币间隔太短")     // 投币间隔太短
	CoinOverMax         = New(34005, " 超过单个视频投币上限") // 超过单个视频投币上限

	// account
	AccountOverdue     = New(35001, " token过期 -658")    // token过期 -658
	AccountNotLogin    = New(35002, " 未登录 -400")        // 未登录 -400
	AccountInexistence = New(35003, " 用户不存在 -626")      // 用户不存在 -626
	AccountAKNotFound  = New(35004, " accessKey不存在 -2") // accessKey不存在 -2

	// player-interface
	PLayerPolicyNotExist = New(19001, " 策略不存在") // 策略不存在
	PLayerPolicyNotStart = New(19002, " 策略未开始") // 策略未开始
	PLayerPolicyEnded    = New(19003, " 策略未开始") // 策略未开始

	// creative
	CreativeArchiveAPIErr                    = New(20001, " 投稿API错误")                        // 投稿API错误
	CreativeSearchErr                        = New(20002, " 搜索API错误")                        // 搜索API错误
	CreativeDataErr                          = New(20003, " 大数据API错误")                       // 大数据API错误
	CreativeElecErr                          = New(20004, " 充电API错误")                        // 充电API错误
	CreativeActivityErr                      = New(20005, " 活动API错误")                        // 活动API错误
	CreativeArcServiceErr                    = New(20006, " 稿件服务错误")                         // 稿件服务错误
	CreativeAccServiceErr                    = New(20007, " 账号服务错误")                         // 账号服务错误
	CreativeOrderAPIErr                      = New(20008, " 商单API错误")                        // 商单API错误
	CreativeGeetestErr                       = New(20009, " 验证码错误")                          // 验证码错误
	CreativeFeedbackErr                      = New(20010, " 用户反馈服务错误")                       // 用户反馈服务错误
	CreativeRecommendOverMax                 = New(20011, " 推荐关联的稿件已超过上限(最多8个)")             // 推荐关联的稿件已超过上限(最多8个)
	CreativeRecommendForbid                  = New(20012, " 当前稿件尚未开放浏览，不允许被关联推荐")            // 当前稿件尚未开放浏览，不允许被关联推荐
	CreativeTemplateOverMax                  = New(20013, " 添加的稿件模板数量已超过上限(最多5个)")           // 添加的稿件模板数量已超过上限(最多5个)
	CreativeAccBanned                        = New(20014, " 用户被禁言")                          // 用户被禁言
	CreativeAccNotActivated                  = New(20015, " 用户未激活")                          // 用户未激活
	CreativeTemplateDeleted                  = New(20016, " 稿件模板已删除")                        // 稿件模板已删除
	CreativeArticleRPCErr                    = New(20017, " 文章RPC接口错误")                      // 文章RPC接口错误
	CreativeArticleParamErr                  = New(20018, " 文章投稿参数错误")                       // 文章投稿参数错误
	CreativeArticleCanNotRepeat              = New(20019, " 新增文章同一个标题短时间内不能重复提交")            // 新增文章同一个标题短时间内不能重复提交
	CreativeArticleNotExist                  = New(20020, " 修改的文章不存在")                       // 修改的文章不存在
	CreativeArticleOwnerErr                  = New(20021, " 不是你的文章")                         // 不是你的文章
	CreativeArticleCategoryErr               = New(20022, " 文章分类不合法")                        // 文章分类不合法
	CreativeArticleReprintErr                = New(20023, " 文章转载类型不合法")                      // 文章转载类型不合法
	CreativeArticleTagErr                    = New(20024, " 文章tag不合法")                       // 文章tag不合法
	CreativeArticleCoverErr                  = New(20025, " 文章封面不合法")                        // 文章封面不合法
	CreativeArticleImageURLsErr              = New(20026, " 文章小图片不合法")                       // 文章小图片不合法
	CreativeArticleTitleErr                  = New(20027, " 文章标题不合法")                        // 文章标题不合法
	CreativeArticleContentErr                = New(20028, " 文章内容不合法")                        // 文章内容不合法
	CreativeArticleUploadErr                 = New(20029, " 文件上传失败")                         // 文件上传失败
	CreativeArticleTIDErr                    = New(20030, " 文章模板类型不合法")                      // 文章模板类型不合法
	CreativeArticleImageTypeErr              = New(20031, " 图片类型不合法")                        // 图片类型不合法
	CreativeFansMedalErr                     = New(20032, " 粉丝勋章服务错误")                       // 粉丝勋章服务错误
	CreativeWaterMarkWrongState              = New(20033, " 水印状态参数错误")                       // 水印状态参数错误
	CreativeWaterMarkWrongType               = New(20034, " 水印类型参数错误")                       // 水印类型参数错误
	CreativeWaterMarkWrongPosition           = New(20035, " 水印位置参数错误")                       // 水印位置参数错误
	CreativeWaterMarkCreateFailed            = New(20036, " 水印生成错误")                         // 水印生成错误
	CreativeLiveErr                          = New(20037, " 直播间服务错误")                        // 直播间服务错误
	CreativeLiveNotOpenErr                   = New(20038, " 没有开通直播间")                        // 没有开通直播间
	CreativePhoneNotBind                     = New(20039, " 请先绑定手机")                         // 请先绑定手机
	CreativeAssistErr                        = New(20040, " 协管服务错误")                         // 协管服务错误
	CreativeAssistLogAlreadyRevoc            = New(20041, " 协管日志已被撤销，不允许重复撤销")               // 协管日志已被撤销，不允许重复撤销
	CreativeTagErr                           = New(20042, " 查询tag出错")                        // 查询tag出错
	CreativeDanmuErr                         = New(20043, " 弹幕服务错误")                         // 弹幕服务错误
	CreativeReplyAPIErr                      = New(20044, " 评论API错误")                        // 评论API错误
	CreativeCoinAPIErr                       = New(20045, " 硬币API错误")                        // 硬币API错误
	CreativeUpNoMedalRight                   = New(20046, " Up主没有粉丝勋章资格")                    // Up主没有粉丝勋章资格
	CreativeMissionDeltimeInvalid            = New(20047, " 禁止删除，当前稿件已过所参加活动的可删除截止时间")       // 禁止删除，当前稿件已过所参加活动的可删除截止时间
	CreativeDanmuFilterParamError            = New(20048, " 弹幕过滤参数错误")                       // 弹幕过滤参数错误
	CreativeNotLogin                         = New(20049, " 未登录")                            // 未登录
	CreativeDynamicIntroErr                  = New(20050, " 专栏动态推荐语不合法")                     // 专栏动态推荐语不合法
	CreativeGameOpenAPIErr                   = New(20051, " 游戏开放平台API错误")                    // 游戏开放平台API错误
	CreativePorderForbidShowFront            = New(20052, " 当前稿件的私单已被禁止在前端展现")               // 当前稿件的私单已被禁止在前端展现
	CreativeMusicErr                         = New(20053, " 音乐API错误")                        // 音乐API错误
	CreativeGeetestAPIErr                    = New(20054, " 投稿极验服务不可用")                      // 投稿极验服务不可用
	CreativeAcademyOIDExistErr               = New(20055, " 已存在该稿件，不能重复添加")                  // 已存在该稿件，不能重复添加
	CreativeMessageAPIErr                    = New(20056, " 消息服务API错误")                      // 消息服务API错误
	CreativeLotteryAPIErr                    = New(20057, " 抽奖服务API错误")                      // 抽奖服务API错误
	CreativeSubtitleAPIErr                   = New(20058, " 字幕服务API错误")                      // 字幕服务API错误
	CreativeAcademyH5RecommendErr            = New(20059, " 创作学院h5推荐服务错误")                   // 创作学院h5推荐服务错误
	CreativePayForbidDeleteAfterOpen         = New(20060, " 第一次开放之后禁止修改稿件信息，包括付费和分P，请再等60天") // 第一次开放之后禁止修改稿件信息，包括付费和分P，请再等60天
	CreativePayAPIErr                        = New(20061, " UGC付费查询API错误")                   // UGC付费查询API错误
	CreativeNewcomerMallAPIErr               = New(20062, " 会员购服务错误")                        // 会员购服务错误
	CreativeNewcomerBCoinAPIErr              = New(20063, " B币券服务错误")                        // B币券服务错误
	CreativeNewcomerPendantAPIErr            = New(20064, " 挂件服务错误")                         // 挂件服务错误
	CreativeNewcomerRepeatRewardErr          = New(20065, " 重复领取奖励")                         // 重复领取奖励
	CreativeNewcomerNotCompleteErr           = New(20066, " 任务未完成不可领取奖励")                    // 任务未完成不可领取奖励
	CreativeNewcomerReBindTaskErr            = New(20067, " 重复绑定任务")                         // 重复绑定任务
	CreativeNotUper                          = New(20068, " 用户非up主")                         // 用户非up主
	CreativeNewcomerBigMemberErr             = New(20069, " 大会员服务错误")                        // 大会员服务错误
	CreativeNewcomerReceiveExpireErr         = New(20070, " 奖励已过期")                          // 奖励已过期
	CreativeNewcomerNoTask                   = New(20071, " 任务体系正在配置任务，请稍候")                 // 任务体系正在配置任务，请稍候
	CreativeNewcomerGroupIDErr               = New(20072, " 任务组ID不存在")                       // 任务组ID不存在
	CreativeNewcomerDuplicateGiftRewardIDErr = New(20073, " 任务礼包包含重复任务分类和奖励ID")              // 任务礼包包含重复任务分类和奖励ID
	CreativeAcademyDuplicateSoftIDErr        = New(20075, " 稿件已绑定该软件")                       // 稿件已绑定该软件

	// videoup
	VideoupParamErr                 = New(21001, " 参数错误")                                        // 参数错误
	VideoupTypeidErr                = New(21002, " 该分区不存在")                                      // 该分区不存在
	VideoupCopyrightErr             = New(21003, " 该稿件类型不合法")                                    // 该稿件类型不合法
	VideoupMissionErr               = New(21004, " 该活动不存在")                                      // 该活动不存在
	VideoupTagErr                   = New(21005, " Tag参数不合法")                                    // Tag参数不合法
	VideoupDelayTimeErr             = New(21006, " 定时发布设置错误")                                    // 定时发布设置错误
	VideoupSubmitErr                = New(21007, " 当前时间限制提交")                                    // 当前时间限制提交
	VideoupCoverErr                 = New(21008, " 该封面地址不合法")                                    // 该封面地址不合法
	VideoupTitleErr                 = New(21009, " 标题不合法")                                       // 标题不合法
	VideoupDescErr                  = New(21010, " 描述信息不合法")                                     // 描述信息不合法
	VideoupZeroVideos               = New(21011, " 新增稿件分P不能为空")                                  // 新增稿件分P不能为空
	VideoupCanotRepeat              = New(21012, " 新增稿件同一个标题短时间内不能重复提交")                         // 新增稿件同一个标题短时间内不能重复提交
	VideoupVideoTitleErr            = New(21013, " 分P标题有违禁符号")                                   // 分P标题有违禁符号
	VideoupVideoDescErr             = New(21014, " 分P描述有违禁符号")                                   // 分P描述有违禁符号
	VideoupVideoFilenameErr         = New(21015, " 分P的filename不能为空")                             // 分P的filename不能为空
	VideoupBodyFormatErr            = New(21016, " 代码模式格式错误")                                    // 代码模式格式错误
	VideoupBodyCidErr               = New(21017, " 代码模式cid不为数字")                                 // 代码模式cid不为数字
	VideoupBodyTitleEmpty           = New(21018, " 代码模式title不能为空（如果多于1P")                        // 代码模式title不能为空（如果多于1P
	VideoupBodyForbidCid            = New(21019, " 代码模式存在不可用cid")                                // 代码模式存在不可用cid
	VideoupBodyForbidAllCid         = New(21020, " 代码模式的cid全部不可用")                               // 代码模式的cid全部不可用
	VideoupSourceErr                = New(21021, " 稿件转载来源不能为空")                                  // 稿件转载来源不能为空
	VideoupOrderAPIErr              = New(21022, " 商单API错误")                                     // 商单API错误
	VideoupUperIDNotAllow           = New(21023, " 当前用户尚未加入商单计划")                                // 当前用户尚未加入商单计划
	VideoupOrderIDNotAllow          = New(21024, " 当前商单ID不合法")                                   // 当前商单ID不合法
	VideoupLaunchTimeIllegal        = New(21025, " 当前商单的发布时间不合法")                                // 当前商单的发布时间不合法
	VideoupForbidNoreprint          = New(21026, " 已允许转载的稿件不允许改为禁止转载")                           // 已允许转载的稿件不允许改为禁止转载
	VideoupFilenameCanotRepeat      = New(21027, " 同个稿件内不能重复提交相同的视频")                            // 同个稿件内不能重复提交相同的视频
	VideoupMissionEtimeInvalid      = New(21028, " 该非法活动禁止参加,原因(结束时间不合法)")                       // 该非法活动禁止参加,原因(结束时间不合法)
	VideoupPvodAderTooMany          = New(21029, " 稿件私单广告主过多")                                   // 稿件私单广告主过多
	VideoupPvodAderTooLong          = New(21030, " 稿件私单单个广告主长度已超出限制")                            // 稿件私单单个广告主长度已超出限制
	VideoupPvodWithoutCommercial    = New(21031, " 稿件私单参数错误，商业广告标记为否")                           // 稿件私单参数错误，商业广告标记为否
	VideoupPvodEpyFlowRemark        = New(21032, " 稿件私单参数错误，商业广告类型为空")                           // 稿件私单参数错误，商业广告类型为空
	VideoupPvodForbidFlowID         = New(21033, " 稿件私单参数错误，商业广告类型ID错误")                         // 稿件私单参数错误，商业广告类型ID错误
	VideoupPvodForbidFlowRemark     = New(21034, " 稿件私单参数错误，商业广告类型名称错误")                         // 稿件私单参数错误，商业广告类型名称错误
	VideoupPvodForbidOrderAlready   = New(21035, " 稿件私单参数错误，当前用户已选择有效的商单ID，与私单互斥")               // 稿件私单参数错误，当前用户已选择有效的商单ID，与私单互斥
	VideoupFilenameExpired          = New(21036, " 第(%d)P的文件已超过了限定的时间范围(48小时)，请在删除此分P后再提交")      // 第(%d)P的文件已超过了限定的时间范围(48小时)，请在删除此分P后再提交
	VideoupForbidMultiVideoForTypes = New(21037, " [电影]分区的特殊二级分区[欧美电影,日本电影,国产电影,其他国家]不允许提交多P稿件") // [电影]分区的特殊二级分区[欧美电影,日本电影,国产电影,其他国家]不允许提交多P稿件
	VideoupDynamicErr               = New(21038, " 稿件动态文案不合法")                                   // 稿件动态文案不合法
	VideoupFmDesZero                = New(21051, " 稿件描述长度为零")                                    // 稿件描述长度为零
	VideoupFmDesLenOverLimit        = New(21052, " 稿件描述长度太长，已超过限制")                              // 稿件描述长度太长，已超过限制
	VideoupFmDesIDForbid            = New(21053, " 稿件描述类型不存在或者不匹配")                              // 稿件描述类型不存在或者不匹配
	VideoupFmDesIDNoMatchTypeID     = New(21054, " 稿件描述类型和对应的分区类型不匹配")                           // 稿件描述类型和对应的分区类型不匹配
	VideoupFmDesIDNoMatchCopyright  = New(21055, " 稿件描述类型和对应的创作类型不匹配")                           // 稿件描述类型和对应的创作类型不匹配
	VideoupTagForbid                = New(21057, " 第(%d)个Tag已被封印")                               // 第(%d)个Tag已被封印
	VideoupVideosMaxLimit           = New(21058, " 当前单次操作添加的视频数已经超过限制")                          // 当前单次操作添加的视频数已经超过限制
	VideoupAdBrandIDErr             = New(21059, " 稿件商业推广信息(官方品牌ID)有误，请修正")                      // 稿件商业推广信息(官方品牌ID)有误，请修正
	VideoupAdIndustryIDErr          = New(21060, " 稿件商业推广信息(推广行业)有误，请修正")                        // 稿件商业推广信息(推广行业)有误，请修正
	VideoupAdShowTypeErr            = New(21061, " 稿件商业推广信息(推广形式)有误，请修正")                        // 稿件商业推广信息(推广形式)有误，请修正
	VideoupAdOfficialIndustryIDErr  = New(21062, " 稿件商业推广信息(官方行业)有误，请修正")                        // 稿件商业推广信息(官方行业)有误，请修正
	VideoupAppMustUploadCover       = New(21063, " APP投稿封面不能为空，请添加稿件封面后再继续")                     // APP投稿封面不能为空，请添加稿件封面后再继续
	VideoupFilterServiceErr         = New(21064, " 投稿过滤服务暂不可用")                                  // 投稿过滤服务暂不可用
	VideoupSingleFilterForbid       = New(21065, " 当前输入有敏感信息,请修正")                               // 当前输入有敏感信息,请修正
	VideoupFieldFilterForbid        = New(21066, " %s中包含敏感词")                                    // %s中包含敏感词
	VideoupTagForbidNotJoinMission  = New(21067, " 自定义标签包含不可选的活动tag，请修改后重新提交")                   // 自定义标签包含不可选的活动tag，请修改后重新提交
	VideoupMissionNoMatch           = New(21068, " 当前活动和分区不匹配，请重新选择活动或者分区")                      // 当前活动和分区不匹配，请重新选择活动或者分区
	VideoupChannelReviewNotTrigger  = New(21069, " 当前稿件所属频道不需要回查")                               // 当前稿件所属频道不需要回查
	VideoupAddLimitHalfMin          = New(21070, " 您投稿的频率过快，请稍等30秒")                             // 您投稿的频率过快，请稍等30秒
	VideoupCopyForbidJoinMission    = New(21071, " 转载类型稿件不支持活动参加哦~")                             // 转载类型稿件不支持活动参加哦~
	VideoupEditLocked               = New(21072, " 稿件后台处理中,请10秒后再尝试")                            // 稿件后台处理中,请10秒后再尝试
	VideoupMaxAllVsCntLimit         = New(21073, " 当前总提交视频个数已经超过上限:(%d)")                        // 当前总提交视频个数已经超过上限:(%d)
	VideoupPayCopyrightErr          = New(21074, " 非自制稿件禁止参与UGC内容付费")                            // 非自制稿件禁止参与UGC内容付费
	VideoupPayProtocolLimit         = New(21075, " UGC付费修改之前必须接受当前最新的投稿协议")                      // UGC付费修改之前必须接受当前最新的投稿协议
	VideoupPayAPIErr                = New(21076, " UGC付费服务API错误")                                // UGC付费服务API错误
	VideoupPayPriceErr              = New(21077, " 当前稿件定价错误")                                    // 当前稿件定价错误
	VideoupPayUserNotAllow          = New(21078, " 当前用户暂无参与UGC内容付费的资格")                          // 当前用户暂无参与UGC内容付费的资格
	VideoupPayCommericalLimit       = New(21079, " 禁止付费投稿时同时参与其他商业性投稿活动")                        // 禁止付费投稿时同时参与其他商业性投稿活动
	VideoupPayForbidEditAfterOpen   = New(21080, " 第一次开放之后禁止修改付费信息")                             // 第一次开放之后禁止修改付费信息
	VideoupStaffBlocked             = New(21085, " 因参与者%s隐私设置，暂无法提交")                            // 因参与者%s隐私设置，暂无法提交
	VideoupStaffCopyright           = New(21086, "非自制稿件不可提交多人合作稿件")                              //非自制稿件不可提交多人合作稿件
	VideoupStaffChangeCopyright     = New(21087, "多人合作稿件不可更改稿件类型")                               //多人合作稿件不可更改稿件类型
	VideoupStaffTypeNotExists       = New(21088, "该分区暂未开放多人合作稿件")                                //该分区暂未开放多人合作稿件
	VideoupStaffAuth                = New(21089, "该up主暂未在多人合作稿件公测范围内")                           //该up主暂未在多人合作稿件公测范围内
	VideoupStaffCountLimit          = New(21090, "超出团队成员上限，暂不可提交")                               //超出团队成员上限，暂不可提交
	VideoupStaffMidInvalid          = New(21091, "无效参与者")                                        //无效参与者
	VideoupStaffTitleLength         = New(21092, "参与者类型长度超出上限，暂不可提交")                            //参与者类型长度超出上限，暂不可提交
	VideoupStaffTitleFilter         = New(21093, "职能已经被封印了")                                     //职能已经被封印了
	VideoupStaffMidRepeat           = New(21094, "已经添加过啦，暂不可重复添加")                               //已经添加过啦，暂不可重复添加
	VideoupStaffTitleChar           = New(21096, "参与者类型暂时仅支持中文、英文、数字哦")                          //参与者类型暂时仅支持中文、英文、数字哦
	VideoupStaffTitleShort          = New(21097, "参与者类型长度不足，暂不可提交")                              //参与者类型长度不足，暂不可提交
	VideoupStaffUpSilence           = New(21098, "很抱歉参与者%s的账号目前为封禁状态")                           //很抱歉参与者%s的账号目前为封禁状态
	VideoupStaffChangeTypeCopyright = New(21099, "暂不支持修改分区和转载类型")                                //暂不支持修改分区和转载类型
	VideoupStaffApply404            = New(21081, " 申请单不存在")                                      // 申请单不存在
	VideoupStaffApplyTypeChange     = New(21082, " 申请单性质改变")                                     // 申请单性质改变
	VideoupStaffApplyStateNotMatch  = New(21083, " 申请单操作不允许")                                    // 申请单操作不允许
	VideoupStaffApplyMidNotIn       = New(21084, " 申请单staff不存在")                                 // 申请单staff不存在
	VideoupStaffApplyUpMidBlack     = New(21101, "staff被UP拉黑")                                   //staff被UP拉黑
	VideoupStaffMidSilence          = New(21102, "Staff 被封禁")                                    //Staff 被封禁

	// relation
	RelFollowSelfBanned          = New(22001, " 不能关注自己")                       // 不能关注自己
	RelFollowBlacked             = New(22002, " 被用户拉黑，无法关注")                   // 被用户拉黑，无法关注
	RelFollowAlreadyBlack        = New(22003, " 已经拉黑用户，无法关注")                  // 已经拉黑用户，无法关注
	RelFollowAttrAlreadySet      = New(22004, " 已经设置该属性了")                     // 已经设置该属性了
	RelFollowAttrNotSet          = New(22005, " 未设置该属性，不能取消")                  // 未设置该属性，不能取消
	RelFollowReachTelLimit       = New(22006, " 关注已达上限，答题成为正式会员或者绑定手机号才能继续关注") // 关注已达上限，答题成为正式会员或者绑定手机号才能继续关注
	RelFollowingGuestLimit       = New(22007, " 访客只限制访问前五页")                   // 访客只限制访问前五页
	RelBlackReachMaxLimit        = New(22008, " 黑名单达到上限")                      // 黑名单达到上限
	RelFollowReachMaxLimit       = New(22009, " 关注失败，已达关注上限")                  // 关注失败，已达关注上限
	RelBatchFollowAlreadyBlack   = New(22010, " 部分拉黑用户未成功关注")                  // 部分拉黑用户未成功关注
	RelTagExistNotAllowedWords   = New(22101, " 分组名称存在不允许的字符")                 // 分组名称存在不允许的字符
	RelTagNumLimit               = New(22102, " 分组数量超过限制")                     // 分组数量超过限制
	RelTagLenLimit               = New(22103, " 分组名称长度超过限制")                   // 分组名称长度超过限制
	RelTagNotExist               = New(22104, " 分组不存在")                        // 分组不存在
	RelTagAddFollowingFirst      = New(22105, " 请先公开关注后再添加分组")                 // 请先公开关注后再添加分组
	RelTagExisted                = New(22106, " 分组已存在")                        // 分组已存在
	RelAwardPhoneRequired        = New(22107, " 该账号未通过手机认证")                   // 该账号未通过手机认证
	RelAwardIsBlocked            = New(22108, " 该账号处于封禁状态")                    // 该账号处于封禁状态
	RelAwardInsufficientFollower = New(22109, " 该账号粉丝数不符合满足10000的标准")          // 该账号粉丝数不符合满足10000的标准
	RelAwardGetFailed            = New(22110, " 获取粉丝成就奖励失败")                   // 获取粉丝成就奖励失败
	RelAwardInfoFailed           = New(22111, " 获取成就信息失败")                     // 获取成就信息失败
	RelAwardInsufficientRank     = New(22112, " 很抱歉您的账号为非转正会员")                // 很抱歉您的账号为非转正会员

	// kvo
	KvoTimestampErr  = New(23001, " 时间戳不合法")      // 时间戳不合法
	KvoCheckSumErr   = New(23002, " checksum不合法") // checksum不合法
	KvoDataOverLimit = New(23003, " 数据超过限制")      // 数据超过限制
	KvoNotModified   = New(23004, " 数据没有修改")      // 数据没有修改

	// credit
	CreditLevelLimit    = New(25001, " 风纪委申请等级限制(会员等级≥3)")                      // 风纪委申请等级限制(会员等级≥3)
	CreditIsVerify      = New(25002, " 风纪委申请没有实名认证")                            // 风纪委申请没有实名认证
	CreditIsBlock       = New(25003, " 风纪委申请90天内有封禁记录")                         // 风纪委申请90天内有封禁记录
	CreditNoJoinCase    = New(25004, " 风纪委没有参与案件")                              // 风纪委没有参与案件
	CreditNotJury       = New(25005, " 风纪委不是风纪委成员")                             // 风纪委不是风纪委成员
	CreditJuryExpired   = New(25006, " 风纪委资格过期")                                // 风纪委资格过期
	CreditUnderVoteRate = New(25007, " 风纪委低于投准率")                               // 风纪委低于投准率
	CreditNoCase        = New(25008, " 风纪委没有案件")                                // 风纪委没有案件
	CreditCaseNotExist  = New(25009, " 风纪委案件不存在")                               // 风纪委案件不存在
	CreditCaseLimit     = New(25010, " 风纪委没有权限查看案件")                            // 风纪委没有权限查看案件
	CreditVoteNotExist  = New(25011, " 风纪委投票类型错误")                              // 风纪委投票类型错误
	CreditNovote        = New(25012, " 风纪委不能投票(可能重复投票或case不存在或投票过期vote被job设置为") // 风纪委不能投票(可能重复投票或case不存在或投票过期vote被job设置为
	CreditNoApply       = New(25013, " 风纪委不能重复申请风纪委资格")                         // 风纪委不能重复申请风纪委资格
	CreditCaseMax       = New(25014, " 风纪委每日获取案件到达上限")                          // 风纪委每日获取案件到达上限
	CreditBlack         = New(25015, " 风纪委在黑名单内用户申请风纪委资格")                      // 风纪委在黑名单内用户申请风纪委资格
	CreditNoJuryNum     = New(25016, " 风纪委当日风纪委员名额已发完")                         // 风纪委当日风纪委员名额已发完
	CreditNoblock       = New(25101, " 劳改未被封禁 没有答题资格")                          // 劳改未被封禁 没有答题资格
	CreditQsNumError    = New(25102, " 劳改问题总数不正确")                              // 劳改问题总数不正确
	CreditAnsNumError   = New(25103, " 劳改答案总数不正确")                              // 劳改答案总数不正确
	CreditForeverBlock  = New(25104, " 劳改 永久封禁 不能答题")                           // 劳改 永久封禁 不能答题
	CreditBlockNotExist = New(25105, " 封禁信息不存在")                                // 封禁信息不存在
	CreditBlockExpired  = New(25106, " 封禁信息已过7天有效期")                            // 封禁信息已过7天有效期
	CreditAppealExisted = New(25107, " 申诉已存在")                                  // 申诉已存在

	// dm
	DMFilterIllegalType   = New(36001, " 不支持该屏蔽类型")                // 不支持该屏蔽类型
	DMFilterTooLong       = New(36002, " 屏蔽词超过上限啦（关键词50字，正则200字）") // 屏蔽词超过上限啦（关键词50字，正则200字）
	DMFilterOverMax       = New(36003, " 屏蔽规则超过条数限制")              // 屏蔽规则超过条数限制
	DMFitlerIllegalRegex  = New(36004, " 屏蔽规则正则格式不对")              // 屏蔽规则正则格式不对
	DMFilterExist         = New(36005, " 屏蔽规则已经存在")                // 屏蔽规则已经存在
	DMFilterIsEmpty       = New(36006, " 屏蔽规则不允许为空")               // 屏蔽规则不允许为空
	DMAdvNotAllow         = New(36007, " 不允许购买")                   // 不允许购买
	DMAdvConfirm          = New(36009, " 正在确认中")                   // 正在确认中
	DMAdvBought           = New(36010, " 已购买")                     // 已购买
	DMAdvNoFound          = New(36011, " 高级弹幕购买记录不存在")             // 高级弹幕购买记录不存在
	DMPADMNotOwner        = New(36101, " 别人的弹幕不可以申请弹幕保护~")         // 别人的弹幕不可以申请弹幕保护~
	DMPAUserLevel         = New(36102, " 目前仅限 lv4及以上的用户可以直接申请哦~")  // 目前仅限 lv4及以上的用户可以直接申请哦~
	DMPAUserLimit         = New(36103, " 一个人一天最多只能申请保护100条哦")      // 一个人一天最多只能申请保护100条哦
	DMPADMLimit           = New(36104, " 该弹幕已经申请保护~")              // 该弹幕已经申请保护~
	DMPADMProtected       = New(36105, " 本弹幕已经被保护了~")              // 本弹幕已经被保护了~
	DMNotFound            = New(36106, " 该弹幕已被删除")                 // 该弹幕已被删除
	DMPAFailed            = New(36107, " 申请失败")                    // 申请失败
	DMPoolLimit           = New(36108, " 弹幕池超过大小")                 // 弹幕池超过大小
	DMReportNotExist      = New(36201, " 举报弹幕不存在")                 // 举报弹幕不存在
	DMReportReasonTooLong = New(36202, " 举报原因过长")                  // 举报原因过长
	DMReportReasonError   = New(36203, " 举报原因类型错误")                // 举报原因类型错误
	DMReportExist         = New(36204, " 已举报")                     // 已举报
	DMReportLimit         = New(36205, " 操作过于频繁，请稍后再试")            // 操作过于频繁，请稍后再试
	DMRecallTimeout       = New(36301, " 撤回失败，弹幕发送已过2分钟")          // 撤回失败，弹幕发送已过2分钟
	DMRecallDeleted       = New(36302, " 撤回失败，弹幕已经被删除或撤回")         // 撤回失败，弹幕已经被删除或撤回
	DMRecallLimit         = New(36303, " 撤回失败，今天撤回的机会已经用完")        // 撤回失败，今天撤回的机会已经用完
	DMRecallError         = New(36304, " 撤回失败，服务器出错")              // 撤回失败，服务器出错
	DMAssistNo            = New(36401, " 不是协管")                    // 不是协管
	DMAssistLimit         = New(36402, " 操作次数不足")                  // 操作次数不足
	DMTransferSame        = New(36501, " 弹幕转移源cid和目标cid相等")        // 弹幕转移源cid和目标cid相等
	DMTransferNotFound    = New(36502, " 弹幕转移cid不存在")              // 弹幕转移cid不存在
	DMTransferNotBelong   = New(36503, " 弹幕转移cid不属于该用户")           // 弹幕转移cid不属于该用户
	DMTransferRepet       = New(36504, " 弹幕转移任务重复")                // 弹幕转移任务重复
	DMActSilence          = New(36601, " 弹幕点赞被禁言")                 // 弹幕点赞被禁言
	DMUpgrading           = New(36700, " 系统升级中")                   // 系统升级中
	DMMsgIlleagel         = New(36701, " 弹幕包含被禁止的内容")              // 弹幕包含被禁止的内容
	DMMsgTooLong          = New(36702, " 您的弹幕长度大于100")             // 您的弹幕长度大于100
	DMMsgPubTooFast       = New(36703, " 您发送弹幕的频率过快")              // 您发送弹幕的频率过快
	DMArchiveIlleagel     = New(36704, " 禁止向未审核的视频发送弹幕")           // 禁止向未审核的视频发送弹幕
	DMMsgNoPubPerm        = New(36705, " 您的等级不足，不能发送弹幕")           // 您的等级不足，不能发送弹幕
	DMMsgNoPubTopPerm     = New(36706, " 您的等级不足，不能发送顶端弹幕")         // 您的等级不足，不能发送顶端弹幕
	DMMsgNoPubBottomPerm  = New(36707, " 您的等级不足，不能发送底端弹幕")         // 您的等级不足，不能发送底端弹幕
	DMMsgNoColorPerm      = New(36708, " 您的等级不足，不能发送彩色弹幕")         // 您的等级不足，不能发送彩色弹幕
	DMMsgNoPubAdvancePerm = New(36709, " 您的等级不足，不能发送高级弹幕")         // 您的等级不足，不能发送高级弹幕
	DMMsgNoPubStylePerm   = New(36710, " 您的权限不足，不能发送这种样式的弹幕")      // 您的权限不足，不能发送这种样式的弹幕
	DMForbidPost          = New(36711, " 该视频禁止发送弹幕")               // 该视频禁止发送弹幕
	DMMsgTooLongLevel1    = New(36712, " level 1用户发送弹幕的最大长度为20")   // level 1用户发送弹幕的最大长度为20
	DMNotpayForPost       = New(36713, " 稿件未付费，不能发送弹幕")            // 稿件未付费，不能发送弹幕
	DMProgressTooBig      = New(36714, " 弹幕发送时间不合法")               // 弹幕发送时间不合法
	DMAssistOpToMuch      = New(36715, " 当日操作数量超过上限，请明天再试")        // 当日操作数量超过上限，请明天再试
	DMTaskRegexTooLong    = New(36800, " 任务正则过长")                  // 任务正则过长
	DMTaskRegexIllegal    = New(36801, " 任务正则不合法")                 // 任务正则不合法

	// article
	ArtLikeDupErr            = New(37001, " 重复点赞")                // 重复点赞
	ArtCancelLikeErr         = New(37002, " 取消点赞失败 用户未点赞")        // 取消点赞失败 用户未点赞
	ArtDislikeDupErr         = New(37003, " 重复不喜欢")               // 重复不喜欢
	ArtCancelDislikeErr      = New(37004, " 取消不喜欢失败 用户未不喜欢")      // 取消不喜欢失败 用户未不喜欢
	ArtCreationNoPrivilege   = New(37101, " 创作中心:用户没有权限发文章")      // 创作中心:用户没有权限发文章
	ArtCreationStateErr      = New(37102, " 创作中心:文章状态错误")         // 创作中心:文章状态错误
	ArtCreationIDErr         = New(37103, " 创作中心:文章ID错误")         // 创作中心:文章ID错误
	ArtCreationMIDErr        = New(37104, " 创作中心:非文章作者")          // 创作中心:非文章作者
	ArtCreationDelPendingErr = New(37105, " 创作中心:审核中的文章不能删除")     // 创作中心:审核中的文章不能删除
	ArtCreationDraftFull     = New(37106, " 创作中心:草稿数已达最大上限")      // 创作中心:草稿数已达最大上限
	ArtCreationTplErr        = New(37107, " 创作中心:模板和图片数量不匹配")     // 创作中心:模板和图片数量不匹配
	ArtCreationDraftDeleted  = New(37108, " 创作中心:草稿已被删除，不可编辑")    // 创作中心:草稿已被删除，不可编辑
	ArtCreationArticleFull   = New(37109, " 创作中心：当日投稿数量已到达上限")    // 创作中心：当日投稿数量已到达上限
	ArtUserDisabled          = New(37200, " 用户被封禁 无法操作")          // 用户被封禁 无法操作
	ArtNoCategory            = New(37300, " 文章分区错误")              // 文章分区错误
	ArtApplyPass             = New(37400, " 申请已通过")               // 申请已通过
	ArtApplyReject           = New(37401, " 已经申请处于冷冻期")           // 已经申请处于冷冻期
	ArtApplySubmit           = New(37402, " 已经申请待审")              // 已经申请待审
	ArtApplyClose            = New(37403, " 关闭申请")                // 关闭申请
	ArtApplyFull             = New(37404, " 今日申请名额已满")            // 今日申请名额已满
	ArtApplyVerify           = New(37405, " 用户未实名认证")             // 用户未实名认证
	ArtApplyForbid           = New(37406, " 用户已被封禁")              // 用户已被封禁
	ArtApplyPhone            = New(37407, " 用户未绑定手机")             // 用户未绑定手机
	ArtApplyPhoneVirtual     = New(37408, " 绑定的手机号是虚拟号码")         // 绑定的手机号是虚拟号码
	ArtCannotEditErr         = New(37409, " 文章不能被编辑")             // 文章不能被编辑
	ArtAuthorReject          = New(37410, " 申请被拒绝")               // 申请被拒绝
	ArtNoActivity            = New(37411, " 活动未开始")               // 活动未开始
	ArtMaxListErr            = New(37412, " 达到文集上限 无法再增加新文集")     // 达到文集上限 无法再增加新文集
	ArtListNameErr           = New(37413, " 文集标题不合法")             // 文集标题不合法
	ArtArtAddListErr         = New(37414, " 文章已存在于其他文集或者文章不存在")   // 文章已存在于其他文集或者文章不存在
	ArtAddListLimitErr       = New(37415, " 达到文集文章数量上限 无法再增加新文章") // 达到文集文章数量上限 无法再增加新文章
	ArtPermClosedErr         = New(37416, " 操作失败，你的专栏权限已被关闭")     // 操作失败，你的专栏权限已被关闭
	ArtLevelFailedErr        = New(37417, " 等级未达到要求")             // 等级未达到要求
	ArtMediaExistedErr       = New(37418, " 已经存在长评了")             // 已经存在长评了
	ArtUpdateFullErr         = New(37419, " 重复编辑次数已用完")           // 重复编辑次数已用完

	// Assist
	AssistAlreadyExist             = New(24001, " 当前协管关系已存在,不能重复添加一个用户为另一个用户的协管") // 当前协管关系已存在,不能重复添加一个用户为另一个用户的协管
	AssistNotExist                 = New(24002, " 不存在当前的协管关系")                    // 不存在当前的协管关系
	AssistUserNotExist             = New(24003, " 用户不存在")                         // 用户不存在
	AssistUserNotFollowUp          = New(24004, " 没有关注UP")                        // 没有关注UP
	AssistUserNotRealAuth          = New(24005, " 协管尚未实名认证")                      // 协管尚未实名认证
	AssistOverMaxLimit             = New(24006, " 禁止再添加协管,当前用户的协管数已经超出最大值")       // 禁止再添加协管,当前用户的协管数已经超出最大值
	AssistAlreadyCancel            = New(24007, " 已经撤销过协管操作日志")                   // 已经撤销过协管操作日志
	AssistNotFollowUp              = New(24008, " 当前用户尚未是up主的粉丝，所以不能添加为协管")       // 当前用户尚未是up主的粉丝，所以不能添加为协管
	AssistForbidType               = New(24009, " 禁止添加不允许的业务类型")                  // 禁止添加不允许的业务类型
	AssistForbidAction             = New(24010, " 禁止添加不允许的操作类型")                  // 禁止添加不允许的操作类型
	AssistLogNotExist              = New(24011, " 不存在的操作记录")                      // 不存在的操作记录
	AssistLogOverMaxLimit          = New(24012, " 禁止再添加操作日志,当前协管操作数已超过当天阈值")      // 禁止再添加操作日志,当前协管操作数已超过当天阈值
	AssistOverMaxLimitDailyAddAll  = New(24013, " 禁止再添加骑士,当天已经添加所有用户为骑士超过最大次数")   // 禁止再添加骑士,当天已经添加所有用户为骑士超过最大次数
	AssistOverMaxLimitDailyAddSame = New(24014, " 禁止再添加骑士,当天已经添加同一个用户为骑士超过最大次数")  // 禁止再添加骑士,当天已经添加同一个用户为骑士超过最大次数

	// Member
	UpdateBirthdayFormat   = New(40001, " 出生日期格式不正确")                // 出生日期格式不正确
	UpdateUnameSensitive   = New(40002, " 昵称包含敏感词")                  // 昵称包含敏感词
	UpdateSexError         = New(40003, " 请选择正常的性别")                 // 请选择正常的性别
	UpdateUnameFormat      = New(40004, " 昵称不可包含除-和_以外的特殊字符")        // 昵称不可包含除-和_以外的特殊字符
	UpdateUnameTooLong     = New(40005, " 昵称过长，不能修改")                // 昵称过长，不能修改
	UpdateUnameTooShort    = New(40006, " 昵称过短，不能修改")                // 昵称过短，不能修改
	UpdateUnameMoneyIsNot  = New(40007, " 硬币不足,改昵称需要6个硬币")           // 硬币不足,改昵称需要6个硬币
	UpdateUnameHadVerified = New(40008, " 已过实名验证，不能修改")              // 已过实名验证，不能修改
	UpdateUnameHadLocked   = New(40009, " 昵称已锁定不能修改")                // 昵称已锁定不能修改
	UpdateUnameHadOfficial = New(40010, " 认证账号不得随意修改昵称，如有需要请联系客服娘~") // 认证账号不得随意修改昵称，如有需要请联系客服娘~

	UpdateFaceFormat          = New(40012, " 头像格式错误，允许：png/jpg/jpeg/jp2") // 头像格式错误，允许：png/jpg/jpeg/jp2
	UpdateFaceSize            = New(40013, " 头像超过限制的大小，允许2M")             // 头像超过限制的大小，允许2M
	UpdateUnameRepeated       = New(40014, " 昵称已存在")                      // 昵称已存在
	MemberSignSensitive       = New(40015, " 签名包含敏感词")                    // 签名包含敏感词
	MemberPhoneRequired       = New(40016, " 根据国家实名制认证的相关要求，需要绑定手机号")     // 根据国家实名制认证的相关要求，需要绑定手机号
	MemberRealPhoneRequired   = New(40017, " 根据国家实名制认证的相关要求，需要绑定非虚拟手机号")  // 根据国家实名制认证的相关要求，需要绑定非虚拟手机号
	MemberSignHasEmoji        = New(40021, " 签名不能包含表情图片")                 // 签名不能包含表情图片
	MemberSignOverLimit       = New(40022, " 签名最多支持70个字")                 // 签名最多支持70个字
	BirthdayInfoIsNull        = New(40043, " 该用户没有生日信息 // 答题系统使用")        // 该用户没有生日信息 // 答题系统使用
	MemberUpdate              = New(40050, " 系统维护中")                      // 系统维护中
	MemberBlocked             = New(40051, " 用户被封禁")                      // 用户被封禁
	MemberNameFormatErr       = New(40052, " 用户名不合法")                     // 用户名不合法
	MemberNameOverLimit       = New(40053, " 用户名长度超过限制")                  // 用户名长度超过限制
	MemberNameUnmodify        = New(40054, " 用户名未修改")                     // 用户名未修改
	MemberNameHasEmoji        = New(40055, " 用户名包含emoji")                 // 用户名包含emoji
	MemberNameCoinErr         = New(40056, " 扣除硬币失败")                     // 扣除硬币失败
	MemberUnRealName          = New(40058, " 用户名未实名")                     // 用户名未实名
	MemberCerted              = New(40059, " 用户名包含敏感词")                   // 用户名包含敏感词
	MemberOverLimit           = New(40060, " 批量请求超过限制")                   // 批量请求超过限制
	MemberNotExist            = New(40061, " 用户不存在")                      // 用户不存在
	MemberUpdateBirthdayFaild = New(40071, " 修改生日失败")                     // 修改生日失败
	MemberBirthdayNotAllow    = New(40072, " 生日信息不合法")                    // 生日信息不合法
	MemberBirthdayInfoIsNull  = New(40073, " 该用户没有生日信息")                  // 该用户没有生日信息
	MemberTagsOverLen         = New(40080, " 用户 Tag 不合法")                 // 用户 Tag 不合法
	MemberTagsOverCount       = New(40081, " 用户 Tag 不合法")                 // 用户 Tag 不合法
	SubmitOfficialDocFailed   = New(40083, " 提交官方认证请求失败")                 // 提交官方认证请求失败
	NoOfficialDoc             = New(40084, " 未提交过官方认证请求")                 // 未提交过官方认证请求
	SearchMidOverLimit        = New(40085, " Mid查询数量过大")                  // Mid查询数量过大

	// Answer
	QuestionStrNotAllow     = New(41001, " 分院帽题目不合法")                // 分院帽题目不合法
	QuestionAnsNotAllow     = New(41002, " 分院帽题目答案不合法")              // 分院帽题目答案不合法
	QuestionTipsNotAllow    = New(41003, " 分院帽题目提示不合法")              // 分院帽题目提示不合法
	QuestionTypeNotAllow    = New(41004, " 分院帽题目类型不合法")              // 分院帽题目类型不合法
	AnswerDenied            = New(41010, " 用户答题非法访问")                // 用户答题非法访问
	AnswerTimeExpire        = New(41011, " 用户答题时间已超时")               // 用户答题时间已超时
	AnswerIdsErr            = New(41012, " 用户答题提交题目id不合法")           // 用户答题提交题目id不合法
	AnswerQsNumErr          = New(41013, " 用户答题提交题目数量不合法")           // 用户答题提交题目数量不合法
	AnswerBlock             = New(41014, " 用户自选题提交过快（2分钟内）被封禁12小时")  // 用户自选题提交过快（2分钟内）被封禁12小时
	AnswerSorceZero         = New(41016, " 该用户答题分数为0")               // 该用户答题分数为0
	AnswerGeetestErr        = New(41017, " 答题验证码错误")                 // 答题验证码错误
	AnswerFormalFailed      = New(41018, " 答题转正失败")                  // 答题转正失败
	AnswerBasePassed        = New(41020, " 用户基础题已通过")                // 用户基础题已通过
	AnswerBaseNotPassed     = New(41021, " 用户基础题未通过")                // 用户基础题未通过
	AnswerHistoryNotFound   = New(41023, " 用户答题记录不存在")               // 用户答题记录不存在
	AnswerMidCacheQidsErr   = New(41024, " 获取用户题目id缓存异常")            // 获取用户题目id缓存异常
	AnswerQidDiffRequestErr = New(41025, " 用户答题提交题目ID和实际用户的答题id不一致") // 用户答题提交题目ID和实际用户的答题id不一致
	AnswerMidDBQueErr       = New(41026, " 获取用户DB题目信息异常")            // 获取用户DB题目信息异常
	AnswerCheckFaild        = New(41027, " 基础题检查不通过")                // 基础题检查不通过
	AnswerProNoPass         = New(41031, " 自选题未通过")                  // 自选题未通过
	AnswerCaptchaPassed     = New(41050, " 用户答题验证码已通过")              // 用户答题验证码已通过
	AnswerCaptchaNoPassed   = New(41051, " 用户答题验证码未通过")              // 用户答题验证码未通过
	AnswerTypeIDsErr        = New(41052, " 用户题目类型不合法")               // 用户题目类型不合法
	AnswerGeetestVaErr      = New(41053, " 极验验证异常")                  // 极验验证异常
	AnswerExtraHadPass      = New(41054, " 基础附加题已通过")                // 基础附加题已通过
	AnswerExtraNoPass       = New(41055, " 基础附加题未通过")                // 基础附加题未通过
	AnswerAccCallErr        = New(41090, " 调用账号异常")                  // 调用账号异常
	AnswerNeedBindTel       = New(41091, " 答题需要绑定手机")                // 答题需要绑定手机

	// bfs upload
	BfsUploadCodeErr                     = New(42001, " bfs响应code错误")   // bfs响应code错误
	BfsUploadStatusErr                   = New(42002, " 返回状态错误（非常规捕捉）") // 返回状态错误（非常规捕捉）
	BfsRequestErr                        = New(42400, " bfs参数错误")       // bfs参数错误
	BfsUploadAuthErr                     = New(42401, " 上传验证错误")        // 上传验证错误
	BfsUplaodBucketNotExist              = New(42404, " bucket不存在")     // bucket不存在
	BfsUploadServiceUnavailable          = New(42503, " 服务不可用")         // 服务不可用
	BfsUploadFileTooLarge                = New(42601, " 上传的文件太大")       // 上传的文件太大
	BfsUploadFilePixelError              = New(42602, " 不能获取图片的像素信息")   // 不能获取图片的像素信息
	BfsUploadFilePixelWidthIllegal       = New(42603, " 宽像素不合法")        // 宽像素不合法
	BfsUploadFilePixelHeightIllegal      = New(42604, " 高像素不合法")        // 高像素不合法
	BfsUploadFilePixelAspectRatioIllegal = New(42605, " 像素宽高比不合法")      // 像素宽高比不合法
	BfsUploadFileContentTypeIllegal      = New(42606, " 文件类型不合法")       // 文件类型不合法

	// remote login
	RemoteLoginStatusQueryError = New(43001, " 查询失败") // 查询失败
	RemoteLoginFeedBackError    = New(43002, " 反馈失败") // 反馈失败
	RemoteLoginWarnCloseError   = New(43003, " 关闭失败") // 关闭失败

	// Spy
	SpyEventNotExist          = New(50001, " 反作弊事件类型不存在") // 反作弊事件类型不存在
	SpyServiceNotExist        = New(50002, " 反作弊服务不存在")   // 反作弊服务不存在
	SpyFactorNotExist         = New(50003, " 反作弊因子不存在")   // 反作弊因子不存在
	SpySettingUnknown         = New(50004, " 反作弊配置类型不存在") // 反作弊配置类型不存在
	SpySettingValTypeError    = New(50005, " 反作弊配置值类型错误") // 反作弊配置值类型错误
	SpySettingValueOutOfRange = New(50006, " 反作弊配置值超出范围") // 反作弊配置值超出范围
	SpyRuleNotExist           = New(50007, " 反作弊规则不存在")   // 反作弊规则不存在
	SpyRulesNotMatch          = New(50008, " 反作弊规则不匹配")   // 反作弊规则不匹配

	// filter-service and filter-job
	FilterHitLimitBlack        = New(52001, " 命中黑名单")                 // 命中黑名单
	FilterHitRubLimit          = New(52002, " 超过发送次数")                // 超过发送次数
	FilterLimitTypeNotExist    = New(52003, " 限制类型不存")                // 限制类型不存
	FilterLimitContentNotExist = New(52004, " 限制关键词不存在")              // 限制关键词不存在
	FilterHitStrictLimit       = New(52005, " 命中严格限制")                // 命中严格限制
	FilterIllegalRegexp        = New(52006, " 非法正则")                  // 非法正则
	FilterIllegalArea          = New(52007, " 业务不存在")                 // 业务不存在
	FilterWhiteSampleHit       = New(52010, " 敏感词可能误杀较大")             // 敏感词可能误杀较大
	FilterBlackSampleHit       = New(52011, " 敏感词导致高危内容失效")           // 敏感词导致高危内容失效
	FilterDuplicateContent     = New(52012, " 已存在相同内容敏感词/白名单")        // 已存在相同内容敏感词/白名单
	FilterRegexpError1         = New(52013, " 含有.*容易引起误伤，请换用.{0,10}") // 含有.*容易引起误伤，请换用.{0,10}
	FilterRegexpError2         = New(52014, " 含有||容易引起误伤")            // 含有||容易引起误伤
	FilterInvalidAreaGroupName = New(52020, " 不合法的业务组命名")             // 不合法的业务组命名
	FilterDuplicateAreaGroup   = New(52021, " 业务组重复")                 // 业务组重复
	FilterAreaGroupNotFound    = New(52022, " 业务组不存在")                // 业务组不存在
	FilterInvalidAreaShowName  = New(52023, " 不合法的业务模块命名")            // 不合法的业务模块命名
	FilterInvalidAreaName      = New(52024, " 不合法的业务id命名")            // 不合法的业务id命名
	FilterDuplicateArea        = New(52025, " 业务重复")                  // 业务重复
	FilterInvalidArea          = New(52026, " 业务不存在")                 // 业务不存在
	FilterInvalidAIWhiteUID    = New(52027, " AI过滤白名单重复添加")           // AI过滤白名单重复添加

	// space channel
	ChNameToLong    = New(53001, " 频道名字数超过限制啦")          // 频道名字数超过限制啦
	ChIntroToLong   = New(53002, " 频道简介字数超过限制啦")         // 频道简介字数超过限制啦
	ChMaxArcCount   = New(53003, " 本频道里的视频已经满啦")         // 本频道里的视频已经满啦
	ChMaxCount      = New(53004, " 你创建的频道已经满额了哦")        // 你创建的频道已经满额了哦
	ChFakeAid       = New(53005, " 频道内有失效视频了哦")          // 频道内有失效视频了哦
	ChAidsExist     = New(53006, " 你提交的视频已失效或者频道里已经有了哦") // 你提交的视频已失效或者频道里已经有了哦
	ChNameExist     = New(53007, " 频道名称已经存在了哦")          // 频道名称已经存在了哦
	ChNoArcs        = New(53008, " 频道内没有视频")             // 频道内没有视频
	ChNoArc         = New(53009, " 频道内没有该视频")            // 频道内没有该视频
	ChNameBanned    = New(53010, " 频道名称有敏感词，请重新编写")      // 频道名称有敏感词，请重新编写
	ChIntroBanned   = New(53011, " 频道简介有敏感词，请重新编写")      // 频道简介有敏感词，请重新编写
	SpaceNoShop     = New(53012, " 非营业中商户号")             // 非营业中商户号
	SpaceNoPrivacy  = New(53013, " 用户隐私设置未公开")           // 用户隐私设置未公开
	SpaceFakeAid    = New(53014, " 该稿件已失效")              // 该稿件已失效
	TopReasonLong   = New(53015, " 置顶理由字数超过限制啦")         // 置顶理由字数超过限制啦
	SpaceNoTopArc   = New(53016, " 没有置顶视频")              // 没有置顶视频
	SpaceNotAuthor  = New(53017, " 只能操作自己的稿件")           // 只能操作自己的稿件
	SpaceTextBanned = New(53018, " 提交文本有敏感词")            // 提交文本有敏感词
	SpaceMpMaxCount = New(53019, " 代表作已达上限")             // 代表作已达上限
	SpaceMpExist    = New(53020, " 代表作内已有该视频")           // 代表作内已有该视频
	SpaceMpNoArc    = New(53021, " 代表作内没有该视频")           // 代表作内没有该视频

	// search
	SearchArchiveCheckFailed     = New(54001, " 搜索稿件管理失败")          // 搜索稿件管理失败
	SearchArticleDataFailed      = New(54002, " 搜索专栏数据失败")          // 搜索专栏数据失败
	SearchReplyRecordFailed      = New(54003, " 搜索个人中心评论记录数据失败")    // 搜索个人中心评论记录数据失败
	SearchBlockedListFailed      = New(54010, " 搜索风纪委封禁列表失败")       // 搜索风纪委封禁列表失败
	SearchBlockedPublishFailed   = New(54011, " 搜索公告列表失败")          // 搜索公告列表失败
	SearchBlockedCaseFailed      = New(54012, " 搜索案件列表失败")          // 搜索案件列表失败
	SearchBlockedJuryFailed      = New(54013, " 搜索委员列表失败")          // 搜索委员列表失败
	SearchBlockedOpinionFailed   = New(54014, " 搜索观点列表失败")          // 搜索观点列表失败
	SearchWorkflowGroupFailed    = New(54015, " 工作流获取反馈列表失败")       // 工作流获取反馈列表失败
	SearchWorkflowChaFailed      = New(54016, " 工作流获取用户工单失败")       // 工作流获取用户工单失败
	SearchWorkflowTagFailed      = New(54017, " 工作流获取Tag列表失败")      // 工作流获取Tag列表失败
	SearchWorkflowLogFailed      = New(54018, " 工作流获取日志列表失败")       // 工作流获取日志列表失败
	SearchWorkflowCommonFaild    = New(54019, " 工作流举报列表获取失败")       // 工作流举报列表获取失败
	SearchUpdateIndexFailed      = New(54900, " 更新索引失败")            // 更新索引失败
	SearchDmFailed               = New(54020, " 弹幕列表获取失败")          // 弹幕列表获取失败
	SearchVideoFailed            = New(54021, " 弹幕列表获取失败")          // 弹幕列表获取失败
	SearchMusicSongsFailed       = New(54022, " 音乐审核列表获取失败")        // 音乐审核列表获取失败
	SearchKpiPointFailed         = New(54023, " 搜索kpi评分列表失败")       // 搜索kpi评分列表失败
	SearchWorkflowFeedbackFailed = New(54024, " 工作流获取feedback列表失败") // 工作流获取feedback列表失败
	SearchUNameFailed            = New(54025, " 用户昵称查询失败")          // 用户昵称查询失败
	SearchDmmonitorFailed        = New(54026, " 弹幕监控查询失败")          // 弹幕监控查询失败
	SearchPgcMediaFailed         = New(54027, " pgc影视查询失败")         // pgc影视查询失败
	SearchFeedbackFailed         = New(54028, " 用户反馈查询失败")          // 用户反馈查询失败
	SearchFeedbackReplyFailed    = New(54029, " 用户反馈报告查询失败")        // 用户反馈报告查询失败
	SearchLogAuditFailed         = New(54030, " 审核日志查询失败")          // 审核日志查询失败
	SearchLogAuditOidFailed      = New(54031, " 根据oid查询审核日志查询失败")   // 根据oid查询审核日志查询失败
	SearchLogUserActionFailed    = New(54032, " 用户操作日志查询失败")        // 用户操作日志查询失败
	SearchUserApplyReviewsFailed = New(54033, " 用户头像挂件查询失败")        // 用户头像挂件查询失败

	SearchBusinessFailed   = New(54901, " 后台接口Business报错") // 后台接口Business报错
	SearchAppidFailed      = New(54902, " 后台接口Appid报错")    // 后台接口Appid报错
	SearchBusinessExistErr = New(54903, " 该业务已经存在")        // 该业务已经存在
	SearchAssetExistErr    = New(54904, " 该数据源已经存在")       // 该数据源已经存在
	SearchAppExistErr      = New(54905, " 该应用已经存在")        // 该应用已经存在

	// figure 信用分服务
	FigureNotFound = New(55001, " 未找到用户信用分") // 未找到用户信用分

	// workflow 工作流
	WkfGroupNotFound                  = New(56001, " 未找到工单")           // 未找到工单
	WkfChallNotFound                  = New(56002, " 未找到工单详情")         // 未找到工单详情
	WkfAppealNotUserOwned             = New(56201, " 只能查看或操作自己的申诉")    // 只能查看或操作自己的申诉
	WkfAppealNotAllowDegree           = New(56202, " 当前申诉状态不允许提交满意度")  // 当前申诉状态不允许提交满意度
	WkfBusinessNotFound               = New(56401, " 未找到业务id")         // 未找到业务id
	WkfBusinessNotConsistent          = New(56402, " 工单业务id不一致")       // 工单业务id不一致
	WkfTagNotFound                    = New(56403, " 未找到tid数据")        // 未找到tid数据
	WkfBusinessCallbackConfigNotFound = New(56404, " 未找到业务回调配置")       // 未找到业务回调配置
	WkfBanNotSupportBatchOperate      = New(56405, " 不支持批量封禁账号")       // 不支持批量封禁账号
	WkfBidNotSupportPublicReferee     = New(56406, " 业务不支持移交众裁")       // 业务不支持移交众裁
	WkfBidNotSupportQuerySource       = New(56407, " 业务不支持查询来源")       // 业务不支持查询来源
	WkfBidNotSupportQueryContentState = New(56408, " 业务不支持查询内容状态")     // 业务不支持查询内容状态
	WkfSearchGroupFailed              = New(56501, " es search工单失败")   // es search工单失败
	WkfSearchChallFailed              = New(56502, " es search工单详情失败") // es search工单详情失败
	WkfGetBlockInfoFailed             = New(56503, " 获取封禁信息失败")        // 获取封禁信息失败
	WkfSetPublicRefereeFailed         = New(56504, " 提交众裁失败")          // 提交众裁失败

	// account common
	UserLoginInvalid      = New(61000, " 使用登录状态访问了，并且登录状态无效，客服端可以／需要删除登录状态")           // 使用登录状态访问了，并且登录状态无效，客服端可以／需要删除登录状态
	UserCheckNoPhone      = New(61001, " 根据国家实名制认证的相关要求，您需要绑定手机号，才能继续进行操作")            // 根据国家实名制认证的相关要求，您需要绑定手机号，才能继续进行操作
	UserCheckInvalidPhone = New(61002, " 根据国家实名制认证的相关要求，您需要换绑一个非170/171的手机号，才能继续进行操作") // 根据国家实名制认证的相关要求，您需要换绑一个非170/171的手机号，才能继续进行操作

	// web-interface
	ElecDenied      = New(62001, " 不需要展示充电信息")      // 不需要展示充电信息
	ArchiveDenied   = New(62002, " 稿件不可见")          // 稿件不可见
	ArchivePass     = New(62003, " 稿件已审核通过，等待发布中")  // 稿件已审核通过，等待发布中
	ArchiveChecking = New(62004, " 视频正在审核中，请耐心等待～") // 视频正在审核中，请耐心等待～
	ArchiveNotLogin = New(62005, " 视频不见了？你可以试试登录！") // 视频不见了？你可以试试登录！
	HelpListError   = New(62006, " 智齿列表结果错误")       // 智齿列表结果错误
	HelpDetailError = New(62007, " 智齿详情结果错误")       // 智齿详情结果错误
	HelpSearchError = New(62008, " 智齿搜索结果错误")       // 智齿搜索结果错误
	ArcAppealLimit  = New(62009, " 短时间内请勿重复投诉相同稿件") // 短时间内请勿重复投诉相同稿件

	// playlist
	PlNameTooLong      = New(63001, " 播单标题超出最大限制")    // 播单标题超出最大限制
	PlDescTooLong      = New(63002, " 播单简介超出最大限制")    // 播单简介超出最大限制
	PlMaxCount         = New(63003, " 播单个数超出最大限制")    // 播单个数超出最大限制
	PlCanNotDelDefault = New(63004, " 不能删除默认播单")      // 不能删除默认播单
	PlExist            = New(63005, " 已经存在相同标题的播单")   // 已经存在相同标题的播单
	PlNotExist         = New(63006, " 播单无法访问")        // 播单无法访问
	PlAlreadyDel       = New(63007, " 播单已经删除")        // 播单已经删除
	PlDenied           = New(63008, " 播单暂未开放")        // 播单暂未开放
	PlVideoOverflow    = New(63009, " 播单内视频个数超出最大限制") // 播单内视频个数超出最大限制
	PlVideoExist       = New(63010, " 视频已经添加进此播单")    // 视频已经添加进此播单
	PlVideoAlreadyDel  = New(63011, " 视频已经不属于此播单")    // 视频已经不属于此播单
	PlSortOverflow     = New(63012, " 播单内视频排序不生效")    // 播单内视频排序不生效
	PlFavExist         = New(63013, " 播单已经收藏")        // 播单已经收藏
	PlFavAlreadyDel    = New(63014, " 播单未收藏")         // 播单未收藏
	PlNotUser          = New(63015, " 非创建者不能修改此播单")   // 非创建者不能修改此播单

	// usersuit
	UsersuitInviteLevelLow               = New(64001, " 你还不满足购买激活码的条件哦，升级到Lv5再来吧~") // 你还不满足购买激活码的条件哦，升级到Lv5再来吧~
	UsersuitInviteReachCurrentMonthLimit = New(64002, " 当月邀请码申请数达到上限")              // 当月邀请码申请数达到上限
	UsersuitInviteAlreadyFormal          = New(64003, " 已经转正不能使用邀请码")               // 已经转正不能使用邀请码
	UsersuitInviteCodeNotExists          = New(64004, " 邀请码不存在")                    // 邀请码不存在
	UsersuitInviteCodeUsed               = New(64005, " 邀请码已使用")                    // 邀请码已使用
	UsersuitInviteCodeExpired            = New(64006, " 邀请码已过期")                    // 邀请码已过期
	UsersuitInviteReachMaxGeneLimit      = New(64007, " 超过批量生成邀请码上限（最多1000个）")      // 超过批量生成邀请码上限（最多1000个）

	// pendant 挂件相关
	PendantNotFound          = New(64101, " 挂件不存在")                // 挂件不存在
	PendantCanNotBuy         = New(64102, " 大会员挂件不能购买")            // 大会员挂件不能购买
	PendantAlreadyGet        = New(64103, " 大会员挂件已领取过")            // 大会员挂件已领取过
	PendantGetVIPErr         = New(64104, " 获取大会员信息错误")            // 获取大会员信息错误
	PendantPayErr            = New(64105, " 订单接口错误")               // 订单接口错误
	PendantOrderNotFound     = New(64106, " 订单不存在")                // 订单不存在
	PendantPackageNotFound   = New(64107, " 背包里无此挂件")              // 背包里无此挂件
	PendantPayTypeErr        = New(64108, " 该挂件无此种支付方式")           // 该挂件无此种支付方式
	PendantVIPOverdue        = New(64109, " 大会员过期")                // 大会员过期
	PendantAboveSendMaxLimit = New(64110, " 超过批量发放挂件上限 (最多1000个)") // 超过批量发放挂件上限 (最多1000个)

	// usersuit medal 勋章
	MedalNotFound = New(64201, " 勋章不存在")  // 勋章不存在
	MedalNotGet   = New(64202, " 不拥有该勋章") // 不拥有该勋章
	MedalHasGet   = New(64203, " 已拥有该勋章") // 已拥有该勋章

	// thumbup
	ThumbupBusinessBlankErr = New(65001, "业务id不存在")         //业务id不存在
	ThumbupOriginErr        = New(65002, "origin id 错误")    //origin id 错误
	ThumbupBusinessErr      = New(65003, "未开通此业务")          //未开通此业务
	ThumbupCancelLikeErr    = New(65004, " 取消点赞失败 用户未点赞")   // 取消点赞失败 用户未点赞
	ThumbupCancelDislikeErr = New(65005, " 取消不喜欢失败 用户未不喜欢") // 取消不喜欢失败 用户未不喜欢
	ThumbupDupLikeErr       = New(65006, " 重复点赞")           // 重复点赞
	ThumbupDupDislikeErr    = New(65007, " 重复点踩")           // 重复点踩

	// sms
	SmsTemplateNotExist       = New(66001, " 模版不存在")                 // 模版不存在
	SmsTemplateParamNotEnough = New(66002, " 模版参数不足")                // 模版参数不足
	SmsTemplateCodeExist      = New(66003, " 模版code已存在")             // 模版code已存在
	SmsTemplateParamIllegal   = New(66004, " 模版参数值不合法")              // 模版参数值不合法
	SmsTemplateModifyForbind  = New(66005, " 修改已审核的模版必须提供approver3") // 修改已审核的模版必须提供approver3
	SmsTemplateNotAct         = New(66006, " 模版不是营销短信")              // 模版不是营销短信
	SmsSendBatchOverLimit     = New(66023, " 批量发送超出限制")              // 批量发送超出限制
	SmsSendBothMidAndMobile   = New(66024, " mid,mobile只能传一个参数")     // mid,mobile只能传一个参数
	SmsMobilePatternErr       = New(66031, " 手机号码格式不正确")             // 手机号码格式不正确

	// growup admin and interface and job
	GrowupDisabled                = New(68001, " up主在黑名单")               // up主在黑名单
	GrowupTagForbit               = New(68002, " 不允许操作该标签")              // 不允许操作该标签
	GrowupNotExist                = New(68003, " 有不存在的UP主账号")            // 有不存在的UP主账号
	GrowupAuthorityUserNotFound   = New(68004, " 用户在权限系统中不存在")           // 用户在权限系统中不存在
	GrowupTagAddForbit            = New(68005, " 该标签在本业务分区下已存在")         // 该标签在本业务分区下已存在
	GrowupAuthorityExist          = New(68006, " 用户名/任务组名/角色组名/权限点名已存在") // 用户名/任务组名/角色组名/权限点名已存在
	GrowupBodyTooLarge            = New(68007, " 上传的文件太大")               // 上传的文件太大
	GrowupBodyNotExist            = New(68008, " 上传的文件无内容")              // 上传的文件无内容
	GrowupGetTypeError            = New(68009, " 获取视频全部分区失败")            // 获取视频全部分区失败
	GrowupGetActivityError        = New(68010, " 根据活动ID获取稿件失败")          // 根据活动ID获取稿件失败
	GrowupPriceErr                = New(68020, " 购买大会员价格错误")             // 购买大会员价格错误
	GrowupPriceNotEnough          = New(68021, " 购买大会员余额不足")             // 购买大会员余额不足
	GrowupBuyErr                  = New(68022, " 激励兑换购买失败")              // 激励兑换购买失败
	GrowupGoodsNotExist           = New(68023, " 激励兑换商品不存在")             // 激励兑换商品不存在
	GrowupGoodsTimeErr            = New(68024, " 不在兑换时间内")               // 不在兑换时间内
	GrowupArchiveNotYours         = New(68101, " 稿件不属于此人")               // 稿件不属于此人
	GrowupSubTidNotExist          = New(68102, " 此二级分类不存在")              // 此二级分类不存在
	GrowupActivityCountNotEnough  = New(68103, " 活动列表不存在")               // 活动列表不存在
	GrowupRecommendUpNotExist     = New(68104, " 推荐UP主列表不存在")            // 推荐UP主列表不存在
	GrowupUpInfoNotExist          = New(68105, " UP主信息不存在")              // UP主信息不存在
	GrowupRecommendUpInfoNotExist = New(68106, " 推荐UP主信息不存在")            // 推荐UP主信息不存在
	GrowupTidNotExist             = New(68107, " 此一级分类不存在")              // 此一级分类不存在
	GrowupSpecialAwardJoined      = New(68201, " 已参加过专项奖")               // 已参加过专项奖
	GrowupSpecialAwardUnqualified = New(68202, " 没有资格参加专项奖")             // 没有资格参加专项奖

	// Vip
	VipBatchIDErr                   = New(69001, " 大会员资源批次不存在")             // 大会员资源批次不存在
	VipPoolIDErr                    = New(69002, " 大会员资源池不存在")              // 大会员资源池不存在
	VipBusinessErr                  = New(69003, " 大会员资源与业务方不匹配")           // 大会员资源与业务方不匹配
	VipAppkeyExitErr                = New(69004, " appkey已存在")              // appkey已存在
	VipChangeHistoryErr             = New(69005, " 该订单已消费")                 // 该订单已消费
	VipBatchNotEnoughErr            = New(69006, " 资源池数量不足")                // 资源池数量不足
	VipBatchTTLErr                  = New(69007, " 资源池有效时间无效")              // 资源池有效时间无效
	VipMidErr                       = New(69008, " mid无效")                  // mid无效
	VipRemarkErr                    = New(69009, " 请填写备注")                  // 请填写备注
	VipRelationIDErr                = New(69010, " 请填写关联ID")                // 请填写关联ID
	VipDaysErr                      = New(69011, " 大会员天数不能少于一天")            // 大会员天数不能少于一天
	VipBusinessNameExitErr          = New(69012, " 业务方名称已存在")               // 业务方名称已存在
	VipConfigNotExitErr             = New(69013, " 配置不存在")                  // 配置不存在
	VipUpdateErr                    = New(69014, " 大会员更新失败")                // 大会员更新失败
	VipPoolTTLErr                   = New(69015, " 资源池过期")                  // 资源池过期
	VipBusinessStatusErr            = New(69016, " 业务方无效")                  // 业务方无效
	VipOrderNoErr                   = New(69017, " 大会员订单号有误")               // 大会员订单号有误
	VipOpenErr                      = New(69018, " 大会员资源池开通失败")             // 大会员资源池开通失败
	VipPoolNameErr                  = New(69020, " 资源池名称不能为空")              // 资源池名称不能为空
	VipPoolReasonErr                = New(69021, " 资源池使用事由不能为空")            // 资源池使用事由不能为空
	VipPoolStartTimeErr             = New(69022, " 资源池开始时间无效")              // 资源池开始时间无效
	VipPoolEndTimeErr               = New(69023, " 资源池结束时间无效")              // 资源池结束时间无效
	VipPoolNameExitErr              = New(69025, " 资源池名称已存在")               // 资源池名称已存在
	VipBusinessNotExitErr           = New(69026, " 业务方不存在")                 // 业务方不存在
	VipPoolValidityTimeErr          = New(69027, " 资源池有效期数据有误")             // 资源池有效期数据有误
	VipBatchUnitErr                 = New(69028, " 资源批次单位不能小于等于零且单位小于3660") // 资源批次单位不能小于等于零且单位小于3660
	VipBatchCountErr                = New(69029, " 资源批次总量出错")               // 资源批次总量出错
	VipBatchPlusResouceErr          = New(69030, " 资源批次增量不能小于零")            // 资源批次增量不能小于零
	VipPayWayNotSupport             = New(69031, " vip支付方式不支持")             // vip支付方式不支持
	VipUserInfoNotExit              = New(69032, "不存在该会员信息")                //不存在该会员信息
	VipCaptchaVerifyCheckErr        = New(69033, "参数 captcha 验证失败")         //参数 captcha 验证失败
	VipMonthErr                     = New(69050, "月份输入错误")                  //月份输入错误
	VipMonthPriceErr                = New(69060, "月份价格输入错误")                //月份价格输入错误
	VipMonthsNotFoundErr            = New(69061, " 月份找不到")                  // 月份找不到
	VipBatchPriceErr                = New(69062, "资源批次价格异常")                //资源批次价格异常
	VipCodeIDErr                    = New(69063, "激活码不存在")                  //激活码不存在
	VipPointExchangeErr             = New(69064, " 积分兑换失败")                 // 积分兑换失败
	VipBatchCodeCountErr            = New(69065, "资源激活码批次数量最多20W")          //资源激活码批次数量最多20W
	VipBatchCodeNameErr             = New(69066, "资源批次激活码的名称已存在")           //资源批次激活码的名称已存在
	VipOpenCodeCountErr             = New(69067, "开通查询达到上限20次")             //开通查询达到上限20次
	VipCodeNotExitErr               = New(69068, "激活码不存在")                  //激活码不存在
	VipOpenCodeErr                  = New(69069, "开通失败")                    //开通失败
	VipCodeUsedErr                  = New(69070, "激活码已被使用")                 //激活码已被使用
	VipCodeFrozenErr                = New(69071, "激活码被冻结")                  //激活码被冻结
	VipCodeTTLErr                   = New(69072, "激活码已过期")                  //激活码已过期
	VipCodeErr                      = New(69073, "无效激活码")                   //无效激活码
	VipCodeNotStartErr              = New(69074, "激活码未开始")                  //激活码未开始
	VipOrderNotFoundErr             = New(69075, "订单号不存在")                  //订单号不存在
	VipOrderHadExistErr             = New(69076, "订单号已存在")                  //订单号已存在
	VipPanelPlatNotExitErr          = New(69077, " 价格面板的业务不存在")             // 价格面板的业务不存在
	VipPanelValidOPriceErr          = New(69078, " 价格面板的原价不能小于等于0")         // 价格面板的原价不能小于等于0
	VipPanelConfNotUQErr            = New(69079, " 同类型的价格面板已存在")            // 同类型的价格面板已存在
	VipPanelConfNotExistErr         = New(69080, " 基础价格面版不存在")              // 基础价格面版不存在
	VipPanelValidDPriceErr          = New(69081, " 折扣价不能大于原价")              // 折扣价不能大于原价
	VipPanelProductIDNotNilErr      = New(69082, " IOS的价格面板的产品ID不能为空")      // IOS的价格面板的产品ID不能为空
	VipPanelValidTimeErr            = New(69083, " 面板的有效时间跟已有的配置冲突")        // 面板的有效时间跟已有的配置冲突
	VipPanelSTGeqETErr              = New(69084, " 面板的结束时间必须大于开始时间")        // 面板的结束时间必须大于开始时间
	VipPanelSuperscriptTooLongErr   = New(69085, " 价格面板角标内容过长")             // 价格面板角标内容过长
	VipPanelFirstPriceNotSupportErr = New(69086, " 非自动续费不支持录入首月价格")         // 非自动续费不支持录入首月价格

	VipTipTooLoogErr               = New(69100, "小黄条提示至多允许填入28字")    //小黄条提示至多允许填入28字
	VipTipNotFoundErr              = New(69101, "小黄条记录没有找到")         //小黄条记录没有找到
	VipTipStartTimeCatNotModifyErr = New(69102, "小黄条生效时间已不允许更新")     //小黄条生效时间已不允许更新
	VipTipEndTimeCatNotModifyErr   = New(69103, "小黄条失效时间已不允许更新")     //小黄条失效时间已不允许更新
	VipTipCatNotDeleteErr          = New(69104, "小黄条只允许删除未生效记录")     //小黄条只允许删除未生效记录
	VipTipCatNotExpireErr          = New(69105, "小黄条只允许下线生效中的记录")    //小黄条只允许下线生效中的记录
	VipTipTimeErr                  = New(69106, "小黄条生效时间有误")         //小黄条生效时间有误
	VipOpenOnlyNotVipErr           = New(69201, "仅非大会员使用")           //仅非大会员使用
	VipCodeLimitErr                = New(69202, "激活码查询数量到达上限")       //激活码查询数量到达上限
	VipSuitPirceNotFound           = New(69250, "大会员套餐价格信息未找到")      //大会员套餐价格信息未找到
	VipOrderNoEmptyErr             = New(69251, "订单号不能为空")           //订单号不能为空
	VipOrderInfoNotFoundErr        = New(69252, "未找到订单信息")           //未找到订单信息
	VipOrderStatusPayingErr        = New(69253, "订单状态必需是未支付")        //订单状态必需是未支付
	VipOrderCancelFaildErr         = New(69254, "订单取消失败或已支付")        //订单取消失败或已支付
	VipGetPirceErr                 = New(69298, "获取价格信息失败")          //获取价格信息失败
	VipOrderPirceErr               = New(69299, "订单价格有误")            //订单价格有误
	VipPushGroupLenErr             = New(69301, "group长度过长")         //group长度过长
	VipPushTitleLenErr             = New(69302, "title长度过长")         //title长度过长
	VipPushContentLenErr           = New(69303, "内容长度过长")            //内容长度过长
	VipPushLinkTypeErr             = New(69304, "链接类型错误")            //链接类型错误
	VipPushFmtTimeErr              = New(69305, "推送时间格式错误")          //推送时间格式错误
	VipPushPlatformErr             = New(69306, "平台及版本错误")           //平台及版本错误
	VipPushEffectTimeErr           = New(69307, "生效时间错误")            //生效时间错误
	VipPushDataNotExitErr          = New(69308, "推送数据不存在")           //推送数据不存在
	VipPushDataUpdateErr           = New(69309, "不允许修改")             //不允许修改
	VipPushDataDelErr              = New(69310, "不允许删除")             //不允许删除
	VipPushDataDisableErr          = New(69311, "不允许下架")             //不允许下架
	VipOrderAlreadyHandlerErr      = New(69312, "订单已处理")             //订单已处理
	VipRenewTypeErr                = New(69313, "请先签订契约吧！")          //请先签订契约吧！
	VipPayChannelNotExitErr        = New(69314, "不存在该渠道")            //不存在该渠道
	VipRescissionErr               = New(69315, "解约失败")              //解约失败
	VipBisTypeErr                  = New(69316, "业务方类型错误")           //业务方类型错误
	VipWhiteIPListErr              = New(69317, "不存在ip白名单中")         //不存在ip白名单中
	VipOrderTypeErr                = New(69318, "只能是非续费订单")          //只能是非续费订单
	VipOldOrderErr                 = New(69319, "请选择新支付平台订单")        //请选择新支付平台订单
	VipOrderMoneyErr               = New(69320, "退款金额大于购买总金额")       //退款金额大于购买总金额
	VipOrderStatusErr              = New(69321, "订单状态错误")            //订单状态错误
	VipOrderToMidErr               = New(69322, "赠送大会员不允许退款")        //赠送大会员不允许退款
	VipRefundErr                   = New(69323, "退款失败")              //退款失败
	VipBatchMaxCountErr            = New(69324, "已达到最大开通次数")         //已达到最大开通次数
	VipBatchLimitDayErr            = New(69325, "仅限新用户可用")           //仅限新用户可用
	VipDialogTimeErr               = New(69330, "大会员支付弹框时间错误")       //大会员支付弹框时间错误
	VipDialogConflictErr           = New(69331, "大会员支付弹框时间冲突")       //大会员支付弹框时间冲突
	VipPlatformConfDelErr          = New(69332, "平台下有价格列表或弹框不允许被删除") //平台下有价格列表或弹框不允许被删除
	VipEleOrderCanNotReFundErr     = New(69333, "饿了么联合会员不允许退款")      //饿了么联合会员不允许退款

	VipStartTimeErr      = New(69400, "结束时间必须大于开始时间") //结束时间必须大于开始时间
	VipTitleTooLongErr   = New(69401, "标题过长")         //标题过长
	VipContentTooLongErr = New(69402, "内容过长")         //内容过长

	VipFileUploadFaildErr         = New(69350, "权益图片上传失败")  //权益图片上传失败
	VipFileImgEmptyErr            = New(69351, "权益图片不能为空")  //权益图片不能为空
	VipFileTypeErr                = New(69352, "权益图片格式不合法") //权益图片格式不合法
	VipPrivilegeNameTooLongErr    = New(69353, "权益名称过长")    //权益名称过长
	VipPrivilegeTitleTooLongErr   = New(69354, "权益标题过长")    //权益标题过长
	VipPrivilegeExplainTooLongErr = New(69355, "权益内容过长")    //权益内容过长

	VipAssociateBindPurchasedErr        = New(69450, "该用户已消费，不可换绑")          //该用户已消费，不可换绑
	VipAssociateOpenIDNotExsitErr       = New(69452, "openID 不存在")           //openID 不存在
	VipAssociateGrantOutTradeNoExsitErr = New(69460, " 第三方订单号已存在")           // 第三方订单号已存在
	VipAssociateGrantDurationErr        = New(69461, " 联合权益发放时长不合法")         // 联合权益发放时长不合法
	VipAssociateGrantLimitErr           = New(69462, "用户联合权益超过限制")           //用户联合权益超过限制
	VipAssociatePrizeKeyErr             = New(69470, " 奖品key不合法")            // 奖品key不合法
	VipAssociateGrantDayErr             = New(69471, " 联合权益发放天数不合法")         // 联合权益发放天数不合法
	VipAssociateWhiteIPListErr          = New(69472, " IP限制")                // IP限制
	VipCouponUniqueNoErr                = New(69480, " 奖品发放受限,重复的unique_no") // 奖品发放受限,重复的unique_no

	VipJavaAPIErr                = New(69500, "vip java接口调用失败") //vip java接口调用失败
	VipNotAutoRenewUserErr       = New(69550, "非自动续费用户")        //非自动续费用户
	VipRenewPayErr               = New(69551, "续约扣费失败")         //续约扣费失败
	VipRetryRenewPayNotExpireErr = New(69552, "无法重放未过期用户续约扣款")  //无法重放未过期用户续约扣款
	VipUserUnFrozenErr           = New(69600, "vip 未冻结")        //vip 未冻结
	VipWelfareCodeRunOut         = New(69700, "福利已经被领完了")       //福利已经被领完了
	VipWelfareOffLine            = New(69701, "该福利已下线，逛逛别的吧~")  //该福利已下线，逛逛别的吧~
	VipWelfareNotStart           = New(69702, "该福利还未上线")        //该福利还未上线
	VipWelfareAlreadyReceived    = New(69703, "你已领取过该福利")       //你已领取过该福利
	VipWelfareVipOnly            = New(69704, "只有会员才可领取福利")     //只有会员才可领取福利
	VipWelfareYearVipOnly        = New(69705, "只有年度大会员才可领取该福利") //只有年度大会员才可领取该福利
	VipWelfareCodeUpdConflict    = New(69706, "多人同时获取同一优惠码冲突")  //多人同时获取同一优惠码冲突
	VipWelfareNotExist           = New(69707, "你访问的福利不存在")      //你访问的福利不存在
	VipWelfareUploadMaxErr       = New(69708, "上传文件超过2w行")      //上传文件超过2w行
	VipWelfareRequestErr         = New(69709, "跳转url不能为空!")     //跳转url不能为空!
	VipWelfareUrlUnvalid         = New(69710, "上传图片url不合法")     //上传图片url不合法

	VipActivityNotStart          = New(69900, " 活动未开始,请稍后再来")    // 活动未开始,请稍后再来
	VipActivityHadEnd            = New(69901, " 活动已结束")          // 活动已结束
	VipActivityDeviceNotSupport  = New(69902, " 请将本页面分享至浏览器中购买") // 请将本页面分享至浏览器中购买
	VipActivityAccountNotSupport = New(69903, " 账号未绑定")          // 账号未绑定
	VipActivityProductLimit      = New(69904, " 超过限额，试试其他商品")    // 超过限额，试试其他商品
	VipEleUnionReceivePrizesErr  = New(69905, " 饿了么奖品发放失败")      // 饿了么奖品发放失败
	VipEleUnionBuyCanPurchaseErr = New(69906, " 饿了么联合会员限制购买")    // 饿了么联合会员限制购买
	VipEleUnionReqErr            = New(69907, " 饿了么服务请求失败")      // 饿了么服务请求失败
	VipEleUnionRespCodeErr       = New(69908, " 饿了么服务请求响应有误")    // 饿了么服务请求响应有误
	VipEleUnionBuyCanProductErr  = New(69909, " 不支持购买的商品类型")     // 不支持购买的商品类型
	VipEleUnionBuyGiveFirendErr  = New(69910, " 联合会员不支持好友赠送")    // 联合会员不支持好友赠送
	VipPlatformByIDNotFoundErr   = New(69911, " 未查到平台数据")        // 未查到平台数据
	VipOrderStatusNotSuccessErr  = New(69912, " 订单未支付成功")        // 订单未支付成功
	VipOrderEleVipGrantFaildErr  = New(69913, " 饿了么会员发放失败")      // 饿了么会员发放失败
	VipPriceInfoNotFoundErr      = New(69914, " 价格信息未找到")        // 价格信息未找到
	VipPriceNotEleProductErr     = New(69915, " 非饿了么商品")         // 非饿了么商品
	VipMailReqErr                = New(69916, " 会员购服务请求失败")      // 会员购服务请求失败
	VipMailRespCodeErr           = New(69917, " 会员购服务请求响应有误")    // 会员购服务请求响应有误
	VipActivityOrderNotFoundErr  = New(69918, " 未找到活动订单信息")      // 未找到活动订单信息
	VipAssociateThirdBindErr     = New(69919, " 饿了么绑定不成功")       // 饿了么绑定不成功

	SvenRepeat            = New(70001, " sven 数据重复")              // sven 数据重复
	CanalAddrFmtErr       = New(70002, "canal地址格式错误{host:port}!") //canal地址格式错误{host:port}!
	CanalAddrExist        = New(70003, "canal地址已存在")              //canal地址已存在
	CanalAddrNotFound     = New(70004, "canal地址不存在")              //canal地址不存在
	CanalApplyUpdateErr   = New(70005, "canal申请信息更新失败")           //canal申请信息更新失败
	CanalApplyErr         = New(70006, "canal申请失败")               //canal申请失败
	DatabasesUnmarshalErr = New(70007, "Databases解析失败")           //Databases解析失败
	GetConfigByNameErr    = New(70008, "根据name获取config信息失败")      //根据name获取config信息失败
	DatabusGroupErr       = New(70009, "databus group信息获取失败")     //databus group信息获取失败
	DatabusAppErr         = New(70010, "databus app信息获取失败")       //databus app信息获取失败
	DatabusActionErr      = New(70011, "databus action信息获取失败")    //databus action信息获取失败
	ConfigCreateErr       = New(70012, "配置中心生成配置文件失败")            //配置中心生成配置文件失败
	ConfigUpdateErr       = New(70013, "配置中心更新配置文件失败")            //配置中心更新配置文件失败
	SetConfigIDErr        = New(70014, "canal更新configId失败")       //canal更新configId失败
	CheckMasterErr        = New(70015, "canalcheckMaster验证失败")    //canalcheckMaster验证失败
	ConfigParseErr        = New(70016, "canal config信息解析失败")      //canal config信息解析失败
	NeedInfoErr           = New(70017, " 需求不存在")                  // 需求不存在
	NeedEditErr           = New(70018, " 需求不符合编辑要求")              // 需求不符合编辑要求
	NeedVerifyErr         = New(70019, "需求已审核")                   //需求已审核
	ConfigNotNow          = New(70020, "配置中心配置源文件非最新")            //配置中心配置源文件非最新
	DatabusDuplErr        = New(70021, "databus group重复")         //databus group重复
	// share
	ShareAlreadyAdd = New(71000, " 已经分享过了") // 已经分享过了

	// push
	PushSensitiveWordsErr  = New(72001, " 推送信息中有敏感词")          // 推送信息中有敏感词
	PushUUIDErr            = New(72002, " 调用添加推送任务接口请求重放了")    // 调用添加推送任务接口请求重放了
	PushBizAuthErr         = New(72003, " 业务方调用接口时token校验未通过") // 业务方调用接口时token校验未通过
	PushSilenceErr         = New(72004, " 业务方处于免打扰时间段，不允许推送")  // 业务方处于免打扰时间段，不允许推送
	PushBizForbiddenErr    = New(72005, " 业务方被禁止推送")           // 业务方被禁止推送
	PushUploadInvalidErr   = New(72006, " 上传的文件内容不符合规范")       // 上传的文件内容不符合规范
	PushRecordRepeatErr    = New(72007, " 该记录已经存在")            // 该记录已经存在
	PushServiceBusyErr     = New(72008, " 系统繁忙，请稍后再试")         // 系统繁忙，请稍后再试
	PushServiceFileSizeErr = New(72009, " 图片大小超过限制")           // 图片大小超过限制
	PushServiceFileExtErr  = New(72010, " 图片格式不支持")            // 图片格式不支持
	PushAdminDPNoDataErr   = New(72101, " 数据平台数据未准备好")         // 数据平台数据未准备好

	// coupon
	CouPonUseTooFrequently        = New(73001, " 观影劵使用劵操作太频繁")       // 观影劵使用劵操作太频繁
	CouPonNotEnoughErr            = New(73002, " 观影劵数量不够")           // 观影劵数量不够
	CouPonConsumeFaildErr         = New(73003, " 劵消费失败")             // 劵消费失败
	CouPonTypeNotExistErr         = New(73004, " 劵类型不存在")            // 劵类型不存在
	CouPonBatchNotExistErr        = New(73005, " 劵批次号不存在")           // 劵批次号不存在
	CouPonBatchNotEnoughErr       = New(73006, " 劵批次可发放劵数不足")        // 劵批次可发放劵数不足
	CouPonBatchTimeErr            = New(73007, " 批次开始时间不能大于等于过期时间")  // 批次开始时间不能大于等于过期时间
	CouPonBatchLimitErr           = New(73008, " 超过单人最大可领取劵数")       // 超过单人最大可领取劵数
	CouPonGrantTooFrequently      = New(73009, " 劵领取操作太频繁")          // 劵领取操作太频繁
	CouPonGrantMaxCountErr        = New(73010, " 该批次劵已发放完")          // 该批次劵已发放完
	CouPonHadUseErr               = New(73011, " 代金券已被使用")           // 代金券已被使用
	CouPonTokenNotFoundErr        = New(73012, " 券不存在")              // 券不存在
	CouPonHadBlockErr             = New(73013, " 券已被冻结")             // 券已被冻结
	CouPonHadExpireErr            = New(73014, " 券不在有效期内")           // 券不在有效期内
	CouPonNotFullPriceErr         = New(73015, " 券未满足可使用条件")         // 券未满足可使用条件
	CouPonOrderHadUseErr          = New(73016, " 订单已经使用过代金券")        // 订单已经使用过代金券
	CouPonStateCanNotCancelErr    = New(73017, " 不能解锁非使用中的劵")        // 不能解锁非使用中的劵
	CouPonUseFaildErr             = New(73018, " 劵使用失败")             // 劵使用失败
	CouPonNotifyStateErr          = New(73019, " 发货状态未知")            // 发货状态未知
	CouponAmountErr               = New(73020, " 代金券的满额金额必须小于等于券金额") // 代金券的满额金额必须小于等于券金额
	CouponUpdateFileErr           = New(73021, " 批量发放文件格式有误")        // 批量发放文件格式有误
	CouponBatchSalaryLimitErr     = New(73022, " 批量发放数量超过上限")        // 批量发放数量超过上限
	CouponBatchSalaryCountZeroErr = New(73023, " 批量发放数量不能为0")        // 批量发放数量不能为0
	CouPonPlatformNotSupportErr   = New(73024, " 平台限制使用")            // 平台限制使用
	CouponExpireErr               = New(73025, "请填写过期时间")            //请填写过期时间
	CouponReceiveErr              = New(73026, "券接收失败")              //券接收失败
	CouponBatchBlockErr           = New(73027, "批次信息被冻结")            //批次信息被冻结
	CouponBatchExpireTimeErr      = New(73028, "批次已过期")              //批次已过期
	CouPonProductNotSupportErr    = New(73029, " 商品限制使用")            // 商品限制使用
	CouponInfoStateBlockErr       = New(73100, "当前观影券状态不允许冻结")       //当前观影券状态不允许冻结
	CouponInfoStateUnblockErr     = New(73101, "当前观影券状态不允许解冻")       //当前观影券状态不允许解冻
	CouponSalaryCountErr          = New(73102, "批量发放数量不匹配")          //批量发放数量不匹配
	CouponSalaryHadRunErr         = New(73103, "批量发放任务已开始")          //批量发放任务已开始
	CouponTypeNotSupportErr       = New(73104, "批量发放类型不支持")          //批量发放类型不支持
	CouponNewYearNotStartErr      = New(73200, "元旦活动未开始")            //元旦活动未开始
	CouponNewYearIsEndErr         = New(73201, "元旦活动已结束")            //元旦活动已结束
	CouponNewYearIsOpenErr        = New(73202, "元旦活动卡牌已翻开")          //元旦活动卡牌已翻开
	CouponNewYearGrantErr         = New(73203, "元旦活动卡牌发放失败")         //元旦活动卡牌发放失败

	// coupon code
	CouponCodeVerifyFaildErr   = New(73300, "验证码错误,请重新输入")          //验证码错误,请重新输入
	CouponCodeNotFoundErr      = New(73301, "兑换码不存在,请重新输入")         //兑换码不存在,请重新输入
	CouponCodeUsedErr          = New(73302, "兑换码已使用,请重新输入")         //兑换码已使用,请重新输入
	CouponCodeBlockErr         = New(73303, "兑换码已冻结,请重新输入")         //兑换码已冻结,请重新输入
	CouponCodeLimitByMidErr    = New(73304, "单个用户兑换数量超过上限,请重新输入")   //单个用户兑换数量超过上限,请重新输入
	CouponCodeMaxLimitByMidErr = New(73305, "批次兑换数量超过上限,请重新输入")     //批次兑换数量超过上限,请重新输入
	CouponCodeCanNotUseErr     = New(73306, "兑换码非使用中状态")            //兑换码非使用中状态
	CouponCodeMaxLimitErr      = New(73307, "未设置批次码最大上线，或超过最大限制5W") //未设置批次码最大上线，或超过最大限制5W

	// realname 实名认证
	RealnameCaptureErr         = New(74001, " 实名验证码输入错误")     // 实名验证码输入错误
	RealnameCaptureSendTooMany = New(74002, " 实名验证码发送次数过于频繁") // 实名验证码发送次数过于频繁
	RealnameCaptureInvalid     = New(74003, " 实名验证码未发送或已失效")  // 实名验证码未发送或已失效
	RealnameCaptureErrTooMany  = New(74004, " 实名验证码错误次数过多")   // 实名验证码错误次数过多
	RealnameApplyAlready       = New(74005, " 实名认证已提交申请")     // 实名认证已提交申请
	RealnameInvalidName        = New(74006, " 实名姓名错误")        // 实名姓名错误
	RealnameImageIDErr         = New(74007, " 实名照片错误")        // 实名照片错误
	RealnameImageExpired       = New(74008, " 实名认证照片过期")      // 实名认证照片过期
	RealnameCardNumErr         = New(74009, " 实名证件号码错误")      // 实名证件号码错误
	RealnameCardBindAlready    = New(74010, " 实名证件已绑定")       // 实名证件已绑定
	RealnameAlipayAntispam     = New(74011, " 实名触发芝麻认证防刷")    // 实名触发芝麻认证防刷
	RealnameAlipayErr          = New(74012, " 实名芝麻认证服务错误")    // 实名芝麻认证服务错误
	RealnameAlipayApplyInvalid = New(74013, " 实名芝麻认证申请不合法")   // 实名芝麻认证申请不合法

	// activity
	ActivityNotExist           = New(75001, " 活动不存在")               // 活动不存在
	ActivityNotStart           = New(75002, " 活动没有开始")              // 活动没有开始
	ActivityOverEnd            = New(75003, " 活动已结束")               // 活动已结束
	ActivityHaveGuess          = New(75004, " 活动已竞猜")               // 活动已竞猜
	ActivityNotEnoughCoin      = New(75005, " 硬币不足")                // 硬币不足
	ActivityOverCoin           = New(75006, " 超额投币")                // 超额投币
	ActivityServerTimeout      = New(75007, " 服务超时")                // 服务超时
	ActivityKeyNotExists       = New(75008, " key不存在")              // key不存在
	ActivityKeyBindAlready     = New(75009, " key已绑定")              // key已绑定
	ActivityMidBindAlready     = New(75010, " 用户已绑定")               // 用户已绑定
	ActivityNotBind            = New(75011, " 用户未绑定")               // 用户未绑定
	ActivityIDNotExists        = New(75012, " ID不存在")               // ID不存在
	ActivityAwardAlready       = New(75013, " 奖品已兑换")               // 奖品已兑换
	ActivityNoAward            = New(75014, " 没有奖品")                // 没有奖品
	ActivityNotOwner           = New(75015, " 不是该point点Owner")      // 不是该point点Owner
	ActivityHasUnlock          = New(75016, " 该point点已解锁")          // 该point点已解锁
	ActivityGameResult         = New(75017, " 请选择游戏结果")             // 请选择游戏结果
	ActivityUserAchieveFail    = New(75018, " 用户成就记录获取失败")          // 用户成就记录获取失败
	ActivityUserPointFail      = New(75019, " 用户point记录获取失败")       // 用户point记录获取失败
	ActivityAchieveFail        = New(75020, " 成就列表获取失败")            // 成就列表获取失败
	ActivityPointFail          = New(75021, " point列表获取失败")         // point列表获取失败
	ActivityNoAchieve          = New(75022, " 没有成就")                // 没有成就
	ActivityKeyFail            = New(75023, " key获取失败")             // key获取失败
	ActivityMidFail            = New(75024, " mid获取失败")             // mid获取失败
	ActivityNotAdmin           = New(75025, " 非管理员登录")              // 非管理员登录
	ActivityAddAchieveFail     = New(75026, " 获得成就失败")              // 获得成就失败
	ActivityLackHp             = New(75027, " 您的HP不足")              // 您的HP不足
	ActivityMaxHp              = New(75028, " 您的HP已满")              // 您的HP已满
	ActivityNotAwardAdmin      = New(75029, " 您不是奖品兑换管理员")          // 您不是奖品兑换管理员
	ActivityNotLotteryAdmin    = New(75030, " 您不是抽奖管理员")            // 您不是抽奖管理员
	ActivityLotteryFail        = New(75031, " 抽奖失败")                // 抽奖失败
	ActivityUnlockFail         = New(75032, " 解锁失败")                // 解锁失败
	ActivityNotLotteryAchieve  = New(75033, " 该成就不支持抽奖")            // 该成就不支持抽奖
	ActivityLikeIPFrequence    = New(75034, "点赞活动ip访问过于频繁-访问过快")    //点赞活动ip访问过于频繁-访问过快
	ActivityLikeScoreLower     = New(75035, "账户score过低不支持点赞-异常账号!") //账户score过低不支持点赞-异常账号!
	ActivityLikeLevelLimit     = New(75036, "-用户等级不够!")             //-用户等级不够!
	ActivityLikeMemberLimit    = New(75037, "-用户注册不足7天！")           //-用户注册不足7天！
	ActivityLikeNotStart       = New(75038, "-评分未开始!")              //-评分未开始!
	ActivityLikeHasEnd         = New(75039, "-评分已结束!")              //-评分已结束!
	ActivityLikeHasLike        = New(75040, "-已点赞过!")               //-已点赞过!
	ActivityLikeHasGrade       = New(75041, "-已评过分!")               //-已评过分!
	ActivityLikeRegisterLimit  = New(75042, "-晚于活动限制注册时间!")         //-晚于活动限制注册时间!
	ActivityLikeHasVote        = New(75043, "-已投过票!")               //-已投过票!
	ActivityHasOffLine         = New(75044, "-活动已经下线")              //-活动已经下线
	ActivityLikeHasOffLine     = New(75045, "-活动稿件已经下线")            //-活动稿件已经下线
	ActivityLikeBeforeRegister = New(75046, "-早于活动限制注册时间!")         //-早于活动限制注册时间!
	ActivityHasMissionGroup    = New(75047, "-已发起过活动!")             //-已发起过活动!
	ActivityMGNotYourself      = New(75048, "-不支持给自己助力哦!")          //-不支持给自己助力哦!
	ActivityMissionNotStart    = New(75049, "-助攻未开始!")              //-助攻未开始!
	ActivityMissionHasEnd      = New(75050, "-助攻已结束!")              //-助攻已结束!
	ActivityHasMission         = New(75051, "-已助攻过!")               //-已助攻过!
	ActivityOverMissionLimit   = New(75052, "-超出可助攻上限!")            //-超出可助攻上限!
	ActivityHasAward           = New(75053, "-重复领取")                //-重复领取
	ActivityNotAward           = New(75054, "-非法领取")                //-非法领取
	ActivityTelValid           = New(75055, "-未绑定有效手机号码")           //-未绑定有效手机号码
	ActivityOverDailyScore     = New(75056, "-超过单日投票上线")            //-超过单日投票上线
	ActivityBnjTimeCancel      = New(75057, " 倒计时活动取消")             // 倒计时活动取消
	ActivityBnjResetCD         = New(75058, " 倒计时重置CD")             // 倒计时重置CD
	ActivityBnjTimeFinish      = New(75059, " 倒计时已完成")              // 倒计时已完成
	ActivityBnjNotSub          = New(75060, " 未预约拜年祭活动")            // 未预约拜年祭活动
	ActivityBnjSubLow          = New(75061, " 该宝箱预约人数未达成")          // 该宝箱预约人数未达成
	ActivityBnjHasReward       = New(75062, " 该宝箱奖励已被领取")           // 该宝箱奖励已被领取
	ActivityBnjRewardFail      = New(75063, " 宝箱奖励领取失败")            // 宝箱奖励领取失败
	ActivityRewardConfErr      = New(75064, " 宝箱奖励配置错误")            // 宝箱奖励配置错误
	ActivityMemberBlocked      = New(75065, "-封禁用户无法操作")            //-封禁用户无法操作
	ActivityKfcHasUsed         = New(75066, "-code已经被使用")           //-code已经被使用
	ActivityKfcNotExist        = New(75067, "-code不存在")             //-code不存在
	ActivityKfcNotGiveOut      = New(75068, "-code未发放")             //-code未发放
	ActivityKfcSqlError        = New(75069, "-发生未知错误")              //-发生未知错误

	// TV
	TvUpperNotInList      = New(76001, " up主不存在")                             // up主不存在
	TvDangbeiWrongType    = New(76002, " 当贝页面-内容类型错误")                        // 当贝页面-内容类型错误
	TvDangbeiPageNotExist = New(76003, " 当贝页面不存在")                            // 当贝页面不存在
	TvSearchOrderWrong    = New(76004, " 搜索所提供的order字段有误")                    // 搜索所提供的order字段有误
	TvVideoNotFound       = New(76005, " 视频找不到，文案：啊叻？视频不见啦！看看别的吧～")           // 视频找不到，文案：啊叻？视频不见啦！看看别的吧～
	TvLabelExist          = New(76006, " 所添加的索引标签已存在")                        // 所添加的索引标签已存在
	TvSyncErr             = New(76007, " 同步牌照错误")                             // 同步牌照错误
	TvPGCRankEmpty        = New(76008, " pgc排行榜结果为空")                         // pgc排行榜结果为空
	TvPGCRankNewEPNil     = New(76009, " pgc排行榜结果中newEP为空")                   // pgc排行榜结果中newEP为空
	TvAllDataAuditing     = New(76011, " tv稿件或者season下所有的数据为审核中，即详情页无可展示数据时") // tv稿件或者season下所有的数据为审核中，即详情页无可展示数据时
	TvPriceTimeConflict   = New(76101, " 折扣价时间冲突，请检查已有折扣时间段")                 // 折扣价时间冲突，请检查已有折扣时间段
	TvVipProductExit      = New(76102, " 该产品id已存在")                           // 该产品id已存在
	TvVipProdSyncErr      = New(76103, " 同步产品包到云视听错误,请尝试重试")                  // 同步产品包到云视听错误,请尝试重试
	TVVipSuitTypeConflict = New(76104, " 请先下线已存在的升级产品")                       // 请先下线已存在的升级产品

	// manager
	ManagerTagDelErr     = New(77001, " manager tag不能删除")   // manager tag不能删除
	ManagerUIDNOTExist   = New(77002, " manager管理员不存在")     // manager管理员不存在
	ManagerFlowForbidden = New(77003, " manager 工作流被禁用")    // manager 工作流被禁用
	ManagerTagTypeDelErr = New(77004, " manager tag类型不能删除") // manager tag类型不能删除
	// APP
	AppNotData = New(78000, " app not data") // app not data

	// subtitle
	SubtitleDrfatExist            = New(79001, " 当前已存在未过审字幕")      // 当前已存在未过审字幕
	SubtitleDrfatNotExist         = New(79002, " 当前字幕草稿不存在")       // 当前字幕草稿不存在
	SubtitleDrfatUnSubmit         = New(79003, " 当前字幕未提交")         // 当前字幕未提交
	SubtitleDelUnExist            = New(79004, " 删除不存在的字幕")        // 删除不存在的字幕
	SubtitlePermissionDenied      = New(79005, " 字幕操作权限不足")        // 字幕操作权限不足
	SubtitleDenied                = New(79006, " 视频禁止投稿")          // 视频禁止投稿
	SubtileLanLocked              = New(79007, " 当前语言版本锁定")        // 当前语言版本锁定
	SubtitleUnValid               = New(79008, " 字幕不合法")           // 字幕不合法
	SubtitleWaveFormFailed        = New(79009, " 波形图调用失败")         // 波形图调用失败
	SubtitlePubNotExist           = New(79010, " 发布的字幕不存在")        // 发布的字幕不存在
	SubtitleIllegalLanguage       = New(79011, " 不合法的语言")          // 不合法的语言
	SubtitleNotPublish            = New(79012, " 当前字幕未发布")         // 当前字幕未发布
	SubtitleTimeUnValid           = New(79013, " 字幕时间不合法")         // 字幕时间不合法
	SubtitleSizeLimit             = New(79014, " 字幕超过限制")          // 字幕超过限制
	SubtitleOriginUnValid         = New(79015, " 原字幕不合法")          // 原字幕不合法
	SubtitleLocationUnValid       = New(79016, " 字幕位置不合法")         // 字幕位置不合法
	SubtitleUserBalcked           = New(79017, " 账号黑名单")           // 账号黑名单
	SubtitleStatusUnValid         = New(79018, " 字幕状态不合法")         // 字幕状态不合法
	SubtitleAlreadyHasDraft       = New(79019, " 前存在草稿或者待审核状态的字幕") // 前存在草稿或者待审核状态的字幕
	SubtitleVideoDurationOverFlow = New(79020, " 字幕时间点超过视频时间长度")   // 字幕时间点超过视频时间长度
	SubtitleDuarionMustThanZero   = New(79021, " 字幕的持续时间必须大于0")    // 字幕的持续时间必须大于0

	// MCN
	// --82000~82499 前台错误
	MCNNotAllowed                        = New(82001, " 没有权限操作")                                    // 没有权限操作
	MCNUpCannotBind                      = New(82002, " 无法绑定Up主，已被绑定")                              // 无法绑定Up主，已被绑定
	MCNUpBindUpAlreadyInProgress         = New(82003, " 已存在与up的绑定在审核中")                             // 已存在与up的绑定在审核中
	MCNUpBindUpDuplicatePhone            = New(82004, " 该电话号码已绑定")                                  // 该电话号码已绑定
	MCNUpBindUpDuplicateIDCard           = New(82005, " 该身份证号码已绑定")                                 // 该身份证号码已绑定
	MCNUpBindUpDuplicateCompanyLicenseID = New(82006, " 该营业执照号码已绑定")                                // 该营业执照号码已绑定
	MCNUpBindUpDuplicateCompanyName      = New(82007, " 该公司名称已绑定")                                  // 该公司名称已绑定
	MCNUpBindUpSTimeLtETime              = New(82008, " up主绑定的开始时间必须小于结束时间")                        // up主绑定的开始时间必须小于结束时间
	MCNUpBindUpIsBlocked                 = New(82009, " up主已被封禁")                                   // up主已被封禁
	MCNUpBindUpDateError                 = New(82010, " 日期错误，请检查")                                  // 日期错误，请检查
	MCNStateInvalid                      = New(82011, " MCN状态异常")                                   // MCN状态异常
	MCNUpBindInvalid                     = New(82012, " 该绑定请求已失效")                                  // 该绑定请求已失效
	MCNUpBindInvalidURL                  = New(82013, " 该绑定的站外up主链接错误，请检查")                         // 该绑定的站外up主链接错误，请检查
	MCNUpOutSiteIsNotQualified           = New(82014, " 站外Up主需要满足（1.粉丝数≤100  或  2. 投稿数＜2及90天内未投稿）") // 站外Up主需要满足（1.粉丝数≤100  或  2. 投稿数＜2及90天内未投稿）
	MCNUpBindUpIsBlueUser                = New(82015, " 该up主为蓝V用户，不能被绑定")                           // 该up主为蓝V用户，不能被绑定
	MCNUpSignStateInvalid                = New(82016, " 该Up主签约状态异常")                                // 该Up主签约状态异常
	MCNChangePermissionAlreadyInProgress = New(82020, " 等待UP主授权或运营审核")                              // 等待UP主授权或运营审核
	MCNChangePermissionLackPermission    = New(82021, " 您缺少要申请的权限")                                 // 您缺少要申请的权限
	MCNChangePermissionSamePermission    = New(82022, " 权限没有任何变更")                                  // 权限没有任何变更
	MCNPublicationFailTimeLimit          = New(82030, " 刊例价每月只能更改一次")                               // 刊例价每月只能更改一次
	MCNPermissionInsufficient            = New(82040, " 权限不足")                                      // 权限不足

	// --82500~82999 后台错误
	MCNSignCycleNotUQErr        = New(82501, " mcn签约周期冲突")                          // mcn签约周期冲突
	MCNUnknownFileTypeErr       = New(82502, " 只允许jpg、png、pdf、word格式的文件上传")         // 只允许jpg、png、pdf、word格式的文件上传
	MCNSignUnknownReviewErr     = New(82503, " mcn签约未知审核方式")                        // mcn签约未知审核方式
	MCNSignOnlyReviewOpErr      = New(82504, " 只有待审核中的mcn才能操作")                     // 只有待审核中的mcn才能操作
	MCNUpUnknownReviewErr       = New(82505, " mcn和up主绑定未知审核方式")                    // mcn和up主绑定未知审核方式
	MCNUpOnlyReviewOpErr        = New(82506, " 有待审核中的mcn和up主绑定才能操作")                // 有待审核中的mcn和up主绑定才能操作
	MCNContractFileSize         = New(82507, " mcn合同上传的大小，允许20M")                   // mcn合同上传的大小，允许20M
	MCNCSignUnknownInfoErr      = New(82508, " 未知mcn信息")                            // 未知mcn信息
	MCNCUpUnknownInfoErr        = New(82509, " 未知mcn-up主绑定信息")                      // 未知mcn-up主绑定信息
	MCNUnknownFileExt           = New(82510, " 未知上传文件后缀名")                          // 未知上传文件后缀名
	MCNSignNoOkState            = New(82511, " 处于封禁中、审核中、未申请中、驳回中、签约中、待开启中的状态不可录入") // 处于封禁中、审核中、未申请中、驳回中、签约中、待开启中的状态不可录入
	MCNUpPassOnEffectSign       = New(82512, " up主通过必须是有效的签约状态")                    // up主通过必须是有效的签约状态
	MCNSignIsBlocked            = New(82513, " mcn管理用户已被封禁")                        // mcn管理用户已被封禁
	MCNSignEtimeNLEQNowTime     = New(82514, " mcn签约结束时间不能小于等于当前时间")                // mcn签约结束时间不能小于等于当前时间
	MCNSignStateFlowErr         = New(82515, " mcn状态流转错误,请联系开发人员")                  // mcn状态流转错误,请联系开发人员
	MCNRecommendUpStateFlowErr  = New(82516, " 推荐的up状态流转错误,请联系开发人员")                // 推荐的up状态流转错误,请联系开发人员
	MCNRecommendUpInPool        = New(82517, " 该up主已经被推荐")                          // 该up主已经被推荐
	MCNRecommendUpMidsIsEmpty   = New(82518, " 推荐池被操作的up主不能为空")                     // 推荐池被操作的up主不能为空
	MCNUpPermitUnknownReviewErr = New(82519, " 未知mcn-up的权限变更审核方式")                  // 未知mcn-up的权限变更审核方式
	MCNUpAbnormalDataErr        = New(82520, " 异常数据")                               // 异常数据
	MCNUpPermitStateFlowErr     = New(82521, " mcn-up的权限变更的状态流转错误,请联系开发人员")         // mcn-up的权限变更的状态流转错误,请联系开发人员

	MCNRenewalNotInRangeErr = New(82601, " MCN续约不在范围内") // MCN续约不在范围内
	MCNRenewalAlreadyErr    = New(82602, " MCN已续过约")    // MCN已续过约
	MCNRenewalDateErr       = New(82603, " MCN续约周期错误")  // MCN续约周期错误

	// esports
	EsportsContestNotExist  = New(83001, " 你所订阅的赛程不存在")                  // 你所订阅的赛程不存在
	EsportsContestMaxCount  = New(83002, " 你订阅赛程数已达上限")                  // 你订阅赛程数已达上限
	EsportsContestFavDel    = New(83003, " 该赛程未订阅，不能取消哦~")               // 该赛程未订阅，不能取消哦~
	EsportsContestFavExist  = New(83004, " 该赛程已订阅，不能重复订阅哦~")             // 该赛程已订阅，不能重复订阅哦~
	EsportsContestNotDay    = New(83005, " 仅可订阅15天内的赛事哦~")               // 仅可订阅15天内的赛事哦~
	EsportsContestStart     = New(83006, " 你订阅的赛程已经开始啦~快来直播间观看吧~")       // 你订阅的赛程已经开始啦~快来直播间观看吧~
	EsportsContestEnd       = New(83007, " 你订阅的赛程已经结束啦~可以点击回放和集锦进行观看哦~") // 你订阅的赛程已经结束啦~可以点击回放和集锦进行观看哦~
	EsportsContestFavNot    = New(83008, " 该赛程不可订阅哦~")                   // 该赛程不可订阅哦~
	EsportsActNotExist      = New(83009, " 赛事活动不存在")                     // 赛事活动不存在
	EsportsActVideoNotExist = New(83010, " 赛事活动视频不存在")                   // 赛事活动视频不存在
	EsportsActPointNotExist = New(83011, " 比赛数据不存在")                     // 比赛数据不存在
	EsportsActKnockNotExist = New(83012, " 淘汰赛数据不存在")                    // 淘汰赛数据不存在
	EsportsModNameErr       = New(83050, " 模块名称重复~")                     // 模块名称重复~
	EsportsActModNot        = New(83051, " 模块不属于该赛事活动~")                 // 模块不属于该赛事活动~
	EsportsActModErr        = New(83052, " 模块信息不正确~")                    // 模块信息不正确~
	EsportsModArcErr        = New(83053, " 模块稿件不正确~")                    // 模块稿件不正确~
	EsportsTreeNodeErr      = New(83054, " 节点不属于该赛事详情~")                 // 节点不属于该赛事详情~
	EsportsTreeDetailErr    = New(83055, " 赛事活动详情不存在~")                  // 赛事活动详情不存在~
	EsportsTreeEmptyErr     = New(83056, " 当前没有任何记录，请编辑后提交~")            // 当前没有任何记录，请编辑后提交~
	EsportsMultiEdit        = New(83057, " 节点不能多人同时保存~")                 // 节点不能多人同时保存~
	EsportsArcServerErr     = New(83058, " 稿件服务出错~")                     // 稿件服务出错~
	EsportsContestDataErr   = New(83059, " 比赛数据不正确~")                    // 比赛数据不正确~

	CardNotEffectiveErr    = New(85000, "非有效卡片")      //非有效卡片
	CardEquipNotVipErr     = New(85001, "非大会员不可装备")   //非大会员不可装备
	CardIDNotFoundErr      = New(85003, "不存在的卡片ID")   //不存在的卡片ID
	GroupIDNotFoundErr     = New(85004, "不存在的分类ID")   //不存在的分类ID
	CardFileUploadFaildErr = New(85005, "卡片图片文件上传失败") //卡片图片文件上传失败
	CardNameTooLongErr     = New(85006, "卡片名称过长")     //卡片名称过长
	CardImageEmptyErr      = New(85007, "卡片图片不能为空")   //卡片图片不能为空
	CardNameExistErr       = New(85008, "卡片名称已存在")    //卡片名称已存在
	CardGroupNameExistErr  = New(85009, "卡片组名称已存在")   //卡片组名称已存在

	//passport
	//passport-login 86000~86299
	//passport-user  86300~86599
	//passport-sns	 86600~
	PassportSnsMidAlreadyBindQQ    = New(86600, "mid已绑定QQ号") //mid已绑定QQ号
	PassportSnsMidAlreadyBindWEIBO = New(86601, "mid已绑定微博")  //mid已绑定微博
	PassportSnsQQAlreadyBind       = New(86610, "QQ号已被绑定")   //QQ号已被绑定
	PassportSnsWEIBOAlreadyBind    = New(86611, "微博已经绑定")    //微博已经绑定
	PassportSnsRequestErr          = New(86660, "请求第三方失败")   //请求第三方失败

	//playurl
	PlayURLNotLogin = New(87000, "用户未登录") //用户未登录
	PlayURLNotPay   = New(87001, "稿件未付费") //稿件未付费

	// upcpay
	UGCPayAssetInvalid         = New(88001, " ugcpay 内容无效")            // ugcpay 内容无效
	UGCPayAssetCantBuy         = New(88002, " ugcpay 内容不可购买")          // ugcpay 内容不可购买
	UGCPayOrderInvalid         = New(88003, " ugcpay 订单无效")            // ugcpay 订单无效
	UGCPayOrderNotPay          = New(88004, " ugcpay 订单未付款")           // ugcpay 订单未付款
	UGCPayAssetPaid            = New(88005, " ugcpay 内容已支付")           // ugcpay 内容已支付
	UGCPayDependArchiveErr     = New(88101, " ugcpay archive依赖错误")     // ugcpay archive依赖错误
	UGCPayDependPayPlatformErr = New(88102, " ugcpay payplatform依赖错误") // ugcpay payplatform依赖错误

	ToViewPayUGC  = New(90000, "付费稿件")        //付费稿件
	ToViewOverMax = New(90001, "塞满啦！先看看库存吧~") //塞满啦！先看看库存吧~

	// uprating
	UpRatingNoPermission = New(91000, " 用户无权限")   // 用户无权限
	UpRatingNoData       = New(91001, " 用户无数据")   // 用户无数据
	UpRatingScoreLimit   = New(91002, " 用户分数未达标") // 用户分数未达标

	//aegis
	AegisUniqueAlreadyExist = New(92001, "当前流程%s %s 已存在")      //当前流程%s %s 已存在
	AegisTokenNotAssign     = New(92002, "当前绑定令牌%s不是赋值语句")     //当前绑定令牌%s不是赋值语句
	AegisTokenNotFound      = New(92003, "当前令牌不存在")            //当前令牌不存在
	AegisFlowNoToken        = New(92004, "当前节点没绑定令牌")          //当前节点没绑定令牌
	AegisFlowNotFound       = New(92005, "当前节点不存在")            //当前节点不存在
	AegisFlowDisabled       = New(92006, "当前节点已被禁用")           //当前节点已被禁用
	AegisFlowNoEnableTran   = New(92007, "当前节点所有下游变化不可用")      //当前节点所有下游变化不可用
	AegisFlowNoFromDir      = New(92008, "当前节点上游没有有向线")        //当前节点上游没有有向线
	AegisFlowBinded         = New(92009, "当前流程节点已被关联，请取消后再禁用") //当前流程节点已被关联，请取消后再禁用
	AegisTranBinded         = New(92010, "当前流程变化已被关联，请取消后再禁用") //当前流程变化已被关联，请取消后再禁用
	AegisTranNotFound       = New(92011, "当前变化不存在")            //当前变化不存在
	AegisTranDisabled       = New(92012, "当前变化已被禁用")           //当前变化已被禁用
	AegisTranNoFlow         = New(92013, "当前变化没有可用输出节点")       //当前变化没有可用输出节点
	AegisRunInDiffNet       = New(92014, "当前资源在不同网内运行")        //当前资源在不同网内运行
	AegisNotRunInFlow       = New(92015, "当前资源不在当前节点上运行")      //当前资源不在当前节点上运行
	AegisFlowDiffNet        = New(92016, "节点不在同一网上")           //节点不在同一网上
	AegisNotRunInRange      = New(92017, "当前资源不在该业务或网内运行")     //当前资源不在该业务或网内运行
	AegisNotTriggerFlow     = New(92018, "当前资源没有触发当前节点上的变化")   //当前资源没有触发当前节点上的变化
	AegisDirOrderConflict   = New(92019, "当前有向线顺序(%s)冲突:%s")   //当前有向线顺序(%s)冲突:%s
	AegisNetErr             = New(92020, "当前网配置错误")            //当前网配置错误

	AegisTaskErr         = New(92021, "任务操作错误")       //任务操作错误
	AegisResourceErr     = New(92022, "资源操作失败")       //资源操作失败
	AegisBusinessCfgErr  = New(92023, "业务配置错误")       //业务配置错误
	AegisSearchErr       = New(92024, "搜索接口错误")       //搜索接口错误
	AegisDuplicateErr    = New(92025, "资源重复")         //资源重复
	AegisReservedCfgErr  = New(92026, "保留字段配置错误")     //保留字段配置错误
	AegisBusinessSyncErr = New(92027, "业务回调错误")       //业务回调错误
	AegisTaskBusy        = New(92028, "任务调度繁忙，请重试")   //任务调度繁忙，请重试
	AegisTaskFinish      = New(92029, "任务已被处理，无需再操作") //任务已被处理，无需再操作

	// TV VIP
	TVIPTokenErr             = New(93000, " 查询令牌无效")       // 查询令牌无效
	TVIPTokenExpire          = New(93001, " 查询令牌过期")       // 查询令牌过期
	TVIPPanelNotFound        = New(93002, " 套餐不存在")        // 套餐不存在
	TVIPPanelNotSuitalbe     = New(93003, " 套餐不用")         // 套餐不用
	TVIPNotContracted        = New(93004, " 用户未签约连续包月")    // 用户未签约连续包月
	TVIPRenewTooEarly        = New(93005, " 续费过早")         // 续费过早
	TVIPRenewTooLate         = New(93006, " 续费过晚")         // 续费过晚
	TVIPNotVip               = New(93007, " 非VIP用户")       // 非VIP用户
	TVIPPayOrderExpired      = New(93008, " 订单超时适")        // 订单超时适
	TVIPPayOrderNotFound     = New(93009, "  订单不存在")       //  订单不存在
	TVIPSignErr              = New(93010, " 签名错误")         // 签名错误
	TVIPBuyNumExceeded       = New(93011, " 购买数量超标")       // 购买数量超标
	TVIPYstSystemErr         = New(93012, " 系统异常")         // 系统异常
	TVIPYstRequestErr        = New(93013, " 云视听请求错误")      // 云视听请求错误
	TVIPYstUnknownErr        = New(93014, " 云视听未知异常")      // 云视听未知异常
	TVIPYstUnknownTradeState = New(93015, " 云视听交易状态异常")    // 云视听交易状态异常
	TVIPDupOrderNo           = New(93016, " 重复的订单号")       // 重复的订单号
	TVIPQrDevideUnsupported  = New(93017, " 扫码设备不支持")      // 扫码设备不支持
	TVIPGiveMVipFailed       = New(93018, " 主站大会员赠送失败")    // 主站大会员赠送失败
	TVIPBatchIdNotFound      = New(93019, " 主站大会员批次id不存在") // 主站大会员批次id不存在
	TVIPBuyRateExceeded      = New(93020, " 购买过于频繁")       // 购买过于频繁
	TVIPMVipRateExceeded     = New(93021, " 购买大会员升级套餐频繁")  // 购买大会员升级套餐频繁

	// user
	PasswordEncodeErr = New(94101, "密码格式化异常") //密码格式化异常
	RoleNameErr       = New(94102, "角色名不能为空") //角色名不能为空

)
