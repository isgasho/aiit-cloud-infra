# クラウドインフラ構築特論

Web service name: `infra-control`

## Infra Control System Design

(T.B.D)

## Database

Create Database:

```
CREATE DATABASE "infra-control" OWNER = admin TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'ja_JP.UTF-8' LC_CTYPE = 'ja_JP.UTF-8';
```

## API

http://localhost:8088/api/v1/

### Host
| API            | Method | Path                            |
| -------------- | ------ | ------------------------------- |
| 追加           | POST   | `/hosts`                        |

### Address
| API            | Method | Path                            |
| -------------- | ------ | ------------------------------- |
| 追加           | POST   | `/addresses`                    |

### Instance
| API            | Method | Path                            |
| -------------- | ------ | ------------------------------- |
| 追加           | POST   | `/instances`                    |
| 取得           | GET    | `/instances/{ID}`               |
| 終了           | DELETE | `/instances/{ID}`               |
| すべて取得     | GET    | `/instances?host_id={ID}`       |
| ステータス更新 | PATCH  | `/instances/{ID}/state/{STATE}` |

## Local Development

### How to launch the application

```
$ docker-compose up --build
```

### How to connect the database

```
$ docker exec -it infra-control-db psql -U admin
```

```
$ psql -h 0.0.0.0 -p 5437 -U admin infra-control
```

### How to migrate

Exit from the container, perform migration.

```
$ export POSTGRESQL_URL='postgres://admin:admin@0.0.0.0:5437/infra-control?sslmode=disable'
$ migrate -database ${POSTGRESQL_URL} -path migrations up
```
