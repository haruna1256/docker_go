### Golang 勉強の際に使用したリポジトリ

# Dockerサーバー構築手順まとめ
### 1. Dockerのインストール方法

- Docker for Windowsをwww.docker.comからダウンロード
- インストーラを実行し、PCを再起動
- WSL2のアップデートが必要な場合あり

### 2. サーバー環境構築手順

以下の環境が構築可能：

- Webサーバー
- PHP
- MySQL
- PHPMyAdmin

#### ディレクトリ構成

```
dockersv/
  ├── htdocs/
  ├── php/
  ├── docker-compose.yml
  └── php/Dockerfile
```

#### 構築手順

1. 必要なディレクトリを作成
2. 設定ファイルを配置
3. dockersvディレクトリで「docker compose up -d」を実行
4. コンテナの起動確認（docker container ls）

### 3. 動作確認方法

#### Apache・PHPの確認

- htdocsフォルダにtest.phpを作成
- [http://localhost:8080/test.phpでアクセス](http://localhost:8080/test.php)

#### MySQLの確認

- コンテナ名を確認（dockersv-mysql-1）
- 「Docker container exec -it コンテナ名 bash」でコンテナに接続
- 「mysql -u root -p」でMySQLにログイン（パスワード：123qwecc）

#### PHPMyAdminの確認

- http://localhost:8081でアクセス可能

#### ディレクトリの中に入る
```
docker exec -it go-app sh
```

### 実行
```
/work # go run main.go
```


# Gin
[ドキュメント](https://gin-gonic.com/ja/docs/)

GinはGolangで書かれたWebサーバフレームワーク。<br>

高速で認証などのmiddlewareをサポート。また構造的なルーティングやXMLファイルやHTMLファイルを返すことができます。

#### Ginのインストール
```
go get -u github.com/gin-gonic/gin
```
