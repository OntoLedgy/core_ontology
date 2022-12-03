package objects

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
	"github.com/OntoLedgy/storage_interop_services/code/services/in_memory/sets"
)

type BnogObjects struct {
	ObjectUuid *uuid_service.UUIDs

	RegistryKeyedOnUuid map[*uuid_service.UUIDs]*BnogObjects

	RegistryKeyedOnCkidType map[ckids.BoroObjectCkIds]*BnogObjects

	MatchedObjects []*uuid_service.UUIDs

	OwningRepositoryUuid *uuid_service.UUIDs

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
	uuid *uuid_service.UUIDs,
	owningRepositoryUuid *uuid_service.UUIDs,
	presentationName string) {

	if uuid == nil {
		bnogObject.SetObjectUuid()
	} else {
		bnogObject.ObjectUuid = uuid
	}

	bnogObject.OwningRepositoryUuid = owningRepositoryUuid

	bnogObject.RegistryKeyedOnUuid = make(map[*uuid_service.UUIDs]*BnogObjects)

	bnogObject.RegistryKeyedOnUuid[uuid] = bnogObject

	bnogObject.RegistryKeyedOnCkidType = make(map[ckids.BoroObjectCkIds]*BnogObjects)

	bnogObject.RegistryKeyedOnCkidType[ckids.Objects] = bnogObject

	bnogObject.UmlName = presentationName

}

func (
	bnogObject *BnogObjects) SetObjectUuid() *uuid_service.UUIDs {

	bnogObject.ObjectUuid, _ = uuid_service.GetUUID(1, "")

	return bnogObject.ObjectUuid
}

func (
	bnogObject *BnogObjects) GetObjectUuid() *uuid_service.UUIDs {

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
