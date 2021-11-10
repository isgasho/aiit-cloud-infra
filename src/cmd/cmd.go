package cmd

// Batch: 1 (Pauling)
// 1. Starting の instance を取得する
// 2. Initializing にする
// 3. SSH Key を払い出して keys.data を更新する
// 4. instance をつくって Key を配置する
// 5. Private Key を渡す (メールで？)
// 6. Running にする

// Batch: 2 (Pauling)
// 1. Terminating の instance を取得する
// 2. instance を削除する
// 3. Terminated にする

// Batch: 3 (Daily)
// 1. Terminated で Addresses に紐付いている instance_id を削除する (しばらく経過...など)
