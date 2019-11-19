#!/bin/zsh


if [ $# -eq 1 ]
then
    gofmt -l -w .
    git add -A
    git commit -m $1
else
    echo "コミットに失敗しました."
fi
