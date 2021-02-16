.PHONY: test mocks
test:
	go test -v ./... -bench . -benchmem

mocks: 
	cd db && mockgen -package db -destination ../mocks/db/db.go . Session,Database,Collection,Query

