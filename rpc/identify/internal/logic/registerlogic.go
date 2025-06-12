package logic

import (
	"context"
	"net/mail"
	"regexp"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/identify/identify"
	"github.com/ZChen470/eShop/rpc/identify/internal/svc"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func isValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	if err != nil {
		return false
	}
	regex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(email)
}

func isValidPassword(passowrd string) bool {
	re := regexp.MustCompile(`^[A-Za-z0-9]{6,20}$`)
	return re.MatchString(passowrd)
}

func (l *RegisterLogic) Register(in *identify.RegisterReq) (*identify.CommonResp, error) {
	// todo: add your logic here and delete this line
	if !isValidEmail(in.Email) {
		return &identify.CommonResp{
			Msg:  "邮箱格式错误",
			Code: 10212,
		}, nil
	}
	if len(in.Nickname) < 2 || len(in.Nickname) > 20 {
		return &identify.CommonResp{
			Msg:  "昵称长度必须在2到20个字符之间",
			Code: 10212,
		}, nil
	}
	if !isValidPassword(in.Password) {
		return &identify.CommonResp{
			Msg:  "密码长度必须在6到20个字符之间，只能由数字和字母组成",
			Code: 10212,
		}, nil
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return &identify.CommonResp{
			Msg:  "密码加密失败",
			Code: 10212,
		}, nil
	}
	// 开启事务
	err = l.svcCtx.DB.Transaction(func(tx *gorm.DB) error {
		user := &model.User{
			Email:    in.Email,
			Nickname: in.Nickname,
			Password: string(hashPassword),
		}
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return &identify.CommonResp{
			Msg:  "邮箱或昵称已存在",
			Code: 10212,
		}, nil
	}
	return &identify.CommonResp{
		Msg:  "注册成功",
		Code: 10202,
	}, nil
}
