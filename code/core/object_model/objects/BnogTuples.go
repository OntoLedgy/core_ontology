package objects

type BnogTuples struct {
	BnogObjects
	Immutable_minor_composition_couple_type_boro_object_ckid BoroObjectCkIds
	tuple_placed_objects_dictionary                          map[*BnogTuples]interface{}
}
