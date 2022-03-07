package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/mohamedveron/task_management/api"
	"github.com/mohamedveron/task_management/domains"
	"github.com/mohamedveron/task_management/service"
	"net/http"
)


func main() {

	// make a list of tasks to do the operation instead of db
	tasksDB := make(map[string]domains.Task)

	// make a list of Users to do the operation instead of db
	UsersDB := make(map[string]domains.User)

	serviceLayer := service.NewService(tasksDB, UsersDB)
	server := api.NewServer(serviceLayer)

	// prepare router
	r := chi.NewRouter()
	r.Route("/", func(r chi.Router) {
		r.Mount("/", api.Handler(server))
	})

	srv := &http.Server{
		Handler: r,
		Addr:    ":8080",
	}

	fmt.Println("server starting on port 8080...")

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Println("server failed to start", "error", err)
	}

}
