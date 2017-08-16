package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"log"
	"github.com/DVI-GI-2017/Jira__backend/auth"
	"github.com/DVI-GI-2017/Jira__backend/tools"
)

var CheckDB = GetOnly(
	func(w http.ResponseWriter, r *http.Request) {
		var user auth.Credentials

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, "Error in request")
			log.Printf("%v", err)

			return
		}

		if err := auth.LoginUser(&user); err != nil {
			w.WriteHeader(http.StatusForbidden)

			fmt.Fprint(w, err)
			log.Printf("%v", err)

			return
		}

		token, err := auth.NewToken()

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)

			fmt.Fprintln(w, "Error while signing the token")
			log.Printf("%v", err)

			return
		}

		response := token
		tools.JsonResponse(response, w)
	})
