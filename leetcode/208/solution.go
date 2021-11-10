package leetcode

const N = 2010 * 26

type Trie struct {
	g   [N][26]int
	cnt [N]int
	idx int
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	p := 0
	for _, ch := range word {
		if t.g[p][ch-'a'] == 0 {
			t.idx++
			t.g[p][ch-'a'] = t.idx
		}
		p = t.g[p][ch-'a']
	}
	t.cnt[p]++
}

func (t *Trie) Search(word string) bool {
	p := 0
	for _, ch := range word {
		if t.g[p][ch-'a'] == 0 {
			return false
		}
		p = t.g[p][ch-'a']
	}
	return t.cnt[p] > 0
}

func (t *Trie) StartsWith(prefix string) bool {
	p := 0
	for _, ch := range prefix {
		if t.g[p][ch-'a'] == 0 {
			return false
		}
		p = t.g[p][ch-'a']
	}
	return true
}
