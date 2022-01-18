package mappings

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"
)

type BoroObjectCkidBnogMappings struct {
	registry_keyed_on_ckid map[ckids.BoroObjectCkIds]objects.BnogObjects
	registry_keyed_on_bnog map[*utils.UUIDs]ckids.BoroObjectCkIds
}
