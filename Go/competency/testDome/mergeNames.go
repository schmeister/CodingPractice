package testDome

func UniqueNames(a, b []string) []string {

	m := make(map[string]string)

	for _, n := range a {
		m[n] = n
	}
	for _, n := range b {
		m[n] = n
	}

	var result []string
	for _, n := range m {
		result = append(result, n)
	}

	return result
}
