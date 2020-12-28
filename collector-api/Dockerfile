FROM golang:1.15.6-alpine

LABEL maintainer="Vague Coder <vaguecoder0to.n@gmail.com>"

WORKDIR $GOPATH/src/github.com/VagueCoder/Amazon-Scraper-Collector/collector-api

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8081

CMD ["collector-api"]