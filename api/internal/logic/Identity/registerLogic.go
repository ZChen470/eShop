package Identity

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/identify/identify"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.IdentifyRpcClient.Register(l.ctx, &identify.RegisterReq{
		Email:    req.Email,
		Nickname: req.Nickname,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	return &types.CommonResp{
		Code: r.Code,
		Msg:  r.Msg,
	}, nil
}
