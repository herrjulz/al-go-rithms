package basesystems_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestBasesystems(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Basesystems Suite")
}
