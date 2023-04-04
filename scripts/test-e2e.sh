#!/bin/bash

API_URL=http://localhost:8080

echo "E2E script for User Service"

echo
echo "Running end-to-end testing..."; echo
echo "$API_URL/api/health"
curl $API_URL/api/health; echo;
echo
echo "$API_URL/api/v1/users/:userId"
curl $API_URL/api/v1/users/1; echo;
echo
echo "Finished testing..."
