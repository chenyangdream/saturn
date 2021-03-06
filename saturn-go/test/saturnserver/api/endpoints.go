package api

import (
	"net/http"
	"strings"
)

func post(i *jsonAPIHandler, path string, w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(path, "/saturn/add"):
		i.POSTAdd(w, r)
	case strings.HasPrefix(path, "/saturn/pin"):
		i.POSTPin(w, r)
	case strings.HasPrefix(path, "/saturn/unpin"):
		i.POSTUnpin(w, r)
	default:
		ErrorResponse(w, http.StatusNotFound, "Not Found")
	}
}

func get(i *jsonAPIHandler, path string, w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.HasPrefix(path, "/saturn/peers"):
		i.GETPeers(w, r)
	case strings.HasPrefix(path, "/saturn/peerid"):
		i.GETPeerId(w, r)
	case strings.HasPrefix(path, "/saturn/cat"):
		i.GETFileContent(w, r)
	case strings.HasPrefix(path, "/saturn/pinls"):
		i.GETPinLs(w, r)
	case strings.HasPrefix(path, "/saturn/get"):
		i.GETGet(w, r)
	default:
		ErrorResponse(w, http.StatusNotFound, "Not Found")
	}
}


