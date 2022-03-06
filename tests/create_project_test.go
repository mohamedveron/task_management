package tests

import (
	"github.com/mohamedveron/task_management/domains"
	"github.com/mohamedveron/task_management/service"
	"testing"
)

	var tasksDB = make(map[string]domains.task)
	var UsersDB = make(map[string]domains.User)


	var User1 = domains.User{
		ID:         "11aa",
		FirstName:  "Peter",
		LastName:   "Golm",
		Role:       "manager",
		Email:      "peter_golm@visable.com",
		Department: "Engineering",
	}

	var User2 = domains.User{
		ID:         "33cc",
		FirstName:  "Andreas",
		LastName:   "Litt",
		Role:       "developer",
		Email:      "Andreas_Litt@visable.com",
		Department: "Engineering",
	}

	var User3 = domains.User{
		ID:         "58kk",
		FirstName:  "Nina",
		LastName:   "Wessels",
		Role:       "hr",
		Email:      "Nina_Wessels@visable.com",
		Department: "HR",
	}

	var User4 = domains.User{
		ID:         "t3t3",
		FirstName:  "Mohamed",
		LastName:   "Abdel Mohaimen",
		Role:       "developer",
		Email:      "Mohamed_veron@gmail.com",
		Department: "Engineering",
	}

	// make a list of tasks to do the operation instead of db
	var task1 = domains.task{
		ID:           "11bb",
		Name:         "task management",
		Owner:        domains.User{},
		Progress:     0,
		State:        domains.EnumtaskStatePlanned,
		Participants: []domains.User{},
	}


func TestCreatetask(t *testing.T){

	task := domains.task{
		ID:           "",
		Name:         "task1",
		Owner:        domains.User{},
		Progress:     0,
		State:        "planned",
		Participants: nil,
	}

	service := service.NewService(tasksDB, UsersDB)

	_ , err := service.Createtask(task)

	if err != nil {
		t.Error("can't create task")
	}
}
