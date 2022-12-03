package objects

import (
	utils "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type BnogRepresentations struct {
	BnogTypes
	exemplarRepresentation string
}

func (bnogRepresentation *BnogRepresentations) New(
	uuid *utils.UUIDs,
	owningRepositoryUuid *utils.UUIDs,
	exemplarRepresentation string,
	presentationName string) {

	bnogRepresentation.BnogTypes.New(
		uuid,
		owningRepositoryUuid,
		presentationName)

	bnogRepresentation.RegistryKeyedOnUuid =
		make(
			map[*utils.UUIDs]*BnogObjects)

	bnogRepresentation.RegistryKeyedOnUuid[uuid] =
		&bnogRepresentation.BnogObjects

	bnogRepresentation.exemplarRepresentation =
		exemplarRepresentation

}
