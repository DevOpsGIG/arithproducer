FROM golang:1.8-alpine

COPY . src/github.com/devopsgig/arithproducer

WORKDIR src/github.com/devopsgig/arithproducer

RUN go build

CMD ["./arithproducer", "run"]
