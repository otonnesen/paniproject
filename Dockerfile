FROM golang:1.16

WORKDIR /go/src/app

COPY ./src/ ./

RUN go build

CMD ["./paniproject"]
