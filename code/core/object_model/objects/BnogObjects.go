package objects

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/sets"
)

type BnogObjects struct {
	ObjectUuid *identity_management_services.UUIDs

	RegistryKeyedOnUuid map[*identity_management_services.UUIDs]*BnogObjects

	RegistryKeyedOnCkidType map[ckids.BoroObjectCkIds]*BnogObjects

	MatchedObjects []*identity_management_services.UUIDs

	OwningRepositoryUuid *identity_management_services.UUIDs

	IsNamedBys *sets.Sets[*BnogObjects]

	Names sets.Sets[*BnogObjects]

	Types sets.Sets[*BnogObjects]

	Instances sets.Sets[*BnogObjects]

	Supertypes sets.Sets[*BnogObjects]

	UmlName string

	NamingSpaces sets.Sets[*BnogObjects]
}

func (
	bnogObject *BnogObjects) New(
	uuid *identity_management_services.UUIDs,
	owningRepositoryUuid *identity_management_services.UUIDs,
	presentationName string) {

	if uuid == nil {
		bnogObject.SetObjectUuid()
	} else {
		bnogObject.ObjectUuid = uuid
	}

	bnogObject.OwningRepositoryUuid = owningRepositoryUuid

	bnogObject.RegistryKeyedOnUuid = make(map[*identity_management_services.UUIDs]*BnogObjects)

	bnogObject.RegistryKeyedOnUuid[uuid] = bnogObject

	bnogObject.RegistryKeyedOnCkidType = make(map[ckids.BoroObjectCkIds]*BnogObjects)

	bnogObject.RegistryKeyedOnCkidType[ckids.Objects] = bnogObject

	bnogObject.UmlName = presentationName

}

func (
	bnogObject *BnogObjects) SetObjectUuid() *identity_management_services.UUIDs {

	bnogObject.ObjectUuid, _ = identity_management_services.GetUUID(1, "")

	return bnogObject.ObjectUuid
}

func (
	bnogObject *BnogObjects) GetObjectUuid() *identity_management_services.UUIDs {

	return bnogObject.ObjectUuid

}

func (
	bnogObject *BnogObjects) addToRegistryKeyedOnCkidType(
	boroObjectCkId ckids.BoroObjectCkIds) {

	var ckid_typed_objects *BnogObjects

	if val, ok := bnogObject.RegistryKeyedOnCkidType[boroObjectCkId]; ok {

		ckid_typed_objects = val

	} else {

		ckid_typed_objects = &BnogObjects{}

		bnogObject.RegistryKeyedOnCkidType[boroObjectCkId] =
			ckid_typed_objects

	}

	//ckid_typed_objects

}
