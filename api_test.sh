
# 测试邮箱、手机号码是否存在
curl  http://localhost:9090/v1/ping


curl "http://localhost:9090/v1/auth/signup/phone/exist" -H "Content-Type: application/json" -d "{
    \"phone\": \"18018001800\"
}" -X POST


curl --request POST 'http://localhost:9090/v1/auth/signup/phone/exist' \
--header 'Content-Type: application/json' \
--data-raw '{"phone": "18018001800"}'



curl "http://localhost:9090/v1/auth/verify-codes/captcha" -X POST

curl --request POST 'http://localhost:9090/v1/auth/signup/email/exist' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "sjxiang@qq.com"}'



# 测试发送验证码请求

curl "http://localhost:9090/v1/auth/verify-codes/captcha" -X POST
{
    "captcha_id":"QtD9VBEHz96M2s1qsBdu",
    "captcha_image":"data:image/png;base64,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"
}

base64 还原图片验证码
https://tool.jisuapi.com/base642pic.html 

"207757"


make login_redis
127.0.0.1:6379> get 'Gohub - CaptchaQtD9VBEHz96M2s1qsBdu' 
"207757"


# 