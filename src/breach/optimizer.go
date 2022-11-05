package breach

func Optimize(firstRow []byte, seqs [][]byte, size int) []byte {
	longestSeqs := findLongestSeqs(firstRow, seqs, size)
	result := longestSeqs[0]
	for _, longestSeq := range longestSeqs {
		for _, seq := range seqs {
			left := merge(longestSeq, seq)
			if len(left) > len(result) && len(left) <= size {
				result = left
			}
			right := merge(seq, longestSeq)
			if len(right) > len(result) && len(right) <= size {
				result = right
			}
		}
	}

	return result
}

func findLongestSeqs(firstRow []byte, seqs [][]byte, size int) [][]byte {
	result := make([][]byte, 0)
	longest := 0
	for _, seq := range seqs {
		seqSize := size
		if !contains(firstRow, seq[0]) {
			seqSize--
		}

		if len(seq) >= longest && len(seq) <= seqSize {
			result = append(result, seq)
			longest = len(seq)
		}
	}

	return result
}

func contains(arr []byte, elem byte) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == elem {
			return true
		}
	}
	return false
}

func merge(seq1, seq2 []byte) []byte {
	index := 0
outer:
	for i := 0; i < len(seq1); i++ {
		j := 0
		for ; j < len(seq2); j++ {
			if i+j == len(seq1) {
				index = j
				break outer
			}
			if seq1[i+j] != seq2[j] {
				break
			}
		}
	}

	return append(seq1, seq2[index:]...)
}
