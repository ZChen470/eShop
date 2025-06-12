package logic

import (
	"context"
	"fmt"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/identify/identify"
	"github.com/ZChen470/eShop/rpc/identify/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProfileLogic {
	return &GetProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProfileLogic) GetProfile(in *identify.GetProfileReq) (*identify.UserProfile, error) {
	// todo: add your logic here and delete this line
	var user model.User
	if err := l.svcCtx.DB.Where("user_id = ?", in.UserId).First(&user).Error; err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}
	return &identify.UserProfile{
		UserId:   int64(user.UserId),
		Email:    user.Email,
		Nickname: user.Nickname,
	}, nil
}
