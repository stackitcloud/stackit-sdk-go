package iaas

import "testing"

func TestNetworkIPv4_UnmarshalJSON(t *testing.T) {
	type fields struct {
		Gateway     NullableString
		Nameservers []string
		Prefixes    []string
		PublicIp    *string
	}
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "simple",
			fields: fields{
				Gateway:     NullableString{isSet: false},
				Nameservers: []string{"8.8.8.8"},
				Prefixes:    nil,
				PublicIp:    nil,
			},
			args: args{
				data: []byte(`{"nameservers":["8.8.8.8"], "prefixes": null}`),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			o := &NetworkIPv4{
				Gateway:     tt.fields.Gateway,
				Nameservers: tt.fields.Nameservers,
				Prefixes:    tt.fields.Prefixes,
				PublicIp:    tt.fields.PublicIp,
			}
			if err := o.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
