# 就活サービス管理(仮)のバックエンドリポジトリ

## 開発環境
- Go：1.24.0

## セットアップ手順
1. リポジトリをローカルに取り込む
```
git clone git@github.com:kappy713/job-hunting-service-management-backend.git
```

### 2. Go モジュールの依存関係を解決
```
go mod tidy
```

### 3. 開発用サーバーの起動
```
go run main.go
```

上記コマンド実行後、http://localhost:8080/ にアクセス<br>
※サーバーを停止させたい場合はCtrl+Cを実行
