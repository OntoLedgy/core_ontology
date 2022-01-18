package objects

import "github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"

type BnogElements struct {
	BnogObjects
}

func (element *BnogElements) New(
	uuid *utils.UUIDs,
	owning_repository_uuid *utils.UUIDs) {

	element.Object_uuid = uuid
	element.Owning_repository_uuid = owning_repository_uuid

}
