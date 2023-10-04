package config

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConfigureRegion(t *testing.T) {
	for _, test := range []struct {
		desc            string
		cfg             *Configuration
		regionEnvVar    string
		expectedServers ServerConfigurations
		isValid         bool
	}{
		{
			desc: "valid_region",
			cfg: &Configuration{
				Region: "eu01",
				Servers: ServerConfigurations{
					ServerConfiguration{
						URL: "https://some-api.api.{region}stackit.cloud",
						Variables: map[string]ServerVariable{
							"region": {
								DefaultValue: "eu01",
								EnumValues: []string{
									"eu01.",
								},
							},
						},
					},
				},
			},
			regionEnvVar: "",
			expectedServers: ServerConfigurations{
				ServerConfiguration{
					URL: "https://some-api.api.eu01.stackit.cloud",
				},
			},
			isValid: true,
		},
		{
			desc: "valid_global",
			cfg: &Configuration{
				Region: "",
				Servers: ServerConfigurations{
					ServerConfiguration{
						URL: "https://some-api.api.{region}stackit.cloud",
						Variables: map[string]ServerVariable{
							"region": {
								DefaultValue: "",
								EnumValues:   []string{},
							},
						},
					},
				},
			},
			regionEnvVar: "",
			expectedServers: ServerConfigurations{
				ServerConfiguration{
					URL: "https://some-api.api.stackit.cloud",
				},
			},
			isValid: true,
		},
		{
			desc: "valid_global_with_specified_region",
			cfg: &Configuration{
				Region: "eu01",
				Servers: ServerConfigurations{
					ServerConfiguration{
						URL: "https://some-api.api.{region}stackit.cloud",
						Variables: map[string]ServerVariable{
							"region": {
								DefaultValue: "",
								EnumValues:   []string{},
							},
						},
					},
				},
			},
			regionEnvVar: "",
			expectedServers: ServerConfigurations{
				ServerConfiguration{
					URL: "https://some-api.api.stackit.cloud",
				},
			},
			isValid: true,
		},
		{
			desc: "env_var_valid_region",
			cfg: &Configuration{
				Region: "",
				Servers: ServerConfigurations{
					ServerConfiguration{
						URL: "https://some-api.api.{region}stackit.cloud",
						Variables: map[string]ServerVariable{
							"region": {
								DefaultValue: "eu01",
								EnumValues: []string{
									"eu01.",
								},
							},
						},
					},
				},
			},
			regionEnvVar: "eu01",
			expectedServers: ServerConfigurations{
				ServerConfiguration{
					URL: "https://some-api.api.eu01.stackit.cloud",
				},
			},
			isValid: true,
		},
		{
			desc: "invalid_region",
			cfg: &Configuration{
				Region: "some region",
				Servers: ServerConfigurations{
					ServerConfiguration{
						URL: "https://some-api.api.{region}stackit.cloud",
						Variables: map[string]ServerVariable{
							"region": {
								DefaultValue: "eu01",
								EnumValues:   []string{},
							},
						},
					},
				},
			},
			regionEnvVar:    "",
			expectedServers: nil,
			isValid:         false,
		},
		{
			desc: "no_region",
			cfg: &Configuration{
				Region: "",
				Servers: ServerConfigurations{
					ServerConfiguration{
						URL: "https://some-api.api.{region}stackit.cloud",
						Variables: map[string]ServerVariable{
							"region": {
								DefaultValue: "eu01",
								EnumValues:   []string{},
							},
						},
					},
				},
			},
			regionEnvVar:    "",
			expectedServers: nil,
			isValid:         false,
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			if test.regionEnvVar != "" {
				t.Setenv("STACKIT_REGION", test.regionEnvVar)
			}

			err := ConfigureRegion(test.cfg)

			if err != nil && test.isValid {
				t.Fatalf("Test returned error on valid test case: %v", err)
			}

			if err == nil && !test.isValid {
				t.Fatalf("Test didn't return error on invalid test case")
			}

			if test.isValid && !cmp.Equal(test.cfg.Servers, test.expectedServers) {
				t.Fatalf("The server is wrong. Expected %+v, got %+v", test.expectedServers, test.cfg.Servers)
			}
		})
	}
}
