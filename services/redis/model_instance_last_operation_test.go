/*
STACKIT Redis API

The STACKIT Redis API provides endpoints to list service offerings, manage service instances and service credentials within STACKIT portal projects.

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package redis

import (
	"testing"
)

// isEnum

func TestInstanceLastOperationState_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"in progress"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"succeeded"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 3`,
			args: args{
				src: []byte(`"failed"`),
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
			v := InstanceLastOperationState("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// isEnum

func TestInstanceLastOperationTypes_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"create"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"update"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 3`,
			args: args{
				src: []byte(`"delete"`),
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
			v := InstanceLastOperationTypes("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
