// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/bank"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/card"
	"github.com/cfif1982/cards/internal/infrastructure/swagger/restapi/operations/user"
)

//go:generate swagger generate server --target ..\..\swagger --name Cards --spec ..\..\..\..\swagger.yaml --principal interface{}

func configureFlags(api *operations.CardsAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CardsAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "api_key" header is set
	if api.APIKeyAuth == nil {
		api.APIKeyAuth = func(token string) (interface{}, error) {
			return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
		}
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()

	if api.BankAddBankHandler == nil {
		api.BankAddBankHandler = bank.AddBankHandlerFunc(func(params bank.AddBankParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bank.AddBank has not yet been implemented")
		})
	}
	if api.CardAddCardHandler == nil {
		api.CardAddCardHandler = card.AddCardHandlerFunc(func(params card.AddCardParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation card.AddCard has not yet been implemented")
		})
	}
	if api.UserAddUserHandler == nil {
		api.UserAddUserHandler = user.AddUserHandlerFunc(func(params user.AddUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.AddUser has not yet been implemented")
		})
	}
	if api.BankDeleteBankHandler == nil {
		api.BankDeleteBankHandler = bank.DeleteBankHandlerFunc(func(params bank.DeleteBankParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bank.DeleteBank has not yet been implemented")
		})
	}
	if api.CardDeleteCardHandler == nil {
		api.CardDeleteCardHandler = card.DeleteCardHandlerFunc(func(params card.DeleteCardParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation card.DeleteCard has not yet been implemented")
		})
	}
	if api.UserDeleteUserHandler == nil {
		api.UserDeleteUserHandler = user.DeleteUserHandlerFunc(func(params user.DeleteUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.DeleteUser has not yet been implemented")
		})
	}
	if api.BankGetBankByUUIDHandler == nil {
		api.BankGetBankByUUIDHandler = bank.GetBankByUUIDHandlerFunc(func(params bank.GetBankByUUIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bank.GetBankByUUID has not yet been implemented")
		})
	}
	if api.CardGetCardByNumberHandler == nil {
		api.CardGetCardByNumberHandler = card.GetCardByNumberHandlerFunc(func(params card.GetCardByNumberParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation card.GetCardByNumber has not yet been implemented")
		})
	}
	if api.UserGetUserByUUIDHandler == nil {
		api.UserGetUserByUUIDHandler = user.GetUserByUUIDHandlerFunc(func(params user.GetUserByUUIDParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.GetUserByUUID has not yet been implemented")
		})
	}
	if api.BankUpdateBankHandler == nil {
		api.BankUpdateBankHandler = bank.UpdateBankHandlerFunc(func(params bank.UpdateBankParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation bank.UpdateBank has not yet been implemented")
		})
	}
	if api.CardUpdateCardHandler == nil {
		api.CardUpdateCardHandler = card.UpdateCardHandlerFunc(func(params card.UpdateCardParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation card.UpdateCard has not yet been implemented")
		})
	}
	if api.UserUpdateUserHandler == nil {
		api.UserUpdateUserHandler = user.UpdateUserHandlerFunc(func(params user.UpdateUserParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation user.UpdateUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
