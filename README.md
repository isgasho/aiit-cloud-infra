# クラウドインフラ構築特論
Web service name: `infra-control` 

## Infra Control System Design
## Database


## API


## Local Development
### How to launch the applicaiton
```
$ docker-compose up --build
```

### How to connect the database
```
$ docker exec -it infra-control-db psql -U admin
```

### How to migrate
Exit from the container, perform migration.
```
$ export POSTGRESQL_URL='postgres://admin:admin@0.0.0.0:5435/infra-control?sslmode=disable'
$ migrate -database ${POSTGRESQL_URL} -path migrations up
```
