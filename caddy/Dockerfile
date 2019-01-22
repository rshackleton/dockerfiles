FROM golang:alpine as builder
RUN apk add git
RUN go get -u -v -d github.com/caddyserver/builds
RUN go get -u -v -d github.com/mholt/caddy
WORKDIR /go/src/github.com/mholt/caddy
RUN git checkout $(git describe --tags `git rev-list --tags --max-count=1`)
WORKDIR /go/src/github.com/mholt/caddy/caddy
RUN go run build.go
RUN wget -qO /bin/healthcheck $(wget -qO- https://api.github.com/repos/oscartbeaumont/dockerfiles/releases/latest | grep "healthcheck-linux-amd64" | grep browser_download_url | cut -d '"' -f 4) && chmod +x /bin/healthcheck

FROM oscartbeaumont/scratch
COPY --from=builder /go/src/github.com/mholt/caddy/caddy/caddy /bin/caddy
ENV CADDYPATH=/root/.caddy
VOLUME /root/.caddy
EXPOSE 8080 8443
COPY --from=builder /bin/healthcheck /bin/healthcheck
HEALTHCHECK CMD ["/bin/healthcheck", "-ignore-tls", "http://localhost:8080", "https://localhost:8443"]
ENTRYPOINT ["/bin/caddy"]
CMD ["-conf=/etc/Caddyfile", "-log=stdout", "-agree=true", "-http-port=8080", "-https-port=8443", "-http2"]
