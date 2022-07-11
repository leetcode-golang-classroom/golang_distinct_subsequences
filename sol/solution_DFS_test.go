package sol

import "testing"

func BenchmarkTestDFS(b *testing.B) {
	s := "babgbag"
	t := "bag"
	for idx := 0; idx < b.N; idx++ {
		numDistinctDFS(s, t)
	}
}
func Test_numDistinctDFS(t *testing.T) {
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
			if got := numDistinctDFS(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("numDistinctDFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
