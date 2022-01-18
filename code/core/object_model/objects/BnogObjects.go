package objects

import (
	"github.com/OntoLedgy/core_ontology/code/core/construction_operations"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type BoroObjectCkIds int

type BnogObjects struct {
	Object_uuid *utils.UUIDs

	Registry_keyed_on_uuid map[*utils.UUIDs]interface{}

	Registry_keyed_on_ckid_type map[BoroObjectCkIds]interface{}

	Matched_objects []*utils.UUIDs

	Owning_repository_uuid *utils.UUIDs

	Is_named_bys *construction_operations.Sets[string]

	Names construction_operations.Sets[string]

	Types construction_operations.Sets[string]

	Instances construction_operations.Sets[string]

	Supertypes construction_operations.Sets[string]

	Uml_name string

	Naming_spaces construction_operations.Sets[string]
}

func (object *BnogObjects) New(
	uuid *utils.UUIDs,
	owning_repository_uuid *utils.UUIDs,
	presentation_name string) {

	if uuid == nil {
		object.Set_object_uuid()
	} else {
		object.Object_uuid = uuid
	}

	object.Owning_repository_uuid = owning_repository_uuid

	object.Registry_keyed_on_uuid = make(map[*utils.UUIDs]interface{})

	object.Registry_keyed_on_uuid[uuid] = object

	object.Uml_name = presentation_name

}

func (object *BnogObjects) Set_object_uuid() *utils.UUIDs {

	object.Object_uuid, _ = utils.GetUUID(1, "")

	return object.Object_uuid
}

func (object *BnogObjects) Get_object_uuid() *utils.UUIDs {

	return object.Object_uuid

}

func (object *BnogObjects) add_to_registry_keyed_on_ckid_type(
	boro_object_ck_id BoroObjectCkIds) {

	var ckid_typed_objects interface{}

	if val, ok := object.Registry_keyed_on_ckid_type[boro_object_ck_id]; ok {

		ckid_typed_objects = val

	} else {

		ckid_typed_objects = construction_operations.New[string]()

		object.Registry_keyed_on_ckid_type[boro_object_ck_id] = &ckid_typed_objects

	}

	//ckid_typed_objects

}
