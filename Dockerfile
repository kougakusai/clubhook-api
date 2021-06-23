# 開発用ステージ
FROM golang:1.16 AS dev
WORKDIR /go/src/app
COPY . .
# 使用するモジュールのインストール
RUN go install github.com/labstack/echo && \
  go install gorm.io/gorm && \
  go install github.com/99designs/gqlgen@latest && \
  go install github.com/go-delve/delve/cmd/dlv@latest && \
  go install golang.org/x/tools/gopls@latest

# ビルド用ステージ
FROM golang:1.16 AS build
WORKDIR /go/src/app
# Herokuのrelease phaseでログを残すためにcurlを入れる
RUN apt-get update && apt-get install curl
COPY . .
RUN go mod download && CGO_ENABLED=0 GOOS=linux go build -o main main.go

# 実行用ステージ
FROM scratch AS prod
COPY --from=build /go/src/app/main /bin/main
ENTRYPOINT ["/bin/main"]
