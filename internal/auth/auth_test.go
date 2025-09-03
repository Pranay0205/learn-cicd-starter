package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	validHeader := make(http.Header)
	validHeader.Set("Authorization", "ApiKey abc123")

	wrongPrefixHeader := make(http.Header)
	wrongPrefixHeader.Set("Authorization", "Bearer abc123")

	missingKeyHeader := make(http.Header)
	missingKeyHeader.Set("Authorization", "ApiKey")

	tests := map[string]struct {
		input   http.Header
		want    string
		wantErr bool
	}{
		"valid_api_key": {
			input:   validHeader,
			want:    "abc123",
			wantErr: false,
		},
		"no_auth_header": {
			input:   make(http.Header),
			want:    "",
			wantErr: true,
		},
		"wrong_prefix": {
			input:   wrongPrefixHeader,
			want:    "",
			wantErr: true,
		},
		"missing_key": {
			input:   missingKeyHeader,
			want:    "",
			wantErr: true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := GetAPIKey(tc.input)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got none")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if got != tc.want {
				t.Errorf("got %q, want %q", got, tc.want)
			}
		})
	}
}
