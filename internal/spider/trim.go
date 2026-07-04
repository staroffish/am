package spider

import "strings"

func trimName(name string) string {
	return strings.Trim(name, " \n\u200b \t")
}
