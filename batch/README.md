# Command-line Tool
## Design
### Create Instance
1. Instance 作成 API を Call する

```
$ go run main.go create --name sample-instance-1 --size 10240
$ go run main.go create --name sample-instance-2 --size 10240
$ go run main.go create --name sample-instance-3 --size 10240
$ go run main.go create --name sample-instance-4 --size 10240
$ go run main.go create --name sample-instance-5 --size 10240
$ go run main.go create --name sample-instance-6 --size 10240
$ go run main.go create --name sample-instance-7 --size 10240
$ go run main.go create --name sample-instance-8 --size 10240
```

### Run Instance (Pauling)
3. Starting の Instance を取得する
4. Initializing にする
5. SSH Key を払い出して keys.data を更新する
6. Instance を作る
7. Instance に Key を配置する
8. Running にする
9. Private Key と設定情報を渡す
   1. 鍵ファイルとテキストファイルを任意の場所に保存する
   2. (できれば) hosts にメールアドレス項目を追加して送信する

```
$ go run main.go running
```

### Terminate Instance (Pauling)
実際に Instance を削除し、Instance テーブルの State を更新する。

1. Terminating の Instance を取得する
2. Instance を削除する
3. Terminated にする

```
$ go run main.go terminated
```

### Release Addresses (Daily)
1. Instance が Terminated で Addresses に紐付いている instance_id を削除する

```
$ go run main.go release
```

### Terminate Instance
Database の情報を削除する。

1. Instance 削除 API を Call する

```
$ go run main.go delete --instance_id 1
```

## Development

Install packages:
```
$ cd batch
$ go get -u github.com/spf13/cobra/cobra
```

Command-line initialize:
```
$ cobra init --pkg-name infra-control --viper=false
```

SubCommand added:
```
$ cobra add create
$ cobra add running
$ cobra add terminated
$ cobra add release
$ cobra add delete
```
