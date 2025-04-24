package cmd

import "strings"

func BuildUse(args ...string) string {
	switch len(args) {
	case 0:
		return ""

	case 1:
		return strings.TrimSpace(args[0])

	default:
		var result = strings.TrimSpace(args[0])

		for i := 1; i < len(args); i++ {
			if s := strings.TrimSpace(args[i]); s != "" {
				result += " " + s
			}
		}

		return result
	}
}
