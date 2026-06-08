package model

import (
	"context"
	"fmt"
	"speedsterApi/app/user/user/pb/pb"
	"speedsterApi/common/errno"
	"speedsterApi/common/errorx"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUserModel = (*customSysUserModel)(nil)

type (
	// SysUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUserModel.
	SysUserModel interface {
		sysUserModel
		withSession(session sqlx.Session) SysUserModel
		FindByAccountAndPW(ctx context.Context, account string, password string) (*SysUser, error)
		SelectUserList(ctx context.Context, req *pb.UserListReq) (total int64, userList []*pb.UserItem, err error)
	}

	customSysUserModel struct {
		*defaultSysUserModel
	}
)

// NewSysUserModel returns a model for the database table.
func NewSysUserModel(conn sqlx.SqlConn) SysUserModel {
	return &customSysUserModel{
		defaultSysUserModel: newSysUserModel(conn),
	}
}

func (m *customSysUserModel) withSession(session sqlx.Session) SysUserModel {
	return NewSysUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customSysUserModel) FindByAccountAndPW(ctx context.Context, username string, password string) (*SysUser, error) {
	var user SysUser
	// 1. 查询所有字段
	query := fmt.Sprintf("SELECT * FROM %s WHERE \"username\" = $1 AND \"password\" = $2 LIMIT 1", m.table)
	// 2. 将查询结果扫描到 user 结构体中
	// 注意：这里假设 QueryRowCtx 支持直接扫描到结构体，如果不支持，需要手动传入 &user.Id, &user.Account
	err := m.conn.QueryRowCtx(ctx, &user, query, username, password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *customSysUserModel) SelectUserList(ctx context.Context, req *pb.UserListReq) (total int64, userList []*pb.UserItem, err error) {
	//var userList []*SysUser
	var temp []*SysUser

	// 使用 squirrel 构建基础的 SELECT 语句
	//builder := squirrel.Select(sysUserRows).From(m.table)
	builder := squirrel.Select().From(m.table).PlaceholderFormat(squirrel.Dollar)

	// 动态添加 WHERE 条件
	if req.Username != "" {
		//builder = builder.Where("username LIKE ?", "%"+req.Usarname+"%")
		builder = builder.Where(squirrel.Like{"username": fmt.Sprintf("%%%s%%", req.Username)})
	}
	if req.Status != 0 {
		//builder = builder.Where("status = ?", req.Status)
		builder = builder.Where(squirrel.Eq{"status": req.Status})
	}

	if req.Phone != "" {
		//builder = builder.Where("phone LIKE ?", "%"+req.Phone+"%")
		builder = builder.Where(squirrel.Like{"phone": fmt.Sprintf("%%%s%%", req.Phone)})
	}

	if req.Email != "" {
		//builder = builder.Where("email LIKE ?", "%"+req.Email+"%")
		builder = builder.Where(squirrel.Like{"email": fmt.Sprintf("%%%s%%", req.Email)})
	}

	if req.Nickname != "" {
		//builder = builder.Where("nickname LIKE ?", "%"+req.Nickname+"%")
		builder = builder.Where(squirrel.Like{"nickname": fmt.Sprintf("%%%s%%", req.Nickname)})
	}

	countBuilder := builder.Columns("COUNT(*)")
	countQuery, countValues, err := countBuilder.ToSql()
	if err != nil {
		return 0, nil, errorx.New(errno.ErrPgsqlFailed)
	}
	logx.Infof("========count sql:%s======%s===", countQuery, countValues)

	err = m.conn.QueryRowCtx(ctx, &total, countQuery, countValues...)
	if err != nil {
		logx.Errorf("err:%v", err)
		return 0, nil, errorx.New(errno.ErrPgsqlFailed)
	}

	// 如果总数为0，直接返回，不用查列表了
	if total == 0 {
		return 0, []*pb.UserItem{}, nil
	}

	// 添加分页和排序
	offset := (req.PageNo - 1) * req.PageSize
	query, values, err := builder.
		Columns(sysUserRows).
		Limit(uint64(req.PageSize)).
		Offset(uint64(offset)).
		ToSql()
	if err != nil {
		logx.Errorf("err:%v", err)
		return 0, nil, errorx.New(errno.ErrPgsqlFailed)
	}

	// 将 squirrel 生成的 SQL 和参数交给 go-zero 的 sqlx 执行
	logx.Infof("========sql:%s=========", query)
	err = m.conn.QueryRows(&temp, query, values...)
	if err != nil {
		logx.Errorf("------------err:%v", err)
		return 0, nil, errorx.New(errno.ErrPgsqlFailed)
	}

	rspList := make([]*pb.UserItem, 0, len(userList))
	logx.Infof("userListRsp:%v", userList)
	for _, user := range temp {
		// 将 SysUser 转换为 UserListRsp
		var tempPhone, tempNickname, tempEmail, tempAvatar string
		if user.Phone.Valid {
			tempPhone = user.Phone.String
		}
		if user.Avatar.Valid {
			tempAvatar = user.Avatar.String
		}
		if user.Nickname.Valid {
			tempNickname = user.Nickname.String
		}
		if user.Email.Valid {
			tempEmail = user.Email.String
		}

		rspList = append(rspList, &pb.UserItem{
			Username: user.Username,
			Nickname: &tempNickname,
			Phone:    &tempPhone,
			Email:    &tempEmail,
			Status:   &user.Status,
			Avatar:   &tempAvatar,
			//Created_at: user.Created_at,
		})
	}

	return total, rspList, nil

}
