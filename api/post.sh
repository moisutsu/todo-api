#!/bin/zsh

curl -X POST -H "Content-Type: application/json" -d '{"date":"2019-11-19", "body":"TypeScriptの勉強", "is_finished":0}' http://localhost:8080/todo-api/tasks
