package fxcollection

import (
	"testing"

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
				"item2": 3,
			},
			wantErr: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			resWithOverride := ReverseMapWithOverride(tt.input)
			require.Equal(t, tt.want, resWithOverride)

			res, err := ReverseMap(tt.input)
			resX := ReverseMapX(tt.input)
			if tt.wantErr {
				require.Error(t, err)
				require.Error(t, resX.AsError())
			} else {
				require.NoError(t, err)
				require.NoError(t, resX.AsError())
				require.Equal(t, tt.want, res)
				require.Equal(t, tt.want, *resX.Unwrap())
			}
		})
	}
}
