package testing

import (
	"fmt"
	sets "github.com/OntoLedgy/storage_interop_services/code/services/in_memory/sets"
	"testing"
)

func TestSetOperations(t *testing.T) {

	test_set := make(sets.Sets[string])

	test_set.Add("tesdt")

	fmt.Println("test set: ", test_set)

}
