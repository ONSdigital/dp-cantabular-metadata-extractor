package api

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/log.go/v2/log"
)

func MockMetadataHandler(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		resp, err := http.Get("http://localhost:2112")
		if err != nil {
			log.Error(ctx, "get request failed", err)
			http.Error(w, "Failed to get mock metadata", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		bodyResp, err := ioutil.ReadAll(resp.Body)

		if err != nil {
			log.Error(ctx, "failed to read the data", err)
			http.Error(w, "Failed to read the response body", http.StatusInternalServerError)
			return
		}
		// jsonResponse, err := json.Marshal(bodyResp)
		// if err != nil {
		// 	log.Error(ctx, "marshalling response failed", err)
		// 	http.Error(w, "Failed to marshall json response", http.StatusInternalServerError)
		// 	return
		// }
		_, err = w.Write(bodyResp)
		if err != nil {
			log.Error(ctx, "writing response failed", err)
			http.Error(w, "Failed to write http response", http.StatusInternalServerError)
			return
		}
	}
}
