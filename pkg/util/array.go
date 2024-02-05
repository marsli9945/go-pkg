package util

func SplitArray(arr []string, chunkSize int) [][]string {
	var result [][]string

	for i := 0; i < len(arr); i += chunkSize {
		end := i + chunkSize

		if end > len(arr) {
			end = len(arr)
		}

		result = append(result, arr[i:end])
	}

	return result
}
