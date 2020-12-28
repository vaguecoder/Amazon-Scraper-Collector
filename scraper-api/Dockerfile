FROM golang:1.15.6-alpine

LABEL maintainer="Vague Coder <vaguecoder0to.n@gmail.com>"

WORKDIR $GOPATH/src/github.com/VagueCoder/Amazon-Scraper-Collector/scraper-api

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 8080

CMD ["scraper-api"]