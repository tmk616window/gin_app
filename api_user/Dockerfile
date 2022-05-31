# ベースとなるDockerイメージ指定
FROM golang:latest
# コンテナ内に作業ディレクトリを作成
# RUN mkdir /go/src
# コンテナログイン時のディレクトリ指定
WORKDIR /go/api_user

# GoApp起動
# CMD ["go","run","main.go"]
# ホストのファイルをコンテナの作業ディレクトリに移行
ADD . /go/api_user
