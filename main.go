package main

import (
	"fmt"
	"go-casbin-http-authrozation/User"
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

	log.Print("监听端口 :8081")
	log.Fatal(http.ListenAndServe(
		":8081",
		User.SessionManager.Use(User.Authorizor(
			User.AuthEnforcer, // casbin 鉴权
			User.Authorized,   // 授权用户
		)(mux)),
	))

}

///

func loginHandler(users User.Items) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.PostFormValue("name")

		session := User.SessionManager.Load(r)

		user, err := users.FindByName(name)
		if err != nil {
			writeError(http.StatusBadRequest, "WRONG_CREDENTIALS", w, err)
			return
		}
		// 创建 token 值
		if err := session.RenewToken(w); err != nil {
			writeError(http.StatusInternalServerError, "内部错误", w, err)
			return
		}
		_ = session.PutInt(w, "id", user.ID)
		_ = session.PutString(w, "role", user.Role)
		writeSuccess("SUCCESS", w)
	})
}

func logoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := User.SessionManager.Load(r)
		if err := session.Destroy(w); err != nil {
			writeError(http.StatusInternalServerError, "内部错误", w, err)
			return
		}
		writeSuccess("SUCCESS", w)
	})
}

func currentMemberHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := User.SessionManager.Load(r)
		uid, err := session.GetInt("id")
		if err != nil {
			writeError(http.StatusInternalServerError, "内部错误", w, err)
			return
		}
		writeSuccess(fmt.Sprintf("当前用户ID: %d", uid), w)
	})
}

func memberRoleHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := User.SessionManager.Load(r)
		role, err := session.GetString("role")
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
