FROM golang:1.19.1-alpine

RUN apk update && apk add git
WORKDIR /app
# ホストのファイルをコンテナの作業ディレクトリに移行
COPY . /app

# パッケージのインポート
RUN go get -u golang.org/x/tools/cmd/goimports