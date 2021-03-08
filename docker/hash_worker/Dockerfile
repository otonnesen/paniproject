FROM golang:1.16

WORKDIR /go/src/app

ENV PORT=5000

COPY ./src/ ./

RUN go build

CMD ["./hash_worker"]
