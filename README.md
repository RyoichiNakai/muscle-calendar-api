# muscle-calendar-api

## アーキテクチャ

レイヤードアーキテクチャを採用。
レイヤードアーキテクチャとは？は以下を参考にすること。

- <https://tech.yyh-gl.dev/blog/go_web_api/>
- <https://qiita.com/tono-maron/items/345c433b86f74d314c8d>

上記のアーキテクチャを使うためのディレクトリ採用にしている。
詳しくは以下のディレクトリ構成に記載している。

## ディレクトリ構成

```bash
todo:
```

## go の環境

- バージョン: 1.21
- フレームワーク: gin

## ローカルでの開発

dockerを用いる。具体的なユースケースは以下。

### ローカルPCでコンテナを立ち上げる

```bash
docker compose up
```

[air](https://github.com/cosmtrek/air)によるホットリロードを導入し、コードを編集・保存すると自動で反映されるようになっている。

### 立ち上げたコンテナをシャットダウン

```bash
docker compose down
```

### localhostにアクセス

以下は/v1/healthにリクエスト結果を確認するためのコマンドは以下

```bash
curl localhost:8080/v1/health
```

### docker compose で立ち上げたmysqlのデータを確認する

```bash
docker container exec -it muscle-calendar-api-mysql-1 mysql -h localhost -u root -p
```

## ブランチ運用

- mainからfeatureブランチを切って、commitとpushを行う
  - めんどいので、stg環境とかは用意しない。
- featureブランチからmainにマージすれば、自動的に本番環境にデプロイされる
