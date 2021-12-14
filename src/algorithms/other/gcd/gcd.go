package gcd

// GCD - Greatest common divisor algorithm
func GCD(q int, p int) int {
	if p == 0 {
		return q
	}
	return GCD(p, q%p)
}
