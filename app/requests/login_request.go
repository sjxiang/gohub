
// 登录

package requests


type LoginByPasswordRequest struct {
	CaptchaID string `json:"captcha_id,omitempty" valid:"captcha_id"`
}