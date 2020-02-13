package wallet

import (
	"errors"
	"net/http"
	"sync"

	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/wallet/crypto"
)

type Auth interface {
	NewSession(walletname string) (string, error)
	VerifyToken(token string) (string, error)
	Revoke(token string) error
}

type walletPassphrase struct {
	wallet     Wallet
	passphrase string
}

type handler struct {
	log      *logging.Logger
	auth     Auth
	rootPath string

	// wallet name -> wallet
	store map[string]walletPassphrase

	// just to make sure we do not access same file conccurently or the map
	mu sync.RWMutex
}

func newHandler(log *logging.Logger, auth Auth, rootPath string) *handler {
	return &handler{
		log:      log,
		auth:     auth,
		rootPath: rootPath,
		store:    map[string]walletPassphrase{},
	}
}

// return the actual token
func (h *handler) CreateWallet(wallet, passphrase string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	if _, ok := h.store[wallet]; ok {
		return "", ErrWalletAlreadyExist
	}

	w, err := Create(h.rootPath, wallet, passphrase)
	if err != nil {
		return "", err
	}

	h.store[wallet] = walletPassphrase{wallet: *w, passphrase: passphrase}
	return h.auth.NewSession(wallet)
}

func (h *handler) LoginWallet(wallet, passphrase string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()
	// first check if the user own the wallet
	w, err := Read(h.rootPath, wallet, passphrase)
	if err != nil {
		return "", ErrWalletDoesNotExist
	}

	// then store it in the memory store then
	if _, ok := h.store[wallet]; !ok {
		h.store[wallet] = walletPassphrase{wallet: *w, passphrase: passphrase}
	}

	return h.auth.NewSession(wallet)
}

func (h *handler) RevokeToken(token string) error {
	return h.auth.Revoke(token)
}

func (h *handler) GenerateKeypair(token string) (string, error) {
	h.mu.Lock()
	defer h.mu.Unlock()

	wname, err := h.auth.VerifyToken(token)
	if err != nil {
		return "", err
	}

	wp, ok := h.store[wname]
	if !ok {
		// this should never happen as we cannot have a valid session
		// without the actual wallet being loaded in memory but...
		return "", errors.New("could not found wallet")
	}

	kp, err := GenKeypair(crypto.Ed25519)
	if err != nil {
		return "", err
	}

	wp.wallet.Keypairs = append(wp.wallet.Keypairs, *kp)
	_, err = writeWallet(&wp.wallet, h.rootPath, wname, wp.passphrase)
	if err != nil {
		return "", err
	}

	h.store[wname] = wp
	return kp.Pub, nil
}

func (h *handler) ListPublicKeys(token string) ([]Keypair, error) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	wname, err := h.auth.VerifyToken(token)
	if err != nil {
		return nil, err
	}

	wp, ok := h.store[wname]
	if !ok {
		// this should never happen as we cannot have a valid session
		// without the actual wallet being loaded in memory but...
		return nil, errors.New("could not found wallet")
	}

	// copy all keys so we do not propagate private keys
	out := make([]Keypair, 0, len(wp.wallet.Keypairs))
	for _, v := range wp.wallet.Keypairs {
		kp := v
		kp.Priv = ""
		kp.privBytes = []byte{}
		out = append(out, kp)
	}

	return out, nil
}

func (h *handler) SignAndSubmitTx(w http.ResponseWriter, r *http.Request) {
	h.mu.RLock()
	defer h.mu.RUnlock()
}

func (h *handler) SignTx(w http.ResponseWriter, r *http.Request) {
	h.mu.RLock()
	defer h.mu.RUnlock()
}
