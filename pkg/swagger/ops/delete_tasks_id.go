// Code generated by go-swagger; DO NOT EDIT.

package ops

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DeleteTasksIDHandlerFunc turns a function with the right signature into a delete tasks ID handler
type DeleteTasksIDHandlerFunc func(DeleteTasksIDParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn DeleteTasksIDHandlerFunc) Handle(params DeleteTasksIDParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// DeleteTasksIDHandler interface for that can handle valid delete tasks ID params
type DeleteTasksIDHandler interface {
	Handle(DeleteTasksIDParams, interface{}) middleware.Responder
}

// NewDeleteTasksID creates a new http.Handler for the delete tasks ID operation
func NewDeleteTasksID(ctx *middleware.Context, handler DeleteTasksIDHandler) *DeleteTasksID {
	return &DeleteTasksID{Context: ctx, Handler: handler}
}

/*
	DeleteTasksID swagger:route DELETE /tasks/{id} deleteTasksId

Delete task by id
*/
type DeleteTasksID struct {
	Context *middleware.Context
	Handler DeleteTasksIDHandler
}

func (o *DeleteTasksID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewDeleteTasksIDParams()
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
