package numbers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestGcd(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Gcd Suite")
}
