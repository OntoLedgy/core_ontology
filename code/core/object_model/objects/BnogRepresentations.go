package objects

import "github.com/OntoLedgy/storage_interop_services/code/services/databases/utils"

type BnogRepresentations struct {
	BnogTypes
	exemplar_representation string
}

func (representation *BnogRepresentations) New(
	uuid *utils.UUIDs,
	owning_repository_uuid *utils.UUIDs,
	exemplar_representation string,
	presentation_name string) {

	if uuid == nil {
		representation.Set_object_uuid()
	} else {
		representation.Object_uuid = uuid
	}

	representation.Owning_repository_uuid = owning_repository_uuid

	representation.Registry_keyed_on_uuid = make(map[*utils.UUIDs]interface{})

	representation.Registry_keyed_on_uuid[uuid] = representation

	representation.Uml_name = presentation_name

	representation.exemplar_representation = exemplar_representation

}
