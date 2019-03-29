package User

import (
	"github.com/alexedwards/scs"
	"time"
)

var (
	// 创建会话存储
	SessionManager = scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")
)

func init() {
	SessionManager.Lifetime(time.Hour) // Set the maximum session lifetime to 1 hour.
	SessionManager.Persist(true)       // Persist the session after a user has closed their browser.
	SessionManager.Secure(true)        // Set the Secure flag on the session cookie.
}
