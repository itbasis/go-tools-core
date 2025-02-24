package os

import (
	"testing"

	itbasisTestUtils "github.com/itbasis/go-test-utils/v5/ginkgo"
	"github.com/onsi/ginkgo/v2"
)

func TestOS(t *testing.T) {
	itbasisTestUtils.InitGinkgoSuite(t, "OS Suite")
}

func SkipIfMinGW() {
	if IsMinGW() {
		ginkgo.Skip("Test not for MinGW environment")
	}
}

func SkipIfNotMinGW() {
	if !IsMinGW() {
		ginkgo.Skip("Test only for MinGW environment")
	}
}
