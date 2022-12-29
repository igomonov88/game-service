FROM golang:1.19.2-alpine as builder
ENV CGO_ENABLED 0
ARG BUILD_REF

RUN mkdir /service
COPY . /service/
COPY etc/config.yaml /service/etc/config.yaml
WORKDIR /service
RUN go mod download

WORKDIR /service/app/tooling/admin
RUN go build -ldflags "-X main.build=${BUILD_REF}"

WORKDIR /service/app/service/api
RUN go build -ldflags "-X main.build=${BUILD_REF}"

FROM alpine:3.16

COPY --from=builder /service/etc/config.yaml /service/etc/config.yaml
COPY --from=builder /service/app/tooling/admin/admin /service/admin
COPY --from=builder /service/app/service/api/api /service/api

WORKDIR /service
EXPOSE 8080
CMD ["./api"]