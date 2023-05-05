package routers

import (
	"imservice/api/bind2group"
	"imservice/api/closeclient"
	"imservice/api/getonlinelist"
	"imservice/api/register"
	"imservice/api/send2client"
	"imservice/api/send2clients"
	"imservice/api/send2group"
	"net/http"

	"imservice/servers"
)

func Init() {
	registerHandler := &register.Controller{}
	sendToClientHandler := &send2client.Controller{}
	sendToClientsHandler := &send2clients.Controller{}
	sendToGroupHandler := &send2group.Controller{}
	bindToGroupHandler := &bind2group.Controller{}
	getGroupListHandler := &getonlinelist.Controller{}
	closeClientHandler := &closeclient.Controller{}

	http.HandleFunc("/api/register", registerHandler.Run)
	http.HandleFunc("/api/send_to_client", AccessTokenMiddleware(sendToClientHandler.Run))
	http.HandleFunc("/api/send_to_clients", AccessTokenMiddleware(sendToClientsHandler.Run))
	http.HandleFunc("/api/send_to_group", AccessTokenMiddleware(sendToGroupHandler.Run))
	http.HandleFunc("/api/bind_to_group", AccessTokenMiddleware(bindToGroupHandler.Run))
	http.HandleFunc("/api/get_online_list", AccessTokenMiddleware(getGroupListHandler.Run))
	http.HandleFunc("/api/close_client", AccessTokenMiddleware(closeClientHandler.Run))

	servers.StartWebSocket()

	go servers.WriteMessage()
}
