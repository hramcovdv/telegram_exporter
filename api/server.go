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
		m := fmt.Sprintf(`telegram_user_messages_count{userid="%d"} %d`,
			user.ID, user.Messages)

		metrics = append(metrics, m)
	}

	fmt.Fprint(w, strings.Join(metrics, "\n"))
}
