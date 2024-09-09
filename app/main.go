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

var requestDuration = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name:        "http_server_request_duration_seconds_sum",
	Help:        "Histogram of response time for handler in seconds",
	Buckets:     []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
}, []string{"route", "method", "status_code"})

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
	prometheus.MustRegister(requestDuration)

}

type MyRequest struct {
	User string
}

func server(w http.ResponseWriter, r *http.Request) {
	var status string
	var user string
	start := time.Now()
	duration := time.Since(start)
	defer func() {
		userStatus.WithLabelValues(user, status).Inc()
		responseStatus.WithLabelValues(status).Inc()
		// timer := prometheus.NewTimer(requestDuration.WithLabelValues(r.URL.Path))

		// defer timer.ObserveDuration()
		//requestDuration.WithLabelValues(r.URL.Path, "POST", status).Observe(duration.Seconds())

		totalRequests.WithLabelValues(r.URL.Path).Inc()

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
	requestDuration.WithLabelValues(r.URL.Path, "POST", status).Observe(duration.Seconds())
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
		http.Post("http://localhost:8080/", "application/json", requestBody)
		//fmt.Println("request")
		time.Sleep(time.Second * 2)
	}
}
func main() {
	//	router := mux.NewRouter()
	go producer()
	http.Handle("/metrics", promhttp.Handler())
	//router.HandleFunc("/serv", server).Methods("POST")
	http.HandleFunc("/", server)
	fmt.Println("server on port 8080")
	fmt.Println("----acesse--v1.1----")
	http.ListenAndServe(":8080", nil)

	//router.HandleFunc("/serv", func(w http.ResponseWriter, r *http.Request) {

	// 	w.Write([]byte(fmt.Sprintf("sucess |  %s", time.Now())))
	// }).Methods("GET")

	//http.ListenAndServe(":8080", router)

}
