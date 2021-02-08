package main

func myStrCapitalize(src string) (string) {
	var dest string

	for _, v := range src {
		if v >= 97 && v <= 122 {
			dest = dest + string(v - 32)
			v = v - 33
		} else {
			dest = dest + string(v)
		}
	}
	return dest
}