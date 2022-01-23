package factories

import (
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	utils "github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
)

type BnogNameFactories struct{}

func (BnogNameFactories) Create(
	uuid *utils.UUIDs,
	exemplar_representation string,
	owning_repository_uuid *utils.UUIDs,
	presentation_name string) *objects.BnogNames {

	bnog_name :=
		&objects.BnogNames{}

	bnog_name.New(
		uuid,
		owning_repository_uuid,
		exemplar_representation,
		presentation_name)

	return bnog_name

}
