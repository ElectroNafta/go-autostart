// +build !darwin

package autostart

import (
	"strconv"
	"strings"
)

func quote(args []string) string {
	copyArgs := make([]string, len(args))

	for i, v := range args {
		copyArgs[i] = strconv.Quote(v)
	}

	return strings.Join(copyArgs, " ")
}
