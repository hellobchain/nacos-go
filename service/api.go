package service

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/model"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func RegistryRoute(r *handle.LogRouter, registry *RegistryService) {
	logger.Info("Init RegistryRoute")
	// 注册
	r.HandleFunc("/nacos/v1/ns/instance", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		if r.Method == http.MethodPost {
			var ins model.Instance
			if err := json.NewDecoder(r.Body).Decode(&ins); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			if err := registry.Register(ins); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("ok"))
			return
		}
		// 删除（Deregister）也可以用同一个 path，这里简化
		if r.Method == http.MethodDelete {
			serviceName := r.URL.Query().Get("serviceName")
			groupName := r.URL.Query().Get("groupName")
			ip := r.URL.Query().Get("ip")
			port := parseUint(r.URL.Query().Get("port"))
			if err := registry.Repo.Deregister(serviceName, groupName, ip, port); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			w.Write([]byte("ok"))
		}
	}).Methods(http.MethodPost, http.MethodDelete)
	// 发现
	r.HandleFunc("/nacos/v1/ns/instance/list", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		serviceName := r.URL.Query().Get("serviceName")
		groupName := r.URL.Query().Get("groupName")
		list, err := registry.List(serviceName, groupName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		_ = json.NewEncoder(w).Encode(list)
	}).Methods(http.MethodGet)

	// 心跳
	r.HandleFunc("/nacos/v1/ns/instance/beat", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPut {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		serviceName := r.URL.Query().Get("serviceName")
		groupName := r.URL.Query().Get("groupName")
		ip := r.URL.Query().Get("ip")
		port := parseUint(r.URL.Query().Get("port"))
		if err := registry.Heartbeat(serviceName, groupName, ip, port); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("ok"))
	}).Methods(http.MethodPut)
}

// 简易端口解析
func parseUint(s string) uint64 {
	v, _ := strconv.ParseUint(s, 10, 64)
	return v
}
