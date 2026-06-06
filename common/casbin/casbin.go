package casbin

import "github.com/casbin/casbin/v2"

var Enforcer *casbin.Enforcer

func Init() error {
	var err error

	Enforcer, err = casbin.NewEnforcer(
		"../../../common/casbin/model.conf",
	)

	return err
}
