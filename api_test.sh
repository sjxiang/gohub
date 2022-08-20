
# 测试请求 - 服务器连通性

curl  http://localhost:9090/v1/ping



# 测试请求 - 邮箱、手机号码是否存在

curl --request POST 'http://localhost:9090/v1/auth/signup/phone/exist' \
--header 'Content-Type: application/json' \
--data-raw '{"phone": "18018001800"}'


curl --request POST 'http://localhost:9090/v1/auth/signup/email/exist' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "sjxiang@qq.com"}'



# 测试请求 - 发送图片验证码

curl "http://localhost:9090/v1/auth/verify-codes/captcha" -X POST

方式 1、base64 还原图片验证码
方式 2、redis> get 'Gohub:Captcha:{{captcha_id}}' 



# 测试请求 - 发送邮箱验证码（管理员权限白名单）

curl --request POST 'http://localhost:9090/v1/auth/verify-codes/email' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "1535484943@qq.com", "captcha_id": "pass123", "captcha_answer": "123456"}'