/*
STACKIT Observability API

API endpoints for Observability on STACKIT

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package observability

import (
	"testing"
)

// isEnum

func TestProjectInstanceFullStatus_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"CREATING"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"CREATE_SUCCEEDED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 3`,
			args: args{
				src: []byte(`"CREATE_FAILED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 4`,
			args: args{
				src: []byte(`"DELETING"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 5`,
			args: args{
				src: []byte(`"DELETE_SUCCEEDED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 6`,
			args: args{
				src: []byte(`"DELETE_FAILED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 7`,
			args: args{
				src: []byte(`"UPDATING"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 8`,
			args: args{
				src: []byte(`"UPDATE_SUCCEEDED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 9`,
			args: args{
				src: []byte(`"UPDATE_FAILED"`),
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
			v := ProjectInstanceFullStatus("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
