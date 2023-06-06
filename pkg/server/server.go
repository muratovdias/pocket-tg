package server

import (
	"context"
	"github.com/muratovdias/pocket-tg/pkg/repository"
	"github.com/zhashkevych/go-pocket-sdk"
	"log"
	"net/http"
	"strconv"
)

type AuthorizationServer struct {
	server       *http.Server
	pocketClient *pocket.Client
	tokenRepo    repository.TokenRepo
	redirectURL  string
}

func NewAuthorizationServer(pocketClient *pocket.Client, tokenRepo repository.TokenRepo, redirectURL string) *AuthorizationServer {
	return &AuthorizationServer{
		pocketClient: pocketClient,
		tokenRepo:    tokenRepo,
		redirectURL:  redirectURL,
	}
}

func (s *AuthorizationServer) Start() error {
	s.server = &http.Server{
		Addr:    ":8000",
		Handler: s,
	}
	return s.server.ListenAndServe()
}

func (s *AuthorizationServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	chatIDParam := r.URL.Query().Get("chat_id")
	if chatIDParam == "" {
		log.Println("error Chat ID Param")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chatID, err := strconv.ParseInt(chatIDParam, 10, 64)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	requestToken, err := s.tokenRepo.Get(chatID, repository.RequestToken)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	authResp, err := s.pocketClient.Authorize(context.Background(), requestToken)
	if err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := s.tokenRepo.Save(chatID, authResp.AccessToken, repository.AccessToken); err != nil {
		log.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("chat_id: %d\n\trequest_token: %s\n\taccess_token: %s\n", chatID, requestToken, authResp.AccessToken)

	w.Header().Add("Location", s.redirectURL)
	w.WriteHeader(http.StatusMovedPermanently)
	return
}
