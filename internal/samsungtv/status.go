package samsungtv

import (
	"github.com/go-openapi/runtime/middleware"
	models "github.com/home-IoT/api-samsungtv/gen/models"
	ops "github.com/home-IoT/api-samsungtv/gen/restapi/operations"
)

// GetStatus returns the status of the connection
func GetStatus(param ops.GetStatusParams) middleware.Responder {
	status := CheckConnection()
	resp := models.StatusResponse{Host: &configuration.TV.Host, Reachable: &status}
	return ops.NewGetStatusOK().WithPayload(&resp)
}
