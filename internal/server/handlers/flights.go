package handlers

import (
	"errors"
	"server/internal/server/restapi/operations"
	"server/pkg/models/rest"

	"github.com/go-openapi/runtime/middleware"
)

type Flight struct {
	From, To string
}

var (
	ErrPathNotLinked = errors.New("flights path is not linked")
	ErrEmptyPath     = errors.New("no flights in request")
	ErrPathIsCircle  = errors.New("path is a circle, impossible to find start and finish")
)

func GetFlightPath(params operations.GetPathParams) middleware.Responder {
	flights := make([]Flight, 0, len(params.Body.Flights))

	for _, f := range params.Body.Flights {
		flights = append(flights, Flight{
			From: f.From,
			To:   f.To,
		})
	}

	path, err := findPath(flights)
	if err != nil {
		return operations.NewGetPathBadRequest().WithPayload(&rest.Error{
			Code:    400,
			Message: err.Error(),
		})
	}

	return operations.NewGetPathOK().WithPayload(&rest.PathResponse{
		Path: &rest.Flight{From: path.From, To: path.To},
	})
}

func findPath(flights []Flight) (path Flight, err error) {
	if len(flights) == 0 {
		return path, ErrEmptyPath
	}
	flightsMap := make(map[string]int)

	for _, f := range flights {
		flightsMap[f.From]--
		flightsMap[f.To]++
	}

	for loc, visits := range flightsMap {
		switch visits {
		case 0:
			continue
		case 1:
			// there should only one with +1 visit
			if path.To != "" {
				return path, ErrPathNotLinked
			}
			path.To = loc
		case -1:
			// there should only one with +1 visit
			if path.From != "" {
				return path, ErrPathNotLinked
			}
			path.From = loc
		default:
			return path, ErrPathNotLinked
		}
	}

	if path.To == "" || path.From == "" {
		return path, ErrPathIsCircle
	}

	return path, nil
}
