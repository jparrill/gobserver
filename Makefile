FLASK_APP := manage.py
FLASK_ENV := dev

.ONESHELL:
.PHONY: clean install tests run all
.EXPORT_ALL_VARIABLES:

clean:
	rm -rf /tmp/gobserver/*
	rm -f $(GOROOT)/bin/gobserver
	rm -f ./gobserver

install:
	go install ./...

tests:
	go test -v ./...

coverage:
	go test -cover -coverprofile
	go tool cover -func -html c.out

doc:
	godoc

run:
	go run main.go

run_mysql:
	@mkdir -p /tmp/gobserver/mysql
	docker run -d --name gobserver_mysql_db -v /tmp/gobserver/mysql:/var/lib/mysql -p 33060:33060 -p 3306:3306 -p 18080:8080 -e MYSQL_ROOT_PASSWORD=rootpass -e MYSQL_DATABASE=gobserver -e MYSQL_USER=gobserver -e MYSQL_PASSWORD=gobpass mysql

check_mysql:
	mysql -uroot -prootpass -Dgobserver -h127.0.0.1 -P3306 -e 'show global variables like "max_connections"'

clean_mysql:
	docker stop gobserver_mysql_db
	docker rm gobserver_mysql_db
	rm -rf /tmp/gobserver/mysql/*

all: clean install tests run
