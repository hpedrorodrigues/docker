package util

import "strings"

func ImageAndTag(image string) (string, string) {
	var name, tag string

	if strings.Contains(image, ":") {
		imageSlice := strings.Split(image, ":")

		name, tag = imageSlice[0], imageSlice[1]
	} else {
		name, tag = image, "latest"
	}

	return name, tag
}
