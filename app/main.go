package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand/v2"
	"time"

	"net/http"

	//	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// adicionando metricas de status
var userStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_get_user_status_count", // metric name
		Help: "Count of status returned by user.",
	},
	[]string{"user", "status"},
)

// adicionando metricas de total de requests
var totalRequests = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of get requests.",
	},
	[]string{"path"},
)

// adicionando metricas de http response
var responseStatus = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "response_status",
		Help: "Status of HTTP response",
	},
	[]string{"status"},
)

// iniciando metricas
func init() {

	prometheus.MustRegister(userStatus)
	prometheus.MustRegister(totalRequests)
	prometheus.MustRegister(responseStatus)
}

type MyRequest struct {
	User string
}

func server(w http.ResponseWriter, r *http.Request) {
	var status string
	var user string
	defer func() {
		userStatus.WithLabelValues(user, status).Inc()
		responseStatus.WithLabelValues(status).Inc()
		totalRequests.WithLabelValues("/").Inc()
	}()
	var mr MyRequest
	json.NewDecoder(r.Body).Decode(&mr)
	// rand para simular status 400 tbm
	if rand.Float32() > 0.8 {
		status = "400"
	} else {
		status = "200"
	}
	user = mr.User
	// log simples para a aplicação
	log.Println(user, status)
	w.Write([]byte(status))
}

func producer() {
	userPool := []string{"bob", "alice", "jack"}
	for {
		postBody, _ := json.Marshal(MyRequest{
			User: userPool[rand.IntN(len(userPool))],
		})
		requestBody := bytes.NewBuffer(postBody)
		http.Post("http://localhost:8080", "application/json", requestBody)
		//fmt.Println("request")
		time.Sleep(time.Second * 2)
	}
}
func main() {

	go producer()
	http.Handle("/metrics", promhttp.Handler())
	//router.HandleFunc("/serv", server).Methods("POST")
	http.HandleFunc("/", server)
	fmt.Println("server on port 8080")
	fmt.Println("----acesse--v1.1----")
	http.ListenAndServe(":8080", nil)

}
