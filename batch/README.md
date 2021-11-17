# Command-line Tool
## Design
### Create Instance
1. Instance 作成 API を Call する

### Run Instance (Pauling)
3. Starting の Instance を取得する
4. Initializing にする
5. SSH Key を払い出して keys.data を更新する
6. Instance を作る
7. Instance の Key を配置する
8. Private Key を渡す (メールで？)
9. Running にする

### Terminate Instance (Pauling)
1. Terminating の instance を取得する
2. Instance を削除する
3. Terminated にする

### Release Addresses (Daily)
1. Instance が Terminated で Addresses に紐付いている instance_id を削除する

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
```
