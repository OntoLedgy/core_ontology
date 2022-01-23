package objects

import (
	"github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
)

type BnogElements struct {
	BnogObjects
}

func (bnogElement *BnogElements) New(
	uuid *identity_management_services.UUIDs,
	owningRepositoryUuid *identity_management_services.UUIDs) {

	bnogElement.ObjectUuid = uuid
	bnogElement.OwningRepositoryUuid = owningRepositoryUuid

}
