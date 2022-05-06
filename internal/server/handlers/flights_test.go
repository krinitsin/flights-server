package handlers

import (
	"math/rand"

	"github.com/stretchr/testify/assert"

	"testing"
)

func Test_identityController_GetFlightPath(t *testing.T) {
	type want struct {
		path Flight
		err  error
	}
	tests := []struct {
		name string
		args []Flight
		want want
	}{
		{
			"one flight",
			[]Flight{{
				From: "AAA",
				To:   "BBB",
			},
			},
			want{
				path: Flight{
					From: "AAA",
					To:   "BBB",
				},
				err: nil,
			},
		},
		{
			"multiple flights",
			[]Flight{
				{
					From: "EEE",
					To:   "FFF",
				},
				{
					From: "CCC",
					To:   "DDD",
				},
				{
					From: "AAA",
					To:   "BBB",
				},
				{
					From: "DDD",
					To:   "EEE",
				},
				{
					From: "BBB",
					To:   "CCC",
				},
			},
			want{
				path: Flight{
					From: "AAA",
					To:   "FFF",
				},
				err: nil,
			},
		},
		{
			"flights with cycle",
			[]Flight{
				{
					From: "AAA",
					To:   "DDD",
				},
				{
					From: "CCC",
					To:   "AAA",
				},
				{
					From: "BBB",
					To:   "CCC",
				},
				{
					From: "AAA",
					To:   "BBB",
				},
			},
			want{
				path: Flight{
					From: "AAA",
					To:   "DDD",
				},
				err: nil,
			},
		},
		{
			"flights with circle",
			[]Flight{
				{
					From: "AAA",
					To:   "BBB",
				},
				{
					From: "BBB",
					To:   "AAA",
				},
				{
					From: "CCC",
					To:   "DDD",
				},
				{
					From: "DDD",
					To:   "CCC",
				},
			},
			want{
				path: Flight{},
				err:  ErrPathIsCircle,
			},
		},
		{
			"flights with cycle",
			[]Flight{},
			want{
				path: Flight{},
				err:  ErrEmptyPath,
			},
		},
		{
			"flights are not linked ",
			[]Flight{
				{
					From: "AAA",
					To:   "BBB",
				},
				{
					From: "BBB",
					To:   "CCC",
				},
				{
					From: "DDD",
					To:   "EEE",
				},
			},
			want{
				path: Flight{},
				err:  ErrPathNotLinked,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// order shouldn't matter, add some chaos
			for n := 0; n < 100; n++ {
				dest := make([]Flight, len(tt.args))
				perm := rand.Perm(len(tt.args))
				for i, v := range perm {
					dest[v] = tt.args[i]
				}

				got, err := findPath(dest)

				assert.Equal(t, err, tt.want.err)
				// check only in happy cases
				if tt.want.err == nil {
					assert.Equal(t, got, tt.want.path)
				}
			}
		})
	}
}
