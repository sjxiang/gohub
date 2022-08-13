
curl  http://localhost:9090/v1/ping


curl "http://localhost:9090/v1/auth/signup/phone/exist" -H "Content-Type: application/json" -d "{
    \"phone\": \"18018001800\"
}" -X POST
