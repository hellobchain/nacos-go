package httpcode

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func Error(w http.ResponseWriter, message string, code int) {
	errMsg := fmt.Sprintf("{\"code\":%d,\"message\":\"%s\"}", code, message)
	logger.Errorf("[http result] code: %d: msg: %s", code, message)
	http.Error(w, errMsg, http.StatusOK)
}

func Success(w http.ResponseWriter, code int, msg string, data interface{}) {
	w.WriteHeader(http.StatusOK)
	resData := map[string]interface{}{
		"code":    code,
		"message": msg,
		"data":    data,
	}
	resDataStr, err := json.Marshal(resData)
	if err != nil {
		logger.Errorf("[http result] code: %d: msg: %s", code, msg)
	}
	w.Write([]byte(resDataStr))
}
