package modelGenerator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Generate_struct(t *testing.T) {
	edmx, _ := getParsedEdmx()
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
}`, generateModelStruct(peopleSet.getEntityType()))
}

func Test_Generate_definition(t *testing.T) {
	edmx, _ := getParsedEdmx()
	peopleSet := edmx.EntitySets["People"]

	assert.Equal(t, `//goland:noinspection GoUnusedExportedFunction
func NewPersonCollection(wrapper odataClient.Wrapper) odataClient.ODataModelCollection[Person] {
	return modelDefinition[Person]{client: wrapper.ODataClient(), name: "Person", url: "People"}
}`, generateModelDefinition(peopleSet))
}

func Test_Generate_enum(t *testing.T) {
	edmx, _ := getParsedEdmx()
	genderEnum := edmx.EnumTypes["PersonGender"]
	assert.Equal(t, `type PersonGender int64

const (
	Male PersonGender = 0
	Female PersonGender = 1
	Unknown PersonGender = 2
)`, generateEnumStruct(genderEnum))
}
