package rest

import (
	"encoding/json"
	"fmt"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/config"
	"github.com/1ambda/domain-driven-design-go/service-gateway/internal/domain/user"
	e "github.com/1ambda/domain-driven-design-go/service-gateway/internal/exception"
	"github.com/gorilla/sessions"
	"github.com/pkg/errors"
	"net/http"
	"strings"
)

func InjectAuthMiddleware(sessionStore sessions.Store, next http.Handler) http.Handler {
	env := config.Env
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// disable swagger-ui based on the flag
		if strings.HasPrefix(r.URL.Path, "/api/docs") && !env.EnableSwaggerUI {
			message := fmt.Sprintf("Not Found: (%s) %s", r.Method, r.URL.Path)
			err := errors.New(message)
			ex := e.NewNotFoundException(err, message)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(ex.StatusCode())
			json.NewEncoder(w).Encode(ex.ToSwaggerError())
			return
		}

		// whitelist for requests which don't require session check
		if strings.HasPrefix(r.URL.Path, "/api/auth/whoami") ||
			strings.HasPrefix(r.URL.Path, "/api/auth/register") ||
			strings.HasPrefix(r.URL.Path, "/api/auth/login") ||
			strings.HasPrefix(r.URL.Path, "/api/auth/logout") {
			next.ServeHTTP(w, r)
			return
		}

		// disable session check based on `ENABLE_SESSION_CHECK`
		if env.DisableSessionCheck {
			next.ServeHTTP(w, r)
			return
		}

		session, _ := sessionStore.Get(r, user.SessionCookieName)
		if authenticated, _ := user.IsAuthenticated(session); !authenticated {
			message := fmt.Sprintf("Not Authenticated: (%s) %s", r.Method, r.URL.Path)
			err := errors.New(message)
			ex := e.NewUnauthorizedException(err, message)

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(ex.StatusCode())
			json.NewEncoder(w).Encode(ex.ToSwaggerError())
			return
		}

		next.ServeHTTP(w, r)
	})
}
