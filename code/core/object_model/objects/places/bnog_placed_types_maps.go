package places

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
)

type BnogPlacedTypesMaps struct {
	RelationPlacedTypesMap map[ckids.PlaceNumberTypeCkIds]objects.BnogTypes
}

func (bnogPlacedTypesMaps *BnogPlacedTypesMaps) add_tuple_placed_type_to_dictionary(
	placing_ckid ckids.PlaceNumberTypeCkIds,
	relation_placed_type objects.BnogTypes) {

	bnogPlacedTypesMaps.RelationPlacedTypesMap[placing_ckid] =
		relation_placed_type

}

func (bnogPlacedTypesMaps *BnogPlacedTypesMaps) try_get_relation_placed_type_for_place_number_ckid(
	placing_ckid ckids.PlaceNumberTypeCkIds) objects.BnogTypes {

	relation_placed_type := bnogPlacedTypesMaps.RelationPlacedTypesMap[placing_ckid]

	return relation_placed_type
}
