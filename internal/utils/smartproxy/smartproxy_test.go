package smartproxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmartProxyFileTxt_ParseFile(t *testing.T) {
	type fields struct {
		pathTextFile string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				pathTextFile: "../../../files/smart_proxy_data_ip.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSmartProxy(tt.fields.pathTextFile)
			got, err := s.ParseFile()

			assert.NoError(t, err)
			assert.NotEmpty(t, got)
		})
	}
}

func TestSmartProxyFileTxt_GetProxy(t *testing.T) {
	type fields struct {
		pathTextFile string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test 1",
			fields: fields{
				pathTextFile: "../../../files/smart_proxy_data_ip.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewSmartProxy(tt.fields.pathTextFile)
			got, err := s.ParseFile()
			assert.NoError(t, err)
			assert.NotEmpty(t, got)
			proxyIp, err := s.GetProxy(0)
			assert.NoError(t, err)
			assert.NotEmpty(t, proxyIp)
			proxyIp, err = s.GetProxyRandom()
			assert.NoError(t, err)
			assert.NotEmpty(t, proxyIp)
		})
	}
}
