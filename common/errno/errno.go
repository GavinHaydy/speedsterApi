package errno

const (
	Ok = 0

	ErrParam            = 10001
	ErrServer           = 10002
	ErrNonce            = 10003
	ErrTimeStamp        = 10004
	ErrRPCFailed        = 10005
	ErrInvalidToken     = 10006
	ErrMarshalFailed    = 10007
	ErrUnMarshalFailed  = 10008
	ErrMustDID          = 10011
	ErrMustSN           = 10012
	ErrHttpFailed       = 10013
	ErrGenTokenFailed   = 10014
	ErrRedisFailed      = 10100
	ErrMongoFailed      = 10101
	ErrPgsqlFailed      = 10102
	ErrRecordNotFound   = 10103
	ErrSelectDbFailed   = 10104
	ErrUpdateDataFailed = 10105

	ErrSignError       = 20001
	ErrRepeatRequest   = 20002
	ErrMustLogin       = 20003
	ErrAuthFailed      = 20004
	ErrAccountNotFound = 20006
	ErrPasswordFailed  = 20007
	ErrRegisterFailed  = 20008
	ErrTokenTypeFailed = 20009

	ErrCompanyNotRemoveMember   = 20101
	ErrCompanyNotFound          = 20102
	ErrCompanySupUpdatePassword = 20103

	ErrUserDisable         = 20201
	ErrUserNotRole         = 20202
	ErrUserForbidden       = 20203
	ErrYetAccountRegister  = 20204
	ErrYetNicknameRegister = 20205
	ErrYetEmailRegister    = 20206
	ErrYetEmailValid       = 20207
	ErrYetEmailNotFound    = 20208
	ErrYetUserNotFound     = 20209
	ErrYetPhoneRegister    = 20210

	ErrRoleNotDel    = 20401
	ErrRoleNotChange = 20402
	ErrRoleExists    = 20403
	ErrRoleNotExists = 20404
	ErrRoleForbidden = 20405

	ErrFileMaxSize  = 20501
	ErrFileMaxLimit = 20502

	ErrNoticeNameRepeat       = 20603
	ErrNoticeGroupNameRepeat  = 20604
	ErrNoticeWebhookURLRepeat = 20605
	ErrNoticeAppIDRepeat      = 20606
	ErrNoticeRelateNotNull    = 20607
)

// CodeAlertMap 错图码映射错误提示，展示给用户
var CodeAlertMap = map[int]string{
	Ok:                        "成功",
	ErrServer:                 "服务器错误",
	ErrParam:                  "参数校验错误",
	ErrSignError:              "签名错误",
	ErrRecordNotFound:         "数据库记录不存在",
	ErrSelectDbFailed:         "数据查询失败",
	ErrRPCFailed:              "请求下游RPC服务失败",
	ErrInvalidToken:           "无效的token",
	ErrMarshalFailed:          "序列化失败",
	ErrUnMarshalFailed:        "反序列化失败",
	ErrRedisFailed:            "redis操作失败",
	ErrMongoFailed:            "mongo操作失败",
	ErrPgsqlFailed:            "pgsql操作失败",
	ErrMustLogin:              "没有获取到登录态",
	ErrHttpFailed:             "请求下游Http服务失败",
	ErrAuthFailed:             "认证错误",
	ErrYetAccountRegister:     "用户账户已注册",
	ErrUserDisable:            "用户账户已封禁",
	ErrFileMaxSize:            "文件大小超过2M",
	ErrRoleNotDel:             "该角色不可删除",
	ErrRoleNotChange:          "角色不可修改",
	ErrUserNotRole:            "用户不存在角色",
	ErrUserForbidden:          "无权访问",
	ErrRoleExists:             "角色名重复",
	ErrCompanyNotRemoveMember: "不能移除超管",
	ErrYetNicknameRegister:    "昵称已存在",
	ErrYetEmailRegister:       "邮箱已注册",
	ErrYetEmailValid:          "邮箱格式不正确",
	ErrYetPhoneRegister:       "手机号已注册",
	ErrRoleNotExists:          "角色不存在",
	ErrRoleForbidden:          "不能添加该角色，角色较高",
	ErrAccountNotFound:        "账号不存在",
	ErrPasswordFailed:         "密码错误",
	ErrRegisterFailed:         "注册失败",
	ErrYetEmailNotFound:       "邮箱不能为空",
	ErrYetUserNotFound:        "用户不能为空",
	ErrGenTokenFailed:         "token生成失败",
	ErrTokenTypeFailed:        "token类型错误",
	ErrUpdateDataFailed:       "数据修改失败",
}

