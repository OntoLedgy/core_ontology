package factories

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects/places"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
)

type BnogPlaceableTypeFactories struct {
}

func (BnogPlaceableTypeFactories) Create(
	uuid *utils.UUIDs,
	placeable_type_placed_types_map *places.BnogPlacedObjectsMaps,
	owning_repository_uuid *utils.UUIDs,
	presentation_name string) *objects.BnogPlaceableTypes {

	bnog_placeable_type :=
		&objects.BnogPlaceableTypes{}

	bnog_placeable_type.New(
		uuid,
		placeable_type_placed_types_map,
		owning_repository_uuid,
		presentation_name)

	return bnog_placeable_type

}
