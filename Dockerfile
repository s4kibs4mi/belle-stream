FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN mkdir -p /app
WORKDIR /app
COPY build/belle-stream /app
ENTRYPOINT ["./belle-stream"]
