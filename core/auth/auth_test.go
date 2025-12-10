package auth

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"net/http"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/stackitcloud/stackit-sdk-go/core/clients"
	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

func setTemporaryHome(t *testing.T) {
	old := userHomeDir
	t.Cleanup(func() {
		userHomeDir = old
	})
	userHomeDir = func() (string, error) {
		return t.TempDir(), nil
	}
}

func fixtureServiceAccountKey(mods ...func(*clients.ServiceAccountKeyResponse)) *clients.ServiceAccountKeyResponse {
	validUntil := time.Now().Add(time.Hour)
	serviceAccountKeyResponse := &clients.ServiceAccountKeyResponse{
		Active:    true,
		CreatedAt: time.Now(),
		Credentials: &clients.ServiceAccountKeyCredentials{
			Aud: "https://stackit-service-account-prod.apps.01.cf.eu01.stackit.cloud",
			Iss: "stackit@sa.stackit.cloud",
			Kid: uuid.New().String(),
			Sub: uuid.New(),
		},
		ID:           uuid.New(),
		KeyAlgorithm: "RSA_2048",
		KeyOrigin:    "USER_PROVIDED",
		KeyType:      "USER_MANAGED",
		PublicKey:    "...",
		ValidUntil:   &validUntil,
	}
	for _, mod := range mods {
		mod(serviceAccountKeyResponse)
	}
	return serviceAccountKeyResponse
}

// helper function to create a credentials.json file with saKey and private key
func createCredentialsKeyJson(serviceAccountKey, privateKey string) ([]byte, error) {
	tempMap := map[string]interface{}{}
	tempMap["STACKIT_SERVICE_ACCOUNT_KEY"] = serviceAccountKey
	tempMap["STACKIT_PRIVATE_KEY"] = privateKey
	return json.Marshal(tempMap)
}

