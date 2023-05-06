
echo negative tests
echo
curl -v http://localhost:8080/users/123
echo

curl -v --location 'http://localhost:8080/users/7' \
--header 'Content-Type: application/json' \
--data '{
    "first_name": "ivo"
}'

curl --location 'http://localhost:8080/users/1' \
--header 'Content-Type: application/json' \
--data '{
    "id": "1",
    "first_name": "ivo",
    "last_name": "stoyanov"
}'
echo
curl --location 'http://localhost:8080/users/1'
echo
