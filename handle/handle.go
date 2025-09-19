package handle

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

type LogRouter struct {
	*mux.Router
}

func NewLogRouter() *LogRouter {
	return &LogRouter{mux.NewRouter()}
}

// 包装 HandleFunc，注册即打印
func (logRouter *LogRouter) HandleFunc(pattern string, handler http.HandlerFunc) *mux.Route {
	logger.Debugf("[Router] register  %-25s -> %s", pattern, "handler")
	return logRouter.Router.HandleFunc(pattern, handler)
}

// 也可以包装 Handle，用途一样
func (logRouter *LogRouter) Handle(pattern string, handler http.Handler) *mux.Route {
	logger.Debugf("[Router] register  %-25s -> %T", pattern, handler)
	return logRouter.Router.Handle(pattern, handler)
}
