## twitter timeline
タイムラインを表示するコマンドラインツールです。

### 使い方
ツイッターから取得したAPIキーとアクセストークンを環境変数に設定して利用します。

```bash:setup.sh
#!/bin/bash
export CONSUMER_KEY="CONSUMER_KEY"
export CONSUMER_SECRET="SONSUMER_SECRET"
export ACCESS_TOKEN="ACCESS_TOKEN"
export ACCESS_TOKEN_SECRET="ACCESS_TOKEN_SECRET" 
```

後は必要な外部パッケージを`go get`し、
```
$go get github.com/dghubble/oauth1
```
ビルドで
```
$go build
```
使えます。
```
$./twitter_timeline
```
