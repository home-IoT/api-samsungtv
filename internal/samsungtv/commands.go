package samsungtv

import (
	"errors"
	"fmt"
	"strings"

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

const keyPowerOff = "KEY_POWER"

// PostPower turns the TV on or off
func PostPower(param ops.PostPowerParams) middleware.Responder {

	state := strings.ToUpper(param.State)

	eCode := 500
	var err error
	switch state {
	case "ON":
		if configuration.TV.Mac == nil {
			err = errors.New("TV's MAC address is not configured")
		} else {

		}
		err = wakeOnLan(*configuration.TV.Mac)

	case "OFF":
		err = SendKey(keyPowerOff)

	default:
		eCode = 400
		err = fmt.Errorf("invalid power state '%s'", state)
	}

	if err != nil {
		errorObj := models.ErrorResponse(err.Error())
		return ops.NewPostKeyDefault(eCode).WithPayload(errorObj)
	}

	return ops.NewPostKeyAccepted()
}
