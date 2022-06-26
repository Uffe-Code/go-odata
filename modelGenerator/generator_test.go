package modelGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Generate_struct(t *testing.T) {
	edmx, _ := parseEdmx([]byte(edmxXmlString))

	peopleSet := edmx.EntitySets["People"]

	assert.Equal(t, `type Person struct {
	AddressInfo []Location
	Age nullable.Nullable[int64]
	Emails []string
	FavoriteFeature Feature
	Features []Feature
	FirstName string
	Gender PersonGender
	HomeAddress nullable.Nullable[Location]
	LastName nullable.Nullable[string]
	MiddleName nullable.Nullable[string]
	UserName string
}`, generateModelStruct(peopleSet))
}
