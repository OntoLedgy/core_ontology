package factories

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type BnopTypeFactories struct {
}

func (BnopTypeFactories) Create(
	uuid *utils.UUIDs,
	owning_repository_uuids *utils.UUIDs,
	presentation_name string) *objects.BnogTypes {

	bnog_type := &objects.BnogTypes{}

	bnog_type.New(
		uuid,
		owning_repository_uuids,
		presentation_name)

	return bnog_type
}
