package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var userMessages = promauto.NewCounterVec(prometheus.CounterOpts{
	Name: "telegram_user_messages_total",
	Help: "Total messages sent by the user",
}, []string{"userid", "chatid"})

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(addr string) error {
	http.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(addr, nil)
}
