package testing

import (
	"fmt"
	"github.com/OntoLedgy/core_ontology/code/core/factories"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
	"testing"
)

func TestOnticCategoriesObjects(t *testing.T) {

	test_object_factory := &factories.BnogObjectFactories{}

	object_uuid, _ := utils.GetUUID(1, "")
	repository_uuid, _ := utils.GetUUID(1, "")

	test_object := test_object_factory.Create(
		object_uuid,
		repository_uuid,
		"test")

	fmt.Printf("BnogObjects name: %s \n"+
		"uuid:%s \n"+
		"owning repository uuid:%s \n"+
		"registry keyed on uuids: %s ",
		test_object.Uml_name,
		test_object.Get_object_uuid().UUID.String(),
		test_object.Owning_repository_uuid,
		test_object.Registry_keyed_on_uuid)

}

func TestOnticCategoriesNames(t *testing.T) {

	test_name_factory := &factories.BnogNameFactories{}

	object_uuid, _ := utils.GetUUID(1, "")
	repository_uuid, _ := utils.GetUUID(1, "")

	test_name := test_name_factory.Create(
		object_uuid,
		"name",
		repository_uuid,
		"test_name")

	fmt.Printf("BnogObjects name: %s \n"+
		"uuid:%s \n"+
		"owning repository uuid:%s \n"+
		"registry keyed on uuids: %s ",
		test_name.Uml_name,
		test_name.Get_object_uuid().UUID.String(),
		test_name.Owning_repository_uuid,
		test_name.Registry_keyed_on_uuid)

}
