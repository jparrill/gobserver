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

all: clean install tests run


#####################
####### MYSQL #######
#####################

mysql: clean_mysql run_mysql sleep_mysql check_mysql configure_mysql

run_mysql:
	@echo " -----> Raising up a MySQL Server: "
	@mkdir -p /tmp/gobserver/mysql
	@docker run -d --name gobserver_mysql_db -v /tmp/gobserver/mysql:/var/lib/mysql -p 33060:33060 -p 3306:3306 -p 18080:8080 -e MYSQL_ROOT_PASSWORD=rootpass -e MYSQL_DATABASE=gobserver -e MYSQL_USER=gobuser -e MYSQL_PASSWORD=gobpass mysql

configure_mysql:
	@echo " -----> Configuring User and Grants on MySQL Server..."
	@MYSQL_PWD=rootpass mysql -uroot -h127.0.0.1 -P3306 -e "GRANT CREATE ON *.* TO 'gobuser'@'%';"
	@MYSQL_PWD=rootpass mysql -uroot -h127.0.0.1 -P3306 -e "GRANT ALL PRIVILEGES ON gobserver.* TO 'gobuser'@'%';"
	@echo " -----> Checking the new user configured..."
	@MYSQL_PWD=gobpass mysql -ugobuser -Dgobserver -h127.0.0.1 -P3306 -e 'show global variables like "max_connections"'

check_mysql:
	@echo " -----> Checking MySQL Server connectivity..."
	@MYSQL_PWD=rootpass mysql -uroot -Dgobserver -h127.0.0.1 -P3306 -e 'show global variables like "max_connections"'

clean_mysql:
	@echo " -----> Cleaning up a MySQL Container..."
	@docker stop gobserver_mysql_db
	@docker rm gobserver_mysql_db
	@echo " -----> Cleaning up a MySQL files..."
	@rm -rf /tmp/gobserver/mysql/*

sleep_mysql:
	@echo " -----> Waiting for MYSQL Container to raise up..."
	@sleep 30
