package objects

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type BnogTypePlaces struct {
	BnogObjects
	placingType   *BnogTypes
	placedType    *BnogTypes
	typePlaceCkid ckids.BoroObjectCkIds
}

func (bnogTyplePlaces *BnogTypePlaces) New(
	uuid *utils.UUIDs,
	placingType *BnogTypes,
	placedType *BnogTypes,
	typePlaceCkid ckids.BoroObjectCkIds,
	owningRepositoryUuid *utils.UUIDs) {

	bnogTyplePlaces.BnogObjects.New(
		uuid,
		owningRepositoryUuid,
		"")

	bnogTyplePlaces.placingType =
		placingType

	bnogTyplePlaces.placedType =
		placedType

	bnogTyplePlaces.typePlaceCkid =
		typePlaceCkid

}
