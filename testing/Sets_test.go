package testing

import (
	"fmt"

	"github.com/OntoLedgy/core_ontology/code/core/construction_operations"
	"testing"
)

func TestSetOperations(t *testing.T) {

	//set_index := new(string)

	test_set := make(construction_operations.Sets[string])

	test_set.Add("tesdt")

	fmt.Println("test set: ", test_set)

}
