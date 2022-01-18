package objects

import (
	"github.com/OntoLedgy/core_ontology/code/core/construction_operations"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type BoroObjectCkIds int

type BnogObjects struct {
	Object_uuid *utils.UUIDs

	registry_keyed_on_uuid map[*utils.UUIDs]interface{}

	registry_keyed_on_ckid_type map[BoroObjectCkIds]*interface{}

	matched_objects []*utils.UUIDs
}

func (object *BnogObjects) Set_object_uuid() *utils.UUIDs {

	object.Object_uuid, _ = utils.GetUUID(1, "")

	return object.Object_uuid
}

func (object *BnogObjects) Get_object_uuid() *utils.UUIDs {

	return object.Object_uuid

}

func (object *BnogObjects) New() *utils.UUIDs {

	object.Object_uuid, _ = utils.GetUUID(1, "")

	return object.Object_uuid
}

func (object *BnogObjects) add_to_registry_keyed_on_ckid_type(
	boro_object_ck_id BoroObjectCkIds) {

	var ckid_typed_objects interface{}

	if val, ok := object.registry_keyed_on_ckid_type[boro_object_ck_id]; ok {

		ckid_typed_objects = val

	} else {

		ckid_typed_objects = construction_operations.New[string]()

		object.registry_keyed_on_ckid_type[boro_object_ck_id] = &ckid_typed_objects

	}

	//ckid_typed_objects

}
