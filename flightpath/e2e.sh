#!/bin/bash

echo 'error "bind: address already in use" is expected if the service is already running'

export DEBUG=TRUE
export USE_GRAPH=graphlib
#export USE_GRAPH=simple
#export USE_GRAPH=kahn

pkill fligths
go build -o ./flights cmd/main.go

./flights &

sleep 2

echo
echo 'Negative test [["SAN", "SFO"], ["SFO", "LAS"], ["LAS", "SAN"]]'
echo 'Expected response: status 400, invalid argument '
echo
curl  -X  POST http://localhost:8080/calculate \
      -H  "Content-Type: application/json" \
      -d  '[["SAN", "SFO"], ["SFO", "LAS"], ["LAS", "SAN"]]'

echo
echo 'Negative test [["foo"], [1, 2]] invalid json'
echo 'Expected response: status 400, Unmarshal error '
echo
curl  -X  POST http://localhost:8080/calculate \
      -H  "Content-Type: application/json" \
      -d  '[["foo"], [1, 2]]'

echo
echo 'Negative test: Empty input'
echo 'Expected response: status 400'
echo
curl  -X  POST http://localhost:8080/calculate \
      -d  '{}'

echo
echo 'Negative test: GET'
echo 'Expected response: status 405 - not allowed'
echo
curl  http://localhost:8080/calculate

echo
echo 'health check - expected 200 OK and the version'
echo
curl  http://localhost:8080/health
echo

echo
echo 'MAIN TEST: Sending request [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'
echo 'Expected response: status 200,  ["SFO","EWR"]'
echo
curl  -X  POST http://localhost:8080/calculate \
      -H  "Content-Type: application/json" \
      -d  '[["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]]'
echo
