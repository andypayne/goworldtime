#!/usr/bin/env bash

PORT=4040
URL=http://localhost:${PORT}/times
#URL=https://postman-echo.com/post

curl -v \
  -X POST \
  -d  @test_data.json \
  -H 'Content-Type: application/json' \
  ${URL}

