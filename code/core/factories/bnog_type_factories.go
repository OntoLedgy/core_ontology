package factories

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type BnogTypeFactories struct {
}

func (BnogTypeFactories) Create(
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
