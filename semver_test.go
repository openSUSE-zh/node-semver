package semver

import "testing"

func TestNewSemver(t *testing.T) {
	cases := []string{"1.0.0-alpha.1+build20200204.1", "1.0.0-alpha", "1.0.0", "1.0", "1", "0.9.0-1.2.3", "1.0.26-1", "2.0.7-bindist-testing", "1.0.0-prerelease-1", "7.0.0-beta.0-ranges"}
	answers := []Semver{{"1", "0", "0", "alpha.1", "build20200204.1"}, {"1", "0", "0", "alpha", ""}, {"1", "0", "0", "", ""}, {"1", "0", "0", "", ""}, {"1", "0", "0", "", ""}, {"0", "9", "0", "1.2.3", ""}, {"1", "0", "26", "1", ""}, {"2", "0", "7", "bindist-testing", ""}, {"1", "0", "0", "prerelease-1", ""}, {"7", "0", "0", "beta.0-ranges", ""}}

	for i, j := range cases {
		if NewSemver(j).Equal(answers[i]) {
			t.Logf("NewSemver %s succeed", j)
		} else {
			t.Errorf("NewSemver %s failed, expected %v, got %v", j, answers[i], NewSemver(j))
		}
	}
}

func TestSemverToString(t *testing.T) {
	ver := "1.0.0-alpha.1+build20200204.1"
	ver1 := NewSemver(ver)
	if ver1.String() == ver {
		t.Logf("Semver to String succeed")
	} else {
		t.Errorf("Semver to string failed, expected %s, got %s", ver, ver1.String())
	}
}

func TestGt(t *testing.T) {
	cases := [][]Semver{{NewSemver("2.3.4"), NewSemver("1.2.3-beta.2")}, {NewSemver("2.3.4-alpha.5"), NewSemver("1.2.3-beta.4")}, {NewSemver("1.2.3-beta.4"), NewSemver("1.2.3-beta.3")}}
	for _, v := range cases {
		if v[0].gt(v[1], false) {
			t.Logf("Test %s.gt(%s) succeed", v[0], v[1])
		} else {
			t.Errorf("Test %s.gt(%s) failed, expected true, got false", v[0], v[1])
		}
	}
}
