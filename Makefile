dev:
	./air -c .air.toml

dev-init:
	go build -o ./tmp/main ./cmd/main.go

container-logs-follow:
	docker container logs container-tsu-backup --follow