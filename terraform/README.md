# golang aws sqs example terraform


## リソースの作成
```
$ export BUCKET_NAME="your s3 bukect name tfstate"
$ export BUCKET_REGION="your s3 bucket region for tfstate
$ bin/apply $BUCKET_NAME $BUCKET_REGION
```

## リソースの削除
```
# IAMユーザーのアクセスキーを全て削除してること
$ export BUCKET_NAME="your s3 bukect name tfstate"
$ export BUCKET_REGION="your s3 bucket region for tfstate
$ bin/destory $BUCKET_NAME $BUCKET_REGION
```
