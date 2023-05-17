package rand

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		setup   func()
		check   func(*testing.T, []byte)
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "filled_by_random",
			args: args{
				b: make([]byte, 10),
			},
			setup: func() {
				Reset()
			},
			check: func(t *testing.T, slice []byte) {
				hasNonZero := false
				for i := range slice {
					if slice[i] > 0 {
						hasNonZero = true
						break
					}
				}
				assert.True(t, hasNonZero)
			},
			want:    10,
			wantErr: assert.NoError,
		},
		{
			name: "filled_by_random_only_len_of_slice",
			args: args{
				b: make([]byte, 5, 10),
			},
			setup: func() {
				Reset()
			},
			check: func(t *testing.T, slice []byte) {
				hasNonZero := false
				for i := range slice {
					if slice[i] > 0 {
						hasNonZero = true
						break
					}
				}
				assert.True(t, hasNonZero)
			},
			want:    5,
			wantErr: assert.NoError,
		},
		{
			name: "filled_manually",
			args: args{
				b: make([]byte, 10),
			},
			setup: func() {
				Reset()
				Set([]byte{1, 7, 33, 41, 29, 77, 51, 28, 19, 66})
			},
			check: func(t *testing.T, slice []byte) {
				assert.Equal(t, []byte{1, 7, 33, 41, 29, 77, 51, 28, 19, 66}, slice)
			},
			want:    10,
			wantErr: assert.NoError,
		},
		{
			name: "filled_manually_only_len_of_slice",
			args: args{
				b: make([]byte, 5, 10),
			},
			setup: func() {
				Reset()
				Set([]byte{1, 7, 33, 41, 29, 77, 51, 28, 19, 66})
			},
			check: func(t *testing.T, slice []byte) {
				assert.Equal(t, []byte{1, 7, 33, 41, 29}, slice)
			},
			want:    5,
			wantErr: assert.NoError,
		},
		{
			name: "filled_manually_next_values",
			args: args{
				b: make([]byte, 10),
			},
			setup: func() {
				Reset()
				Set([]byte{1, 7, 33, 41, 29, 77, 51, 28, 19, 66, 84, 31, 90, 92, 99})

				tmp := make([]byte, 5)
				_, _ = Read(tmp)
			},
			check: func(t *testing.T, slice []byte) {
				assert.Equal(t, []byte{77, 51, 28, 19, 66, 84, 31, 90, 92, 99}, slice)
			},
			want:    10,
			wantErr: assert.NoError,
		},
		{
			name: "err_not_enough_data_in_buffer",
			args: args{
				b: make([]byte, 10),
			},
			setup: func() {
				Reset()
				Set([]byte{1, 7, 33, 41})
			},
			check: func(t *testing.T, slice []byte) {
				assert.Equal(t, []byte{1, 7, 33, 41, 0, 0, 0, 0, 0, 0}, slice)
			},
			want:    4,
			wantErr: assert.Error,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.setup()
			defer Reset()

			got, err := Read(tt.args.b)

			tt.wantErr(t, err)
			assert.Equal(t, tt.want, got)
			tt.check(t, tt.args.b)
		})
	}
}
