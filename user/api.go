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
	// 登录
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
		if req.Username == "" || req.Password == "" {
			httpcode.Error(w, "username or password required", http.StatusBadRequest)
			return
		}
		token, uuid, err := as.Login(r.Context(), req.Username, req.Password)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", map[string]string{"accessToken": token, "uuid": uuid})
	}).Methods(http.MethodPost)
	// 获取用户信息
	r.HandleFunc(constant.USER_INFO, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		username := r.Header.Get(constant.SRC_USER)
		r.Header.Del(constant.SRC_USER)
		user, err := as.GetUserInfo(r.Context(), username)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", user)
	}).Methods(http.MethodGet)
	// 修改密码
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
		if req.Username == "" || req.OldPassword == "" || req.NewPassword == "" {
			httpcode.Error(w, "username or password required", http.StatusBadRequest)
			return
		}
		if err := as.ChangePassword(r.Context(), req.Username, req.OldPassword, req.NewPassword); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", "true")
	}).Methods(http.MethodPost)
	// 注册
	r.HandleFunc(constant.REGISTER_USER, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Username == "" || req.Password == "" || req.Role == "" {
			httpcode.Error(w, "username or password or role required", http.StatusBadRequest)
			return
		}
		if req.Role == constant.ROLE_ADMIN {
			httpcode.Error(w, "role must not be admin", http.StatusBadRequest)
			return
		}
		u, _ := as.GetUserInfo(r.Context(), req.Username)
		if u != nil {
			httpcode.Error(w, "user already exists", http.StatusBadRequest)
			return
		}
		if err := as.Register(r.Context(), req.Username, req.Password, req.Role); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", "true")
	}).Methods(http.MethodPost)
	// 修改
	r.HandleFunc(constant.UPDATE_USER, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var req struct {
			Username string `json:"username"`
			Password string `json:"password"`
			Role     string `json:"role"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if req.Username == "" || req.Password == "" {
			httpcode.Error(w, "username or password required", http.StatusBadRequest)
			return
		}
		u, _ := as.GetUserInfo(r.Context(), req.Username)
		if u == nil {
			httpcode.Error(w, "user not exist", http.StatusBadRequest)
			return
		}
		if u.Role == constant.ROLE_ADMIN {
			httpcode.Error(w, "permission denied", http.StatusBadRequest)
			return
		}
		if err := as.Update(r.Context(), req.Username, req.Password, req.Role); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", true)
	}).Methods(http.MethodPost)
	// 删除
	r.HandleFunc(constant.DELETE_USER, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		username := r.URL.Query().Get("username")
		if username == "" {
			httpcode.Error(w, "username is empty", http.StatusBadRequest)
			return
		}
		u, _ := as.GetUserInfo(r.Context(), username)
		if u == nil {
			httpcode.Error(w, "user not exist", http.StatusBadRequest)
			return
		}
		if u.Role == constant.ROLE_ADMIN {
			httpcode.Error(w, "permission denied", http.StatusBadRequest)
			return
		}
		if err := as.Delete(r.Context(), username); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", "true")
	}).Methods(http.MethodDelete)
	// 获取用户列表
	r.HandleFunc(constant.GET_USER_LIST, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		item, err := as.List(req.Context())
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", item)
	}).Methods(http.MethodGet)

	return r
}
