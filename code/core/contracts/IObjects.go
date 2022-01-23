package contracts

import "github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"

//TODO review how this wil be used.

type IObjects interface {
	New(uuid *identity_management_services.UUIDs,
		owningRepositoryUuid *identity_management_services.UUIDs,
		presentationName string)
}
