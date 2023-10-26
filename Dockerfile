FROM golang:1.21-alpine AS builder
ARG VERSION

WORKDIR /bin

COPY / .
RUN apk add build-base && go mod download && go build -o tifling \
    -ldflags "-X main.Version=$VERSION" main.go

FROM scratch

COPY --from=builder /bin/tifling /

CMD [ "/tifling" ]
