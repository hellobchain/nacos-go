package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/hellobchain/nacos-go/httpcode"
	"github.com/hellobchain/nacos-go/pkg/utils"
	"github.com/hellobchain/wswlog/wlogging"
)

// Logger 简单请求日志
var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/v1/auth/login" {
			next.ServeHTTP(w, r)
			return
		}
		bearer := r.Header.Get("Authorization")
		if bearer == "" || !strings.HasPrefix(bearer, "Bearer ") {
			httpcode.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		tokenStr := bearer[7:]
		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return []byte(utils.GetSecret()), nil
		})
		if err != nil || !token.Valid {
			httpcode.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().ID()
		start := time.Now()

		// 包装 ResponseWriter 以捕获状态码
		lw := &logResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}
		next.ServeHTTP(lw, r)

		latency := time.Since(start)
		logger.Infof("| %10d | %3d | %13v | %15s | %-7s %s",
			requestID,
			lw.statusCode,
			latency,
			getClientIP(r),
			r.Method,
			r.URL.Path,
		)
	})
}

// 捕获状态码
type logResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (w *logResponseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

// 获取客户端 IP（简单取 X-Real-IP / X-Forwarded-For / RemoteAddr）
func getClientIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	return r.RemoteAddr
}

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// 简单 * 放行
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, HEAD, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// 预检请求直接返回
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
