package contracts

import "github.com/OntoLedgy/ol_common_services/code/services/identification_services/uuid_service"

//TODO review how this wil be used.

type IObjects interface {
	New(uuid *uuid_service.UUIDs,
		owningRepositoryUuid *uuid_service.UUIDs,
		presentationName string)
}
