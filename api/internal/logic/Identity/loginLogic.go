package Identity

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/identify/identify"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.IdentifyRpcClient.Login(
		l.ctx,
		&identify.LoginReq{
			Email:    req.Email,
			Password: req.Password,
		},
	)
	if err != nil {
		return nil, err
	}
	return &types.LoginResp{
		AccessToken: r.AccessToken,
		ExpireAt:    r.ExpireAt,
	}, nil
}
