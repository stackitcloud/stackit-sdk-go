package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"testing"

	"github.com/google/uuid"
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

const saKeyStrPattern = `{
	"active": true,
	"createdAt": "2023-03-23T18:26:20.335Z",
	"credentials": {
	  "aud": "https://test-url.com",
	  "iss": "test@test.cloud",
	  "kid": "%s",
	  "sub": "%s"
	},
	"id": "%s",
	"keyAlgorithm": "RSA_2048",
	"keyOrigin": "USER_PROVIDED",
	"keyType": "USER_MANAGED",
	"publicKey": "...",
	"validUntil": "2024-03-22T18:05:41Z"
}`

var saKey = fmt.Sprintf(saKeyStrPattern, uuid.New().String(), uuid.New().String(), uuid.New().String())

// Error cases are tested in the noAuth, TokenAuth and DefaultAuth functions
func TestSetupAuth(t *testing.T) {
	for _, test := range []struct {
		desc     string
		config   *config.Configuration
		setToken bool
		setPath  bool
		isValid  bool
	}{
		{
			desc:     "default_config",
			config:   nil,
			setToken: true,
			setPath:  false,
			isValid:  true,
		},
		{
			desc:     "valid_path_to_file",
			config:   nil,
			setToken: false,
			setPath:  true,
			isValid:  true,
		},
		{
			desc: "custom_config_token",
			config: &config.Configuration{
				Token: "token",
			},
			setToken: false,
			setPath:  false,
			isValid:  true,
		},
		{
			desc: "custom_config_path",
			config: &config.Configuration{
				CredentialsFilePath: "test_resources/test_credentials_bar.json",
			},
			setToken: false,
			setPath:  false,
			isValid:  true,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if test.setToken {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "test-token")
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "")
			}

			if test.setPath {
				t.Setenv("STACKIT_CREDENTIALS_PATH", "test_resources/test_credentials_bar.json")
			} else {
				t.Setenv("STACKIT_CREDENTIALS_PATH", "test-path")
			}

			t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", "test-email")

			authRoundTripper, err := SetupAuth(test.config)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && authRoundTripper == nil {
				t.Fatalf("Roundtripper returned is nil for valid test case")
			}
		})
	}
}

func TestReadCredentials(t *testing.T) {
	for _, test := range []struct {
		desc               string
		path               string
		pathEnv            string
		credentialType     string
		isValid            bool
		expectedCredential string
	}{
		{
			desc:               "valid_path_argument_token",
			path:               "test_resources/test_credentials_bar.json",
			pathEnv:            "",
			credentialType:     tokenCredentialType,
			isValid:            true,
			expectedCredential: "bar_token",
		},
		{
			desc:               "valid_path_argument_sa_key",
			path:               "test_resources/test_credentials_bar.json",
			pathEnv:            "",
			credentialType:     serviceAccountKeyCredentialType,
			isValid:            true,
			expectedCredential: "bar_sa_key",
		},
		{
			desc:               "valid_path_argument_private_key",
			path:               "test_resources/test_credentials_bar.json",
			pathEnv:            "",
			credentialType:     privateKeyCredentialType,
			isValid:            true,
			expectedCredential: "bar_private_key",
		},
		{
			desc:               "valid_path_env",
			path:               "",
			pathEnv:            "test_resources/test_credentials_bar.json",
			credentialType:     tokenCredentialType,
			isValid:            true,
			expectedCredential: "bar_token",
		},
		{
			desc:               "path_over_env",
			path:               "test_resources/test_credentials_bar.json",
			pathEnv:            "test_resources/test_credentials_foo.json",
			credentialType:     tokenCredentialType,
			isValid:            true,
			expectedCredential: "bar_token",
		},
		{
			desc:               "invalid_structure",
			path:               "test_resources/test_invalid_structure.json",
			credentialType:     tokenCredentialType,
			pathEnv:            "",
			isValid:            false,
			expectedCredential: "",
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			t.Setenv("STACKIT_CREDENTIALS_PATH", test.pathEnv)

			var credential string
			credentials, err := readCredentialsFile(test.path)
			if err == nil {
				credential, err = readCredential(test.credentialType, credentials)
			}

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && credential != test.expectedCredential {
				t.Fatalf("Token is not correct. Expected %s, got %s", test.expectedCredential, credential)
			}
		})
	}
}

