package User

import (
	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
)

var (
	AuthEnforcer *casbin.Enforcer
)

const (
	CasbinModel = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && (r.act == p.act || p.act == "*")
`
)

func init() {
	// 加载 casbin 鉴权规则
	if enforcer, err := casbin.NewEnforcerSafe(
		casbin.NewModel(CasbinModel),
		gormadapter.NewAdapter(
			"mysql",
			"root:123321000@tcp(127.0.0.1:3306)/",
		),
	); err != nil {
		panic(err)
	} else {
		AuthEnforcer = enforcer
	}

	if err := AuthEnforcer.LoadPolicy(); err != nil {
		panic(err)
	}

	insertPolicy()
}

func insertPolicy() {
	AuthEnforcer.AddPolicy("admin", "/*", "*")
	AuthEnforcer.AddPolicy("anonymous", "/login", "*")
	AuthEnforcer.AddPolicy("member", "/logout", "*")
	AuthEnforcer.AddPolicy("member", "/member/*", "*")

	if err := AuthEnforcer.SavePolicy(); err != nil {
		panic(err)
	}
}
