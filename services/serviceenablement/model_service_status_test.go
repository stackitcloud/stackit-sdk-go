/*
STACKIT Service Enablement API

STACKIT Service Enablement API

API version: 2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package serviceenablement

import (
	"testing"
)

// isEnum

func TestServiceStatusEnablement_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"REQUEST"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"AUTO"`),
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
			v := ServiceStatusEnablement("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// isEnum

func TestServiceStatusLifecycle_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"FLEX"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"PROJECT"`),
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
			v := ServiceStatusLifecycle("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// isEnum

func TestServiceStatusScope_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"PRIVATE"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"PUBLIC"`),
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
			v := ServiceStatusScope("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// isEnum

func TestServiceStatusState_UnmarshalJSON(t *testing.T) {
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
				src: []byte(`"ENABLED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 2`,
			args: args{
				src: []byte(`"ENABLING"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 3`,
			args: args{
				src: []byte(`"DISABLED"`),
			},
			wantErr: false,
		},
		{
			name: `success - possible enum value no. 4`,
			args: args{
				src: []byte(`"DISABLING"`),
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
			v := ServiceStatusState("")
			if err := v.UnmarshalJSON(tt.args.src); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
