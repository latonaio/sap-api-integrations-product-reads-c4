# syntax = docker/dockerfile:experimental
# Build Container
FROM golang:1.18 as builder

ENV GO111MODULE on
ENV GOPRIVATE=github.com/latonaio
WORKDIR /go/src/github.com/latonaio

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o sap-api-integrations-product-reads-c4

# Runtime Container
FROM alpine:3.14
RUN apk add --no-cache libc6-compat
ENV SERVICE=sap-api-integrations-product-reads-c4 \
    APP_DIR="${AION_HOME}/${POSITION}/${SERVICE}"

WORKDIR ${AION_HOME}

COPY --from=builder /go/src/github.com/latonaio/sap-api-integrations-product-reads-c4 .
COPY --from=builder /go/src/github.com/latonaio/sample.json .

CMD ["./sap-api-integrations-product-reads-c4"]