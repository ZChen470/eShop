package Identity

import (
	"context"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/identify/identify"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateProfileLogic) UpdateProfile(req *types.UpdateProfileReq) (resp *types.CommonResp, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.IdentifyRpcClient.UpdateProfile(
		l.ctx,
		&identify.UpdateProfileReq{
			UserId:   l.ctx.Value("userId").(int64),
			Nickname: req.Nickname,
		},
	)
	if err != nil {
		return nil, err
	}
	return &types.CommonResp{
		Code: r.Code,
		Msg:  r.Msg,
	}, nil
}
