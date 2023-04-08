package api

import (
	"fmt"
	"net/http"
	"strings"
)

type Server struct {
	bot *Bot
}

func NewServer(bot *Bot) *Server {
	return &Server{bot: bot}
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/metrics", s.GetMetrics)

	return http.ListenAndServe(addr, nil)
}

func (s *Server) GetMetrics(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var metrics []string
	for _, user := range s.bot.users {
		m := fmt.Sprintf(`telegram_user_messages_total{userid="%d"} %d`,
			user.ID, user.Messages)

		metrics = append(metrics, m)
	}

	if len(metrics) > 0 {
		metrics_head := []string{
			`# HELP telegram_user_messages_total Total messages sent by the user`,
			`# TYPE telegram_user_messages_total counter`,
		}

		metrics = append(metrics_head, metrics...)
	}

	fmt.Fprint(w, strings.Join(metrics, "\n"))
}
