/*
SKE-API

The SKE API provides endpoints to create, update, delete clusters within STACKIT portal projects and to trigger further cluster management tasks.

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ske

import (
	"testing"
)

// isEnum

func TestRuntimeErrorCode_UnmarshalJSON(t *testing.T) {
	type args struct {
		src []byte
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: `success - possible enum value no. 1`,
			args: args{
				src: []byte(`"SKE_UNSPECIFIED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"SKE_TMP_AUTH_ERROR"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 3`,
			args: args{
				src: []byte(`"SKE_QUOTA_EXCEEDED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 4`,
			args: args{
				src: []byte(`"SKE_OBSERVABILITY_INSTANCE_NOT_FOUND"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 5`,
			args: args{
				src: []byte(`"SKE_RATE_LIMITS"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 6`,
			args: args{
				src: []byte(`"SKE_INFRA_ERROR"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 7`,
			args: args{
				src: []byte(`"SKE_REMAINING_RESOURCES"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 8`,
			args: args{
				src: []byte(`"SKE_CONFIGURATION_PROBLEM"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 9`,
			args: args{
				src: []byte(`"SKE_UNREADY_NODES"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 10`,
			args: args{
				src: []byte(`"SKE_API_SERVER_ERROR"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 11`,
			args: args{
				src: []byte(`"SKE_DNS_ZONE_NOT_FOUND"`),
			},
			wantErr: false,
		},
		{
			name: "fail",
			args: args{
				src: []byte("\"FOOBAR\""),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := RuntimeErrorCode("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
