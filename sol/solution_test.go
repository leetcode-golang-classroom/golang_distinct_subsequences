package sol

import "testing"

func BenchmarkTest(b *testing.B) {
	s := "babgbag"
	t := "bag"
	for idx := 0; idx < b.N; idx++ {
		numDistinct(s, t)
	}
}
func Test_numDistinct(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "s = \"rabbbit\", t = \"rabbit\"",
			args: args{s: "rabbbit", t: "rabbit"},
			want: 3,
		},
		{
			name: "s = \"babgbag\", t = \"bag\"",
			args: args{s: "babgbag", t: "bag"},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDistinct(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinct() = %v, want %v", got, tt.want)
			}
		})
	}
}
