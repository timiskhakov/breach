package breach

//func Optimize(seqs [][]byte, size int) [][]byte {
//	result := make([][]byte, 0)
//	longestSeqs := findLongestSeqs(seqs, size)
//	for _, longestSeq := range longestSeqs {
//
//	}
//
//	return result
//}
//
func findLongestSeqs(seqs [][]byte, size int) [][]byte {
	result := make([][]byte, 0)
	longest := 0
	for _, seq := range seqs {
		if len(seq) >= longest && len(seq) <= size {
			result = append(result, seq)
			longest = len(seq)
		}
	}

	return result
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
