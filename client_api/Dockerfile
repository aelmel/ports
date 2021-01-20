FROM golang:alpine AS build
RUN mkdir /go/src/client-api
ADD ./ /go/src/client-api
WORKDIR /go/src/client-api
RUN apk update && apk add --no-cache git && apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main cmd/main.go

FROM scratch AS runtime
COPY --from=build /main /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8000
CMD ["/main"]


