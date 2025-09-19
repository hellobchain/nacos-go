package config

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/wswlog/wlogging"
)

var logger = wlogging.MustGetFileLoggerWithoutName(nil)

func ConfigRoute(r *handle.LogRouter, svc *Service) {
	logger.Debug("init config route")
	// 发布 / 更新配置
	r.HandleFunc("/v1/cs/configs", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var item ConfigItem
		if err := json.NewDecoder(req.Body).Decode(&item); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := svc.Publish(req.Context(), item); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("true"))
	}).Methods("POST")

	// 获取配置
	r.HandleFunc("/v1/cs/configs", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		dataId := req.URL.Query().Get("dataId")
		group := req.URL.Query().Get("group")
		tenant := req.URL.Query().Get("tenantId")
		item, err := svc.Get(req.Context(), dataId, group, tenant)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		_ = json.NewEncoder(w).Encode(item)
	}).Methods("GET")

	// 删除配置
	r.HandleFunc("/v1/cs/configs", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodDelete {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		dataId := req.URL.Query().Get("dataId")
		group := req.URL.Query().Get("group")
		tenant := req.URL.Query().Get("tenantId")
		if err := svc.Delete(req.Context(), dataId, group, tenant); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write([]byte("true"))
	}).Methods("DELETE")

	// 监听配置（简易版：30s 长轮询）
	r.HandleFunc("/v1/cs/configs/listener", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
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
					_ = json.NewEncoder(w).Encode(map[string]string{"content": item.Content, "md5": item.Md5})
					return
				}
			}
		}
	}).Methods("POST")
}
