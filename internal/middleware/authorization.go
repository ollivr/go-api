package middleware

import (
	"errors"
	log "github.com/sirupsen/logrus"
	"goapi/api"
	"goapi/internal/tools"
	"net/http"
)

var UnauthorizedError = errors.New("Invalid Username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InternalErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails
		loginDetails = (*database).GetUserLoginDetails(username)

		if loginDetails == nil || (token != (*loginDetails).AuthToken) {
			log.Error(UnauthorizedError)
			api.RequestErrorHandler(w, UnauthorizedError)
			return
		}
		next.ServeHTTP(w, r)
	})
}
