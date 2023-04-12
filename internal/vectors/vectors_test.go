package vectors

import (
	"reflect"
	"testing"
)

func TestConvertPlayerToVector(t *testing.T) {
	// Test cases
	testCases := []struct {
		player Player
		want   []float32
	}{
		{Player{Level: 10, Kost: 0.76, Rank: 3, RankPoints: 500}, []float32{10.0, 0.76, 3.0, 500.0}},
		{Player{Level: 5, Kost: 10.5, Rank: 1, RankPoints: 1000}, []float32{5.0, 10.5, 1.0, 1000.0}},
		{Player{Level: 20, Kost: 50.0, Rank: 5, RankPoints: 250}, []float32{20.0, 50.0, 5.0, 250.0}},
	}

	// Loop through test cases
	for _, testCase := range testCases {
		got := ConvertPlayerToVector(testCase.player)

		// Check if the output is correct
		if !reflect.DeepEqual(got, testCase.want) {
			t.Errorf("ConvertPlayerToVector(%v) = %v, want %v", testCase.player, got, testCase.want)
		}
	}
}
