ARGS :=
create-m:
	migrate create -ext sql -dir pkg/database/migrations -seq $(ARGS)
