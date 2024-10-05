# go-practice
Goのキャッチアップリポジトリ

## コマンド
go.modの作成
```
go mod init <モジュールdir>
```

指定したディレクトリ配下全てのテストの実行
```
go test ./...
```

詳細表示
```
go test -v ./...
```

カバレッジ出力
```
go test -cover ./...
```

ソースコードの整形
```
go fmt
```

## 参考
- 自作Moduleのimport
  - https://qiita.com/taku-yamamoto22/items/4d6f9ff8451a0b86997b
