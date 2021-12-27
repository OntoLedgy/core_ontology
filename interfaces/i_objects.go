package interfaces

import "database_manager/utils"

type IObjects interface {
	Set_object_uuid() *utils.UUID
	Get_object_uuid() *utils.UUID
}
