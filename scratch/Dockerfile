FROM alpine as builder
RUN apk add --no-cache \
      ca-certificates \
      tzdata
RUN update-ca-certificates
RUN adduser --system --disabled-password --no-create-home --uid 1000 user

FROM scratch
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
USER user
