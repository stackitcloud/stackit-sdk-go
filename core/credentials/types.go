package credentials

import (
	"time"

	"github.com/google/uuid"
)

// serviceAccountJSON is the API response when creating a STACKIT service account key.
type serviceAccountJSON struct {
	Active       bool                          `json:"active"`
	CreatedAt    time.Time                     `json:"createdAt"`
	Credentials  *serviceAccountKeyCredentials `json:"credentials"`
	ID           uuid.UUID                     `json:"id"`
	KeyAlgorithm string                        `json:"keyAlgorithm"`
	KeyOrigin    string                        `json:"keyOrigin"`
	KeyType      string                        `json:"keyType"`
	PublicKey    string                        `json:"publicKey"`
	ValidUntil   *time.Time                    `json:"validUntil,omitempty"`
}

// serviceAccountKeyCredentials contains the credential fields embedded in a serviceAccountJSON.
type serviceAccountKeyCredentials struct {
	Aud           string    `json:"aud"`
	Iss           string    `json:"iss"`
	Kid           string    `json:"kid"`
	PrivateKey    *string   `json:"privateKey,omitempty"`
	Sub           uuid.UUID `json:"sub"`
	TokenEndpoint string    `json:"tokenEndpoint"`
}
