package places

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
)

type BnogPlacedTypesMaps struct {
	RelationPlacedTypesMap map[ckids.PlaceNumberTypeCkIds]interface{} //objects.BnogTypes
}

func (bnogPlacedTypesMaps *BnogPlacedTypesMaps) add_tuple_placed_type_to_dictionary(
	placing_ckid ckids.PlaceNumberTypeCkIds,
	relation_placed_type interface{}) { //objects.BnogTypes

	bnogPlacedTypesMaps.RelationPlacedTypesMap[placing_ckid] =
		relation_placed_type

}

func (bnogPlacedTypesMaps *BnogPlacedTypesMaps) try_get_relation_placed_type_for_place_number_ckid(
	placing_ckid ckids.PlaceNumberTypeCkIds) interface{} { //objects.BnogTypes

	relation_placed_type := bnogPlacedTypesMaps.RelationPlacedTypesMap[placing_ckid]

	return relation_placed_type
}
