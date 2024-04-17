package uploads

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

//	type UserHandler struct {
//		emailExp    *regexp.Regexp
//		passwordExp *regexp.Regexp
//	}
//
//	func (u *UserHandler) RegisterRoutes(server *gin.Engine) {
//		ug := server.Group("/user")
//		ug.POST("/login", u.Login)
//	}
//
//	func NewUserHandler() *UserHandler {
//		const (
//			emailRegexPattern    = "^\\w+([-+.]\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$"
//			passwordRegexPattern = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[$@$!%*#?&])[A-Za-z\d$@$!%*#?&]{8,}$`
//		)
//		emailExp := regexp.MustCompile(emailRegexPattern, regexp.None)
//		passwordExp := regexp.MustCompile(passwordRegexPattern, regexp.None)
//		return &UserHandler{
//			emailExp:    emailExp,
//			passwordExp: passwordExp,
//		}
//	}
//
//	type LoginRequest struct {
//		Email string `json:"email"`
//		Pwd   string `json:"pwd"`
//	}
//
//	func (u *UserHandler) Login(ctx *gin.Context) {
//		var req LoginRequest
//		if err := ctx.ShouldBindJSON(&req); err != nil {
//			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "参数不正确!"})
//			return
//		}
//
//		// 校验邮箱和密码是否为空
//		if req.Email == "" || req.Pwd == "" {
//			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱或密码不能为空"})
//			return
//		}
//
//		// 正则校验邮箱
//		ok, err := u.emailExp.MatchString(req.Email)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "系统错误!"})
//			return
//		}
//		if !ok {
//			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱格式不正确"})
//			return
//		}
//
//		// 校验密码格式
//		ok, err = u.passwordExp.MatchString(req.Pwd)
//		if err != nil {
//			ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "系统错误!"})
//			return
//		}
//		if !ok {
//			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "密码必须大于8位，包含数字、特殊字符"})
//			return
//		}
//
//		// 校验邮箱和密码是否匹配特定的值来确定登录成功与否
//		if req.Email != "123@qq.com" || req.Pwd != "hello#world123" {
//			ctx.JSON(http.StatusBadRequest, gin.H{"msg": "邮箱或密码不匹配!"})
//			return
//		}
//
//		ctx.JSON(http.StatusOK, gin.H{"msg": "登录成功!"})
//	}
//
//	func InitWebServer(userHandler *UserHandler) *gin.Engine {
//		server := gin.Default()
//		userHandler.RegisterRoutes(server)
//		return server
//	}
//
// ReqParam API请求参数
type ReqParam struct {
	X int `json:"x"`
}

// Result API返回结果
type Result struct {
	Value int `json:"value"`
}

func GetResultByAPI(x, y int) int {
	p := &ReqParam{X: x}
	b, _ := json.Marshal(p)

	// 调用其他服务的API
	resp, err := http.Post(
		"http://your-api.com/post",
		"application/json",
		bytes.NewBuffer(b),
	)
	if err != nil {
		return -1
	}
	body, _ := ioutil.ReadAll(resp.Body)
	var ret Result
	if err := json.Unmarshal(body, &ret); err != nil {
		return -1
	}
	// 这里是对API返回的数据做一些逻辑处理
	return ret.Value + y
}
