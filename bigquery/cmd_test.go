package bigquery

import "testing"

func Test_tokenize(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 string
		want2 string
	}{
		{"a.b.c", args{"a.b.c"}, "a", "b", "c"},
		{"a:b.c", args{"a:b.c"}, "a", "b", "c"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tokenize(tt.args.t)
			if got != tt.want {
				t.Errorf("tokenize() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("tokenize() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("tokenize() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
