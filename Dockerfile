FROM golang:alpine as builder

ENV USER=appuser
ENV UID=1000
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

WORKDIR /usr/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -buildvcs=false -o /usr/local/bin/app


FROM alpine AS final
LABEL author="Cajually <me@caj.me>"


COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /usr/local/bin/app /usr/local/bin/app

WORKDIR /

USER appuser:appuser
ENTRYPOINT ["/usr/local/bin/app"]

EXPOSE 8080
