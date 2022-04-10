#!/bin/sh

COMMAND=$1
DEFAULTAG="go-echo-server"

function docker_build() {
	read -p "Tag Name (default: $DEFAULTAG) : " $TAG
	if [ -z $TAG ]; then
		TAG=`echo $DEFAULTAG`
	fi
	docker build -t $TAG .
}

function go_build() {
	go build -o ./build/go-echo-server ./cmd/main.go
}

function docker_compose() {
	docker compose up --build
	docker compose down
}

function go_run() {
	export MYSQL_DSN="root:animelovers@tcp(localhost:3306)/crud_go?charset=utf8&parseTime=True&loc=Local"
	go run ./cmd/main.go
	unset MYSQL_DSN
}

function help() {
printf "Command:
=> go_build\t\t\tBuild go application to binary
=> go_run\t\t\tRun go application
=> docker_build\t\t\tBuild docker container
=> docker_compose\t\tBuild and run docker container
=> help\t\t\t\tShow this help text\n
example: ./run.sh go_build\n
"
return 0
}

if [[ -z $COMMAND ]]; then
	help
fi
$COMMAND
