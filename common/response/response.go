package response

import (
	"net/http"

	"speedsterApi/common/errno"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// Response 泛型版本（推荐）
type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data T      `json:"data,omitempty"`
}

// ================== 核心方法 ==================

func display[T any](w http.ResponseWriter, r *http.Request, code int, data T) {
	var msg string
	lang := r.Form.Get("Accept-Language")
	if lang == "en" {
		if m, ok := errno.CodeMsgMap[code]; ok {
			msg = m
		}
	} else {
		if m, ok := errno.CodeAlertMap[code]; ok {
			msg = m
		}
	}

	resp := Response[T]{
		Code: code,
		Msg:  msg,
		Data: data,
	}
	httpx.OkJson(w, resp)
	//httpx.WriteJson(w, 20007, resp) //  100 < code < 999
}

// ================== 成功 ==================

func Success(w http.ResponseWriter, r *http.Request) {
	display(w, r, errno.Ok, struct{}{})
}

func SuccessWithData[T any](w http.ResponseWriter, r *http.Request, data T) {
	display(w, r, errno.Ok, data)
}

// ================== 错误 ==================

func Error(w http.ResponseWriter, r *http.Request, code int) {
	display(w, r, code, struct{}{})
}

func ErrorWithCode(w http.ResponseWriter, r *http.Request, code int) {
	display(w, r, code, struct{}{})
}

func ErrorWithCodeAndData[T any](w http.ResponseWriter, r *http.Request, code int, data T) {
	display(w, r, code, data)
}

func ErrorWithMsgAndData[T any](w http.ResponseWriter, r *http.Request, code int, data T) {

	display(w, r, code, data)
}
