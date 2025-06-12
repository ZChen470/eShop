package Basket

import (
	"net/http"

	"github.com/ZChen470/eShop/api/internal/logic/Basket"
	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateItemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := Basket.NewUpdateItemLogic(r.Context(), svcCtx)
		resp, err := l.UpdateItem(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
