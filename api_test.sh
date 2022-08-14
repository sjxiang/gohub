
curl  http://localhost:9090/v1/ping


curl "http://localhost:9090/v1/auth/signup/phone/exist" -H "Content-Type: application/json" -d "{
    \"phone\": \"18018001800\"
}" -X POST


curl --request POST 'http://localhost:9090/v1/auth/signup/phone/exist' \
--header 'Content-Type: application/json' \
--data-raw '{"phone": "18018001800"}'


curl --request POST 'http://localhost:9090/v1/auth/signup/email/exist' \
--header 'Content-Type: application/json' \
--data-raw '{"email": "sjxiang@qq.com"}'