func TestDefaultAuth(t *testing.T) {
	for _, test := range []struct {
		desc     string
		setToken bool
		isValid  bool
	}{
		{
			desc:     "valid_case",
			setToken: true,
			isValid:  true,
		},
		{
			desc:     "no_token",
			setToken: false,
			isValid:  false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if test.setToken {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "test-token")
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "")
			}
			t.Setenv("STACKIT_CREDENTIALS_PATH", "test-path")
			t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", "test-email")

			// Get the default authentication client and ensure that it's not nil
			authClient, err := DefaultAuth(nil)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && authClient == nil {
				t.Fatalf("Client returned is nil for valid test case")
			}
		})
	}
}

func TestTokenAuth(t *testing.T) {
	for _, test := range []struct {
		desc    string
		token   string
		isValid bool
	}{
		{
			desc:    "valid_case",
			token:   "token",
			isValid: true,
		},
		{
			desc:    "no_token",
			token:   "",
			isValid: false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "")
			t.Setenv("STACKIT_CREDENTIALS_PATH", "test-path")

			cfg := &config.Configuration{
				Token: test.token,
			}
			// Get the token authentication client and ensure that it's not nil
			authClient, err := TokenAuth(cfg)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && authClient == nil {
				t.Fatalf("Client returned is nil for valid test case")
			}
		})
	}
}

func TestKeyAuth(t *testing.T) {
	privateKey, err := generatePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key for testing")
	}

	for _, test := range []struct {
		desc              string
		serviceAccountKey string
		privateKey        string
		isValid           bool
	}{
		{
			desc:              "valid_case",
			serviceAccountKey: saKey,
			privateKey:        string(privateKey),
			isValid:           true,
		},
		{
			desc:              "no_sa_key",
			serviceAccountKey: "",
			privateKey:        string(privateKey),
			isValid:           false,
		},
		{
			desc:              "no_private_key",
			serviceAccountKey: "no_sa_key",
			privateKey:        "",
			isValid:           false,
		},
		{
			desc:              "no_keys",
			serviceAccountKey: "",
			privateKey:        "",
			isValid:           false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", "")
			t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", "")
			t.Setenv("STACKIT_PRIVATE_KEY", "")
			t.Setenv("STACKIT_PRIVATE_KEY_PATH", "")

			cfg := &config.Configuration{
				ServiceAccountKey: test.serviceAccountKey,
				PrivateKey:        test.privateKey,
			}
			// Get the key authentication client and ensure that it's not nil
			authClient, err := KeyAuth(cfg)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && authClient == nil {
				t.Fatalf("Client returned is nil for valid test case")
			}
		})
	}
}

func TestNoAuth(t *testing.T) {
	for _, test := range []struct {
		desc string
	}{
		{
			desc: "valid_case",
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			// Get the default authentication client and ensure that it's not nil
			authClient, err := NoAuth(nil)
			if err != nil {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if authClient == nil {
				t.Fatalf("Client returned is nil for valid test case")
			}
		})
	}
}

func TestGetServiceAccountEmail(t *testing.T) {
	for _, test := range []struct {
		description   string
		cfg           *config.Configuration
		envEmailSet   bool
		path          string
		expectedEmail string
		isValid       bool
	}{
		{
			description: "custom_config",
			cfg: &config.Configuration{
				ServiceAccountEmail: "test_email",
			},
			path:          "",
			expectedEmail: "test_email",
			isValid:       true,
		},
		{
			description:   "config_over_env_var",
			cfg:           &config.Configuration{},
			envEmailSet:   true,
			path:          "",
			expectedEmail: "env_email",
			isValid:       true,
		},
		{
			description: "env_variable",
			cfg: &config.Configuration{
				ServiceAccountEmail: "test_email",
			},
			envEmailSet:   true,
			path:          "",
			expectedEmail: "test_email",
			isValid:       true,
		},
		{
			description:   "path",
			cfg:           &config.Configuration{},
			envEmailSet:   false,
			path:          "test_resources/test_credentials_bar.json",
			expectedEmail: "bar_email",
			isValid:       true,
		},
		{
			description:   "invalid_structure",
			cfg:           &config.Configuration{},
			envEmailSet:   false,
			path:          "test_resources/test_invalid_structure.json",
			expectedEmail: "",
			isValid:       false,
		},
	} {
		t.Run(test.description, func(t *testing.T) {
			if test.envEmailSet {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", "env_email")
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", "")
			}
			t.Setenv("STACKIT_CREDENTIALS_PATH", test.path)
			got := getServiceAccountEmail(test.cfg)
			if (got != "") && !test.isValid {
				t.Errorf("getServiceAccountEmail() did not return empty value for invalid test case")
				return
			}
			if got != test.expectedEmail {
				t.Errorf("getServiceAccountEmail() = %v, want %v", got, test.expectedEmail)
			}
		})
	}
}

