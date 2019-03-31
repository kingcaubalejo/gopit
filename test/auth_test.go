package test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAuthUser(t *testing.T) {
	Describe("User Authentication", func(){
		Context("initially", func(){
			It("has no username", func() {})
			It("has no password", func() {})
			Specify("Username and Password initialized to nil", func() {})
		})
	})
}