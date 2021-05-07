# 2021/02/28最新version
FROM golang:1.16.0

# コンテナ内にディレクトリを作成
RUN mkdir -p /go/src/github.com/kougakusai/clubhook-backend

# コンテナログイン時のディレクトリを設定
WORKDIR /go/src/github.com/kougakusai/clubhook-backend

# ホストのファイルをコンテナにコピー
ADD . /go/src/github.com/kougakusai/clubhook-backend

# 使用するモジュールのインストール
RUN go get github.com/labstack/echo && go get gorm.io/gorm && go get github.com/99designs/gqlgen
