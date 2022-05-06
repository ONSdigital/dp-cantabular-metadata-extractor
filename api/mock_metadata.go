package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"

	// "fmt"
	"io/ioutil"
	"net/http"

	"github.com/ONSdigital/log.go/v2/log"
	"strings"
	// "github.com/gorilla/mux"
)

type graphQLRequest struct {
	Query     string `json:"query"`
	Variables string `json:"vars"`
}

func GetMockMetadata(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		resp, err := http.Get("http://localhost:2112")
		if err != nil {
			log.Error(ctx, "get request failed", err)
			http.Error(w, "Failed to get mock metadata", http.StatusInternalServerError)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")

		defer resp.Body.Close()

		bytesRespBody, errRead := ioutil.ReadAll(resp.Body)
		if errRead != nil {
			log.Error(ctx, "failed to read the data", errRead)
			http.Error(w, "Failed to read the response body", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(bytesRespBody)
		if err != nil {
			log.Error(ctx, "writing response failed", err)
			http.Error(w, "Failed to write http response", http.StatusInternalServerError)
			return
		}
	}
}

func GetCantabularMetadata(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := req.Context()
		jsonData := map[string]string{
			"query": `{
					service {
					  tables {
						name
						label
						description
						datasetName
						vars
					  }
					}
				  }`,
		}
		gqlQueryMarshalled, err := json.Marshal(jsonData)
		if err != nil {
			log.Error(ctx, "marshalling the gql failed", err)
			http.Error(w, "Failed to marshall the gql query", http.StatusInternalServerError)
			return
		}
		request, err := http.NewRequest("POST", "http://localhost:8492/graphql", bytes.NewBuffer(gqlQueryMarshalled))
		fmt.Println("working request", request)
		if err != nil {
			log.Error(ctx, "get request failed", err)
			http.Error(w, "Failed to get cantabular metadata", http.StatusInternalServerError)
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Error(ctx, "request failed", err)
			http.Error(w, "an error occured, failed to return a response", http.StatusInternalServerError)
		}

		defer response.Body.Close()

		bytesRespBody, errRead := ioutil.ReadAll(response.Body)
		if errRead != nil {
			log.Error(ctx, "failed to read the data", errRead)
			http.Error(w, "Failed to read the response body", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(bytesRespBody)
		if err != nil {
			log.Error(ctx, "writing response failed", err)
			http.Error(w, "Failed to write http response", http.StatusInternalServerError)
			return
		}
	}
}

func GetCantabularMetadataWithVars(ctx context.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ctx := req.Context()
		query := `query($dataset:String!) {
					dataset(name: $dataset) {
						variables(names: ["Age"]) {
							edges {
								node {
									name
									label
									categories {
										totalCount
									}
								}
							}
						}
					  }
					}`

		variables := `vars{"dataset":"Teaching-Dataset"}`

		gqlQueryMarshalled, err := json.Marshal(graphQLRequest{Query: query, Variables: variables})

		if err != nil {
			log.Error(ctx, "marshalling the gql query failed", err)
			http.Error(w, "Failed to marshall the gql query", http.StatusInternalServerError)
			return
		}
		request, err := http.NewRequest("POST", "http://localhost:8492/graphql", strings.NewReader(string(gqlQueryMarshalled)))

		if err != nil {
			log.Error(ctx, "get request failed", err)
			http.Error(w, "Failed to get cantabular metadata", http.StatusInternalServerError)
		}

		client := &http.Client{}
		response, err := client.Do(request)
		if err != nil {
			log.Error(ctx, "request failed", err)
			http.Error(w, "an error occured, failed to return a response", http.StatusInternalServerError)
		}

		defer response.Body.Close()

		bytesRespBody, errRead := ioutil.ReadAll(response.Body)
		if errRead != nil {
			log.Error(ctx, "failed to read the data", errRead)
			http.Error(w, "Failed to read the response body", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(bytesRespBody)
		if err != nil {
			log.Error(ctx, "writing response failed", err)
			http.Error(w, "Failed to write http response", http.StatusInternalServerError)
			return
		}
	}
}

// const QueryDimensionsByName = `
// query($dataset: String!, $variables: [String!]!) {
// 	dataset(name: $dataset) {
// 		variables(names: $variables) {
// 			edges {
// 				node {
// 					name
// 					label
// 					categories {
// 						totalCount
// 					}
// 				}
// 			}
// 		}
// 	}
// }`

// // GetDimensionsByNameRequest holds the request variables required from the
// // caller for making a request to obtain dimensions (Cantabular variables) by name
// // POST [cantabular-ext]/graphql
// type GetDimensionsByNameRequest struct {
// 	Dataset        string
// 	DimensionNames []string
// }
