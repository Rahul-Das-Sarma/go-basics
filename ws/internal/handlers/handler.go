package handlers

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)

	if err != nil {
		log.Println(err)

	}
}

var WsChannel = make(chan WsPayload)

var clients = make(map[*WsConnection]string)

type WsConnection struct {
	*websocket.Conn
}
type WsJsonResponse struct {
	Action         string   `json:"action"`
	Message        string   `json:"message"`
	MessageType    string   `json:"messageType"`
	ConnectedUsers []string `json:"connected_users"`
}

type WsPayload struct {
	Action   string        `json:"action"`
	Username string        `json:"username"`
	Message  string        `json:"message"`
	Conn     *WsConnection `json:"-"`
}

func WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	var response WsJsonResponse
	response.Message = `<em><strong>Connected to server</strong></em>`

	conn := WsConnection{Conn: ws}
	clients[&conn] = ""

	err = ws.WriteJSON((response))
	if err != nil {
		log.Println(err)
	}
	go ListenForWs(&conn, WsChannel)
}

func ListenToWsChannel() {
	var response WsJsonResponse
	for {
		e := <-WsChannel

		switch e.Action {
		case "username":
			clients[e.Conn] = e.Username
			users := getUserList()
			response.Action = "list_users"
			response.ConnectedUsers = users
			broadCastToAll(response)
		}
	}
}

func getUserList() []string {
	var userList []string
	for _, x := range clients {
		if x != "" {
			userList = append(userList, x)
		}
	}
	sort.Strings(userList)
	return userList
}
func broadCastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			log.Println("websocket err", err)
			_ = client.Close()
			delete(clients, client)
		}
	}
}
func ListenForWs(conn *WsConnection, messages chan<- WsPayload) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// clients[conn] = ""
			// delete(clients, conn)
			// break
		} else {

			payload.Conn = conn
			WsChannel <- payload
		}
	}
}

func renderPage(w http.ResponseWriter, templ string, data jet.VarMap) error {
	view, err := views.GetTemplate(templ)
	if err != nil {
		log.Println(err)
		return err
	}
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
