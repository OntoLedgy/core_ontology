package testing

import (
	"fmt"
	"github.com/OntoLedgy/core_ontology/code/bnog_io"
	"github.com/OntoLedgy/core_ontology/code/core/factories"
	uuids "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
	"testing"
)

func TestOnticCategories_Objects(t *testing.T) {

	testObjectFactory := &factories.BnogObjectFactories{}

	objectUuid, _ := uuids.GetUUID(1, "")
	repositoryUuid, _ := uuids.GetUUID(1, "")

	testObject := testObjectFactory.Create(
		objectUuid,
		repositoryUuid,
		"test")

	fmt.Printf("BnogObjects name: %s \n"+
		"uuid:%s \n"+
		"owning repository uuid:%s \n"+
		"registry keyed on uuids: %s ",
		testObject.UmlName,
		testObject.GetObjectUuid().UUID.String(),
		testObject.OwningRepositoryUuid,
		testObject.RegistryKeyedOnUuid[objectUuid].ObjectUuid.String())

	bnog_io.WriteBnog(
		"./testing.gob",
		testObject)

	bnog_io.ReadBnog(
		"./testing.gob")
}

func TestOnticCategories_Names(t *testing.T) {

	testNameFactory := &factories.BnogNameFactories{}

	objectUuid, _ := uuids.GetUUID(1, "")
	repositoryUuid, _ := uuids.GetUUID(1, "")

	testName := testNameFactory.Create(
		objectUuid,
		"name",
		repositoryUuid,
		"test_name")

	fmt.Printf("BnogObjects name: %s \n"+
		"uuid:%s \n"+
		"owning repository uuid:%s \n"+
		"registry keyed on uuids: %s ",
		testName.UmlName,
		testName.GetObjectUuid().UUID.String(),
		testName.OwningRepositoryUuid,
		testName.RegistryKeyedOnUuid[objectUuid].ObjectUuid.UUID.String())

}

func TestOnticCategories_PlaceableTypes(t *testing.T) {

	testPlaceableTypeFactory := &factories.BnogPlaceableTypeFactories{}

	objectUuid, _ := uuids.GetUUID(1, "")
	repositoryUuid, _ := uuids.GetUUID(1, "")

	//TODO - Create two types, and add to placeable type dictionary.

	testPlaceableType := testPlaceableTypeFactory.Create(
		objectUuid,
		nil,
		repositoryUuid,
		"test_placeable_type")

	fmt.Printf("BnogObjects name: %s \n"+
		"uuid:%s \n"+
		"owning repository uuid:%s \n"+
		"registry keyed on uuids: %s ",
		testPlaceableType.UmlName,
		testPlaceableType.GetObjectUuid().UUID.String(),
		testPlaceableType.OwningRepositoryUuid,
		testPlaceableType.RegistryKeyedOnUuid[objectUuid].ObjectUuid.UUID)

}
