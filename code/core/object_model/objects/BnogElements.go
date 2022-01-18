package objects

import "github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"

type BnogElements struct {
	BnogObjects
}

func (element *BnogElements) New(
	uuid *utils.UUIDs,
	owning_repository_uuid *utils.UUIDs,
	presentation_name string) {

	//element.BnogObjects.New()
	//element.add_to_registry_keyed_on_ckid_type(
	//	BoroObjectCkIds.Elements)

}
