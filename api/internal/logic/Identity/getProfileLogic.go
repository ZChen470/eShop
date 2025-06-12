package Identity

import (
	"context"
	"strconv"

	"github.com/ZChen470/eShop/api/internal/svc"
	"github.com/ZChen470/eShop/api/internal/types"
	"github.com/ZChen470/eShop/rpc/identify/identify"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetProfileLogic) GetProfile() (resp *types.UserProfile, err error) {
	// todo: add your logic here and delete this line
	r, err := l.svcCtx.IdentifyRpcClient.GetProfile(
		l.ctx,
		&identify.GetProfileReq{
			UserId: l.ctx.Value("userId").(int64),
		},
	)
	if err != nil {
		return nil, err
	}
	return &types.UserProfile{
		UserId:    strconv.Itoa(int(r.UserId)),
		Email:     r.Email,
		Nickname:  r.Nickname,
		CreatedAt: "",
	}, nil
}
