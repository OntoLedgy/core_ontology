package factories

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
)

type BnogTypePlaceFactories struct {
}

func (BnogTypePlaceFactories) Create(
	uuid *utils.UUIDs,
	placing_type *objects.BnogTypes,
	placed_type *objects.BnogTypes,
	type_place_ckid ckids.BoroObjectCkIds,
	owning_repository_uuid *utils.UUIDs) *objects.BnogTypePlaces {

	bnog_typleplace := &objects.BnogTypePlaces{}

	bnog_typleplace.New(
		uuid,
		placing_type,
		placed_type,
		type_place_ckid,
		owning_repository_uuid)

	return bnog_typleplace

}
