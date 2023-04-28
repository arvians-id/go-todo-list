migrate-docker:
	migrate -path database/migrations -database "mysql://root:root@tcp(host.docker.internal:3306)/todo4" -verbose ${verbose}

migrate-dev:
	migrate -path migrations -database "mysql://root:@tcp(localhost:3306)/todo4" -verbose ${verbose}

create-table:
	migrate create -ext sql -dir migrations -seq ${table}

run-docker:
	docker run -d --name my_app -e MYSQL_HOST=host.docker.internal -e MYSQL_USER=root -e MYSQL_PASSWORD=root -e MYSQL_DBNAME=todo4 -p 3030:3030 arvians/go-todo-list:latest

run-dev:
	go run cmd/server/main.go