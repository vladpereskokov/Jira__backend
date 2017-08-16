package routes

import (
	"fmt"
	"github.com/DVI-GI-2017/Jira__backend/handlers"
	"net/http"
)

func NewRouter() (http.Handler, error) {
	const apiRoot = "/api/v1"

	mux := http.NewServeMux()

	for _, r := range routeList {
		mux.HandleFunc(fmt.Sprintf("%s%s", apiRoot, r.Pattern), r.HandlerFunc)
	}

	return handlers.Logger(mux), nil
}