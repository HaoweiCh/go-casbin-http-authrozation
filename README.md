# casbin-http-role-exampe

简单且实用的 HTTP 鉴权体系

更新历史

* 2021-11-05T11:11:48+0800
  * 升级依赖
  * 升级 go 版本
  * casbin 规则独立成文件, (通过 embed 特性编译时嵌入)
* 2019.6.21 
  * 更新 http 接口文件
  * 使用Go Module  做包管理工具
    * 推荐 Module Proxy   https://goproxy.io

## 依赖

* [casbin](https://github.com/casbin/casbin) for role-based HTTP Authorization
* [scs](https://github.com/alexedwards/scs)  for session handling.
* redis 会话 token 存储
* mariadb 用户和规则存储 
  * 记得修改 rule.go 里面的sql 服务器帐号密码


## Run with

```bash
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