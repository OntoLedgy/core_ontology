package objects

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects/places"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type BnogPlaceableTypes struct {
	BnogTypes
	PlaceableTypePlacedTypesMap *places.BnogPlacedObjectsMaps
}

func (
	bnogPlaceableType *BnogPlaceableTypes) New(
	uuid *utils.UUIDs,
	placeableTypePlacedTypesMap *places.BnogPlacedObjectsMaps,
	owningRepositoryUuid *utils.UUIDs,
	presentationName string) {

	bnogPlaceableType.BnogTypes.New(
		uuid,
		owningRepositoryUuid,
		presentationName)

	bnogPlaceableType.RegistryKeyedOnCkidType =
		make(
			map[ckids.BoroObjectCkIds]*BnogObjects)

	bnogPlaceableType.RegistryKeyedOnCkidType[ckids.PlaceableTypes] =
		&bnogPlaceableType.BnogObjects

	bnogPlaceableType.PlaceableTypePlacedTypesMap =
		&places.BnogPlacedObjectsMaps{}

	bnogPlaceableType.PlaceableTypePlacedTypesMap =
		placeableTypePlacedTypesMap

}
