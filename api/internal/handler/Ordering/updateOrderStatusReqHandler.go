package Ordering

import (
	"net/http"

	"github.com/ZChen470/eShop/api/internal/logic/Ordering"
	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UpdateOrderStatusReqHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateOrderStatusReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := Ordering.NewUpdateOrderStatusReqLogic(r.Context(), svcCtx)
		resp, err := l.UpdateOrderStatusReq(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
