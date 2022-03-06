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

	// make a list of Users to do the operation instead of db
	User1 := domains.User{
		ID:         "11aa",
		FirstName:  "Peter",
		LastName:   "Golm",
		Email:      "peter_golm@visable.com",
	}

	User2 := domains.User{
		ID:         "33cc",
		FirstName:  "Andreas",
		LastName:   "Litt",
		Email:      "Andreas_Litt@visable.com",
	}

	// make a list of tasks to do the operation instead of db
	task1 := domains.Task{
		ID:           "11bb",
		Name:         "task management",
		Owner:        domains.User{},
		State:        domains.EnumtaskStatePlanned,
	}

	tasksDB := make(map[string]domains.Task)
	UsersDB := make(map[string]domains.User)

	UsersDB["11aa"] = User1
	UsersDB["33cc"] = User2

	tasksDB["11bb"] = task1

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
