package fx

import (
	"testing"

	fxerror "github.com/fredsh/go-fxtend/pkg/fx-error"
	"github.com/stretchr/testify/require"
)

func TestReverseMap(t *testing.T) {

	cases := []struct {
		name    string
		input   map[int]string
		want    map[string]int
		wantErr bool
	}{
		{
			name:    "empty map produce an empty map",
			input:   map[int]string{},
			want:    map[string]int{},
			wantErr: false,
		},
		{
			name: "single item map produce reversed single item map",
			input: map[int]string{
				1: "item1",
			},
			want: map[string]int{
				"item1": 1,
			},
			wantErr: false,
		},
		{
			name: "multiple items map without duplicate values produce reversed multiple items map",
			input: map[int]string{
				1: "item1",
				2: "item2",
				3: "item3",
			},
			want: map[string]int{
				"item1": 1,
				"item2": 2,
				"item3": 3,
			},
			wantErr: false,
		},
		{
			name: "multiple items map with duplicate values produce error or reversed multiple items map with overridden value",
			input: map[int]string{
				1: "item1",
				2: "item2",
				3: "item2",
			},
			want: map[string]int{
				"item1": 1,
				"item2": 3, // or 2, order cannot be guaranteed
			},
			wantErr: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resWithOverride := MapReverseOverride(tt.input)

			res, err := MapReverse(tt.input)
			resX := MapReverseX(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				require.ErrorAs(t, resX.AsError(), &fxerror.ErrDuplicateValue)
				for k := range tt.want {
					require.Contains(t, resWithOverride, k)
				}
			} else {
				require.Equal(t, tt.want, resWithOverride)

				require.NoError(t, err)
				require.NoError(t, resX.AsError())
				require.Equal(t, tt.want, res)
				require.Equal(t, tt.want, resX.Unwrap())
			}
		})
	}
}
