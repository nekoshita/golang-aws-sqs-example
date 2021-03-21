# golang-aws-sqs-example

## これは何？
AWSのSQSの実装例
https://aws.amazon.com/jp/sqs/

## 事前準備
- go v1.16.2
- terraform
- AWSアカウント
  - S3バケット

## リソースの作成
```
$ cd terraform
```
残りの手順は[こちら](./terraform/README.md#リソースの作成)を参照

## リソースの削除
```
$ cd terraform
```
残りの手順は[こちら](./terraform/README.md#リソースの削除)を参照

## 実行
```
# 事前に[リソースを作成](#リソースの作成)をしておく
# AWSコンソールでIAMユーザーのアクセスキーIDとシークレットアクセスキーを発行しておく
$ export AWS_ACCESS_KEY_ID=xxxx
$ export AWS_SECRET_ACCESS_KEY=xxxx
$ go run main.go -sqs-queue-url "your-sqs-url" -sqs-region "your-sqs-region"
```

## AWS SQSについて
- Messageの内容を既読にせずに覗き見ることはできるのか？
  - できない
  - https://stackoverflow.com/questions/10100522/how-to-peek-at-messages-in-the-queue

- MessageGroupIDとは
  - Messages that belong to the same message group are always processed one by one, in a strict order relative to the message group
  - 同じグループに所属するメッセージは1キューずつしか取り出せない

- Messageの最大保有期間はあるか？
  - ある
  - https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/sqs_queue#message_retention_seconds
  - デフォルトは4日間
  - 最短1分から最長14日

- 可視性タイムアウトとは？
  - 一度キューを取得すると、そのキューは「処理中」のようなステータスになる
  - 「処理中」の間は、そのキューは他の誰にも処理されない
  - デフォルトは30秒間
  - 最短0秒から最長12時間

- 動作確認
  - キューをsendすると、「利用可能なメッセージ」に1追加される
  - キューをrecieveする以外で、キューのmessageを見る方法はない
  - キューをrecieveすると、「処理中のメッセージ」ってステータスになる
  - 「処理中のメッセージ」のキューは、他のでrevieveしてもrecieveされることはない
  - `visibility_timeout_seconds`がすぎたら、「利用可能なメッセージ」に戻る
  - `maxReceiveCount`を超えると、deadletterキューにメッセージが移動される
