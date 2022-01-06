FROM golang:1.17-alpine as builder
WORKDIR /go/src/scooter_service
COPY ./scooter_service .
ENV GO111MODULE=on
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/scooter_service

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
WORKDIR /usr/bin
COPY --from=builder /go/src/scooter_service/build/scooter_service ./scooter_service
ENTRYPOINT [ "/usr/bin/scooter_service" ]