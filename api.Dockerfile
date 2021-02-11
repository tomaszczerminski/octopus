FROM golang:1.15 AS builder

COPY ${PWD} /app
WORKDIR /app

RUN CGO_ENABLED=0 go build -ldflags '-s -w -extldflags "-static"' -o /app/api *.go

FROM scratch
LABEL MAINTAINER = czerminski.tomasz@gmail.com

COPY --from=builder /app /app

WORKDIR /app
EXPOSE 8080

CMD ["./api"]