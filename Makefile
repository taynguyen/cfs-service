

.PHONY: dev
dev:
	DB_CONNECTION_STRING="cfsdev:cfsdev@/cfsservice?parseTime=true" \
	go run cmd/main.go

reset-db-dev:
	./dev/reset-db.sh