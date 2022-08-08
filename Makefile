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
	go test -v -cover ./...

coverage:
	go test -cover -coverprofile=c.out ./...
	go tool cover -html=c.out

doc:
	godoc

run:
	go run gobserver/main.go

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
	@echo " -----> Checking the new user..."
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

######################
##### PostgreSQL #####
######################

pgsql: clean_pgsql run_pgsql sleep_pgsql check_pgsql configure_pgsql

run_pgsql:
	@echo " -----> Raising up a PostgreSQL Server: "
	@mkdir -p /tmp/gobserver/pgsql
	@docker run -d --name gobserver_pgsql_db -v /tmp/gobserver/pgsql:/var/lib/postgresql/data -p 5432:5432 -e PGDATA=/var/lib/postgresql/data/pgdata -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rootpass postgres

configure_pgsql:
	@echo " -----> Configuring User and Grants on PostgreSQL Server..."
	@PGPASSWORD=rootpass psql -v -w -U root -h127.0.0.1 -p5432 -c "CREATE DATABASE gobserver;"
	@PGPASSWORD=rootpass psql -v -w -U root -h127.0.0.1 -p5432 -c "CREATE USER gobuser with encrypted password 'gobpass';"
	@PGPASSWORD=rootpass psql -v -w -U root -h127.0.0.1 -p5432 -c "GRANT ALL privileges on database gobserver to gobuser;"
	@echo " -----> Checking the new user..."
	@PGPASSWORD=gobpass pg_isready -dgobserver -Ugobuser -h127.0.0.1 -p5432

check_pgsql:
	@echo " -----> Checking PostgreSQL Server connectivity..."
	@PGPASSWORD=gobpass pg_isready -dgobserver -Ugobuser -h127.0.0.1 -p5432

clean_pgsql:
	@echo " -----> Cleaning up a PostgreSQL Container..."
	@docker stop gobserver_pgsql_db
	@docker rm gobserver_pgsql_db
	@echo " -----> Cleaning up a PostgreSQL files..."
	@rm -rf /tmp/gobserver/pgsql

sleep_pgsql:
	@echo " -----> Waiting for PostgreSQL Container to raise up..."
	@sleep 30
