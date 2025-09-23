package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hellobchain/nacos-go/constant"
	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/httpcode"
	"github.com/hellobchain/nacos-go/model"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func RegistryRoute(r *handle.LogRouter, registry *RegistryService) {
	logger.Info("Init RegistryRoute")
	// 注册
	r.HandleFunc(constant.REGISTER_SERVICE_ROUTER, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodDelete {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.Method == http.MethodPost {
			var ins model.Instance
			if err := json.NewDecoder(r.Body).Decode(&ins); err != nil {
				httpcode.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := registry.Register(ins); err != nil {
				httpcode.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			httpcode.Success(w, http.StatusOK, "success", nil)
			return
		}
		// 删除（Deregister）也可以用同一个 path，这里简化
		if r.Method == http.MethodDelete {
			serviceName := r.URL.Query().Get("serviceName")
			groupName := r.URL.Query().Get("groupName")
			ip := r.URL.Query().Get("ip")
			port := parseUint(r.URL.Query().Get("port"))
			if err := registry.Repo.Deregister(serviceName, groupName, ip, port); err != nil {
				httpcode.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			httpcode.Success(w, http.StatusOK, "success", nil)
			return
		}
	}).Methods(http.MethodPost, http.MethodDelete)
	// 发现
	r.HandleFunc(constant.LIST_SERVICE_ROUTER, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		serviceName := r.URL.Query().Get("serviceName")
		groupName := r.URL.Query().Get("groupName")
		list, err := registry.List(serviceName, groupName)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", list)
	}).Methods(http.MethodGet)

	// 心跳
	r.HandleFunc(constant.HEARTBEAT_SERVICE_ROUTER, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		serviceName := r.URL.Query().Get("serviceName")
		groupName := r.URL.Query().Get("groupName")
		ip := r.URL.Query().Get("ip")
		port := parseUint(r.URL.Query().Get("port"))
		if err := registry.Heartbeat(serviceName, groupName, ip, port); err != nil {
			httpcode.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", nil)
	}).Methods(http.MethodPut)
}

// 简易端口解析
func parseUint(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}
