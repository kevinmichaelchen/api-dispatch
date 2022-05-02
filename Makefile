.PHONY: all
all:
	$(MAKE) gen-proto
	$(MAKE) gen-models

.PHONY: gen-proto
gen-proto:
	buf generate idl

.PHONY: gen-models
gen-models:
	sqlboiler psql --output internal/models