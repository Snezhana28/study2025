package uniq

import (
	"flag"
	"reflect"
	"testing"
)

func TestGetOptions(t *testing.T) {
	tests := []struct {
		name     string
		args     []string
		expected Options
		wantErr  bool
	}{
		{
			name:     "только -c",
			args:     []string{"-c"},
			expected: Options{c: true},
		},
		{
			name:    "одновременно -c и -d",
			args:    []string{"-c", "-d"},
			wantErr: true,
		},
		{
			name:     "пропуск -f 2",
			args:     []string{"-f", "2"},
			expected: Options{f: 2},
		},
		{
			name:     "поддержка -s 3",
			args:     []string{"-s", "3"},
			expected: Options{s: "3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flagSet := flag.NewFlagSet("test", flag.ContinueOnError)

			var opts Options
			flagSet.BoolVar(&opts.c, "c", false, "")
			flagSet.BoolVar(&opts.d, "d", false, "")
			flagSet.BoolVar(&opts.u, "u", false, "")
			flagSet.BoolVar(&opts.i, "i", false, "")
			flagSet.IntVar(&opts.f, "f", 0, "")
			flagSet.StringVar(&opts.s, "s", "", "")

			err := flagSet.Parse(tt.args)
			if err != nil && !tt.wantErr {
				t.Errorf("parse error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			count := 0
			if opts.c {
				count++
			}
			if opts.d {
				count++
			}
			if opts.u {
				count++
			}
			if (count > 1) != tt.wantErr {
				t.Errorf("conflict check failed: gotErr=%v, wantErr=%v", count > 1, tt.wantErr)
			}

			if !tt.wantErr && !reflect.DeepEqual(opts, tt.expected) {
				t.Errorf("Options = %+v, want %+v", opts, tt.expected)
			}
		})
	}
}

func TestBuildKey(t *testing.T) {
	tests := []struct {
		line    string
		options Options
		want    string
	}{
		{"I love music", Options{f: 1, s: "", i: false}, "love_music"},
		{"Thanks", Options{f: 1, s: "1", i: false}, "hanks"},
		{"We love music of Kartik", Options{f: 3, s: "4", i: false}, "artik"},
	}

	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			got := buildKey(tt.line, tt.options)
			if got != tt.want {
				t.Errorf("buildKey(%q, %+v) = %q, want %q", tt.line, tt.options, got, tt.want)
			}
		})
	}
}
