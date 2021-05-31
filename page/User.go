package page

//User 用户注册信息结构体
type User struct {
	Phone      string `json:"phone" binding:"required"`        // 电话号码
	UserPass   string `json:"user_pass" binding:"required"`    // 用户密码
	ReUserPass string `json:"re_user_pass" binding:"required"` // 用户密码
	Code       string `json:"code" binding:"required"`         // 验证码
}

//User 用户登录
type UserLogin struct {
	Phone    string `json:"phone" binding:"required"`     // 电话号码
	UserPass string `json:"user_pass" binding:"required"` // 用户密码
	Code     string `json:"code" binding:"required"`      // 验证码
}
