package Catalog

import (
	"net/http"

	"github.com/ZChen470/eShop/api/internal/logic/Catalog"
	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func ListProductsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := Catalog.NewListProductsLogic(r.Context(), svcCtx)
		resp, err := l.ListProducts()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
