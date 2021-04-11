

.PHONY: dev
dev:
	DB_CONNECTION_STRING="cfsdev:cfsdev@/cfsservice?parseTime=true" \
	PUBLIC_KEY_PATH="$(shell pwd)/dev/jwtRS256.key.pub" \
	go run cmd/main.go

reset-db-dev:
	./dev/reset-db.sh