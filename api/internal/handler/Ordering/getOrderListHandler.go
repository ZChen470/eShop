package Ordering

import (
	"net/http"

	"github.com/ZChen470/eShop/api/internal/logic/Ordering"
	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetOrderListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := Ordering.NewGetOrderListLogic(r.Context(), svcCtx)
		resp, err := l.GetOrderList()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
