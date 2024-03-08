request for an otp
curl -H "Content-Type: application/json" -X POST -d '{"phone_number": "+xxxxxxxxxx"}' http://localhost:8000/otp

verify the otp
curl -H "Content-Type: application/json" -X POST -d '{"phone_number": "+xxxxxxxxxx",code:"xxxxx"}' http://localhost:8000/verifyotp