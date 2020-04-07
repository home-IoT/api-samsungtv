// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/home-IoT/api-samsungtv/gen/restapi/operations"
	"github.com/home-IoT/api-samsungtv/internal/samsungtv"
)

//go:generate swagger generate server --target ../../gen --name Samsungtv --spec ../../api/server.yml

func configureFlags(api *operations.SamsungtvAPI) {
	api.CommandLineOptionsGroups = samsungtv.CommandLineOptionsGroups
}

func configureAPI(api *operations.SamsungtvAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	samsungtv.Configure(api)

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.GetStatusHandler = operations.GetStatusHandlerFunc(samsungtv.GetStatus)
	api.PostKeyHandler = operations.PostKeyHandlerFunc(samsungtv.PostKey)
	api.PostPowerHandler = operations.PostPowerHandlerFunc(samsungtv.PostPower)

	if api.GetStatusHandler == nil {
		api.GetStatusHandler = operations.GetStatusHandlerFunc(func(params operations.GetStatusParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetStatus has not yet been implemented")
		})
	}
	if api.PostKeyHandler == nil {
		api.PostKeyHandler = operations.PostKeyHandlerFunc(func(params operations.PostKeyParams) middleware.Responder {
			return middleware.NotImplemented("operation .PostKey has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {
		// samsungtv.CloseConnection()
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
