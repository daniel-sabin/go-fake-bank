package server

import (
	"context"
	"errors"
	"net/http"
	"regexp"
	"strings"
)

type ContextKey string

const ParamFromUrl ContextKey = "paramUrl"

type Router struct {
	http.ServeMux
}

func (r *Router) HandleWithParam(route string, toHandle http.HandlerFunc) {
	routeToHandle := routeToHandle(route)
	handler := withRegex(regexToGetParam(route), toHandle)

	r.Handle(routeToHandle, handler)
}

func routeToHandle(route string) string {
	var routeToHandle string
	r, _ := regexp.Compile("(.*){param}")
	found := r.FindStringSubmatch(route)

	if len(found) > 1 {
		routeToHandle = found[1]
	}

	return routeToHandle
}

func withRegex(regex string, next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		paramFromUrl, err := paramFromUrl(r.URL.Path, regex)

		if err != nil {
			sendNotFound(w)
			return
		}

		ctx := context.WithValue(r.Context(), ParamFromUrl, paramFromUrl)

		next(w, r.WithContext(ctx))
	}
}

func paramFromUrl(url string, regex string) (string, error) {
	r, _ := regexp.Compile(regex)
	found := r.FindStringSubmatch(url)

	if len(found) == 0 {
		return "", errors.New("error in valueFromUrl")
	}

	return found[1], nil
}

func regexToGetParam(route string) string {
	return strings.Replace(route, "{param}", "(.*?)", 1)
}

func sendNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("404 page not found"))
}
