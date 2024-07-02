### Install 
```
# Golang-migrate - https://github.com/golang-migrate/migrate/tree/master/cmd/migrate
$ scoop install migrate
```

### Migrations
```
# Initialize
$ migrate create -ext sql -dir db/migration -seq init_schema

# Create migration
# migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up
```

### SQLC
- Init command 
```
# generate sqlc.yaml
$ sqlc init

# on makefile create script

# install pq
$ go get github.com/lib/pq

# run test(TestCreateAccount - method name)(chose one)
$ go test -timeout 30s github.com/TheDidEv/simplebank/db/sqlc -run ^TestCreateAccount || go test -v -cover ./... || make test
```

