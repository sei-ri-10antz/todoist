package http

import (
	"context"
	"log"
	"net"
	"net/http"
	"strconv"

	"github.com/sei-ri-10antz/todoist/db"
)

type Server struct {
	Host string
	Port int
}

func (s *Server) Serve(ctx context.Context) {
	db, err := db.NewClient(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// new http handlers
	handler := NewRouter(Service{
		Store: &Store{
			UsersStore: db.UsersStore(),
			TasksStore: db.TasksStore(),
		},
	})

	// new http server
	srv := &http.Server{
		Handler: handler,
	}

	// listening
	lis, err := net.Listen("tcp", net.JoinHostPort(s.Host, strconv.Itoa(s.Port)))
	if err != nil {
		log.Fatal(err)
	}
	defer lis.Close()

	log.Println("Serving at", lis.Addr())

	srv.Serve(lis)
}
