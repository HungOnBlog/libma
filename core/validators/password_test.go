package validators

import (
	"errors"
	"testing"
)

func TestIsPasswordValid(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want error
	}{
		{
			name: "empty string",
			arg:  "",
			want: errors.New("password must be between 8 and 32 characters"),
		},
		{
			name: "with less than 8 characters",
			arg:  "1234567",
			want: errors.New("password must be between 8 and 32 characters"),
		},
		{
			name: "with more than 32 characters",
			arg:  "123456789012345678901234567890123",
			want: errors.New("password must be between 8 and 32 characters"),
		},
		{
			name: "with no lowercase letter",
			arg:  "1234567890123456789012345678901",
			want: errors.New("password must contain at least one lowercase letter"),
		},
		{
			name: "with no uppercase letter",
			arg:  "1234567890123456789012345678901a",
			want: errors.New("password must contain at least one uppercase letter"),
		},
		{
			name: "with no number",
			arg:  "Aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
			want: errors.New("password must contain at least one number"),
		},
		{
			name: "with no symbol",
			arg:  "Aaaaaaaaaaaaaaaaaaaaaa1",
			want: errors.New("password must contain at least one symbol"),
		},
		{
			name: "valid password",
			arg:  "thevalidpassword1!",
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPasswordValid(tt.arg); got != nil && got.Error() != tt.want.Error() {
				t.Errorf("IsPasswordValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
