package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"project/api-gateway/database"
	"project/api-gateway/helper"
	"project/api-gateway/model"
	"project/api-gateway/service"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type ChatController struct {
	service service.Service
	logger  *zap.Logger
	rdb     database.Cacher
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins. Modify as per your security needs.
		return true
	},
}
var broadcast = make(chan string)

func NewChatController(service service.Service, logger *zap.Logger, rdb database.Cacher) *ChatController {
	return &ChatController{service, logger, rdb}
}

func (ctrl *ChatController) Websocket(c *gin.Context) {
	var message model.Message
	username := c.MustGet("email").(string)
	if username == "" {
		BadResponse(c, "unauthorized", http.StatusUnauthorized)
	}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()
	roomId := c.Param("id")
	uintRoomId, err := helper.Uint(roomId)
	message.Sender = username
	message.RoomId = uintRoomId
	if err != nil {
		log.Println("ERROR PARSING UINT")
		return
	}
	pubsub := ctrl.rdb.Subcribe("room:" + roomId)
	defer pubsub.Close()
	go func() {
		for {
			payload, err := pubsub.ReceiveMessage(context.Background())
			if err != nil {
				log.Println("Failed Received Message Redis: ", err)
				return
			}
			// log.Println(payload.Payload, "++++++")
			// log.Printf("%+v<<<<<<--------\n", message)
			err = conn.WriteMessage(websocket.TextMessage, []byte(payload.Payload))
			if err != nil {
				log.Println("Write error:", err)
				break
			}
		}
	}()
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}
		chat := string(msg)
		err = json.Unmarshal(msg, &message)
		if err != nil {
			log.Println(err)
			return
		}
		err = ctrl.service.Chat.SaveMessage(&message)
		if err != nil {
			log.Println(err, "DARI SAVE MESSAGE HANDLER")
			return
		}

		ctrl.rdb.Publish("room:"+roomId, chat)
	}
}
func (ctrl *ChatController) GetRoomMessages(c *gin.Context) {
	query := c.Query("page")
	var page uint
	var err error
	if query != "" {
		page, err = helper.Uint(query)
		if err != nil {
			BadResponse(c, err.Error(), http.StatusBadRequest)
			return
		}
	}
	param := c.Param("id")
	roomId, err := helper.Uint(param)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := ctrl.service.Chat.GetRoomMessages(roomId, page)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get History Message Succes", http.StatusOK, res)
}
func (ctrl *ChatController) GetAllParticipants(c *gin.Context) {
	param := c.Param("id")
	roomId, err := helper.Uint(param)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := ctrl.service.Chat.GetRoomParticipants(roomId)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get Participants Success", http.StatusOK, res)
}
func (ctrl *ChatController) AddParticipants(c *gin.Context) {
	param := c.Param("id")
	roomId, err := helper.Uint(param)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	var participant model.Participant
	if err := c.ShouldBindJSON(&participant); err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := ctrl.service.Chat.AddRoomParticipant(uint64(roomId), participant.Email)
	if err != nil {
		BadResponse(c, err.Error(), http.StatusBadRequest)
		return
	}
	GoodResponseWithData(c, "Get Participants Success", http.StatusOK, res)
}
