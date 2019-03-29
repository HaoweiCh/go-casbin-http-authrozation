# casbin-http-role-exampe

简单且实用的 HTTP 鉴权体系

* [casbin](https://github.com/casbin/casbin) for role-based HTTP Authorization
* [scs](https://github.com/alexedwards/scs)  for session handling.


Run with

```bash
dep ensure
go run main.go
```

Which starts a server at `http://localhost:8080` with the following routes:

* `POST /login` - accessible if not logged in
   * takes `name` as a form-data parameter - there is no password
   * Valid Users: 
     * `Admin` ID: `1`, Role: `admin`
     * `Sabine` ID: `2`, Role: `member`
     * `Sepp` ID: `3`, Role: `member`
* `POST /logout` - accessible if logged in
* `GET /member/current` - accessible if logged in as a member
* `GET /member/role` - accessible if logged in as a member
* `GET /admin/stuff` - accessible if logged in as an admin

# 项目思路及来源
* https://studygolang.com/articles/12323