package openapi

import (
	"github.com/scrohde/openapi/objects"
	"reflect"
	"testing"
)

func TestCompile(t *testing.T) {
	type args struct {
		rawURL string
	}
	tests := []struct {
		name    string
		args    args
		want    *objects.OpenAPI
		wantErr bool
	}{
		{
			name:    "test1",
			args:    args{rawURL: "file:///home/srohde/src/icebrg-event-api/api/v3/openapi.yaml"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compile(tt.args.rawURL)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Compile() got = %v, want %v", got, tt.want)
			}
		})
	}
}
