package Basket

import (
	"net/http"

	"github.com/ZChen470/eShop/api/internal/logic/Basket"
	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func GetCartHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := Basket.NewGetCartLogic(r.Context(), svcCtx)
		resp, err := l.GetCart()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
