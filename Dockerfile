# 2021/02/28最新version
FROM golang:1.16.0

# コンテナ内にディレクトリを作成
RUN mkdir -p /go/src/github.com/kougakusaiHPTeam/clubhook-api

# コンテナログイン時のディレクトリを設定
WORKDIR /go/src/github.com/kougakusaiHPTeam/clubhook-api

# ホストのファイルをコンテナにコピー
ADD . /go/src/github.com/kougakusaiHPTeam/clubhook-api

# 使用するモジュールのインストール
RUN go get github.com/labstack/echo && \
  go get github.com/jinzhu/gorm
