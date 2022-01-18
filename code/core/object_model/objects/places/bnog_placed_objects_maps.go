package places

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
)

type BnogPlacedObjectsMaps struct {
	RelationPlacedObjectMap map[ckids.PlaceNumberTypeCkIds]objects.BnogObjects
}

func (bnogPlacedObjectsMaps *BnogPlacedObjectsMaps) add_tuple_placed_object_to_dictionary(
	placing_ckid ckids.PlaceNumberTypeCkIds,
	relation_placed_object objects.BnogObjects) {

	bnogPlacedObjectsMaps.RelationPlacedObjectMap[placing_ckid] =
		relation_placed_object

}

func (bnogPlacedObjectsMaps *BnogPlacedObjectsMaps) try_get_relation_placed_object_for_place_number_ckid(
	placing_ckid ckids.PlaceNumberTypeCkIds) objects.BnogObjects {

	relation_placed_object := bnogPlacedObjectsMaps.RelationPlacedObjectMap[placing_ckid]

	return relation_placed_object
}
