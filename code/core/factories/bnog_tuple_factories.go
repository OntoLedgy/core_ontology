package factories

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects/places"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type BnogTupleFactories struct {
}

func (BnogTupleFactories) Create(
	uuid *utils.UUIDs,
	tuple_placed_objects_map *places.BnogPlacedObjectsMaps,
	immutable_minor_composition_couple_type_boro_object_ckid ckids.BoroObjectCkIds,
	owning_repository_uuid *utils.UUIDs) *objects.BnogTuples {

	bnog_tuple := &objects.BnogTuples{}

	bnog_tuple.New(
		uuid,
		tuple_placed_objects_map,
		immutable_minor_composition_couple_type_boro_object_ckid,
		owning_repository_uuid)

	return bnog_tuple

}
