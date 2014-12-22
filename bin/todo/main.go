package main

import (
	"github.com/gorilla/mux"
	"gopkg.in/unrolled/render.v1"
	"net/http"
	"strconv"
	"todo"
)

var (
	tService todo.TasksService
)

func main() {
	c := &TasksController{Render: render.New(render.Options{})}

	r := mux.NewRouter().StrictSlash(false)

	// Tasks Collection
	tasks := r.Path("/").Subrouter()
	tasks.Methods("GET").Handler(c.Action(c.Index))
	tasks.Methods("POST").Handler(c.Action(c.Create))

	// Task Singular
	task := r.PathPrefix("/{id:[0-9]+}").Subrouter()
	task.Methods("GET").Handler(c.Action(c.Show))
	task.Methods("PUT", "POST").Handler(c.Action(c.Update))
	task.Methods("DELETE").Handler(c.Action(c.Delete))

	http.ListenAndServe(":8080", r)
}

type Action func(rw http.ResponseWriter, r *http.Request) error

// This is our Base Controller
type AppController struct{}

// The action function helps with error handling in a controller
func (c *AppController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

type TasksController struct {
	AppController
	*render.Render
}

func (c *TasksController) Index(rw http.ResponseWriter, r *http.Request) error {
	tasks, _ := tService.GetAll()
	c.JSON(rw, 200, tasks)
	return nil
}

func (c *TasksController) Show(rw http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.Atoi(mux.Vars(r)["id"])
	task, _ := tService.Get(id)

	c.JSON(rw, 200, task)
	return nil
}

func (c *TasksController) Create(rw http.ResponseWriter, r *http.Request) error {
	tasks, _ := tService.GetAll()
	c.JSON(rw, 200, tasks)
	return nil
}

func (c *TasksController) Update(rw http.ResponseWriter, r *http.Request) error {
	tasks, _ := tService.GetAll()
	c.JSON(rw, 200, tasks)
	return nil
}

func (c *TasksController) Delete(rw http.ResponseWriter, r *http.Request) error {
	tasks, _ := tService.GetAll()
	c.JSON(rw, 200, tasks)
	return nil
}
