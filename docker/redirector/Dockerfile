FROM golang:1.16

WORKDIR /go/src/app

ENV PORT=5001

COPY ./src/ ./

RUN go build

CMD ["./redirector"]
