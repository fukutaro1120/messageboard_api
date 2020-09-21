package main

import (
	"context"
	"database/sql"
	"fmt"
	"message-board-api/infra/mysql"
	"message-board-api/router"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	db, err := sql.Open("mysql", "ito:pass1234@tcp(127.0.0.1:3306)/messageboard")
	if err != nil {
		fmt.Printf("%v", err)
	}
	routerConfig := &router.Config{
		MySQLArticleRepo: mysql.NewArticleRepository(db),
	}
	router := router.New(routerConfig)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080), // TODO: logger
		Handler: router,
	}

	go func() {
		fmt.Println("Server starting ...") // TODO: logger
		if err := srv.ListenAndServe(); err != nil {
			fmt.Printf("%v", err) // TODO: logger
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	fmt.Println("Server stopping ...") // TODO: logger

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Printf("Server shutdown: %v", err) // TODO: logger
	}
	fmt.Println("Server stopped") // TODO: logger
}
