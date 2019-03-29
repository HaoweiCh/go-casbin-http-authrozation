package User

import (
	"github.com/alexedwards/scs/engine/memstore"
	"github.com/alexedwards/scs/session"
	"time"
)

var (
	// 创建会话存储
	SessionManager = session.Manage(
		memstore.New(30*time.Minute),        // 30 分钟自动清理失效数据
		session.IdleTimeout(30*time.Minute), // 空闲失效时间，超过指定时间无操作 token 失效
		session.Persist(true),               // 持久化
		session.Secure(true),                // 加密
	)
)
