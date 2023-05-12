package api

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	labels = []string{"userid", "chatid"}

	userMessages = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_messages_total",
		Help: "Total messages sent by the user",
	}, labels)

	userAnimations = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_animations_total",
		Help: "Total animations sent by the user",
	}, labels)

	userAudios = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_audios_total",
		Help: "Total audios sent by the user",
	}, labels)

	userDocuments = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_documents_total",
		Help: "Total documents sent by the user",
	}, labels)

	userPhotos = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_photos_total",
		Help: "Total photos sent by the user",
	}, labels)

	userStickers = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_stickers_total",
		Help: "Total stickers sent by the user",
	}, labels)

	userVideos = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_videos_total",
		Help: "Total videos sent by the user",
	}, labels)

	userVideoNotes = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_video_notes_total",
		Help: "Total video notes sent by the user",
	}, labels)

	userVoices = promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "telegram_user_voices_total",
		Help: "Total voices sent by the user",
	}, labels)
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start(addr string) error {
	http.Handle("/metrics", promhttp.Handler())

	return http.ListenAndServe(addr, nil)
}
