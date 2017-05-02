package main

func hIndex(citations []int) int {
	len := len(citations)
	if len == 0 {
		return 0
	}

	h := len
	for ; h > 0; h-- {
		count := 0
		for i := len - 1; i >= 0; i-- {
			if citations[i] >= h {
				count++
			}

			if count == h {
				return h
			}
		}
	}

	return h
}
