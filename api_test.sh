
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
    "captcha_image":"data:image/png;base64,"
}

base64 还原图片验证码
https://tool.jisuapi.com/base642pic.html 

"207757"


make login_redis
127.0.0.1:6379> get 'Gohub - CaptchaQtD9VBEHz96M2s1qsBdu' 
"207757"


# 



curl "http://localhost:9090/v1/auth/verify-codes/captcha" -X POST



# 
curl --request POST 'http://localhost:9090/v1/auth/verify-codes/email' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "1535484943@qq.com", "captcha_id": "pass123", "captcha_answer": "123456"}'