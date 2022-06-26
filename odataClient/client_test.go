package odataClient

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestODataClient_AddHeader(t *testing.T) {
	client := New("http://test.api/")
	client.AddHeader("X-Foo", "Bar")
	assert.Equal(t, "Bar", client.(*oDataClient).headers["x-foo"])
}
