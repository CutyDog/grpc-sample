.PHONY: migrate

# マイグレーションコマンドのビルド
migrate-build:
	go build -o migrate cmd/migrate/main.go

# 全てのマイグレーションを実行
migrate-all: migrate-build
	./migrate -all

# 指定したマイグレーションまで実行
migrate-to:
	@read -p "マイグレーション対象のマイグレーションIDを入力してください: " migration_id; \
	./migrate -to $$migration_id

# マイグレーション状態を確認
migrate-status: migrate-build
	./migrate -status