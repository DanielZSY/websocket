FROM alpine:latest

WORKDIR /app

COPY ./deploy/ ./

EXPOSE 7988

CMD ["./chatroom", "-f", "./config/config.yaml"]
