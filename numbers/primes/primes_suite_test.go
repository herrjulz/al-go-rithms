package primes_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrimes(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Primes Suite")
}
