package validators

import "testing"

func TestIsEmailValid(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "empty string",
			arg:  "",
			want: false,
		},
		{
			name: "invalid email",
			arg:  "test",
			want: false,
		},
		{
			name: "valid email",
			arg:  "thevalidemail@gmail.com",
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmailValid(tt.arg); got != tt.want {
				t.Errorf("IsEmailValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
