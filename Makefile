migrate:
	migrate -path migrations -database "mysql://root:@tcp(localhost:3306)/todo4" -verbose ${verbose}

create-table:
	migrate create -ext sql -dir migrations -seq ${table}