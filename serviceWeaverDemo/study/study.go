package study

import (
	"context"
	"encoding/json"
	"github.com/ServiceWeaver/weaver"
	"github.com/go-chi/chi"
	"github.com/timliudream/go-test/serviceWeaverDemo/interview"
	"github.com/timliudream/go-test/serviceWeaverDemo/model"
	"github.com/timliudream/go-test/serviceWeaverDemo/working"
	"net/http"
)

type Server struct {
	weaver.Implements[weaver.Main]

	handler http.Handler

	interviewService weaver.Ref[interview.Service]
	working          weaver.Ref[working.Service]

	studyapi weaver.Listener `weaver:"studyapi"`
}

func (s *Server) Init(ctx context.Context) error {
	s.Logger(ctx).Info("Init")
	r := chi.NewRouter()
	r.Route("/api/study", func(r chi.Router) {
		r.Post("/", s.Study)
		r.Get("/{type}", s.GetKnowledges)
	})
	s.handler = r
	return nil
}

var ctx = context.Background()

func (s *Server) Study(writer http.ResponseWriter, request *http.Request) {
	var golang model.Golang
	err := json.NewDecoder(request.Body).Decode(&golang)
	if err != nil {
		http.Error(writer, "Invalid Course Data", 500)
		return
	}
	golang.Goroutine = "GMP"
	golang.Channel = "blocking"
	if err := s.interviewService.Get().MakeInterview(ctx, golang); err != nil {
		s.Logger(ctx).Error(
			"interview failed",
			"error:", err,
		)
		http.Error(writer, "interview failed", http.StatusInternalServerError)
		return
	}
	golang.Goroutine = "fixing goroutine bug"
	golang.Channel = "non-blocking"
	// send notification using notificationService component
	if err := s.working.Get().Working(ctx, golang); err != nil {
		s.Logger(ctx).Error(
			"fixing bug failed",
			"error:", err,
		)
	}
	s.Logger(ctx).Info("those bugs has been fixed")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func (s *Server) GetKnowledges(writer http.ResponseWriter, request *http.Request) {
	s.Logger(ctx).Info("get golang knowledge")
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
}

func Serve(ctx context.Context, s *Server) error {
	s.Logger(ctx).Info("StudyApi listener available.", "addr:", s.studyapi)
	httpServer := &http.Server{
		Handler: s.handler,
	}
	httpServer.Serve(s.studyapi)
	return nil
}
