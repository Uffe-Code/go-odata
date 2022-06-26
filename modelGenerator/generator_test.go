package modelGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Generate_struct(t *testing.T) {
	edmx, _ := parseEdmx([]byte(edmxXmlString))

	peopleSet := edmx.EntitySets["People"]

	assert.Equal(t, `type Person struct {
	AddressInfo []interface{}
	Age nullable.Nullable[int64]
	Emails []string
	FavoriteFeature interface{}
	Features []interface{}
	FirstName string
	Gender interface{}
	HomeAddress nullable.Nullable[interface{}]
	LastName nullable.Nullable[string]
	MiddleName nullable.Nullable[string]
	UserName string
}`, generateModelStruct(peopleSet))
}
