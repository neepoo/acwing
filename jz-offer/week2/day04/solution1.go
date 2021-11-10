package day04

func qmi(a float64, b int) float64{
	var res float64 = 1
	for b > 0{
		if b & 1 == 1{
			res = res * a
		}
		b >>=1
		a *= a
	}

	return res
}

func Power(base float64, exponent int) float64 {
	if exponent > 0{
		return qmi(base,exponent)
	}else if exponent == 0{
		return 1
	}else {
		return 1 / qmi(base, -exponent)
	}
}