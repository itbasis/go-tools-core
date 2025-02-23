package os

import "os"

func IsMinGW() bool {
	_, ok := os.LookupEnv("MSYSTEM")

	return ok
}
