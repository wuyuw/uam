package logic

import (
	"context"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"math/rand"
	"time"
	"uam/services/model"
	"uam/services/model/auth"
	"uam/services/model/user"
	"uam/tools/constants"

	"uam/services/rpc/internal/svc"
	"uam/services/rpc/pb/uamrpc"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddAuthLocalLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddAuthLocalLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddAuthLocalLogic {
	return &AddAuthLocalLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AddAuthLocal 添加本地认证记录
func (l *AddAuthLocalLogic) AddAuthLocal(in *uamrpc.AddAuthLocalReq) (resp *uamrpc.AddAuthLocalResp, err error) {
	defer func() {
		if err != nil {
			l.Logger.Error(err.Error())
		}
	}()
	authLocal, err := l.svcCtx.AuthLocalModel.FindOneByUsername(l.ctx, in.Username)
	if err != nil && err != model.ErrNotFound {
		return nil, errors.Wrap(err, constants.MsgDBErr)
	}
	if authLocal != nil {
		return nil, errors.New("用户名已存在")
	}
	//uid := idgen.NextId()
	var uid int
	for i := 0; i < 10; i++ {
		uid = l.randUid(10, 100000)
		_, err := l.svcCtx.UserModel.FindOneByUid(l.ctx, int64(uid))
		if err != nil {
			if err != model.ErrNotFound {
				return nil, errors.Wrap(err, constants.MsgDBErr)
			}
			if err = l.AddAuthLocalInTx(int64(uid), in.Username, in.Password, in.Salt); err != nil {
				return nil, err
			}
			return &uamrpc.AddAuthLocalResp{}, nil
		}
	}
	return nil, errors.New("操作失败，请重试")
}

func (l *AddAuthLocalLogic) AddAuthLocalInTx(uid int64, username, password, salt string) error {
	return model.ExecInTx(l.svcCtx.DB, func(tx *gorm.DB) error {
		if err := l.svcCtx.UserModel.InsertOne(l.ctx, tx, &user.User{
			Uid:      uid,
			Nickname: username,
		}); err != nil {
			return err
		}
		if err := l.svcCtx.AuthLocalModel.InsertOne(l.ctx, tx, &auth.AuthLocal{
			Uid:      uid,
			Username: username,
			Password: password,
			Salt:     salt,
		}); err != nil {
			return err
		}
		return nil
	})
}

func (l *AddAuthLocalLogic) randUid(minNum, maxNum int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(maxNum-minNum) + minNum + maxNum
}
