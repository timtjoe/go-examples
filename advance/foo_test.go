package advance

import "testing"

func TestAdd(t *testing.T) {
	cases := []struct {
		name		string 
		a,b 		int 
		want 		int
	}{
			{"two positives", 2, 3, 5},
			{"with zero", 0, 5, 5},
			{"two negatives", -2, -3, -5},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Add(tc.a, tc.b);
			got != tc.want {
				t.Errorf("got %d,  want %d", got, tc.want)
			}
		})
	}
}