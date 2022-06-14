.PHONY: all
all:
	$(MAKE) gen-proto
	$(MAKE) gen-models

.PHONY: gen-proto
gen-proto:
	rm -rf internal/idl
	buf breaking idl --against '.git#branch=main,subdir=idl'
	buf mod update idl
	buf lint idl
	buf format idl -w
	buf generate idl

.PHONY: gen-models
gen-models:
	sqlboiler psql --output internal/models

.PHONY: migrate-up
migrate-up:
	migrate -path ./schema -database postgres://postgres:postgres@localhost:5432/dispatch\?sslmode=disable up

.PHONY: migrate-down
migrate-down:
	migrate -path ./schema -database postgres://postgres:postgres@localhost:5432/dispatch\?sslmode=disable down
