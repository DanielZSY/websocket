FROM golang:1.19 AS build-dist

WORKDIR /chat/release

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o websocket main.go


FROM centos:latest as prod

WORKDIR /chat/service

COPY --from=build-dist /chat/release /chat/service
COPY --from=build-dist /chat/release/conf /chat/service/conf

EXPOSE 7990 7991

CMD ["/chat/service/websocket","-c","./conf/app.ini"]
