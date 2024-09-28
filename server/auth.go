package server

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/arsmoriendy/opor/gql-srv/db"
	"github.com/arsmoriendy/opor/gql-srv/internal/loglvl"
)

var ErrAuthHeaderFmt = errors.New("authentication header has an invalid format")
var ErrMismatchAuthScheme = errors.New("mismatched authentication header scheme")

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// check auth header
		aHdr := r.Header.Get("Authorization")
		if aHdr == "" {
			w.Header().Add("WWW-Authenticate", "Basic realm=\"API Access\"")
			w.WriteHeader(401)
			return
		}

		// parse auth header
		_, password, err := parseBasicAuth(aHdr)
		if err != nil {
			w.WriteHeader(400)
			if loglvl.LogLvl >= loglvl.TRACE {
				log.Println(err)
			}
			return
		}

		// authorize uuid
		exists, err := db.UuidValid(password)
		if err != nil {
			w.WriteHeader(400)
			if loglvl.LogLvl >= loglvl.TRACE {
				log.Println(err)
			}
			return
		}
		if !exists {
			w.WriteHeader(403)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func parseBasicAuth(authHeader string) (username string, password string, err error) {
	scheme, b64user, found := strings.Cut(authHeader, " ")
	if !found {
		return username, password, fmt.Errorf("%w: %v", ErrAuthHeaderFmt, authHeader)
	}
	if scheme != "Basic" {
		return username, password, fmt.Errorf("%w: %v", ErrMismatchAuthScheme, scheme)
	}

	user, err := base64.StdEncoding.DecodeString(b64user)
	if err != nil {
		return
	}

	username, password, found = strings.Cut(string(user), ":")
	if !found {
		return username, password, fmt.Errorf("%w: %v", ErrAuthHeaderFmt, user)
	}
	return
}
