# muscle-calendar-api

## goの環境構築

versionは`1.21.0`
フレームワークはginを使うか
swaggerでapiテストを自動的にするようにしよう
テストも書くようにする

swag initのコマンド

```bash
swag init -g ./cmd/muscle-calendar-api/main.go -o ./docs
```

## アーキテクチャ

レイヤードアーキテクチャを採用

これを参考にするとわかりやすいかも

[https://tech.yyh-gl.dev/blog/go_web_api/]
[https://qiita.com/tono-maron/items/345c433b86f74d314c8d]

## コードを書く上で注意する点

以下の点に注意してコードを書くようにすること

- レイヤードアーキテクチャ
- SOLIDの原則

## todo

1. レイヤードアーキテクチャについて学び直す
2. 簡単なDDDを実装する(簡単なのでドメインモデル貧血症になるのは仕方ない)
  - 調査の過程でドメインモデル貧血症になると何がダメなのかを調べる
3. airについてを学び直す(なぜかhot reloadされないので修正)
4. Readmeをまとめる

## docker コマンド

### mysqlに入るためのコマンド

```bash
docker container exec -it muscle-calendar-api-mysql-1 mysql -uroot -hlocalhost -p
```
