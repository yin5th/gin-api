package e

var MsgFlag = map[int]string {
    SUCCESS: "ok",
    ERROR: "fail",
    INVALID_PARAMS: "请求参数错误",
    ERROR_AUTH_CHECK_TOKEN_FAIL: "Token鉴权失败",
    ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Tokend鉴权超时",
    ERROR_AUTH_GENERATE_TOKEN_FAIL: "Token生成失败",
    ERROR_AUTH: "Token错误",
}

func GetMsg(code int) string {
    msg, ok := MsgFlag[int]
    if ok {
        return msg
    }
    return MsgFlag[ERROR]
}
