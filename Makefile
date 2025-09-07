start:
	air -c .air.toml
migrate:
	sql-migrate up -config=configs/dbconfig.yml
migrate-down:
	sql-migrate down -config=configs/dbconfig.yml
migration:
	./scripts/make-migration.sh
wire:
	wire ./...
swagger:
	swag init -g cmd/main.go -d ./ --parseDependency --parseInternal
test:
	go test ./... -v