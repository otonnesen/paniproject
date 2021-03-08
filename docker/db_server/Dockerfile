FROM golang:1.16

WORKDIR /go/src/app

ENV PORT=3000

COPY ./src/ ./

RUN go build

CMD ["./db_server"]
