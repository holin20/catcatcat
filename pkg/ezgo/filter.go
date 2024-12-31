package ezgo

func FilterEmpty(slice []string) []string {
	var ret []string
	for _, s := range slice {
		if s != "" {
			ret = append(ret, s)
		}
	}
	return ret
}
