package factories

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type BnogElementFactories struct {
}

func (
	BnogElementFactories) Create(
	uuid *utils.UUIDs,
	owning_repository_uuids *utils.UUIDs) *objects.BnogElements {

	bnog_element :=
		&objects.BnogElements{}

	bnog_element.New(
		uuid,
		owning_repository_uuids)

	return bnog_element
}
