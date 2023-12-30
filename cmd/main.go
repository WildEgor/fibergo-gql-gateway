package main

import (
	server "github.com/WildEgor/fibergo-gql-gateway/internal"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	Start()
	Shutdown()
}

func Start() {
	srv, _ := server.NewServer()
	log.Printf("Server is listening on PORT: %s", srv.AppConfig.Port)

	addr := ":" + srv.AppConfig.Port

	if err := srv.App.Listen(addr); err != nil {
		log.Panicf("[CRIT] Unable to start server. Reason: %v", err)
	}
}

func Shutdown() {
	// block main thread - wait for shutdown signal
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		log.Println()
		log.Println(sig)
		done <- true
	}()

	log.Println("[Main] Awaiting signal")
	<-done
	log.Println("[Main] Stopping consumer")
}
