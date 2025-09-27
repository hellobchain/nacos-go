package user

import (
	"encoding/json"
	"net/http"

	"github.com/hellobchain/nacos-go/constant"
	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/httpcode"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func AuthRoute(r *handle.LogRouter, as *AuthUserService) http.Handler {
	logger.Info("Init AuthRoute")
	r.HandleFunc(constant.AUTH_LOGIN, func(w http.ResponseWriter, r *http.Request) {
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
		httpcode.Success(w, http.StatusOK, "success", map[string]string{"accessToken": token, "uuid": uuid})
	}).Methods(http.MethodPost)

	r.HandleFunc(constant.USER_INFO, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		userName := r.Header.Get(constant.SRC_USER)
		r.Header.Del(constant.SRC_USER)
		user, err := as.GetUserInfo(r.Context(), userName)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", user)
	}).Methods(http.MethodGet)
	r.HandleFunc(constant.USER_INFO, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username    string `json:"username"`
			OldPassword string `json:"oldPassword"`
			NewPassword string `json:"newPassword"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := as.ChangePassword(r.Context(), req.Username, req.OldPassword, req.NewPassword); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", "true")
	}).Methods(http.MethodPost)
	return r
}
