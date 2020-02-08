package semver

import (
	"testing"
)

func TestComparatorEqual(t *testing.T) {
	c := NewComparator(">=", "1.2.3")
	c1 := c
	if c.Equal(c1) {
		t.Logf("Test Comparator Equal succeed")
	} else {
		t.Errorf("Test Comparator Equal failed, expected true, got false")
	}
}

func TestComplete(t *testing.T) {
	cases := []string{"1", "1.2", "1.2.3"}
	answers := []int{1, 2, 3}
	for i, v := range cases {
		if complete(v) == answers[i] {
			t.Logf("Test complete() succeed")
		} else {
			t.Errorf("Test complete(%s) failed, expected %d, got %d", v, answers[i], complete(v))
		}
	}
}

func TestParseHyphen(t *testing.T) {
	cases := []string{"1.2 - 2.0", "1 - 2", "1.2.3 - 2.0.0"}
	answers := [][]Comparator{{NewComparator(">=", "1.2.0"), NewComparator("<", "2.1.0")}, {NewComparator(">=", "1.0.0"), NewComparator("<", "3.0.0")}, {NewComparator(">=", "1.2.3"), NewComparator("<=", "2.0.0")}}
	for i, v := range cases {
		m, n := parseHyphen(v)
		if m.Equal(answers[i][0]) && n.Equal(answers[i][1]) {
			t.Logf("Test parseHyphen(%s) succeed", v)
		} else {
			t.Errorf("Test parseHyphen(%s) failed, expected %v %v, got %v %v", v, answers[i][0], answers[i][1], m, n)
		}
	}
}

func TestParseX(t *testing.T) {
	cases := []string{"*", "1", "1.2", "1.x", "1.1.x"}
	answers := [][]Comparator{{NewComparator(">=", "0"), {}}, {NewComparator(">=", "1.0.0"), NewComparator("<", "2.0.0")}, {NewComparator(">=", "1.2.0"), NewComparator("<", "1.3.0")}, {NewComparator(">=", "1.0.0"), NewComparator("<", "2.0.0")}, {NewComparator(">=", "1.1.0"), NewComparator("<", "1.2.0")}}
	for i, v := range cases {
		m, n := parseX(v)
		if m.Equal(answers[i][0]) && n.Equal(answers[i][1]) {
			t.Logf("Test parseX(%s) succeed", v)
		} else {
			t.Errorf("Test parseX(%s) failed, expected %v %v, got %v %v", v, answers[i][0], answers[i][1], m, n)
		}
	}
}

func TestParseTilde(t *testing.T) {
	cases := []string{"1", "1.2", "1.2.3-beta.2"}
	answers := [][]Comparator{{NewComparator(">=", "1.0.0"), NewComparator("<", "2.0.0")}, {NewComparator(">=", "1.2.0"), NewComparator("<", "1.3.0")}, {NewComparator(">=", "1.2.3-beta.2"), NewComparator("<", "1.3.0")}}
	for i, v := range cases {
		m, n := parseTilde(v)
		if m.Equal(answers[i][0]) && n.Equal(answers[i][1]) {
			t.Logf("Test parseTilde(%s) succeed", v)
		} else {
			t.Errorf("Test parseTilde(%s) failed, expected %v %v, got %v %v", v, answers[i][0], answers[i][1], m, n)
		}
	}
}

func TestParseCaret(t *testing.T) {
	cases := []string{"1.2.3", "0.2.3", "0.0.3-beta.2", "1.x", "0.0.x", "0.x"}
	answers := [][]Comparator{{NewComparator(">=", "1.2.3"), NewComparator("<", "2.0.0")}, {NewComparator(">=", "0.2.3"), NewComparator("<", "0.3.0")}, {NewComparator(">=", "0.0.3-beta.2"), NewComparator("<", "0.0.4")}, {NewComparator(">=", "1.0.0"), NewComparator("<", "2.0.0")}, {NewComparator(">=", "0.0.0"), NewComparator("<", "0.1.0")}, {NewComparator(">=", "0.0.0"), NewComparator("<", "1.0.0")}}
	for i, v := range cases {
		m, n := parseCaret(v)
		if m.Equal(answers[i][0]) && n.Equal(answers[i][1]) {
			t.Logf("Test parseCaret(%s) succeed", v)
		} else {
			t.Errorf("Test parseCaret(%s) failed, expected %v %v, got %v %v", v, answers[i][0], answers[i][1], m, n)
		}
	}
}

func TestComparatorSatisfy(t *testing.T) {
	c := []Comparator{NewComparator(">", "1.2.3-beta.2"), NewComparator("<=", "1.2.3"), NewComparator("=", "1.2.3-beta.2")}
	c1 := []Semver{NewSemver("1.2.4-beta.2"), NewSemver("0.0.1"), NewSemver("1.2.3-beta.2")}
	answers := []bool{false, true, true}
	for i, v := range c {
		if v.Satisfy(c1[i]) == answers[i] {
			t.Logf("Test Comparator %s.Satisfy(%s) succeed", v.String(), c1[i].String())
		} else {
			t.Errorf("Test Comparator %s.Satisfy(%s) failed, expected %t, got %t", v.String(), c1[i].String(), answers[i], v.Satisfy(c1[i]))
		}
	}

}
