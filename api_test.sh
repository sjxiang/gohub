

# 测试文档 

前置条件：
1. docker-compose 装好
2. 运行 make container_open，启动 mysql、redis 容器
3. 运行 make container_net，查看 mysql、redis 的网络地址
4. 修改 .env 相关环境配置参数


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


# 测试请求 - 用户注册（Email + 邮件验证码）


curl --request POST 'http://localhost:9090/v1/auth/signup/using-email' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "1535484943@qq.com", "name": "sjxiang", "password": "123456", "password_confirm": "123456", "verify_code": "654321"}'


curl --request POST 'http://localhost:9090/v1/auth/signup/using-email' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "153548494@test.com", "name": "sjxiang12", "password": "123456", "password_confirm": "123456", "verify_code": "654321"}'



2022-08-22 00:50:23	
validators/custom_rules.go:54	
'Database Query	
{
    "sql": "SELECT count(*) FROM `users` WHERE name = 'sjxiang'", 
    "time": "0.568 ms", 
    "rows": 1
}

2022-08-22 00:50:23	
validators/custom_rules.go:54	
'Database Query	
{
    "sql": "SELECT count(*) FROM `users` WHERE email = '1535484943@qq.com'", 
    "time": "0.559 ms", 
    "rows": 1
}

2022-08-22 00:50:23
verifycode/verifycode.go:68	
验证码	{"检查验证码": "{\"1535484943@qq.com\":\"654321\"}"}

2022-08-22 00:50:23
user/user_util.go:25	
'Database Query	
{
    "sql": "INSERT INTO `users` (`name`,`email`,`phone`,`password`,`created_at`,`updated_at`) VALUES ('sjxiang','1535484943@qq.com','','123456','2022-08-22 00:50:23.9','2022-08-22 00:50:23.9')", 
    "time": "7.971 ms", 
    "rows": 1
}


2022-08-22 00:50:23	
middlewares/logger.go:97	
HTTP 访问日志	
{   
    "status": 201, 
    "request": "POST /v1/auth/signup/using-email", 
    "query": "", 
    "ip": "127.0.0.1", 
    "user-agent": "curl/7.81.0", 
    "errors": "", 
    "time": "9.653 ms", 
    "Request Body": "", 
    "Response Body": "{\"data\":{\"id\":5,\"name\":\"sjxiang\",\"created_at\":\"2022-08-22T00:50:23.9+08:00\",\"updated_at\":\"2022-08-22T00:50:23.9+08:00\"}}"
}
