package identifier_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestIdentifier(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Identifier Suite")
}
