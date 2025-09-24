FROM golang:1.21-alpine3.20

WORKDIR /app

RUN go mod init example.com/gorilla && go get -u github.com/gorilla/mux

COPY main.go main.go

CMD [ "go", "run", "main.go" ]