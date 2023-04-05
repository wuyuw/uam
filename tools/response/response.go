package response

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type RespOpt struct {
	Code int
	Msg  string
}

var (
	RespSuccess   = RespOpt{Code: 0, Msg: "操作成功"}
	RespFailed    = RespOpt{Code: 1, Msg: "操作失败"}
	RespArgsErr   = RespOpt{Code: 10400, Msg: "参数错误"}
	RespNologin   = RespOpt{Code: 10401, Msg: "用户未登录"}
	RespForbidden = RespOpt{Code: 10403, Msg: "无访问权限"}
	RespSvcErr    = RespOpt{Code: 10500, Msg: "服务异常"}

	RespInvalidCli = RespOpt{Code: 10401, Msg: "非法客户端"}
	RespAuthFail   = RespOpt{Code: 10401, Msg: "客户端认证失败"}
)

type Body struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, code int, msg string, data interface{}) {
	httpx.OkJson(w, Body{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(w http.ResponseWriter) {
	Response(w, RespSuccess.Code, RespSuccess.Msg, map[string]interface{}{})
}

func OkWithMsg(w http.ResponseWriter, msg string) {
	Response(w, RespSuccess.Code, msg, map[string]interface{}{})
}

func OkWithData(w http.ResponseWriter, data interface{}) {
	Response(w, RespSuccess.Code, RespSuccess.Msg, data)
}

func Fail(w http.ResponseWriter) {
	Response(w, RespFailed.Code, RespFailed.Msg, map[string]interface{}{})
}

func FailWithMsg(w http.ResponseWriter, msgs ...string) {
	var msg string
	if len(msgs) == 0 {
		msg = RespFailed.Msg
	} else {
		msg = fmt.Sprintf("%s: %s", RespFailed.Msg, strings.Join(msgs, ";"))
	}
	Response(w, RespFailed.Code, msg, map[string]interface{}{})
}

// FailByArgsErr 参数错误
func FailByArgsErr(w http.ResponseWriter, msgs ...string) {
	var msg string
	if len(msgs) == 0 {
		msg = RespArgsErr.Msg
	} else {
		msg = fmt.Sprintf("%s: %s", RespArgsErr.Msg, strings.Join(msgs, ";"))
	}
	Response(w, RespArgsErr.Code, msg, map[string]interface{}{})
}

// FailBySvcErr 服务异常
func FailBySvcErr(w http.ResponseWriter, msgs ...string) {
	var msg string
	if len(msgs) == 0 {
		msg = RespSvcErr.Msg
	} else {
		msg = fmt.Sprintf("%s: %s", RespSvcErr.Msg, strings.Join(msgs, ";"))
	}
	Response(w, RespSvcErr.Code, msg, map[string]interface{}{})
}

// FailByNoLogin 用户未登录
func FailByNoLogin(w http.ResponseWriter, data interface{}) {
	Response(w, RespNologin.Code, RespNologin.Msg, data)
}

// FailByForbidden 无访问权限
func FailByForbidden(w http.ResponseWriter, msgs ...string) {
	var msg string
	if len(msgs) == 0 {
		msg = RespForbidden.Msg
	} else {
		msg = fmt.Sprintf("%s: %s", RespForbidden.Msg, strings.Join(msgs, ";"))
	}
	Response(w, RespForbidden.Code, msg, map[string]interface{}{})
}

// FailByInvalidClient 非法的客户端
func FailByInvalidClient(w http.ResponseWriter, msgs ...string) {
	var msg string
	if len(msgs) == 0 {
		msg = RespInvalidCli.Msg
	} else {
		msg = fmt.Sprintf("%s: %s", RespInvalidCli.Msg, strings.Join(msgs, ";"))
	}
	Response(w, RespInvalidCli.Code, msg, map[string]interface{}{})
}

// FailByAuthFail 客户端认证失败
func FailByAuthFail(w http.ResponseWriter, msgs ...string) {
	var msg string
	if len(msgs) == 0 {
		msg = RespAuthFail.Msg
	} else {
		msg = fmt.Sprintf("%s: %s", RespAuthFail.Msg, strings.Join(msgs, ";"))
	}
	Response(w, RespAuthFail.Code, msg, map[string]interface{}{})
}
