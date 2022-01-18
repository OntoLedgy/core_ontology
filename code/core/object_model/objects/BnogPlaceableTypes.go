package objects

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type BnogPlaceableTypes struct {
	BnogTypes
	placeable_type_placed_types_map map[*BnogPlaceableTypes]interface{}
}

func (placeableType *BnogPlaceableTypes) New(
	uuid *utils.UUIDs,
	placeable_type_placed_types_map map[*BnogPlaceableTypes]interface{},
	owning_repository_uuid *utils.UUIDs,
	presentation_name string) {

	if uuid == nil {
		placeableType.Set_object_uuid()
	} else {
		placeableType.Object_uuid = uuid
	}

	placeableType.Owning_repository_uuid = owning_repository_uuid

	placeableType.Registry_keyed_on_ckid_type = make(map[BoroObjectCkIds]interface{})

	placeableType.Registry_keyed_on_ckid_type[ckids.PlaceableTypes] = placeableType

	placeableType.placeable_type_placed_types_map = make(map[*BnogPlaceableTypes]interface{})

	placeableType.placeable_type_placed_types_map = placeable_type_placed_types_map

	placeableType.Uml_name = presentation_name

}
