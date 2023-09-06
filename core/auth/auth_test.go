package auth

import (
	"testing"

	"github.com/stackitcloud/stackit-sdk-go/core/config"
)

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

func TestReadTokenFromCredentialsFile(t *testing.T) {
	for _, test := range []struct {
		desc          string
		path          string
		pathEnv       string
		isValid       bool
		expectedToken string
	}{
		{
			desc:          "valid_path_argument",
			path:          "test_resources/test_credentials_bar.json",
			pathEnv:       "",
			isValid:       true,
			expectedToken: "bar_token",
		},
		{
			desc:          "valid_path_env",
			path:          "",
			pathEnv:       "test_resources/test_credentials_bar.json",
			isValid:       true,
			expectedToken: "bar_token",
		},
		{
			desc:          "path_over_env",
			path:          "test_resources/test_credentials_bar.json",
			pathEnv:       "test_resources/test_credentials_foo.json",
			isValid:       true,
			expectedToken: "bar_token",
		},
		{
			desc:          "invalid_structure",
			path:          "test_resources/test_invalid_structure.json",
			pathEnv:       "",
			isValid:       false,
			expectedToken: "",
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			t.Setenv("STACKIT_CREDENTIALS_PATH", test.pathEnv)

			token, err := readTokenFromCredentialsFile(test.path)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && token != test.expectedToken {
				t.Fatalf("Token is not correct. Expected %s, got %s", test.expectedToken, token)
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
			// Get the token authentication client and ensure that it's not nil
			authClient, err := TokenAuth(test.token, nil)

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
