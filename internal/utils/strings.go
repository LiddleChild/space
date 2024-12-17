package utils

func AlignString(list []string, whitespace int) []int {
	mxLen := 0
	for _, s := range list {
		mxLen = max(mxLen, len(s))
	}

	lens := make([]int, len(list))

	for i := range len(lens) {
		lens[i] = mxLen - len(list[i]) + whitespace
	}

	return lens
}
