package objects

import (
	"github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"
)

type BnogElements struct {
	BnogObjects
}

func (bnogElement *BnogElements) New(
	uuid *uuid_service.UUIDs,
	owningRepositoryUuid *uuid_service.UUIDs) {

	bnogElement.ObjectUuid = uuid
	bnogElement.OwningRepositoryUuid = owningRepositoryUuid

}
