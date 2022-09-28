package utils

func MergeMaps(l ...map[string]string) map[string]string {
	res := make(map[string]string)

	for _, v := range l {
		if nil != v {
			for lKey, lValue := range v {
				res[lKey] = lValue
			}
		}
	}
	return res
}
