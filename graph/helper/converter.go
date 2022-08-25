package helper

func ConvertAsToPas(items []string) []*string {
	result := make([]*string, 0)
	for i := range items {
		result = append(result, &items[i])
	}
	return result
}
