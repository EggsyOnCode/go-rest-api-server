package middleware

import (
	"errors"
	"net/http"

	"github.com/EggsyOnCode/go-rest-api-server/api"
	"github.com/EggsyOnCode/go-rest-api-server/internal/tools"
	log "github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("you are not authorized to view this info")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username = r.URL.Query().Get("username")
		var token = r.Header.Get("Authorization")

		if username == "" || token == "" {
			log.Error(UnAuthorizedError)
			api.ResponseErrorHandler(w, UnAuthorizedError)
		}

		var database *tools.DatabaseInterface
		database, _err := tools.NewDB()
		if _err != nil {
			api.InternalServerError(w, _err)
			return
		}

		var logindetails *tools.LoginDetails
		logindetails = (*database).GetLoginDetails(username)

		if logindetails == nil || (token != (*&logindetails.AuthToken)) {
			log.Error(UnAuthorizedError)
			api.ResponseErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
