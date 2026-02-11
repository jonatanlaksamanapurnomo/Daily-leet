type SegTree struct {
	minVal []int
	maxVal []int
	lazy   []int
	n      int // range [0, n]
}

func NewSegTree(n int) *SegTree {
	sz := 4 * (n + 1)
	st := &SegTree{
		minVal: make([]int, sz),
		maxVal: make([]int, sz),
		lazy:   make([]int, sz),
		n:      n,
	}
	// Semua value sudah 0 dari make(), tidak perlu build eksplisit
	// Tapi kita tetap perlu set leaf identity agar query benar
	st.build(1, 0, n)
	return st
}

func (st *SegTree) build(nd, lo, hi int) {
	st.lazy[nd] = 0
	st.minVal[nd] = 0
	st.maxVal[nd] = 0
	if lo == hi {
		return
	}
	mid := (lo + hi) / 2
	st.build(nd*2, lo, mid)
	st.build(nd*2+1, mid+1, hi)
}

func (st *SegTree) pull(nd int) {
	l, r := nd*2, nd*2+1
	st.minVal[nd] = min(st.minVal[l], st.minVal[r])
	st.maxVal[nd] = max(st.maxVal[l], st.maxVal[r])
}

func (st *SegTree) push(nd int) {
	if st.lazy[nd] != 0 {
		for _, c := range []int{nd * 2, nd*2 + 1} {
			st.minVal[c] += st.lazy[nd]
			st.maxVal[c] += st.lazy[nd]
			st.lazy[c] += st.lazy[nd]
		}
		st.lazy[nd] = 0
	}
}

// RangeAdd: tambahkan val ke semua balance[i] di range [ql, qr]
func (st *SegTree) RangeAdd(nd, lo, hi, ql, qr, val int) {
	if ql > qr || ql > hi || qr < lo {
		return
	}
	if ql <= lo && hi <= qr {
		st.minVal[nd] += val
		st.maxVal[nd] += val
		st.lazy[nd] += val
		return
	}
	st.push(nd)
	mid := (lo + hi) / 2
	st.RangeAdd(nd*2, lo, mid, ql, qr, val)
	st.RangeAdd(nd*2+1, mid+1, hi, ql, qr, val)
	st.pull(nd)
}

// QueryLeftmostZero: cari index terkecil di [ql, qr] dimana value == 0
// Return -1 kalau tidak ada
func (st *SegTree) QueryLeftmostZero(nd, lo, hi, ql, qr int) int {
	if ql > qr || ql > hi || qr < lo {
		return -1
	}
	// PRUNE: kalau semua value di range ini positif atau negatif, tidak ada 0
	if st.minVal[nd] > 0 || st.maxVal[nd] < 0 {
		return -1
	}
	if lo == hi {
		if lo >= ql && lo <= qr && st.minVal[nd] == 0 {
			return lo
		}
		return -1
	}
	st.push(nd)
	mid := (lo + hi) / 2
	// Cek kiri dulu (mau index terkecil)
	res := st.QueryLeftmostZero(nd*2, lo, mid, ql, qr)
	if res != -1 {
		return res
	}
	return st.QueryLeftmostZero(nd*2+1, mid+1, hi, ql, qr)
}

// ============================================================
// SOLUSI UTAMA
// ============================================================
//
// KONSEP:
// Kita maintain balance[l] = distinctEven(l..r) - distinctOdd(l..r)
// untuk setiap possible starting index l, dengan r = current position.
//
// Saat kita proses nums[r]:
//   - Cari prev = last index dimana nums[r] muncul (-1 kalau belum)
//   - Untuk l di [prev+1, r]: nums[r] jadi distinct baru di subarray [l..r]
//   - Kalau even: balance[l] += 1 untuk l ∈ [prev+1, r]
//   - Kalau odd:  balance[l] -= 1 untuk l ∈ [prev+1, r]
//
// Ini adalah RANGE UPDATE → pakai segment tree!
// Lalu QUERY: cari leftmost l di [0, r] dimana balance[l] == 0
//
// Complexity: O(n log n)
// ============================================================

func longestBalanced(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	st := NewSegTree(n) // range [0, n]

	lastSeen := make(map[int]int) // val -> last index
	ans := 0

	for r := 0; r < n; r++ {
		val := nums[r]

		prev, exists := lastSeen[val]
		if !exists {
			prev = -1
		}

		// Range update: untuk l ∈ [prev+1, r], nums[r] menjadi distinct baru
		updateLeft := prev + 1
		updateRight := r

		if val%2 == 0 {
			st.RangeAdd(1, 0, n, updateLeft, updateRight, 1)
		} else {
			st.RangeAdd(1, 0, n, updateLeft, updateRight, -1)
		}

		lastSeen[val] = r

		// Cari leftmost l di [0, r] dimana balance[l] == 0
		l := st.QueryLeftmostZero(1, 0, n, 0, r)
		if l != -1 {
			length := r - l + 1
			if length > ans {
				ans = length
			}
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}