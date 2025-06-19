package http

import (
	"danmaku/danmaku_reply/model"
	"danmaku/danmaku_reply/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

type Server struct {
	service *service.Service
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// 允许跨域连接（可按需修改）
		return true
	},
}

func (s *Server) handleWebSocket(c *gin.Context) {
	platform := c.Query("platform")
	roomID := c.Query("room_id")
	userID := c.Query("user_id")

	if platform == "" || roomID == "" || userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Missing platform or room_id",
		})
		return
	}

	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	//s.service.
	if err := s.service.FetchRoomDanmaku(c, userID, roomID, platform); err != nil {
		log.Println("WebSocket fetch room danmaku failed:", err)
		return
	}
	s.service.AddWsConn(c, userID, platform, roomID, conn)
	defer s.service.Remo
	log.Printf("New WebSocket connection:userID=%s, platform=%s, room_id=%s\n", userID, platform, roomID)

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error:", err)
			break
		}

		log.Printf("Received message: %s", msg)
	}

}

func StartServer(s *service.Service, config *model.Config) (h *Server, err error) {
	h = &Server{
		service: s,
	}
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "pong",
		})
	})

	router.POST("/fetch/start", func(c *gin.Context) {
		var param model.FetchDanmakuParam
		if err := c.ShouldBindJSON(&param); err != nil {
			log.Printf("BindJSON err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := s.FetchRoomDanmaku(c, param.UserID, param.RoomID, param.Platform); err != nil {
			log.Printf("FetchRoomDanmaku err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{})
	})

	router.POST("/fetch/stop", func(c *gin.Context) {
		var param model.FetchDanmakuParam
		if err := c.ShouldBindJSON(&param); err != nil {
			log.Printf("BindJSON err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := s.StopFetchRoomDanmaku(c, param.UserID, param.RoomID, param.Platform); err != nil {
			log.Printf("FetchRoomDanmaku err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}

		c.JSON(http.StatusOK, gin.H{})
	})

	router.GET("/ws", h.handleWebSocket)

	g := router.Group("/room")
	g.GET("/list", func(c *gin.Context) {
		userIDStr := c.Query("user_id")
		userID, err := strconv.ParseInt(userIDStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		rooms, err := s.RoomsByUserID(c, userID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"data": gin.H{
				"rooms": rooms,
			},
		})
	})

	g.POST("/sub", func(c *gin.Context) {
		var param model.SubRoomParam
		if err := c.ShouldBindJSON(&param); err != nil {
			log.Printf("BindJSON err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := s.SubRoom(c, &param); err != nil {
			log.Printf("SubRoom err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	g.POST("/delete", func(c *gin.Context) {
		var param model.DeleteRoomParam
		if err := c.ShouldBindJSON(&param); err != nil {
			log.Printf("BindJSON err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := s.DeleteRoom(c, &param); err != nil {
			log.Printf("SubRoom err %v\n", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
	})

	go func() {
		err := router.Run(fmt.Sprintf("%s:%d", config.Http.Host, config.Http.Port))
		if err != nil {
			log.Printf("gin router Run err %v\n", err)
		}
	}()
	return
}

func (s *Server) Close() (err error) {
	return
}
