FROM golang:latest

WORKDIR /app

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 80 81

ENTRYPOINT go mod init; CompileDaemon --build="go build -mod=mod god" --command=./god
