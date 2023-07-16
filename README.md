## Dockerの起動コマンド
docker-compose up

別のターミナル
go run ./cmd/migrate/migration.go
docker-compose exec dev-postgres bash
// コンテナの中に入る

psql -h dev-postgres -p 5432 -U todo-user -d todos
// パスワード入力する

// 一覧確認

\dt

別のターミナルで
go run ./cmd/main.go

// 上記でサーバーを起動させる

別のターミナルで
curlコマンドでリクエストを送る

