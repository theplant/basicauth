package basicauth

import "net/http"
import "fmt"

func BasicAuth(realm string, validate func(username, password string, r *http.Request) (err error)) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return &basicAuth{h: handler, v: validate, realm: realm}
	}
}

type basicAuth struct {
	h     http.Handler
	v     func(username, password string, r *http.Request) (err error)
	realm string
}

func (ba *basicAuth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	u, p, _ := r.BasicAuth()

	err := ba.v(u, p, r)
	if err == nil {
		ba.h.ServeHTTP(w, r)
		return
	}

	w.Header().Set("WWW-Authenticate", fmt.Sprintf(`Basic realm="%s"`, ba.realm))
	w.WriteHeader(401)
	w.Write([]byte("401 Unauthorized\n"))
}
