package routes

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/api-gateway/infra"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func NewRoutes(ctx infra.ServiceContext) {
	r := gin.Default()

	r.POST("/register", ctx.Ctl.AuthHandler.Register)
	r.POST("/login", ctx.Ctl.AuthHandler.Login)
	r.PUT("/otp/:id", ctx.Ctl.AuthHandler.ValidateOtp)

	r.Use(ctx.Middleware.Auth())
	r.GET("/users", ctx.Ctl.UserHandler.GetAllUsers)
	r.PUT("/profile", ctx.Ctl.UserHandler.UpdateProfile)

	contactRoutes := r.Group("/user/contacts")
	{
		contactRoutes.POST("/", ctx.Ctl.ContactHandler.Add)
		contactRoutes.DELETE("/", ctx.Ctl.ContactHandler.Remove)
	}

	chatRoutes := r.Group("/user/chats")
	{
		chatRoutes.GET("/:id/ws", ctx.Ctl.ChatHandler.Websocket)
		chatRoutes.GET("/:id/messages", ctx.Ctl.ChatHandler.GetRoomMessages)
		chatRoutes.GET("/:id/participants", ctx.Ctl.ChatHandler.GetAllParticipants)
		chatRoutes.POST("/:id/participants", ctx.Ctl.ChatHandler.AddParticipants)
	}

	gracefulShutdown(ctx, r.Handler())
}

func gracefulShutdown(ctx infra.ServiceContext, handler http.Handler) {
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", ctx.Cfg.ServerIp, ctx.Cfg.ServerPort),
		Handler: handler,
	}

	if ctx.Cfg.ShutdownTimeout == 0 {
		launchServer(srv, ctx.Cfg.ServerPort)
		return
	}

	go func() {
		launchServer(srv, ctx.Cfg.ServerPort)
	}()

	// Wait for interrupt signal to gracefully shut down the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	appContext, cancel := context.WithTimeout(context.Background(), time.Duration(ctx.Cfg.ShutdownTimeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(appContext); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching appContext.Done(). timeout of ShutdownTimeout seconds.
	select {
	case <-appContext.Done():
		log.Println(fmt.Sprintf("timeout of %d seconds.", ctx.Cfg.ShutdownTimeout))
	}
	log.Println("Server exiting")
}

func launchServer(server *http.Server, port string) {
	// service connections
	log.Println("Listening and serving HTTP on", port)
	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("listen: %s\n", err)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow all origins. Modify as per your security needs.
		return true
	},
}
