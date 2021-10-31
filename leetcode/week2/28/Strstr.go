package main

const PrimeRK = 16777619

func HashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*PrimeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, PrimeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

func strStr(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	if len(s) < len(substr) {
		return -1
	}
	hashss, pow := HashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*PrimeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= PrimeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}

func strStr1(s, pa string) int {
	lp := len(pa)
	ls := len(s)
	var P uint64 = 131
	var pHash uint64
	if lp == 0 {
		return 0
	}
	// 保证模板串长度小于等于s
	if ls < lp {
		return -1
	}
	p := make([]uint64, ls+10)
	h := make([]uint64, ls+10)
	p[0] = 1
	for idx, ch := range s {
		p[idx+1] = p[idx] * P
		h[idx+1] = h[idx]*P + uint64(ch-'a'+1) // a:1,b:2,...
	}
	for _, ch := range pa{
		pHash = pHash*P + uint64(ch-'a'+1)
	}
	for i := 0; i +lp <= ls ; i++ {
		currentHash := h[i+lp] - h[i]*p[lp]
		if currentHash == pHash{
			return i
		}
	}
	return -1
}
