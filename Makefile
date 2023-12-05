build:
	@go build -o bin/app
run: build
	@./bin/app
drop:
	@go run scripts/drop.go