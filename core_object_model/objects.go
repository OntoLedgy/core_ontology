package core_object_model

import (
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type Objects struct {
	Object_uuid *utils.UUIDs
}

func (object Objects) Set_object_uuid() *utils.UUIDs {

	object.Object_uuid, _ = utils.GetUUID(1, "")

	return object.Object_uuid
}

func (object Objects) Get_object_uuid() *utils.UUIDs {

	return object.Object_uuid

}
