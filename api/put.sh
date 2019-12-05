#!/bin/zsh

curl -X PUT -H "Content-Type: application/json" -d '{"date":"2019-11-19", "body":"TypeScript", "is_finished":true}' http://localhost:8080/todo-api/tasks/1
