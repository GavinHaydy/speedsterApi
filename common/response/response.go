package response

import (
	"net/http"

	"speedsterApi/common/errno"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// 泛型版本（推荐）
type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data,omitempty"`
}

// ================== 核心方法 ==================

func display[T any](w http.ResponseWriter, code int, msg string, data T) {
	resp := Response[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	httpx.OkJson(w, resp)
}

// ================== 成功 ==================

func Success(w http.ResponseWriter) {
	display(w, errno.Ok, errno.CodeAlertMap[errno.Ok], struct{}{})
}

func SuccessWithMsg(w http.ResponseWriter, msg string) {
	display(w, errno.Ok, msg, struct{}{})
}

func SuccessWithData[T any](w http.ResponseWriter, data T) {
	display(w, errno.Ok, errno.CodeAlertMap[errno.Ok], data)
}

// ================== 错误 ==================

func Error(w http.ResponseWriter, err error) {
	display(w, errno.ErrServer, err.Error(), struct{}{})
}

func ErrorWithCode(w http.ResponseWriter, code int, msg string) {
	if m, ok := errno.CodeAlertMap[code]; ok {
		msg = m + " " + msg
	}
	display(w, code, msg, struct{}{})
}

func ErrorWithMsg(w http.ResponseWriter, code int) {
	display(w, code, errno.CodeAlertMap[code], struct{}{})
}

func ErrorWithCodeAndData[T any](w http.ResponseWriter, code int, data T) {
	display(w, code, errno.CodeAlertMap[code], data)
}

func ErrorWithMsgAndData[T any](w http.ResponseWriter, code int, msg string, data T) {
	if m, ok := errno.CodeAlertMap[code]; ok {
		msg = m + " " + msg
	}
	display(w, code, msg, data)
}
