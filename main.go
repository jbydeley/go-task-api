// Copyright 2014 The cli Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/gorilla/mux"
	tasks "github.com/jbydeley/go-task-lib"
	"gopkg.in/unrolled/render.v1"
	"net/http"
	"strconv"
)

var tService tasks.TasksService

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
type BaseController struct{}

// The action function helps with error handling in a controller
func (c *BaseController) Action(a Action) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Access-Control-Allow-Origin", "http://localhost:8081")
		if err := a(rw, r); err != nil {
			http.Error(rw, err.Error(), 500)
		}
	})
}

type TasksController struct {
	BaseController
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
