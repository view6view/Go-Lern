package dp

import "testing"

func TestLongestPalindrome(t *testing.T) {
	if ans := longestPalindrome("aacabdkacaa"); ans != "aca" {
		t.Errorf("aacabdkacaa expected be aca, but %s got", ans)
	}

	if ans := longestPalindrome("babad"); ans != "bab" {
		t.Errorf("babad expected be bab, but %s got", ans)
	}

	if ans := longestPalindrome("cbbd"); ans != "bb" {
		t.Errorf("cbbd expected be bb, but %s got", ans)
	}
}
