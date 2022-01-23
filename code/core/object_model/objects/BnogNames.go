package objects

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
	sets "github.com/OntoLedgy/storage_interop_services/code/services/in_memory/sets"
)

type BnogNames struct {
	BnogRepresentations
	NamingSpaces *sets.Sets[string]
}

func (bnogName *BnogNames) New(
	uuid *identity_management_services.UUIDs,
	owningRepositoryUuid *identity_management_services.UUIDs,
	exemplarRepresentation string,
	presentationName string) {

	bnogName.BnogRepresentations.New(
		uuid,
		owningRepositoryUuid,
		exemplarRepresentation,
		presentationName)

	bnogName.addToRegistryKeyedOnCkidType(
		ckids.Names)
}
