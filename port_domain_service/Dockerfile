FROM golang:alpine AS build
RUN mkdir /go/src/port_domain_service
ADD ./ /go/src/port_domain_service
WORKDIR /go/src/port_domain_service
RUN apk update && apk add --no-cache git && apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /main cmd/main.go

FROM scratch AS runtime
COPY --from=build /main /
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 50051
CMD ["/main"]


