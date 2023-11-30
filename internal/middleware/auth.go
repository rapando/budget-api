package middleware

import (
	"encoding/base64"
	"github.com/rapando/budget-api/pkg/log"
	"net/http"
	"os"
	"strings"
	"time"
)

func AuthMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var start = time.Now()
		var decodedCredentials []byte
		var err error

		defer func() {
			log.Infof("%s %50s took %s", r.Method, r.URL, time.Since(start))
		}()

		var authHeader = strings.TrimPrefix(
			r.Header.Get("Authorization"),
			"Basic ",
		)
		decodedCredentials, err = base64.StdEncoding.DecodeString(authHeader)
		if err != nil {
			log.Warnf("invalid credentials format")
			Response(w, http.StatusUnauthorized, UnauthorizedResponse)
			return
		}
		if os.Getenv("AUTH") == string(decodedCredentials) {
			log.Infof("credentials ok")
			next.ServeHTTP(w, r)
			return
		} else {
			log.Warnf("credentials not ok")
			Response(w, http.StatusUnauthorized, UnauthorizedResponse)
			return
		}
	})
}
