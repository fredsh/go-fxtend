package fxcollection

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestToMap(t *testing.T) {
	type item struct {
		x string
		y int
	}
	keySelector := func(i item) int {
		return i.y
	}

	cases := []struct {
		name    string
		input   []item
		want    map[int]item
		wantErr bool
	}{
		{
			name:    "empty slice produce an empty map",
			input:   []item{},
			want:    map[int]item{},
			wantErr: false,
		},
		{
			name: "single item slice produce a single item map",
			input: []item{
				{x: "item1", y: 1},
			},
			want: map[int]item{
				1: {x: "item1", y: 1},
			},
			wantErr: false,
		},
		{
			name: "multiple item without duplicate keys produce a map",
			input: []item{
				{x: "item1", y: 1},
				{x: "item2", y: 2},
				{x: "item3", y: 3},
			},
			want: map[int]item{
				1: {x: "item1", y: 1},
				2: {x: "item2", y: 2},
				3: {x: "item3", y: 3},
			},
			wantErr: false,
		},
		{
			name: "multiple item without duplicate keys produce an error or map with overridden value",
			input: []item{
				{x: "item1", y: 1},
				{x: "item2", y: 2},
				{x: "item2bis", y: 2},
			},
			want: map[int]item{
				1: {x: "item1", y: 1},
				2: {x: "item2bis", y: 2},
			},
			wantErr: true,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			preparedToMapWithOverride := PrepareToMapWithOverride(keySelector)
			resWithOverride := ToMapWithOverride(tt.input, keySelector)
			resPreparedWithOverride := preparedToMapWithOverride(tt.input)
			require.Equal(t, tt.want, resWithOverride)
			require.Equal(t, tt.want, resPreparedWithOverride)

			preparedToMap := PrepareToMap(keySelector)
			preparedToMapX := PrepareToMapX(keySelector)
			res, err := ToMap(tt.input, keySelector)
			resPrepared, errPrepared := preparedToMap(tt.input)
			resPreparedX := preparedToMapX(tt.input)
			resX := ToMapX(tt.input, keySelector)
			if tt.wantErr {
				require.Error(t, err)
				require.Error(t, errPrepared)
				require.Error(t, resX.AsError())
				require.Error(t, resPreparedX.AsError())
			} else {
				require.NoError(t, err)
				require.NoError(t, errPrepared)
				require.NoError(t, resX.AsError())
				require.NoError(t, resPreparedX.AsError())
				require.Equal(t, tt.want, res)
				require.Equal(t, tt.want, resPrepared)
				require.Equal(t, tt.want, *resX.Unwrap())
				require.Equal(t, tt.want, *resPreparedX.Unwrap())
			}
		})
	}
}
