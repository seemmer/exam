package uint

import (
	"github.com/seemmer/exam/entity"
	"testing"
	"github.com/asaskevich/govalidator"
	."github.com/onsi/gomega"
)

func TestMember(t *testing.T){
	g:= NewGomegaWithT(t)

	t.Run("valid Member", func(t *testing.T){
		member := entity.Member{
			Phone: "0123456789",
			Password: "",
			Url: "https://example.com",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})

	t.Run("invalid phone", func(t *testing.T){
		member := entity.Member{
			Phone: "012345",
			Password: "",
			Url: "https://example.com",
		}

		ok, err := govalidator.ValidateStruct(member)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())
		g.Expect(err.Error()).To(Equal("Phone is not 10 digits"))
	})
}