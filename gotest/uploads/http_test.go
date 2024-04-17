package uploads

import (
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestUserHandler_Login(t *testing.T) {
//	// 定义测试用例
//	testCases := []struct {
//		name     string
//		reqBody  string
//		wantCode int
//		wantBody string
//	}{
//		{
//			name:     "登录成功",
//			reqBody:  `{"email": "123@qq.com", "pwd": "hello#world123"}`,
//			wantCode: http.StatusOK,
//			wantBody: `{"msg": "登录成功!"}`,
//		},
//		{
//			name:     "参数不正确",
//			reqBody:  `{"email": "123@qq.com", "pwd": "hello#world123",}`,
//			wantCode: http.StatusBadRequest,
//			wantBody: `{"msg": "参数不正确!"}`,
//		},
//		{
//			name:     "邮箱或密码为空",
//			reqBody:  `{"email": "", "pwd": ""}`,
//			wantCode: http.StatusBadRequest,
//			wantBody: `{"msg": "邮箱或密码不能为空"}`,
//		},
//		{
//			name:     "邮箱格式不正确",
//			reqBody:  `{"email": "invalidemail", "pwd": "hello#world123"}`,
//			wantCode: http.StatusBadRequest,
//			wantBody: `{"msg": "邮箱格式不正确"}`,
//		},
//		{
//			name:     "密码格式不正确",
//			reqBody:  `{"email": "123@qq.com", "pwd": "invalidpassword"}`,
//			wantCode: http.StatusBadRequest,
//			wantBody: `{"msg": "密码必须大于8位，包含数字、特殊字符"}`,
//		},
//		{
//			name:     "邮箱或密码不匹配",
//			reqBody:  `{"email": "123123@qq.com", "pwd": "hello#world123"}`,
//			wantCode: http.StatusBadRequest,
//			wantBody: `{"msg": "邮箱或密码不匹配!"}`,
//		},
//	}
//
//	for _, tc := range testCases {
//		t.Run(tc.name, func(t *testing.T) {
//			// 创建一个 gin 的上下文
//			server := gin.Default()
//			h := NewUserHandler()
//			h.RegisterRoutes(server)
//			// mock 创建一个 http 请求
//			req, err := http.NewRequest(
//				http.MethodPost,                     // 请求方法
//				"/user/login",                       // 请求路径
//				bytes.NewBuffer([]byte(tc.reqBody)), // 请求体
//			)
//			// 断言没有错误
//			assert.NoError(t, err)
//			// 设置请求头
//			req.Header.Set("Content-Type", "application/json")
//			// 创建一个响应
//			resp := httptest.NewRecorder()
//			// 服务端处理请求
//			server.ServeHTTP(resp, req)
//			// 断言响应码和响应体
//			assert.Equal(t, tc.wantCode, resp.Code)
//			// 断言 JSON 字符串是否相等
//			assert.JSONEq(t, tc.wantBody, resp.Body.String())
//		})
//	}
//}

func TestGetResultByAPI(t *testing.T) {
	defer gock.Off() // 测试执行后刷新挂起的mock

	// mock 请求外部api时传参x=1返回100
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 1}).
		Reply(200).
		JSON(map[string]int{"value": 100})

	// 调用我们的业务函数
	res := GetResultByAPI(1, 1)
	// 校验返回结果是否符合预期
	assert.Equal(t, res, 101)

	// mock 请求外部api时传参x=2返回200
	gock.New("http://your-api.com").
		Post("/post").
		MatchType("json").
		JSON(map[string]int{"x": 2}).
		Reply(200).
		JSON(map[string]int{"value": 200})

	// 调用我们的业务函数
	res = GetResultByAPI(2, 2)
	// 校验返回结果是否符合预期
	assert.Equal(t, res, 202)

	assert.True(t, gock.IsDone()) // 断言mock被触发
}
