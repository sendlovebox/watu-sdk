// Package api defines implementations of endpoints and calls
package api

import (
	"context"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/rs/zerolog"

	"github.com/sendlovebox/watu-sdk/model"
)

// RemoteCalls abstracted definition of supported functions
//
//go:generate mockgen -source api.go -destination ./mock_remote_calls.go -package api RemoteCalls
type RemoteCalls interface {
	Auth(secretKey, publicKey, iv, encryptionKey string) error
	RunInSandboxMode()

	// Bills
	GetBillGroups(ctx context.Context) ([]*model.BillGroup, error)
	GetBillsByGroup(ctx context.Context, group string) ([]*model.Bill, error)
	GetBillInformation(ctx context.Context, billChannel string) (model.Bill, error)
	GetBillTypes(ctx context.Context, billChannel string) ([]*model.BillType, error)
	ValidateBill(ctx context.Context, body interface{}) (interface{}, error)
	VendBill(ctx context.Context, body interface{}) (model.VendBillResponse, error)
}

// Call object
type Call struct {
	client        *resty.Client
	logger        zerolog.Logger
	sandboxMode   bool
	baseURL       string
	secretKey     string
	publicKey     string
	iv            string
	encryptionKey string
}

// New initialises the object Call
func New(z *zerolog.Logger, c *resty.Client, baseURL string) RemoteCalls {
	c.SetTimeout(15 * time.Second)
	call := &Call{
		client:  c,
		logger:  z.With().Str("sdk", "watu").Logger(),
		baseURL: baseURL,
	}
	return RemoteCalls(call)
}

// RunInSandboxMode this forces Call functionalities to run in sandbox mode for relevant logic/API consumption
func (c *Call) RunInSandboxMode() {
	c.client.SetDebug(true)
	c.client.EnableTrace()
	c.sandboxMode = true
}

// Auth to set authorization token
func (c *Call) Auth(secretKey, publicKey, iv, encryptionKey string) error {
	c.secretKey = secretKey
	c.publicKey = publicKey
	c.iv = iv
	c.encryptionKey = encryptionKey

	return nil
}
