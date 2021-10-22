# ベースとなるDockerイメージ指定
FROM golang:latest

ENV GO111MODULE=on
# コンテナ内に作業ディレクトリを作成
RUN mkdir /go/src/work
# コンテナログイン時のディレクトリ指定
WORKDIR /go/src/work
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/src/work

RUN go get github.com/go-sql-driver/mysql 
RUN go get github.com/labstack/echo/middleware 
RUN go get github.com/jinzhu/gorm 