// Error cases are tested in the NoAuth, KeyAuth, TokenAuth and DefaultAuth functions
func TestSetupAuth(t *testing.T) {
	privateKey, err := generatePrivateKey()
	if err != nil {
		t.Fatalf("Generating private key: %s", err)
	}

	// Create a temporary file
	privateKeyFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		err := os.Remove(privateKeyFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	// Write some text to the file
	_, errs = privateKeyFile.WriteString(string(privateKey))
	if errs != nil {
		t.Fatalf("Writing private key to temporary file: %s", err)
	}

	defer func() {
		err := privateKeyFile.Close()
		if err != nil {
			t.Fatalf("Closing temporary file: %s", err)
		}
	}()

	// Create a temporary file
	saKeyFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		err := os.Remove(saKeyFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	// Write the service account key to the file
	saKey, err := json.Marshal(fixtureServiceAccountKey())
	if err != nil {
		t.Fatalf("unmarshalling service account key: %s", err)
	}
	_, errs = saKeyFile.WriteString(string(saKey))
	if errs != nil {
		t.Fatalf("Writing private key to temporary file: %s", err)
	}

	defer func() {
		err := saKeyFile.Close()
		if err != nil {
			t.Fatalf("Closing temporary file: %s", err)
		}
	}()

	// create a wif assertion file
	wifAssertionFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		_ = wifAssertionFile.Close()
		err := os.Remove(wifAssertionFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		Subject:   "sub",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("Removing temporary file: %s", err)
	}

	_, errs = wifAssertionFile.WriteString(string(token))
	if errs != nil {
		t.Fatalf("Writing wif assertion to temporary file: %s", err)
	}

	// create a credentials file with saKey and private key
	credentialsKeyFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		_ = credentialsKeyFile.Close()
		err := os.Remove(credentialsKeyFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	credKeyJson, err := createCredentialsKeyJson(string(saKey), privateKey)
	if err != nil {
		t.Fatalf("createCredentialsKeyJson: %s", err)
	}
	_, errs = credentialsKeyFile.WriteString(string(credKeyJson))
	if errs != nil {
		t.Fatalf("Writing credentials json to temporary file: %s", err)
	}

	for _, test := range []struct {
		desc                        string
		config                      *config.Configuration
		setToken                    bool
		setWorkloadIdentity         bool
		setKeys                     bool
		setKeyPaths                 bool
		setCredentialsFilePathToken bool
		setCredentialsFilePathKey   bool
	}{
		{
			desc:                "wif_config",
			config:              nil,
			setWorkloadIdentity: true,
		},
		{
			desc:                        "token_config",
			config:                      nil,
			setToken:                    true,
			setCredentialsFilePathToken: false,
		},
		{
			desc:                        "key_config",
			config:                      nil,
			setKeys:                     true,
			setCredentialsFilePathToken: false,
		},
		{
			desc:                        "key_config_path",
			config:                      nil,
			setKeys:                     false,
			setKeyPaths:                 true,
			setCredentialsFilePathToken: false,
		},
		{
			desc:                      "key_config_credentials_path",
			config:                    nil,
			setKeys:                   false,
			setKeyPaths:               false,
			setCredentialsFilePathKey: true,
		},
		{
			desc:                        "valid_path_to_file",
			config:                      nil,
			setToken:                    false,
			setCredentialsFilePathToken: true,
		},
		{
			desc: "custom_config_token",
			config: &config.Configuration{
				Token: "token",
			},
			setToken:                    false,
			setCredentialsFilePathToken: false,
		},
		{
			desc: "custom_config_path",
			config: &config.Configuration{
				CredentialsFilePath: "test_resources/test_credentials_bar.json",
			},
			setToken:                    false,
			setCredentialsFilePathToken: false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			setTemporaryHome(t)
			if test.setKeys {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", string(saKey))
				t.Setenv("STACKIT_PRIVATE_KEY", privateKey)
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", "")
				t.Setenv("STACKIT_PRIVATE_KEY", "")
			}

			if test.setKeyPaths {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", saKeyFile.Name())
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", privateKeyFile.Name())
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", "")
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", "")
			}

			if test.setToken {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "test-token")
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "")
			}

			if test.setCredentialsFilePathToken {
				t.Setenv("STACKIT_CREDENTIALS_PATH", "test_resources/test_credentials_bar.json")
			} else if test.setCredentialsFilePathKey {
				t.Setenv("STACKIT_CREDENTIALS_PATH", credentialsKeyFile.Name())
			} else {
				t.Setenv("STACKIT_CREDENTIALS_PATH", "")
			}

			if test.setWorkloadIdentity {
				t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", wifAssertionFile.Name())
			} else {
				t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", "")
			}

			t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", "test-email")

			authRoundTripper, err := SetupAuth(test.config)

			if err != nil {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if authRoundTripper == nil {
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
		credentialType     credentialType
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
			setTemporaryHome(t)
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
	privateKey, err := generatePrivateKey()
	if err != nil {
		t.Fatalf("Generating private key: %s", err)
	}

	// Create a temporary file
	privateKeyFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		err := os.Remove(privateKeyFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	// Write some text to the file
	_, errs = privateKeyFile.WriteString(string(privateKey))
	if errs != nil {
		t.Fatalf("Writing private key to temporary file: %s", err)
	}

	defer func() {
		err := privateKeyFile.Close()
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	// Create a temporary file
	saKeyFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		_ = saKeyFile.Close()
		err := os.Remove(saKeyFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	// Write the service account key to the file
	saKey, err := json.Marshal(fixtureServiceAccountKey())
	if err != nil {
		t.Fatalf("unmarshalling service account key: %s", err)
	}
	_, errs = saKeyFile.WriteString(string(saKey))
	if errs != nil {
		t.Fatalf("Writing private key to temporary file: %s", err)
	}

	// create a wif assertion file
	wifAssertionFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		_ = wifAssertionFile.Close()
		err := os.Remove(wifAssertionFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute)),
		Subject:   "sub",
	}).SignedString([]byte("test"))
	if err != nil {
		t.Fatalf("Removing temporary file: %s", err)
	}

	_, errs = wifAssertionFile.WriteString(string(token))
	if errs != nil {
		t.Fatalf("Writing wif assertion to temporary file: %s", err)
	}

	// create a credentials file with saKey and private key
	credentialsKeyFile, errs := os.CreateTemp("", "temp-*.txt")
	if errs != nil {
		t.Fatalf("Creating temporary file: %s", err)
	}
	defer func() {
		_ = credentialsKeyFile.Close()
		err := os.Remove(credentialsKeyFile.Name())
		if err != nil {
			t.Fatalf("Removing temporary file: %s", err)
		}
	}()

	credKeyJson, err := createCredentialsKeyJson(string(saKey), privateKey)
	if err != nil {
		t.Fatalf("createCredentialsKeyJson: %s", err)
	}
	_, errs = credentialsKeyFile.WriteString(string(credKeyJson))
	if errs != nil {
		t.Fatalf("Writing credentials json to temporary file: %s", err)
	}

	for _, test := range []struct {
		desc                      string
		setToken                  bool
		setKeyPaths               bool
		setKeys                   bool
		setCredentialsFilePathKey bool
		setWorkloadIdentity       bool
		isValid                   bool
		expectedFlow              string
	}{
		{
			desc:         "token",
			setToken:     true,
			isValid:      true,
			expectedFlow: "token",
		},
		{
			desc:                "wif_precedes_key_precedes_token",
			setToken:            true,
			setKeyPaths:         true,
			setWorkloadIdentity: true,
			isValid:             true,
			expectedFlow:        "wif",
		},
		{
			desc:         "key_precedes_token",
			setToken:     true,
			setKeyPaths:  true,
			isValid:      true,
			expectedFlow: "key",
		},
		{
			desc:     "no_credentials",
			setToken: false,
			isValid:  false,
		},
		{
			desc:         "use keys via environment",
			setKeys:      true,
			setToken:     false,
			isValid:      true,
			expectedFlow: "key",
		},
		{
			desc:                      "use keys via credentials file",
			setKeys:                   false,
			setToken:                  false,
			setCredentialsFilePathKey: true,
			isValid:                   true,
			expectedFlow:              "key",
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			setTemporaryHome(t)
			if test.setKeyPaths {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", saKeyFile.Name())
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", privateKeyFile.Name())
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", "")
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", "")
			}

			if test.setKeys {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", string(saKey))
				t.Setenv("STACKIT_PRIVATE_KEY", privateKey)
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", "")
				t.Setenv("STACKIT_PRIVATE_KEY", "")
			}

			if test.setCredentialsFilePathKey {
				t.Setenv("STACKIT_CREDENTIALS_PATH", credentialsKeyFile.Name())
			} else {
				t.Setenv("STACKIT_CREDENTIALS_PATH", "test-path")
			}

			if test.setToken {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "test-token")
			} else {
				t.Setenv("STACKIT_SERVICE_ACCOUNT_TOKEN", "")
			}

			if test.setWorkloadIdentity {
				t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", wifAssertionFile.Name())
			} else {
				t.Setenv("STACKIT_FEDERATED_TOKEN_FILE", "")
			}

			t.Setenv("STACKIT_SERVICE_ACCOUNT_EMAIL", "test-email")

			// Get the default authentication client and ensure that it's not nil
			authClient, err := DefaultAuth(nil)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid {
				if authClient == nil {
					t.Fatalf("Client returned is nil for valid test case")
				}
				switch test.expectedFlow {
				case "token":
					if _, ok := authClient.(*clients.TokenFlow); !ok {
						t.Fatalf("Expected token flow, got %s", reflect.TypeOf(authClient))
					}
				case "key":
					if _, ok := authClient.(*clients.KeyFlow); !ok {
						t.Fatalf("Expected key flow, got %s", reflect.TypeOf(authClient))
					}
				case "wif":
					if _, ok := authClient.(*clients.WorkloadIdentityFederationFlow); !ok {
						t.Fatalf("Expected key flow, got %s", reflect.TypeOf(authClient))
					}
				}
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
			setTemporaryHome(t)
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
	includedPrivateKey, err := generatePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key for testing")
	}
	configuredPrivateKey, err := generatePrivateKey()
	if err != nil {
		t.Fatalf("Failed to generate private key for testing")
	}

	for _, test := range []struct {
		desc                 string
		serviceAccountKey    *clients.ServiceAccountKeyResponse
		includedPrivateKey   *string
		configuredPrivateKey string
		envVarPrivateKey     string
		expectedPrivateKey   string
		isValid              bool
	}{
		{
			desc:                 "configured_private_key",
			serviceAccountKey:    fixtureServiceAccountKey(),
			configuredPrivateKey: string(configuredPrivateKey),
			expectedPrivateKey:   string(configuredPrivateKey),
			isValid:              true,
		},
		{
			desc:               "included_private_key",
			serviceAccountKey:  fixtureServiceAccountKey(),
			includedPrivateKey: &includedPrivateKey,
			expectedPrivateKey: includedPrivateKey,
			isValid:            true,
		},
		{
			desc:                 "empty_configured_use_included_private_key",
			serviceAccountKey:    fixtureServiceAccountKey(),
			includedPrivateKey:   &includedPrivateKey,
			configuredPrivateKey: "",
			expectedPrivateKey:   includedPrivateKey,
			isValid:              true,
		},
		{
			desc:                 "configured_over_included_private_key",
			serviceAccountKey:    fixtureServiceAccountKey(),
			includedPrivateKey:   &includedPrivateKey,
			configuredPrivateKey: configuredPrivateKey,
			expectedPrivateKey:   configuredPrivateKey,
			isValid:              true,
		},
		{
			desc:                 "no_sa_key",
			serviceAccountKey:    nil,
			configuredPrivateKey: configuredPrivateKey,
			isValid:              false,
		},
		{
			desc:                 "missing_private_key",
			serviceAccountKey:    fixtureServiceAccountKey(),
			configuredPrivateKey: "",
			isValid:              false,
		},
		{
			desc:                 "no_keys",
			serviceAccountKey:    nil,
			configuredPrivateKey: "",
			isValid:              false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			setTemporaryHome(t)
			t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY", "")
			t.Setenv("STACKIT_SERVICE_ACCOUNT_KEY_PATH", "")
			t.Setenv("STACKIT_PRIVATE_KEY", "")
			t.Setenv("STACKIT_PRIVATE_KEY_PATH", "")

			var saKey string
			if test.serviceAccountKey != nil {
				test.serviceAccountKey.Credentials.PrivateKey = test.includedPrivateKey

				saKeyBytes, err := json.Marshal(test.serviceAccountKey)
				if err != nil {
					t.Fatalf("unmarshalling service account key: %s", err)
				}

				saKey = string(saKeyBytes)
			}

			cfg := &config.Configuration{
				ServiceAccountKey: saKey,
				PrivateKey:        test.configuredPrivateKey,
			}
			// Get the key authentication client and ensure that it's not nil
			authClient, err := KeyAuth(cfg)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid {
				if authClient == nil {
					t.Fatalf("Client returned is nil for valid test case")
				}

				keyFlow, ok := authClient.(*clients.KeyFlow)
				if !ok {
					t.Fatalf("Could not convert authClient to KeyFlow")
				}
				if keyFlow.GetConfig().PrivateKey != test.expectedPrivateKey {
					t.Fatalf("The private key is wrong: expected %s, got %s", test.expectedPrivateKey, keyFlow.GetConfig().PrivateKey)
				}
			}
		})
	}
}

func TestKeyAuthPemInsteadOfJsonKeyErrorHandling(t *testing.T) {
	cfg := &config.Configuration{
		ServiceAccountKey: "-----BEGIN PRIVATE KEY",
	}
	_, err := KeyAuth(cfg)
	if err == nil {
		t.Fatalf("error expected")
	}
	if !strings.HasSuffix(err.Error(), "Please provide it in JSON format.") {
		t.Fatalf("expected to end with JSON format hint: %s", err)
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
			setTemporaryHome(t) // Get the default authentication client and ensure that it's not nil
			authClient, err := NoAuth()
			if err != nil {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if authClient == nil {
				t.Fatalf("Client returned is nil for valid test case")
			}
		})
	}
}

func TestNoAuthWithConfig(t *testing.T) {
	for _, test := range []struct {
		desc string
	}{
		{
			desc: "valid_case",
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			setTemporaryHome(t) // Get the default authentication client and ensure that it's not nil
			authClient, err := NoAuth(&config.Configuration{HTTPClient: http.DefaultClient})
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
			setTemporaryHome(t)
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

func generatePrivateKey() (string, error) {
	// Generate a new RSA key pair with a size of 2048 bits
	privKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", err
	}

	// Encode the private key in PEM format
	privKeyPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privKey),
	}

	// Print the private and public keys
	return string(pem.EncodeToMemory(privKeyPEM)), nil
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
			name:                 "env_private_key_path",
			cfg:                  &config.Configuration{},
			envPrivateKeyPathSet: true,
			wantErr:              false,
			expectedKey:          "key",
		},
		{
			name:                "credentials_file_private_key_path",
			cfg:                 &config.Configuration{},
			credentialsFilePath: "test_resources/test_credentials_foo.json",
			wantErr:             false,
			expectedKey:         "foo_key",
		},
		{
			name: "cfg_private_key_precedes_path",
			cfg: &config.Configuration{
				PrivateKey:     "cfg_key",
				PrivateKeyPath: "test_resources/test_string_key.txt",
			},
			wantErr:     false,
			expectedKey: "cfg_key",
		},
		{
			name: "cfg_private_key_precedes_env",
			cfg: &config.Configuration{
				PrivateKey: "cfg_key",
			},
			envPrivateKeyPathSet: true,
			wantErr:              false,
			expectedKey:          "cfg_key",
		},
		{
			name: "cfg_private_key_precedes_creds_file",
			cfg: &config.Configuration{
				PrivateKey: "cfg_key",
			},
			credentialsFilePath: "test_resources/test_credentials_foo.json",
			wantErr:             false,
			expectedKey:         "cfg_key",
		},
		{
			name:                 "env_private_key_precedes_creds_file",
			cfg:                  &config.Configuration{},
			envPrivateKeyPathSet: true,
			credentialsFilePath:  "test_resources/test_credentials_foo.json",
			wantErr:              false,
			expectedKey:          "key",
		},
		{
			name:        "no_private_key_provided",
			cfg:         &config.Configuration{},
			wantErr:     true,
			expectedKey: "",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			setTemporaryHome(t)
			t.Setenv("STACKIT_CREDENTIALS_PATH", test.credentialsFilePath)

			if test.envPrivateKeyPathSet {
				t.Setenv("STACKIT_PRIVATE_KEY_PATH", "test_resources/test_string_key.txt")
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
			name:                        "env_sa_key_path",
			cfg:                         &config.Configuration{},
			envServiceAccountKeyPathSet: true,
			wantErr:                     false,
			expectedKey:                 "key",
		},
		{
			name:                "credentials_file_sa_key_path",
			cfg:                 &config.Configuration{},
			credentialsFilePath: "test_resources/test_credentials_foo.json",
			wantErr:             false,
			expectedKey:         "foo_key",
		},
		{
			name: "cfg_sa_key_precedes_path",
			cfg: &config.Configuration{
				ServiceAccountKey:     "cfg_key",
				ServiceAccountKeyPath: "test_resources/test_string_key.txt",
			},
			wantErr:     false,
			expectedKey: "cfg_key",
		},
		{
			name: "cfg_sa_key_precedes_env",
			cfg: &config.Configuration{
				ServiceAccountKey: "cfg_key",
			},
			envServiceAccountKeyPathSet: true,
			wantErr:                     false,
			expectedKey:                 "cfg_key",
		},
		{
			name: "cfg_sa_key_precedes_creds_file",
			cfg: &config.Configuration{
				ServiceAccountKey: "cfg_key",
			},
			credentialsFilePath: "test_resources/test_credentials_foo.json",
			wantErr:             false,
			expectedKey:         "cfg_key",
		},
		{
			name:                        "env_sa_key_precedes_creds_file",
			cfg:                         &config.Configuration{},
			envServiceAccountKeyPathSet: true,
			credentialsFilePath:         "test_resources/test_credentials_foo.json",
			wantErr:                     false,
			expectedKey:                 "key",
		},
		{
			name:        "no_sa_key_provided",
			cfg:         &config.Configuration{},
			wantErr:     true,
			expectedKey: "",
		},
	} {
		t.Run(test.name, func(t *testing.T) {
			setTemporaryHome(t)
			t.Setenv("STACKIT_CREDENTIALS_PATH", test.credentialsFilePath)

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
