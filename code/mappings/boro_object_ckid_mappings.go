package mappings

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
)

type BoroObjectCkidBnogMappings struct {
	registryKeyedOnCkid map[ckids.BoroObjectCkIds]*objects.BnogObjects
	registryKeyedOnBnog map[*objects.BnogObjects]ckids.BoroObjectCkIds
	Ckid                ckids.BoroObjectCkIds
	bnogObject          *objects.BnogObjects
}

func (
	boroObjectCkidBnogMappings *BoroObjectCkidBnogMappings) New(
	boroObjectCkid ckids.BoroObjectCkIds,
	bnogObject *objects.BnogObjects) {

	boroObjectCkidBnogMappings.Ckid = boroObjectCkid

	boroObjectCkidBnogMappings.bnogObject =
		bnogObject

	boroObjectCkidBnogMappings.registryKeyedOnCkid[boroObjectCkidBnogMappings.Ckid] =
		boroObjectCkidBnogMappings.bnogObject

	boroObjectCkidBnogMappings.registryKeyedOnBnog[bnogObject] =
		boroObjectCkidBnogMappings.Ckid

}

func (
	boroObjectCkidBnogMappings *BoroObjectCkidBnogMappings) getBnopType(
	boroObjectCkid ckids.BoroObjectCkIds) *objects.BnogObjects {

	bnogObject :=
		boroObjectCkidBnogMappings.registryKeyedOnCkid[boroObjectCkid]

	return bnogObject
}

func (boroObjectCkidBnogMappings *BoroObjectCkidBnogMappings) get_ckid(
	bnogObject *objects.BnogObjects) ckids.BoroObjectCkIds {

	boro_object_ckid :=
		boroObjectCkidBnogMappings.registryKeyedOnBnog[bnogObject]

	return boro_object_ckid

}
