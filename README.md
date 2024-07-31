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
### Docker container
```
# Create docker image
$ docker build -t simplebank:latest .
# Run container 
$ docker run --name siplebankapp -p 8080:8080 simplebank:latest / (or image id)

# Run gin_mode
$ docker run --name siplebankapp -p 8080:8080 -e GIN_MODE=release simplebank:latest / (or image id)

# connect container to db
# get ip from postgres container
$ docker container inspect [container name] -> "IPAddress": "172.17.0.2"

# start container agane
docker run --name siplebankapp -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@172.17.0.2/simple_bank?sslmode=disable" simplebank:latest / (or image id)

-----------------------------------------------------------

# Or we can no use postgres ip, for this input next command
# we can se bridge nework
$ docker network ls

# we can see more details about this network
# docker network inspect [OPTIONS] NETWORK [NETWORK...]
$ docker network inspect bridge

# we need to create new network that don`t write postgres ip manually
$ docker network create bank-network -> 26aac71aa1bab84e0daffd18cf5c1bf0083e0a1e4355f1cd6e264d7fd2599e7e

# connect to this network
# docker network connect [OPTIONS] NETWORK CONTAINER (simplebank container - postgres image)
$ docker network connect bank-network simplebank

# check network
$ docker network inspect bank-network

# after this we our project container can connect to DB on one network and creade DB (make postgres)
$ docker run --name simplebankapp --network bank-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@simplebank/simple_bank?sslmode=disable" simplebank:latest
```