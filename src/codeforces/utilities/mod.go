package main

/** modular multiplicative inverse **/
/*
https://en.wikipedia.org/wiki/Fermat%27s_little_theorem
*/
func mod_inverse(x int64, mod int64) int64 {
	return mod_power(x, mod-2, mod)
}

func mod_power(x int64, y int64, mod int64) int64 {
	if y == 0 {
		return 1
	}
	p := mod_power(x, y/2, mod) % mod
	p = (p * p) % mod
	if y%2 == 1 {
		p = (p * x) % mod
	}
	return p
}
