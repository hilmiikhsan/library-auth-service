test:
	go test -v ./... -cover

run:
	go run main.go serve-http

hot:
	@echo " >> Installing gin if not installed"
	@go install github.com/codegangsta/gin@latest
	@gin -i -p 9002 -a 9090 serve-http

goose-create:
# example : make goose-create name=create_users_table
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
ifndef name
	$(error Usage: make goose-create name=<table_name>)
else
	@goose -dir scripts/migrations/sql create $(name) sql
endif

goose-up:
# example : make goose-up
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir scripts/migrations/sql postgres "host=localhost user=postgres password=password dbname=library_auth sslmode=disable" up

goose-down:
# example : make goose-down
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir scripts/migrations/sql postgres "host=localhost user=postgres password=password dbname=library_auth sslmode=disable" down

goose-status:
# example : make goose-status
	@echo " >> Installing goose if not installed"
	@go install github.com/pressly/goose/v3/cmd/goose@latest
	@goose -dir scripts/migrations/sql postgres "host=localhost user=postgres password=password dbname=library_auth sslmode=disable" status