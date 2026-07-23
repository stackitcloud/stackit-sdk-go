package identity

import (
	"time"

	"github.com/google/uuid"
)

// ServiceAccountJson is the API response when creating a STACKIT service account key.
type ServiceAccountJson struct {
	Active       bool                          `json:"active"`
	CreatedAt    time.Time                     `json:"createdAt"`
	Credentials  *ServiceAccountKeyCredentials `json:"credentials"`
	ID           uuid.UUID                     `json:"id"`
	KeyAlgorithm string                        `json:"keyAlgorithm"`
	KeyOrigin    string                        `json:"keyOrigin"`
	KeyType      string                        `json:"keyType"`
	PublicKey    string                        `json:"publicKey"`
	ValidUntil   *time.Time                    `json:"validUntil,omitempty"`
}

// ServiceAccountKeyCredentials contains the credential fields embedded in a ServiceAccountKeyResponse.
type ServiceAccountKeyCredentials struct {
	Aud           string    `json:"aud"`
	Iss           string    `json:"iss"`
	Kid           string    `json:"kid"`
	PrivateKey    *string   `json:"privateKey,omitempty"`
	Sub           uuid.UUID `json:"sub"`
	TokenEndpoint string    `json:"tokenEndpoint"`
}
