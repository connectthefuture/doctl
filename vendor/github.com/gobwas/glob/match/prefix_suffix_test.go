package match

import (
	"reflect"
	"testing"
)

func TestPrefixSuffixIndex(t *testing.T) {
	for id, test := range []struct {
		prefix   string
		suffix   string
		fixture  string
		index    int
		segments []int
	}{
		{
			"a",
			"c",
			"abc",
			0,
			[]int{3},
		},
		{
			"f",
			"f",
			"fffabfff",
			0,
			[]int{1, 2, 3, 6, 7, 8},
		},
		{
			"ab",
			"bc",
			"abc",
			0,
			[]int{3},
		},
	} {
		p := PrefixSuffix{test.prefix, test.suffix}
		index, segments := p.Index(test.fixture)
		if index != test.index {
			t.Errorf("#%d unexpected index: exp: %d, act: %d", id, test.index, index)
		}
		if !reflect.DeepEqual(segments, test.segments) {
			t.Errorf("#%d unexpected segments: exp: %v, act: %v", id, test.segments, segments)
		}
	}
}

func BenchmarkIndexPrefixSuffix(b *testing.B) {
	m := PrefixSuffix{"qew", "sqw"}
	for i := 0; i < b.N; i++ {
		m.Index(bench_pattern)
	}
}
