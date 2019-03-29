package User

import (
	"github.com/casbin/casbin"
)

var (
	AuthEnforcer *casbin.Enforcer
)

func init() {
	// 加载 casbin 鉴权规则
	if enforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv"); err != nil {
		panic(err)
	} else {
		AuthEnforcer = enforcer
	}
}
