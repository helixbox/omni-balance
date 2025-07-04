FROM golang:1.24.0-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG VERSION
ARG COMMIT_MESSAGE
ARG COMMIT_TIME

COPY . .
# garble build
RUN CGO_ENABLED=0 go build -a -ldflags \
            "-X 'main.version=$VERSION' \
            -X 'main.commitMessage=$COMMIT_MESSAGE' \
            -X 'main.commitTime=$COMMIT_TIME' -s -w -extldflags '-static'" \
        -o omni-balance ./cmd

FROM gruebel/upx:latest AS upx
COPY --from=builder /app/omni-balance /omni-balance
RUN upx --best --lzma /omni-balance

FROM alpine:latest AS prod
COPY --from=upx /omni-balance /omni-balance
ENTRYPOINT ["/omni-balance"]
