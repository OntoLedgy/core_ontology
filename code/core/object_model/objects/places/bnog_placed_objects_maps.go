package places

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
)

type BnogPlacedObjectsMaps struct {
	RelationPlacedObjectMap map[ckids.PlaceNumberTypeCkIds]interface{} //*objects.BnogObjects
}

func (bnogPlacedObjectsMaps *BnogPlacedObjectsMaps) Add_tuple_placed_object_to_dictionary(
	placing_ckid ckids.PlaceNumberTypeCkIds,
	relation_placed_object interface{}) { //*objects.BnogObjects

	bnogPlacedObjectsMaps.RelationPlacedObjectMap[placing_ckid] =
		relation_placed_object

}

func (bnogPlacedObjectsMaps *BnogPlacedObjectsMaps) try_get_relation_placed_object_for_place_number_ckid(
	placing_ckid ckids.PlaceNumberTypeCkIds) interface{} { //*objects.BnogObjects

	relation_placed_object := bnogPlacedObjectsMaps.RelationPlacedObjectMap[placing_ckid]

	return relation_placed_object
}
