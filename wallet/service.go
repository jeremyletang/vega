package wallet

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"code.vegaprotocol.io/vega/logging"
	"github.com/rs/cors"
)

type Service struct {
	*http.ServeMux

	cfg     *Config
	log     *logging.Logger
	s       *http.Server
	handler *handler
}

func NewService(log *logging.Logger, cfg *Config, rootPath string) (*Service, error) {
	// ensure the folder exist
	if err := EnsureBaseFolder(rootPath); err != nil {
		return nil, err
	}

	auth, err := newAuth(log, rootPath)
	if err != nil {
		return nil, err
	}
	handler := newHandler(log, auth, rootPath)

	s := &Service{
		ServeMux: http.NewServeMux(),
		log:      log,
		cfg:      cfg,
		handler:  handler,
	}

	s.HandleFunc("/api/v1/health", s.health)
	s.HandleFunc("/api/v1/create", s.createWallet)
	s.HandleFunc("/api/v1/login", s.login)
	s.HandleFunc("/api/v1/revoke", extractToken(s.revoke))
	s.HandleFunc("/api/v1/gen-keys", extractToken(s.generateKeypair))
	s.HandleFunc("/api/v1/list-keys", extractToken(s.listPublicKeys))

	return s, nil
}

func (s *Service) Start() error {
	s.s = &http.Server{
		Addr:    fmt.Sprintf("%s:%v", s.cfg.IP, s.cfg.Port),
		Handler: cors.AllowAll().Handler(s), // middlewar with cors
	}

	s.log.Info("starting wallet http server", logging.String("address", s.s.Addr))
	return s.s.ListenAndServe()
}

func (s *Service) Stop() error {
	return s.s.Shutdown(context.Background())
}

func (s *Service) createWallet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	// unmarshal request
	req := struct {
		Wallet     string `json:"wallet"`
		Passphrase string `json:"passphrase"`
	}{}
	if err := unmarshalBody(r, &req); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	// validation
	if len(req.Wallet) <= 0 {
		writeError(w, newError("missing wallet field"), http.StatusBadRequest)
		return
	}
	if len(req.Passphrase) <= 0 {
		writeError(w, newError("missing passphrase field"), http.StatusBadRequest)
		return
	}

	token, err := s.handler.CreateWallet(req.Wallet, req.Passphrase)
	if err != nil {
		writeError(w, newError(err.Error()), http.StatusForbidden)
		return
	}
	writeSuccess(w, token, http.StatusOK)
}

func (s *Service) login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
		return
	}
	req := struct {
		Wallet     string `json:"wallet"`
		Passphrase string `json:"passphrase"`
	}{}
	if err := unmarshalBody(r, &req); err != nil {
		writeError(w, err, http.StatusBadRequest)
		return
	}

	// validation
	if len(req.Wallet) <= 0 {
		writeError(w, newError("missing wallet field"), http.StatusBadRequest)
		return
	}
	if len(req.Passphrase) <= 0 {
		writeError(w, newError("missing passphrase field"), http.StatusBadRequest)
		return
	}

	token, err := s.handler.LoginWallet(req.Wallet, req.Passphrase)
	if err != nil {
		writeError(w, newError(err.Error()), http.StatusForbidden)
		return
	}
	writeSuccess(w, token, http.StatusOK)
}

func (s *Service) revoke(t string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
		return
	}

	err := s.handler.RevokeToken(t)
	if err != nil {
		writeError(w, newError(err.Error()), http.StatusForbidden)
		return
	}

	writeSuccess(w, true, http.StatusOK)
}

func (s *Service) generateKeypair(t string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
		return
	}

	pubKey, err := s.handler.GenerateKeypair(t)
	if err != nil {
		writeError(w, newError(err.Error()), http.StatusForbidden)
		return
	}

	writeSuccess(w, pubKey, http.StatusOK)
}

func (s *Service) listPublicKeys(t string, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
		return
	}

	keys, err := s.handler.ListPublicKeys(t)
	if err != nil {
		writeError(w, newError(err.Error()), http.StatusForbidden)
		return
	}

	writeSuccess(w, keys, http.StatusOK)
}

func (h *Service) signAndSubmitTx(w http.ResponseWriter, r *http.Request) {
}

func (h *Service) signTx(w http.ResponseWriter, r *http.Request) {

}

func (h *Service) health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
}

func unmarshalBody(r *http.Request, into interface{}) error {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return ErrInvalidRequest
	}
	return json.Unmarshal(body, into)
}

func writeError(w http.ResponseWriter, e error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	buf, _ := json.Marshal(e)
	w.Write(buf)
}

type successResponse struct {
	Data interface{}
}

func writeSuccess(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	buf, _ := json.Marshal(successResponse{data})
	w.Write(buf)
}

var (
	ErrInvalidRequest        = newError("invalid request")
	ErrInvalidMethod         = newError("invalid method")
	ErrInvalidOrMissingToken = newError("invalid or missing token")
)

type HttpError struct {
	ErrorStr string `json:"error"`
}

func (e HttpError) Error() string {
	return e.ErrorStr
}

func newError(e string) HttpError {
	return HttpError{
		ErrorStr: e,
	}
}
