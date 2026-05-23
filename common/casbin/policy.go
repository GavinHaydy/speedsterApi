package casbin

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

func LoadPolicy(conn sqlx.SqlConn) error {
	Enforcer.ClearPolicy()

	sql := `
	SELECT
		r.code,
		p.path,
		p.method
	FROM sys_role_permission rp
	JOIN role r
		ON rp.role_id = r.id
	JOIN sys_permission p
		ON rp.permission_id = p.id
	WHERE r.status = 1
	  AND p.status = 1
	`

	var list []Policy

	err := conn.QueryRowsCtx(
		context.Background(),
		&list,
		sql,
	)
	if err != nil {
		return err
	}

	for _, v := range list {
		_, _ = Enforcer.AddPolicy(
			v.RoleCode,
			v.Path,
			v.Method,
		)
	}

	return nil
}
