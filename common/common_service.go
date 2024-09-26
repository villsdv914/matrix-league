package common

import (
	"context"
	"log"
	"net/http"
	"strconv"
)

func WriteResponse(ctx context.Context, w http.ResponseWriter, contentType string, status int, data []byte) {
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Length", strconv.Itoa(len(data)))
	w.Header().Set("Connection", "close")
	w.WriteHeader(status)
	_, err := w.Write(data)
	if nil != err {
		log.Fatal(ctx, err, "Not able to write data to stream")
	}
}
