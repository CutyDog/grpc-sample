# gRPC Sample Project

Go + gRPC + GraphQL のマイクロサービスアーキテクチャのサンプルプロジェクト

## 📋 プロジェクト概要

このプロジェクトは、以下の技術スタックを使用したマイクロサービスアーキテクチャの実装例です：

- **gRPC**: マイクロサービス間の通信
- **GraphQL**: クライアント向けのAPIゲートウェイ
- **Protocol Buffers**: スキーマ定義とコード生成
- **Buf**: Protocol Buffersのビルドツール
- **Docker Compose**: 開発環境の構築
- **PostgreSQL**: データベース
- **Air**: ホットリロード開発サーバー

## 🏗️ アーキテクチャ

```
クライアント
    ↓ (HTTP/GraphQL)
GraphQL Service (Port 3000)
    ↓ (gRPC)
Account Service (Port 9090)
    ↓ (GORM)
PostgreSQL Database
```

## 📁 ディレクトリ構成

```
grpc-sample/
├── proto/                      # Protocol Buffers定義
│   ├── account/v1/            # Accountサービスのスキーマ
│   │   └── account.proto
│   ├── gen/                   # 生成されたGoコード（.gitignore対象）
│   ├── buf.yaml               # Buf設定
│   ├── buf.gen.yaml          # コード生成設定
│   └── go.mod                # protoモジュール
│
├── services/
│   ├── account/              # Account gRPCサービス
│   │   ├── cmd/
│   │   │   ├── server/       # サーバーエントリーポイント
│   │   │   └── migrate/      # マイグレーションツール
│   │   ├── internal/
│   │   │   ├── db/           # データベース接続
│   │   │   ├── model/        # データモデル
│   │   │   ├── repo/         # リポジトリ層
│   │   │   └── server/       # gRPCサーバー実装
│   │   ├── Dockerfile.dev    # 開発用Dockerfile
│   │   └── .air.toml         # Airホットリロード設定
│   │
│   └── graphql/              # GraphQL APIゲートウェイ
│       ├── cmd/server/       # サーバーエントリーポイント
│       ├── client/           # gRPCクライアント
│       ├── graph/            # GraphQL実装
│       │   ├── model/        # GraphQLモデル
│       │   ├── resolver.go   # リゾルバーベース
│       │   └── schema.resolvers.go  # リゾルバー実装
│       ├── schema.graphql    # GraphQLスキーマ
│       ├── gqlgen.yml        # gqlgen設定
│       └── Dockerfile.dev    # 開発用Dockerfile
│
└── docker-compose.yml        # Docker Compose設定
```

## 🚀 セットアップ

### 前提条件

- Docker & Docker Compose
- Go 1.24以上（ローカル開発時）
- Buf CLI（ローカルでproto生成する場合）

### 環境構築

1. **リポジトリをクローン**

```bash
git clone <repository-url>
cd grpc-sample
```

2. **Protoファイルからコード生成（初回のみ）**

```bash
cd proto
buf generate
```

3. **サービスを起動**

```bash
docker-compose up
```

または、特定のサービスのみ起動：

```bash
docker-compose up account    # Accountサービスのみ
docker-compose up graphql    # GraphQLサービスのみ
```

## 📡 エンドポイント

| サービス | エンドポイント | 説明 |
|---------|--------------|------|
| GraphQL API | http://localhost:3000/ | GraphQL Playground |
| GraphQL Query | http://localhost:3000/query | GraphQLクエリエンドポイント |
| Account gRPC | localhost:9090 | gRPCサービス（内部通信用） |

## 🔧 開発ワークフロー

### Protoファイルを更新する

1. `proto/account/v1/account.proto` を編集
2. コード生成を実行：

```bash
cd proto
buf generate
```

3. サービスを再起動：

```bash
docker-compose restart account graphql
```

### GraphQLスキーマを更新する

1. `services/graphql/schema.graphql` を編集
2. コード生成を実行：

```bash
cd services/graphql
go run github.com/99designs/gqlgen generate
```

3. リゾルバーを実装：`graph/schema.resolvers.go`

### データベースマイグレーション

マイグレーションを実行：

```bash
docker exec -it <container-name> ./migrate -all
```

マイグレーション状態を確認：

```bash
docker exec -it <container-name> ./migrate -status
```

## 📝 使用例

### GraphQLクエリ例

GraphQL Playground (http://localhost:3000/) で以下のクエリを実行：

```graphql
query GetAccount {
  account(uid: "firebase-uid-here") {
    id
    uid
    createdAt
    updatedAt
  }
}
```

### gRPCクライアント例（grpcurl）

```bash
# サービス一覧を取得
grpcurl -plaintext localhost:9090 list

# Accountを取得（UIDで検索）
grpcurl -plaintext -d '{"uid": "firebase-uid"}' \
  localhost:9090 account.v1.AccountService/GetAccount
```

## 🛠️ 技術スタック

### バックエンド

- **Go 1.24**: プログラミング言語
- **gRPC**: マイクロサービス間通信
- **Protocol Buffers**: IDL（インターフェース定義言語）
- **GraphQL (gqlgen)**: クライアント向けAPI
- **GORM**: ORMライブラリ
- **PostgreSQL**: リレーショナルデータベース

### 開発ツール

- **Buf**: Protocol Buffersのビルドツール
- **Air**: Goのホットリロードツール
- **Docker Compose**: コンテナオーケストレーション

## 📚 主要な概念

### リポジトリパターン

データアクセス層を抽象化し、ビジネスロジックとデータベースアクセスを分離：

```go
// インターフェース定義
type AccountRepository interface {
    FindByID(ctx context.Context, id int64) (*model.Account, error)
    FindByUID(ctx context.Context, uid string) (*model.Account, error)
}

// 実装
type accountRepository struct {
    db *gorm.DB
}
```

### Protocol Buffers & Buf

- `.proto`ファイルでスキーマを定義
- `buf generate`でGoコードを自動生成
- `buf.yaml`でリント設定
- `buf.gen.yaml`でコード生成設定

### Go Modules依存関係管理

各サービスは独立したGoモジュール：

```
services/account/go.mod       # Accountサービス
services/graphql/go.mod       # GraphQLサービス
proto/go.mod                  # Proto生成コード
```

`replace`ディレクティブでローカルprotoモジュールを参照：

```go
replace github.com/CutyDog/grpc-sample/proto => ../../proto
```

## 🐛 トラブルシューティング

### Protoコード生成でエラーが出る

```bash
# bufを再インストール
go install github.com/bufbuild/buf/cmd/buf@latest

# 生成ファイルを削除して再生成
rm -rf proto/gen
cd proto && buf generate
```

### コンテナが起動しない

```bash
# コンテナとボリュームを完全に削除
docker-compose down -v

# イメージを再ビルド
docker-compose build --no-cache

# 再起動
docker-compose up
```

### データベース接続エラー

```bash
# データベースの状態を確認
docker-compose logs account-db

# データベースコンテナを再起動
docker-compose restart account-db
```

## 📄 ライセンス

MIT License

## 👥 コントリビューター

開発中のサンプルプロジェクトです。

## 🔗 参考リンク

- [gRPC Documentation](https://grpc.io/docs/)
- [Protocol Buffers](https://protobuf.dev/)
- [Buf](https://buf.build/)
- [gqlgen](https://gqlgen.com/)
- [GORM](https://gorm.io/)

