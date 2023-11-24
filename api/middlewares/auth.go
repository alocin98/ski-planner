package middlewares

import (
	"net/http"

	"github.com/alocin98/ski-planner-api/handlers"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/julienschmidt/httprouter"
)

func WithAuth(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		token, err := providers.FbAuthClient.VerifyIDToken(r.Context(), r.Header.Get("AccessToken"))
		if err != nil {
			handlers.AuthorizationResponse("unauthorized", w)
			return
		}
		r.Header.Add("uid", token.UID)

		// call registered handler
		n(w, r, ps)
	}
}