// CodeMsgMap 错误码映射错误信息，不展示给用户
var CodeMsgMap = map[int]string{
	Ok:                          "success",
	ErrServer:                   "internal server error",
	ErrParam:                    "param error",
	ErrSignError:                "signature error",
	ErrRepeatRequest:            "repeat request",
	ErrNonce:                    "nonce error",
	ErrTimeStamp:                "timestamp error",
	ErrRecordNotFound:           "record not found",
	ErrSelectDbFailed:           "select database failed",
	ErrRPCFailed:                "rpc failed",
	ErrInvalidToken:             "invalid token",
	ErrMarshalFailed:            "marshal failed",
	ErrUnMarshalFailed:          "unmarshal failed",
	ErrRedisFailed:              "redis operate failed",
	ErrMongoFailed:              "mongo operate failed",
	ErrPgsqlFailed:              "mysql operate failed",
	ErrMustLogin:                "must login",
	ErrMustDID:                  "must DID",
	ErrMustSN:                   "must SN",
	ErrHttpFailed:               "http failed",
	ErrAuthFailed:               "auth failed",
	ErrYetAccountRegister:       "ErrYetAccountRegister",
	ErrUserDisable:              "ErrUserDisable",
	ErrFileMaxSize:              "ErrFileMaxSize",
	ErrRoleNotDel:               "ErrRoleNotDel",
	ErrRoleNotChange:            "ErrRoleNotChange",
	ErrUserNotRole:              "ErrUserNotRole",
	ErrUserForbidden:            "ErrUserForbidden",
	ErrRoleExists:               "ErrRoleExists",
	ErrCompanyNotRemoveMember:   "ErrCompanyNotRemoveMember",
	ErrYetNicknameRegister:      "ErrYetNicknameRegister",
	ErrYetEmailRegister:         "ErrYetEmailRegister",
	ErrYetEmailValid:            "ErrYetEmailValid",
	ErrYetPhoneRegister:         "ErrYetPhoneRegister",
	ErrRoleNotExists:            "ErrRoleNotExists",
	ErrRoleForbidden:            "ErrRoleForbidden",
	ErrAccountNotFound:          "ErrAccountNotFound",
	ErrPasswordFailed:           "ErrPasswordFailed",
	ErrRegisterFailed:           "ErrRegisterFailed",
	ErrFileMaxLimit:             "ErrFileMaxLimit",
	ErrCompanyNotFound:          "ErrCompanyNotFound",
	ErrNoticeNameRepeat:         "ErrNoticeNameRepeat",
	ErrNoticeGroupNameRepeat:    "ErrNoticeGroupNameRepeat",
	ErrNoticeWebhookURLRepeat:   "ErrNoticeWebhookURLRepeat",
	ErrCompanySupUpdatePassword: "ErrCompanySupUpdatePassword",
	ErrNoticeAppIDRepeat:        "ErrNoticeAppIDRepeat",
	ErrNoticeRelateNotNull:      "ErrNoticeRelateNotNull",
	ErrYetEmailNotFound:         "ErrYetEmailNotFound",
	ErrYetUserNotFound:          "ErrYetUserNotFound",
	ErrGenTokenFailed:           "ErrGenTokenFailed",
	ErrTokenTypeFailed:          "ErrTokenTypeFailed",
	ErrUpdateDataFailed:         "ErrUpdateDataFailed",
}
