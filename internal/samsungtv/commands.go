package samsungtv

import (
	"github.com/go-openapi/runtime/middleware"
	models "github.com/home-IoT/api-samsungtv/gen/models"
	ops "github.com/home-IoT/api-samsungtv/gen/restapi/operations"
)

// PostKey sends a key to the receiver
func PostKey(param ops.PostKeyParams) middleware.Responder {

	err := SendKey(param.Key)
	if err != nil {
		errorObj := models.ErrorResponse(err.Error())
		return ops.NewPostKeyDefault(502).WithPayload(errorObj)
	}

	return ops.NewPostKeyAccepted()
}
