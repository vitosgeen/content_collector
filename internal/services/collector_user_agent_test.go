package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_getRandomInt(t *testing.T) {
	type args struct {
		min int
		max int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "get random int max less than min",
			args: args{
				min: 10,
				max: 0,
			},
			want: 0,
		},
		{
			name: "get random int max equal min",
			args: args{
				min: 10,
				max: 10,
			},
			want: 10,
		},
		{
			name: "get random int max greater than min",
			args: args{
				min: 10,
				max: 20,
			},
			want: 15,
		},
	}

	got := getRandomInt(tests[0].args.min, tests[0].args.max)
	assert.Equal(t, tests[0].want, got)

	got = getRandomInt(tests[1].args.min, tests[1].args.max)
	assert.Equal(t, tests[1].want, got)

	got = getRandomInt(tests[2].args.min, tests[2].args.max)
	assert.GreaterOrEqual(t, got, tests[2].args.min)
	assert.LessOrEqual(t, got, tests[2].args.max)
}
