# API Server
http://localhost:8080/api/v1/

## API
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
| すべて取得     | GET    | `/instances?state={STATE}`     |
| ステータス更新 | PATCH  | `/instances/{ID}/state/{STATE}` |

## Database

Create Database:

```
CREATE DATABASE "infra-control" OWNER = admin TEMPLATE = template0 ENCODING = 'UTF8' LC_COLLATE = 'ja_JP.UTF-8' LC_CTYPE = 'ja_JP.UTF-8';
```

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

### Test data

Hello World:
```
$ curl -X GET http://localhost:8080/api/v1/
```

Hosts:
```
POST http://localhost:8080/api/v1/hosts

{
    "name": "host-A",
    "limit": 1048576
}

{
    "name": "host-B",
    "limit": 1048576
}
```


Addresses:
```
POST http://localhost:8080/api/v1/address

{
    "ip_address": "10.1.1.11",
    "mac_address": "d0:ec:cb:e6:74:e8"
}

{
    "ip_address": "10.1.1.12",
    "mac_address": "9b:03:d5:85:46:21"
}

{
    "ip_address": "10.1.1.13",
    "mac_address": "79:cd:89:d7:4c:ac"
}

{
    "ip_address": "10.1.1.14",
    "mac_address": "63:d9:b2:66:49:a0"
}

{
    "ip_address": "10.1.1.15",
    "mac_address": "94:ec:83:9d:0c:90"
}

{
    "ip_address": "10.1.2.11",
    "mac_address": "25:81:30:3b:b4:01"
}

{
    "ip_address": "10.1.2.12",
    "mac_address": "b2:61:ea:2e:7d:c6"
}

{
    "ip_address": "10.1.2.13",
    "mac_address": "d0:de:1b:36:b7:7b"
}

{
    "ip_address": "10.1.2.14",
    "mac_address": "e7:ad:d7:0b:09:57"
}

{
    "ip_address": "10.1.2.15",
    "mac_address": "97:dd:17:37:82:32"
}
```
