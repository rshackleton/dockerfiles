# Healthcheck
A Super Simple Go Lang Application For Monitoring HTTP/HTTPS Endpoints Using The Dockerfile `HEALTHCHECK` Instruction.

# Usage
Add the following to your Dockerfile to get started:
```bash
RUN wget -qO /bin/healthcheck $(wget -qO- https://api.github.com/repos/oscartbeaumont/dockerfiles/releases/latest | grep "healthcheck-linux-amd64" | grep browser_download_url | cut -d '"' -f 4) && chmod +x /bin/healthcheck
HEALTHCHECK CMD ["/bin/healthcheck", "-ignore-tls", "http://localhost:8080", "https://localhost:8443"]
```
You can add as many urls to the command as you want and you can remove the `-ignore-tls` flag if your https urls all have valid https certificates.
