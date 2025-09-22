package httpcode

import (
	"fmt"
	"net/http"
)

func Error(w http.ResponseWriter, message string, code int) {
	errMsg := fmt.Sprintf("{\"code\":%d,\"message\":\"%s\"}", code, message)
	http.Error(w, errMsg, http.StatusOK)
}

func Success(w http.ResponseWriter, data interface{}) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%v", data)
}
