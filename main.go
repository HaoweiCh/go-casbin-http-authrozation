package main

import (
	"fmt"
	"github.com/alexedwards/scs/session"
	"haowei.ch/casbin-http-role-example/User"
	"log"
	"net/http"
)

func main() {
	// 定义路径
	mux := http.NewServeMux()
	mux.HandleFunc("/login", loginHandler(User.Authorized))
	mux.HandleFunc("/logout", logoutHandler())
	mux.HandleFunc("/member/current", currentMemberHandler())
	mux.HandleFunc("/member/role", memberRoleHandler())
	mux.HandleFunc("/admin/stuff", adminHandler())

	log.Print("监听端口 :8080")
	log.Fatal(http.ListenAndServe(
		":8080",
		User.SessionManager(
			User.Authorizor(
				User.AuthEnforcer, // casbin 鉴权
				User.Authorized,   // 授权用户
			)(mux),
		),
	))

}

///

func loginHandler(users User.Items) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")
		user, err := users.FindByName(name)
		if err != nil {
			writeError(http.StatusBadRequest, "WRONG_CREDENTIALS", w, err)
			return
		}
		// 创建 token 值
		if err := session.RegenerateToken(r); err != nil {
			writeError(http.StatusInternalServerError, "内部错误", w, err)
			return
		}
		_ = session.PutInt(r, "id", user.ID)
		_ = session.PutString(r, "role", user.Role)
		writeSuccess("SUCCESS", w)
	})
}

func logoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := session.Renew(r); err != nil {
			writeError(http.StatusInternalServerError, "内部错误", w, err)
			return
		}
		writeSuccess("SUCCESS", w)
	})
}

func currentMemberHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid, err := session.GetInt(r, "id")
		if err != nil {
			writeError(http.StatusInternalServerError, "内部错误", w, err)
			return
		}
		writeSuccess(fmt.Sprintf("当前用户ID: %d", uid), w)
	})
}

func memberRoleHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, err := session.GetString(r, "role")
		if err != nil {
			writeError(http.StatusInternalServerError, "内部错误", w, err)
			return
		}
		writeSuccess(fmt.Sprintf("当前用户角色: %s", role), w)
	})
}

func adminHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeSuccess("你是管理员!", w)
	})
}

///

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("错误: ", err.Error())
	w.WriteHeader(status)
	_, _ = w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(message))
}
