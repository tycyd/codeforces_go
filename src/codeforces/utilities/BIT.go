package main

/** binary indexed tree with range update and query **/
/*
https://cp-algorithms.com/data_structures/fenwick.html
sum[0,i] =
i<l			0
l<=i<=r	 	x*(i-(l-1))
i>r			x*(r-l+1)

sum[0,i] = sum(B1,i)*i - sum(B2,i)
i<l			0*i - 0
l<=i<=r	 	x*i - x*(l-1)
i>r			0*i - (x*(l-1) - x*r)

def range_add(l, r, x):
    add(B1, l, x)
    add(B1, r+1, -x)
    add(B2, l, x*(l-1))
    add(B2, r+1, -x*r))
*/

type BIT_Range struct {
	B1   []int64
	B2   []int64
	size int
}

func (t *BIT_Range) build(sz int) {
	t.size = sz + 2
	t.B1 = make([]int64, t.size)
	t.B2 = make([]int64, t.size)
}

func (t *BIT_Range) add(b []int64, i int, x int64) {
	i++
	for i < t.size {
		b[i] += x
		i += i & -i
	}
}

func (t *BIT_Range) range_add(l int, r int, x int64) {
	t.add(t.B1, l, x)
	t.add(t.B1, r+1, -x)
	t.add(t.B2, l, x*int64(l-1))
	t.add(t.B2, r+1, -x*int64(r))
}

func (t *BIT_Range) sum(b []int64, i int) int64 {
	i++
	var res int64 = 0
	for i > 0 {
		res += b[i]
		i -= i & -i
	}
	return res
}

func (t *BIT_Range) prefix_sum(i int) int64 {
	return t.sum(t.B1, i)*int64(i) - t.sum(t.B2, i)
}

func (t *BIT_Range) range_sum(l int, r int) int64 {
	return t.prefix_sum(r) - t.prefix_sum(l-1)
}
