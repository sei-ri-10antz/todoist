package main

import (
	"context"

	"github.com/sei-ri-10antz/todoist/http"
)

func main() {
	srv := http.Server{
		Host: "0.0.0.0",
		Port: 8888,
	}

	srv.Serve(context.Background())
}
