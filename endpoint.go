package todo

import "github.com/go-kit/kit/endpoint"

// Endpoints collects all endpoints which compose the Todo service
type TodoEndpoints struct {
	GetAllForUserEndPoint endpoint.Endpoint
	GetByIDEndpoint       endpoint.Endpoint
	AddEndpoint           endpoint.Endpoint
	UpdateEndpoint        endpoint.Endpoint
	DeleteEndpoint        endpoint.Endpoint
}
