# クラウドインフラ構築特論
Web service name: `infra-control` 

## Infra Control System Design
## Database
```
$ docker exec -it infra-control-db psql -U admin
```

or

```
$ psql -h 0.0.0.0 -p 5437 -U admin infra-control
```

## API
| API            | Method | Path                            |
| -------------- | ------ | ------------------------------- |
| 追加           | POST   | `/instances`                    |
| 取得           | GET    | `/instances/{ID}`               |
| 終了           | DELETE | `/instances/{ID}`               |
| すべて取得     | GET    | `/instances`                    |
| ステータス更新 | PATCH  | `/instances/{ID}/state/{state}` |

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
$ export POSTGRESQL_URL='postgres://admin:admin@0.0.0.0:5437/infra-control?sslmode=disable'
$ migrate -database ${POSTGRESQL_URL} -path migrations up
```