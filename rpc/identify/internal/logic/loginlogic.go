package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ZChen470/eShop/model"
	"github.com/ZChen470/eShop/rpc/identify/identify"
	"github.com/ZChen470/eShop/rpc/identify/internal/svc"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func getToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat // 签发时间
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func (l *LoginLogic) Login(in *identify.LoginReq) (*identify.LoginResp, error) {
	// todo: add your logic here and delete this line
	auth := l.svcCtx.Config.Auth

	var user model.User
	// 根据 email 查询用户
	if err := l.svcCtx.DB.Where("email = ?", in.Email).Take(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("邮箱或密码错误")
		}
		return nil, err
	}
	// 对密码进行校验
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.Password)); err != nil {
		return nil, fmt.Errorf("邮箱或密码错误")
	}
	// 生成 Token
	userId := int64(user.UserId)
	now := time.Now()
	token, err := getToken(auth.AccessSecret, now.Unix(), auth.AccessExpire, userId)
	if err != nil {
		return nil, fmt.Errorf("生成token失败: %v", err)
	}
	return &identify.LoginResp{
		AccessToken: token,
		ExpireAt:    now.Add(time.Duration(auth.AccessExpire)).Format("2006-01-02 15:04:05"),
	}, nil
}
