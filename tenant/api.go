package tenant

import (
	"encoding/json"
	"net/http"

	"github.com/hellobchain/nacos-go/constant"
	"github.com/hellobchain/nacos-go/handle"
	"github.com/hellobchain/nacos-go/httpcode"
)

func TetantRoute(r *handle.LogRouter, svc *Service) {
	// 1. 获取全部 tenant 列表
	r.HandleFunc(constant.TENANT_ROUTER, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			list, err := svc.List(r.Context())
			if err != nil {
				httpcode.Error(w, err.Error(), http.StatusInternalServerError)
			} else {
				httpcode.Success(w, http.StatusOK, "success", list)
			}
			return
		}
		// 2. 创建 tenant
		if r.Method == http.MethodPost {
			var req struct {
				Tenant string `json:"tenant"`
			}
			_ = json.NewDecoder(r.Body).Decode(&req)
			if req.Tenant == "" {
				httpcode.Error(w, "tenant required", http.StatusBadRequest)
				return
			}
			if err := svc.Create(r.Context(), req.Tenant); err != nil {
				httpcode.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			httpcode.Success(w, http.StatusOK, "success", true)
			return
		}
		// 3. 删除 tenant
		if r.Method == http.MethodDelete {
			tenant := r.URL.Query().Get("tenant")
			if tenant == "" {
				httpcode.Error(w, "tenant required", http.StatusBadRequest)
				return
			}
			if err := svc.Delete(r.Context(), tenant); err != nil {
				httpcode.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			httpcode.Success(w, http.StatusOK, "success", true)
		}
	}).Methods(http.MethodDelete, http.MethodPost, http.MethodGet)
}
