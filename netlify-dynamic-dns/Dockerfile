FROM golang:alpine as builder
RUN adduser -D -g '' go
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
WORKDIR /go/src/github.com/oscartbeaumont/netlify-dynamic-dns
RUN go get -u github.com/pkg/errors
COPY cmd ./cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o netlify-dynamic-dns ./cmd

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /go/src/github.com/oscartbeaumont/netlify-dynamic-dns/netlify-dynamic-dns /
USER go
CMD ["/netlify-dynamic-dns"]
