package testing

import (
	"fmt"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"testing"
)

func TestOnticCategories(t *testing.T) {

	test_object := objects.BnogObjects{}
	test_object.Set_object_uuid()

	types := objects.BnogTypes{}

	types.Set_object_uuid()

	elements := objects.BnogElements{}
	elements.Set_object_uuid()

	placeableTypes := objects.BnogPlaceableTypes{}
	placeableTypes.Set_object_uuid()

	fmt.Println("BnogObjects:" + test_object.Get_object_uuid().UUID.String())
	fmt.Println("BnogTypes:" + types.Get_object_uuid().UUID.String())
	fmt.Println("BnogElements:" + elements.Get_object_uuid().UUID.String())
	fmt.Println("placeable types:" + test_object.Get_object_uuid().UUID.String())

}
