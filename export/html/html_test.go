package html

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/abaldeweg/mission_storage/mission/create"
)

func TestConfig(t *testing.T) {
	file := func() []byte {
		d, err := json.Marshal(create.Logfile{Missions: []create.Mission{}})
		if err != nil {
			t.Fatal(err)
		}

		return d
	}

	export, err := Export(file())

	if err != nil {
		t.Fatal(err)
	}

	if reflect.TypeOf(export).String() != "string" {
		t.Fatalf("%s is not type of %s", reflect.TypeOf(export), "string")
	}
}
