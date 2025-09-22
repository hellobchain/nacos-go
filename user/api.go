package user

import (
	"encoding/json"
	"net/http"

	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/httpcode"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func AuthRoute(r *handle.LogRouter, as *AuthUserService) http.Handler {
	logger.Info("Init AuthRoute")
	r.HandleFunc("/v1/auth/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		token, uuid, err := as.Login(r.Context(), req.Username, req.Password)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"accessToken": "Bearer " + token, "uuid": uuid})
	}).Methods(http.MethodPost)
	return r
}
