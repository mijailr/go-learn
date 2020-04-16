FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git curl

RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.4/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/mijailr/go-learn
WORKDIR $GOPATH/src/github.com/mijailr/go-learn

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure -vendor-only

COPY . .
RUN GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /go/bin/go-learn cmd/main.go

FROM scratch
COPY --from=builder /go/bin/go-learn ./go-learn
ENTRYPOINT ["./go-learn"]