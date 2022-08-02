package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDivision1_StdAssertion(t *testing.T) {
	res, err := division(3.0, 1.0)
	if err != nil {
		t.Errorf("division() returned an error %v where none was expected", err)
	}
	if res != 3.0 {
		t.Errorf("division() = %v, want 3.0", res)
	}
}

func TestDivision1(t *testing.T) {
	res, err := division(3.0, 1.0)
	assert.NoError(t, err)
	assert.Equal(t, 3.0, res)
}

func TestDivision2(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		res, err := division(3.0, 1.0)
		assert.NoError(t, err)
		assert.Equal(t, 3.0, res)
	})
	t.Run("failure", func(t *testing.T) {
		res, err := division(3.0, 0.0)
		assert.EqualError(t, err, "division by zero")
		assert.Equal(t, 0.0, res)
	})
}

func TestDivision3(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := map[string]struct {
		args        args
		want        float64
		expectedErr string
	}{
		"success": {args: args{a: 3.0, b: 1.0}, want: 3.0},
		"failure": {args: args{a: 3.0, b: 0.0}, expectedErr: "division by zero"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got, err := division(tt.args.a, tt.args.b)
			if tt.expectedErr != "" {
				assert.EqualError(t, err, tt.expectedErr)
				assert.Equal(t, 0.0, got)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

var res float64
var err error

func BenchmarkDivision(b *testing.B) {

	// Any initialization code comes here
	var res1 float64
	var err1 error

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		res1, err1 = division(3.0, 1.0)
	}

	res = res1
	err = err1
}
