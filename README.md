# キャリマネのバックエンドリポジトリ

## キャリマネとは？
就活サービスの情報(自己PRやキャリア観、スキルなど)を一元管理するプロダクト<br>
キャリマネで書いたESから複数の就活サービスに対応する項目を生成できる

フロントエンドのリポジトリは[こちら](https://github.com/kappy713/job-hunting-service-management-frontend)

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

### 3. 環境変数の設定
```
DATABASE_URL=YOUR_DATABASE_URL
FRONTEND_URL=YOUR_FRONTEND_URL
GEMINI_API_KEY=YOUR_GEMINI_API_KEY
PORT=YOUR_PORT
```

### 4. 開発用サーバーの起動
※環境変数が正しく設定されていない場合は開発用サーバーが起動しません※
```
make start
// go run ./app/cmd/server/main.go
```

上記コマンド実行後、http://localhost:8080/ にアクセス<br>
※サーバーを停止させたい場合はCtrl+Cを実行