func generatePrivateKey() ([]byte, error) {
	// Generate a new RSA key pair with a size of 2048 bits
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, err
	}

	// Encode the private key in PEM format
	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	// Print the private and public keys
	return pem.EncodeToMemory(privKeyPEM), nil
}

func TestGetPrivateKey(t *testing.T) {
	for _, test := range []struct {
		name                 string
		cfg                  *config.Configuration
		envPrivateKeySet     bool
		envPrivateKeyPathSet bool
		credentialsFilePath  string
		wantErr              bool
		expectedKey          string
	}{
		{
			name: "cfg_private_key",
			cfg: &config.Configuration{
				PrivateKey: "key",
			},
			wantErr:     false,
			expectedKey: "key",
		},
		{
			name: "cfg_private_key_path",
			cfg: &config.Configuration{
				PrivateKeyPath: "test_resources/test_string_key.txt",
			},
			wantErr:     false,
			expectedKey: "key",
		},
		{
			name:             "env_private_key",
			cfg:              &config.Configuration{},
			envPrivateKeySet: true,
			wantErr:          false,
			expectedKey:      "key",
		},
		{
			name:                 "env_private_key_path",
			cfg:                  &config.Configuration{},
			envPrivateKeyPathSet: true,
			wantErr:              false,
			expectedKey:          "key",
		},
		{
			name:                "credentials_file_private_key",
			cfg:                 &config.Configuration{},
			credentialsFilePath: "test_resources/test_credentials_bar.json",
			wantErr:             false,
			expectedKey:         "bar_private_key",
		},
		{
			name:        "no_private_key_provided",
			cfg:         &config.Configuration{},
			wantErr:     true,
			expectedKey: "",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("STACKIT_CREDENTIALS_PATH", test.credentialsFilePath)

			if test.envPrivateKeySet {
				t.Setenv("STACKIT_PRIVATE_KEY", "key")
			} else {
				t.Setenv("STACKIT_PRIVATE_KEY", "")
			}

			if test.envPrivateKeyPathSet {
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", "test_resources/test_invalid_private_key.txt")
			} else {
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", "")
			}

			err := getPrivateKey(test.cfg)
			if (err != nil) != test.wantErr {
				t.Errorf("getPrivateKey() error = %v, wantErr %v", err, test.wantErr)
			} else if test.cfg.PrivateKey != test.expectedKey {
				t.Errorf("getPrivateKey() = %v, want %v", test.cfg.PrivateKey, test.expectedKey)
			}
		})
	}
}

func TestGetServiceAccountKey(t *testing.T) {
	for _, test := range []struct {
		name                        string
		cfg                         *config.Configuration
		envServiceAccountKeySet     bool
		envServiceAccountKeyPathSet bool
		credentialsFilePath         string
		wantErr                     bool
		expectedKey                 string
	}{
		{
			name: "cfg_sa_key",
			cfg: &config.Configuration{
				ServiceAccountKey: "key",
			},
			wantErr:     false,
			expectedKey: "key",
		},
		{
			name: "cfg_sa_key_path",
			cfg: &config.Configuration{
				ServiceAccountKeyPath: "test_resources/test_string_key.txt",
			},
			wantErr:     false,
			expectedKey: "key",
		},
		{
			name:                    "env_sa_key",
			cfg:                     &config.Configuration{},
			envServiceAccountKeySet: true,
			wantErr:                 false,
			expectedKey:             "key",
		},
		{
			name:                        "env_sa_key_path",
			cfg:                         &config.Configuration{},
			envServiceAccountKeyPathSet: true,
			wantErr:                     false,
			expectedKey:                 "key",
		},
		{
			name:                "credentials_file_sa_key",
			cfg:                 &config.Configuration{},
			credentialsFilePath: "test_resources/test_credentials_bar.json",
			wantErr:             false,
			expectedKey:         "bar_sa_key",
		},
		{
			name:        "no_sa_key_provided",
			cfg:         &config.Configuration{},
			wantErr:     true,
			expectedKey: "",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			t.Setenv("STACKIT_CREDENTIALS_PATH", test.credentialsFilePath)

			if test.envServiceAccountKeySet {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", "key")
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", "")
			}

			if test.envServiceAccountKeyPathSet {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", "test_resources/test_string_key.txt")
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", "")
			}

			err := getServiceAccountKey(test.cfg)
			if (err != nil) != test.wantErr {
				t.Errorf("getServiceAccountKey() error = %v, wantErr %v", err, test.wantErr)
			} else if test.cfg.ServiceAccountKey != test.expectedKey {
				t.Errorf("getServiceAccountKey() = %v, want %v", test.cfg.ServiceAccountKey, test.expectedKey)
			}
		})
	}
}
