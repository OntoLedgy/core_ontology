package objects

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects/places"
	"github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
)

type BnogTuples struct {
	BnogObjects
	ImmutableMinorCompositionCoupleTypeBoroObjectCkid ckids.BoroObjectCkIds
	TuplePlacedObjectsMap                             *places.BnogPlacedObjectsMaps
}

func (bnogTuple *BnogTuples) New(
	uuid *identity_management_services.UUIDs,
	tuplePlacedObjectsDictionary *places.BnogPlacedObjectsMaps,
	immutableMinorCompositionCoupleTypeBoroObjectCkid ckids.BoroObjectCkIds,
	owningRepositoryUuid *identity_management_services.UUIDs) {

	bnogTuple.BnogObjects.New(
		uuid,
		owningRepositoryUuid,
		"")

	bnogTuple.ImmutableMinorCompositionCoupleTypeBoroObjectCkid =
		immutableMinorCompositionCoupleTypeBoroObjectCkid

	bnogTuple.TuplePlacedObjectsMap =
		tuplePlacedObjectsDictionary

}
