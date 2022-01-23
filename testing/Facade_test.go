package testing

import (
	"fmt"
	"github.com/OntoLedgy/core_ontology/code/core"
	"github.com/OntoLedgy/core_ontology/code/core/factories"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
	"testing"
)

func TestFacade(t *testing.T) {

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
		test_name.UmlName,
		test_name.GetObjectUuid().UUID.String(),
		test_name.OwningRepositoryUuid,
		test_name.RegistryKeyedOnUuid[object_uuid].ObjectUuid.UUID.String())

	bnogFacade := &core.BnogFacades{}

	name_uuid, _ := utils.GetUUID(1, "")

	test_name2 := bnogFacade.CreateBnogName(
		name_uuid,
		"test",
		repository_uuid,
		"testpresentationName")

	fmt.Printf("BnogObjects name: %s \n"+
		"uuid:%s \n"+
		"owning repository uuid:%s \n"+
		"registry keyed on uuids: %s ",
		test_name2.UmlName,
		test_name2.GetObjectUuid().UUID.String(),
		test_name2.OwningRepositoryUuid,
		test_name2.RegistryKeyedOnUuid[object_uuid].ObjectUuid.UUID.String())

}
