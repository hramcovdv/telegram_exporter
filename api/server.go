package api

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/hramcovdv/telegram_exporter/types"
)

type Server struct {
	bot Telebot
}

type Telebot interface {
	GetUsers() []*types.User
}

func NewServer(bot Telebot) *Server {
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

	var messages, symbols, metrics []string
	for _, user := range s.bot.GetUsers() {
		m := fmt.Sprintf(
			`telegram_user_messages_total{userid="%d",chatid="%d"} %d`,
			user.ID, user.ChatID, user.Messages)

		messages = append(messages, m)

		s := fmt.Sprintf(
			`telegram_user_symbols_total{userid="%d",chatid="%d"} %d`,
			user.ID, user.ChatID, user.Symbols)

		symbols = append(symbols, s)
	}

	if len(messages) > 0 {
		messages = append([]string{
			`# HELP telegram_user_messages_total Total messages sent by the user`,
			`# TYPE telegram_user_messages_total counter`,
		}, messages...)

		metrics = append(metrics, messages...)
	}

	if len(symbols) > 0 {
		symbols = append([]string{
			`# HELP telegram_user_symbols_total Total symbols sent by the user`,
			`# TYPE telegram_user_symbols_total counter`,
		}, symbols...)

		metrics = append(metrics, symbols...)
	}

	fmt.Fprint(w, strings.Join(metrics, "\n"))
}
