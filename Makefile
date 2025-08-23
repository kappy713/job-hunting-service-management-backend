.PHONY: start, migrate

# 開発用サーバーの起動
start:
	go run ./app/cmd/server/main.go

# マイグレーションの実行
migrate:
	go run ./app/cmd/migrate/main.go