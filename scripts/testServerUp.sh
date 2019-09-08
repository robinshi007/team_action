#!/bin/bash

export APP_ENV=test
echo "resetting test database"
rm tmp/team_action_test.db
cp tmp/team_action_test.db.bak tmp/team_action_test.db

echo "starting server"
go run cmd/server/main.go
