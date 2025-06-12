package logic

import (
	"context"
	"fmt"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/identify/identify"
	"github.com/ZChen470/eShop/rpc/identify/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProfileLogic) UpdateProfile(in *identify.UpdateProfileReq) (*identify.CommonResp, error) {
	// todo: add your logic here and delete this line
	if err := l.svcCtx.DB.Model(&model.User{}).Where("user_id = ?", in.UserId).
		Updates(map[string]interface{}{
			"nickname": in.Nickname,
		}).Error; err != nil {
		return nil, fmt.Errorf("更新用户信息失败: %v", err)
	}
	return &identify.CommonResp{
		Msg:  "更新用户信息成功",
		Code: 10201,
	}, nil
}
