// Code generated by go-swagger; DO NOT EDIT.

package ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GetTasksHandlerFunc turns a function with the right signature into a get tasks handler
type GetTasksHandlerFunc func(GetTasksParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn GetTasksHandlerFunc) Handle(params GetTasksParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// GetTasksHandler interface for that can handle valid get tasks params
type GetTasksHandler interface {
	Handle(GetTasksParams, interface{}) middleware.Responder
}

// NewGetTasks creates a new http.Handler for the get tasks operation
func NewGetTasks(ctx *middleware.Context, handler GetTasksHandler) *GetTasks {
	return &GetTasks{Context: ctx, Handler: handler}
}

/*
	GetTasks swagger:route GET /tasks getTasks

# Return a list of tasks

Get a list of tasks for current user
*/
type GetTasks struct {
	Context *middleware.Context
	Handler GetTasksHandler
}

func (o *GetTasks) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewGetTasksParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
