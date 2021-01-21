FROM golang:alpine
RUN apk add --update \
    curl \
    && rm -rf /var/cache/apk/*
RUN mkdir /go/src/client-api
WORKDIR /go/src/client-api
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go build -o /main cmd/main.go

EXPOSE 50051
CMD ["/main"]




