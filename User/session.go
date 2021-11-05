package User

import (
	"time"

	"github.com/alexedwards/scs"
	"github.com/alexedwards/scs/stores/redisstore"
	"github.com/gomodule/redigo/redis"
)

var (
	// 创建会话存储
	SessionManager *scs.Manager
)

func init() {
	SessionManager = scs.NewManager(redisstore.New(&redis.Pool{
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				return nil, err
			}
			// 输入密码
			//if _, err := c.Do("AUTH", password); err != nil {
			//	c.Close()
			//	return nil, err
			//}
			if _, err := c.Do("SELECT", "11"); err != nil {
				c.Close()
				return nil, err
			}
			return c, nil
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}))

	SessionManager.Lifetime(time.Hour) // Set the maximum session lifetime to 1 hour.
	SessionManager.Persist(true)       // Persist the session after a user has closed their browser.
	SessionManager.Secure(true)        // Set the Secure flag on the session cookie.
}
