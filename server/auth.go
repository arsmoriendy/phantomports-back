package server

import (
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/arsmoriendy/opor/gql-srv/db"
	"github.com/arsmoriendy/opor/gql-srv/internal"
	"github.com/arsmoriendy/opor/gql-srv/internal/loglvl"
	"github.com/joho/godotenv"
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
		err = db.UuidValid(password)
		if err != nil {
			if err.Error()[:12] == "invalid UUID" {
				w.WriteHeader(400)
				w.Write([]byte("malformed UUID"))
				return
			}
			if errors.Is(err, db.ErrExpiredUuid) {
				w.WriteHeader(403)
				w.Write([]byte("expired UUID"))
				return
			}
			if errors.Is(err, db.ErrUnregisteredUuid) {
				w.WriteHeader(403)
				w.Write([]byte("unregistered UUID"))
				return
			}
			w.WriteHeader(500)
			if loglvl.LogLvl >= loglvl.TRACE {
				log.Println(err)
			}
			return
		}

		next.ServeHTTP(w, r)
	})
}

// Password to refresh frontend uuid
var refpass = func() string {
	godotenv.Load()
	refpass_env := os.Getenv("REF_FRONT_UUID_PASS")
	if refpass_env == "" {
		panic("$REF_FRONT_UUID_PASS must be set")
	}
	return refpass_env
}()

func RefreshFrontUuid(w http.ResponseWriter, r *http.Request) {
	// check auth header
	aHdr := r.Header.Get("Authorization")
	if aHdr == "" {
		w.Header().Add("WWW-Authenticate", "Basic realm=\"Refresh frontend UUID\"")
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

	// authorize password
	if password != refpass {
		w.WriteHeader(403)
		if loglvl.LogLvl >= loglvl.TRACE {
			log.Println("invalid ref front uuid pass: " + password)
		}
		return
	}

	// register uuid
	uuid, err := db.NewFrontUuid()
	if err != nil {
		w.WriteHeader(500)
		if loglvl.LogLvl >= loglvl.ERROR {
			log.Println("refresh handler: " + err.Error())
		}
		return
	}

	// delete on expire (not guaranteed if process closes)
	del_timer := time.NewTimer(internal.FrontUuidLifetime)
	go func() {
		<-del_timer.C
		err := db.RmUuid(uuid)
		if err != nil && loglvl.LogLvl >= loglvl.ERROR {
			log.Println(err)
		}
	}()

	w.Write([]byte(uuid))
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
