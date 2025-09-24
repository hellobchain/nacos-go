package config

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/hellobchain/nacos-go/constant"
	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/httpcode"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func ConfigRoute(r *handle.LogRouter, svc *Service) {
	logger.Info("init config route")
	// 发布 / 更新配置
	r.HandleFunc(constant.CONFIGS_ROUTER, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var item ConfigItem
		if err := json.NewDecoder(req.Body).Decode(&item); err != nil {
			httpcode.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		item.Tenant = uuid.New().String()
		item.SrcIp = req.Header.Get(constant.SRC_IP)
		req.Header.Del(constant.SRC_IP)
		item.SrcUser = req.Header.Get(constant.SRC_USER)
		req.Header.Del(constant.SRC_USER)
		if err := svc.Publish(req.Context(), item); err != nil {
			httpcode.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", "true")
	}).Methods(http.MethodPost)

	// 获取配置
	r.HandleFunc(constant.CONFIGS_ROUTER, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		dataId := req.URL.Query().Get("dataId")
		group := req.URL.Query().Get("group")
		tenant := req.URL.Query().Get("tenantId")
		item, err := svc.Get(req.Context(), dataId, group, tenant)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", item)
	}).Methods(http.MethodGet)

	// 获取配置
	r.HandleFunc(constant.LIST_CONFIGS, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		dataId := req.URL.Query().Get("dataId")
		group := req.URL.Query().Get("group")
		tenant := req.URL.Query().Get("tenantId")
		item, err := svc.List(req.Context(), dataId, group, tenant)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", item)
	}).Methods(http.MethodGet)

	// 修改配置
	r.HandleFunc(constant.LIST_CONFIGS, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPut {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		dataId := req.URL.Query().Get("dataId")
		group := req.URL.Query().Get("group")
		tenant := req.URL.Query().Get("tenantId")
		item, err := svc.List(req.Context(), dataId, group, tenant)
		if err != nil {
			httpcode.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", item)
	}).Methods(http.MethodPut)

	// 删除配置
	r.HandleFunc(constant.CONFIGS_ROUTER, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		dataId := req.URL.Query().Get("dataId")
		group := req.URL.Query().Get("group")
		tenant := req.URL.Query().Get("tenantId")
		if err := svc.Delete(req.Context(), dataId, group, tenant); err != nil {
			httpcode.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		httpcode.Success(w, http.StatusOK, "success", "true")
	}).Methods(http.MethodDelete)

	// 监听配置（简易版：30s 长轮询）
	r.HandleFunc(constant.LISTEN_CONFIGS, func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			httpcode.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		dataId := req.URL.Query().Get("dataId")
		group := req.URL.Query().Get("group")
		tenant := req.URL.Query().Get("tenantId")
		md5Client := req.URL.Query().Get("md5")

		// 30s 长轮询
		ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
		defer cancel()

		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				// 超时返回空
				w.WriteHeader(http.StatusNotModified)
				return
			case <-ticker.C:
				item, err := svc.Get(ctx, dataId, group, tenant)
				if err == nil && item.Md5 != md5Client {
					httpcode.Success(w, http.StatusOK, "success", map[string]string{"content": item.Content, "md5": item.Md5})
					return
				}
			}
		}
	}).Methods(http.MethodPost)
}
