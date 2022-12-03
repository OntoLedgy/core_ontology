package factories

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type BnogObjectFactories struct {
}

func (BnogObjectFactories) Create(
	uuid *utils.UUIDs,
	owning_repository_uuids *utils.UUIDs,
	presentation_name string) *objects.BnogObjects {

	bnog_object :=
		&objects.BnogObjects{}

	bnog_object.New(
		uuid,
		owning_repository_uuids,
		presentation_name)

	return bnog_object
}
