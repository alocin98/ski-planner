package middlewares

import (
	"fmt"
	"net/http"

	"github.com/alocin98/ski-planner-api/handlers"
	"github.com/alocin98/ski-planner-api/providers"
	"github.com/julienschmidt/httprouter"
)

func Cors(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		next(w, r, ps)
	}
}

func WithAuth(n httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("Authenticating...")
		accessToken, err := r.Cookie("AccessToken")
		if err != nil {
			handlers.AuthorizationResponse("unauthorized", w)
			fmt.Println(err)
			return
			//panic(err)
		}
		token, err := providers.FbAuthClient.VerifyIDToken(r.Context(), accessToken.Value)
		if err != nil {
			handlers.AuthorizationResponse("unauthorized", w)
			return
		}
		r.Header.Add("uid", token.UID)

		// call registered handler
		n(w, r, ps)
	}
}
