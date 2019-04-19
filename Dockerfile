FROM golang:1.12.4 AS builder
ENV GO111MODULE=on
ADD src /go/src/cropse/IPfinder
WORKDIR /go/src/cropse/IPfinder
RUN wget -qO- https://geolite.maxmind.com/download/geoip/database/GeoLite2-City.tar.gz | tar --wildcards -zxvO GeoLite2-City_*/GeoLite2-City.mmdb > GeoLite2-City.mmdb
RUN go mod tidy &&\
    go test &&\
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .

FROM alpine:3.9
# newer version tar needed
RUN mkdir /app && apk add tar
WORKDIR /app

COPY --from=builder /go/src/cropse/IPfinder/GeoLite2-City.mmdb /app/GeoLite2-City.mmdb
COPY --from=builder /go/src/cropse/IPfinder/app /app/app
EXPOSE 7000
ENTRYPOINT ["./app"]