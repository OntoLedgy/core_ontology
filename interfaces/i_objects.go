package interfaces

import (
	"github.com/OntoLedgy/storage_interop_services/databases/utils"
)

type IObjects interface {
	Set_object_uuid() *utils.UUIDs
	Get_object_uuid() *utils.UUIDs
}
