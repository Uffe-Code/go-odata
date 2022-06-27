package date

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Date_to_string(t *testing.T) {
	d, err := Parse("2006-01-02", "2021-10-11")
	assert.NoError(t, err)
	assert.Equal(t, "2021-10-11", d.String())
}

func Test_Json_marshal_date(t *testing.T) {
	d := New(2021, 10, 11)
	assert.Equal(t, "2021-10-11", d.String())
	jsonData, err := json.Marshal(d)
	assert.NoError(t, err)
	assert.Equal(t, `"2021-10-11"`, string(jsonData))

	type res struct {
		StartDate Date `json:"startDate"`
	}
	jsonData, err = json.Marshal(res{d})
	assert.NoError(t, err)
	assert.Equal(t, `{"startDate":"2021-10-11"}`, string(jsonData))
}

func Test_Json_unmarshal_date(t *testing.T) {
	type res struct {
		StartDate Date `json:"startDate"`
	}
	var data res
	err := json.Unmarshal([]byte(`{"startDate":"2021-10-11"}`), &data)
	assert.NoError(t, err)
	assert.Equal(t, "2021-10-11", data.StartDate.String())
}

func Test_Get_time(t *testing.T) {
	d, err := Parse("2006-01-02 03:04:05", "2021-05-15 10:45:00")
	assert.NoError(t, err)
	assert.Equal(t, "2021-05-15 00:00:00 +0000 UTC", d.Time().String())
}
