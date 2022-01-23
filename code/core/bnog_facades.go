package core

import (
	"github.com/OntoLedgy/core_ontology/code/ckids"
	"github.com/OntoLedgy/core_ontology/code/core/factories"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects/places"
	"github.com/OntoLedgy/core_ontology/code/migrations"
	uuids "github.com/OntoLedgy/ol_common_services/code/services/identity_management_services"
	"reflect"
)

type BnogFacades struct{}

func (BnogFacades) WriteBnogObjectToXml(
	xmlFilePath string) {

	migrations.OrchestrateJsonWrite(
		xmlFilePath)
}

func (BnogFacades) CreateBnogObject(
	objectUuid *uuids.UUIDs,
	owningRepositoryUuid *uuids.UUIDs,
	presentationName string) *objects.BnogObjects {

	bnog_object_factory :=
		&factories.BnogObjectFactories{}

	bnog_object :=
		bnog_object_factory.Create(
			objectUuid,
			owningRepositoryUuid,
			presentationName)

	return bnog_object
}

func (BnogFacades) CreateNewBnogType(
	owningRepositoryUuid *uuids.UUIDs) *objects.BnogTypes {

	bnog_type_factory := factories.BnogTypeFactories{}

	bnop_type :=
		bnog_type_factory.Create(
			nil,
			owningRepositoryUuid,
			"")

	return bnop_type
}

func (BnogFacades) CreateBnogType(
	typeUuid *uuids.UUIDs,
	owningRepositoryUuid *uuids.UUIDs,
	presentationName string) *objects.BnogTypes {

	bnog_type_factory := &factories.BnogTypeFactories{}

	bnog_type :=
		bnog_type_factory.Create(
			typeUuid,
			owningRepositoryUuid,
			presentationName)

	return bnog_type
}

func (BnogFacades) CreateNewBnogElement(
	owningRepositoryUuid *uuids.UUIDs) *objects.BnogElements {

	bnog_element_factory := &factories.BnogElementFactories{}

	bnog_element :=
		bnog_element_factory.Create(
			nil,
			owningRepositoryUuid)

	return bnog_element

}

func (BnogFacades) CreateBnogElement(
	elementUuid *uuids.UUIDs,
	owningRepositoryUuid *uuids.UUIDs) *objects.BnogElements {

	bnog_element_factory := &factories.BnogElementFactories{}

	bnog_element :=
		bnog_element_factory.Create(
			elementUuid,
			owningRepositoryUuid)

	return bnog_element

}

func (BnogFacades) CreateNewBnogName(
	exemplarRepresentation string,
	owningRepositoryUuid *uuids.UUIDs) *objects.BnogNames {

	bnog_name_factory := factories.BnogNameFactories{}

	bnog_name :=
		bnog_name_factory.Create(
			nil,
			exemplarRepresentation,
			owningRepositoryUuid,
			"")

	return bnog_name

}

func (BnogFacades) CreateBnogName(
	nameUuid *uuids.UUIDs,
	exemplarRepresentation string,
	owningRepositoryUuid *uuids.UUIDs,
	presentationName string) *objects.BnogNames {

	bnog_name_factory := factories.BnogNameFactories{}

	bnog_name :=
		bnog_name_factory.Create(
			nameUuid,
			exemplarRepresentation,
			owningRepositoryUuid,
			presentationName)

	return bnog_name

}

func (BnogFacades) CreateNewBnogTupleFromTwoPlacedObjects(
	placed1Object *objects.BnogObjects,
	placed2Object *objects.BnogObjects,
	immutableMinorCompositionCoupleTypeBoroObjectCkid ckids.BoroObjectCkIds,
	owningRepositoryUuid *uuids.UUIDs) *objects.BnogTuples {

	tuple_placed_objects_map :=
		&places.BnogPlacedObjectsMaps{}

	tuple_placed_objects_map.Add_tuple_placed_object_to_dictionary(
		ckids.PlaceNumberOnes,
		placed1Object)

	tuple_placed_objects_map.Add_tuple_placed_object_to_dictionary(
		ckids.PlaceNumberTwos,
		placed2Object)

	bnog_tuple_factory := &factories.BnogTupleFactories{}

	bnog_tuple :=
		bnog_tuple_factory.Create(
			nil,
			tuple_placed_objects_map,
			immutableMinorCompositionCoupleTypeBoroObjectCkid,
			owningRepositoryUuid)

	return bnog_tuple

}

func (BnogFacades) CreateBnogTupleFromTwoPlacedObjects(
	tupleUuid *uuids.UUIDs,
	placed1Object *objects.BnogObjects,
	placed2Object *objects.BnogObjects,
	immutableMinorCompositionCoupleTypeBoroObjectCkid ckids.BoroObjectCkIds,
	owningRepositoryUuid *uuids.UUIDs) *objects.BnogTuples {

	tuple_placed_objects_map :=
		&places.BnogPlacedObjectsMaps{}

	tuple_placed_objects_map.Add_tuple_placed_object_to_dictionary(
		ckids.PlaceNumberOnes,
		placed1Object)

	tuple_placed_objects_map.Add_tuple_placed_object_to_dictionary(
		ckids.PlaceNumberTwos,
		placed2Object)

	bnog_tuple_factory := &factories.BnogTupleFactories{}

	bnop_tuple :=
		bnog_tuple_factory.Create(
			tupleUuid,
			tuple_placed_objects_map,
			immutableMinorCompositionCoupleTypeBoroObjectCkid,
			owningRepositoryUuid)
	//TODO

	if immutableMinorCompositionCoupleTypeBoroObjectCkid == ckids.SuperSubTypes {

		placed2Object.Supertypes.Add(
			placed1Object)

		placed1Object.Supertypes.Add(
			placed2Object)
	}

	if immutableMinorCompositionCoupleTypeBoroObjectCkid == ckids.NamedBy {

		placed1Object.IsNamedBys.Add(
			placed2Object)

		placed2Object.Names.Add(
			placed1Object)

	}

	if immutableMinorCompositionCoupleTypeBoroObjectCkid == ckids.TypesInstances {
		placed2Object.Types.Add(
			placed1Object)

		placed1Object.Instances.Add(
			placed2Object)
	}
	names_type := reflect.TypeOf((*objects.BnogNames)(nil))

	if reflect.TypeOf(placed2Object) == reflect.TypeOf(names_type) {
		placed2Object.NamingSpaces.Add(
			placed1Object)

	}

	return bnop_tuple

}

func (BnogFacades) CreateBnogPlaceabletype(
	placeabletypeUuid *uuids.UUIDs,
	owningRepositoryUuid *uuids.UUIDs,
	presentationName string) *objects.BnogPlaceableTypes {

	placeable_type_placed_type_map :=
		&places.BnogPlacedObjectsMaps{}

	bnog_placeable_type_factory := factories.BnogPlaceableTypeFactories{}

	bnog_placeabletype :=
		bnog_placeable_type_factory.Create(
			placeabletypeUuid,
			placeable_type_placed_type_map,
			owningRepositoryUuid,
			presentationName)

	return bnog_placeabletype
}

func (BnogFacades) CreateBnogTypePlace(
	typePlaceUuid *uuids.UUIDs,
	placingType *objects.BnogTypes,
	placedType *objects.BnogTypes,
	typePlaceCkid ckids.BoroObjectCkIds,
	owningRepositoryUuid *uuids.UUIDs) *objects.BnogTypePlaces {

	bnog_type_place_factory := &factories.BnogTypePlaceFactories{}

	bnog_type_place :=
		bnog_type_place_factory.Create(
			typePlaceUuid,
			placingType,
			placedType,
			typePlaceCkid,
			owningRepositoryUuid)

	return bnog_type_place

}

func (BnogFacades) AddBnogPlacedTypeToPlaceabletype(
	placedType *objects.BnogTypes,
	placeNumberCkid ckids.PlaceNumberTypeCkIds,
	placeableType *objects.BnogPlaceableTypes) {

	placeableType.PlaceableTypePlacedTypesMap.RelationPlacedObjectMap[placeNumberCkid] =
		&placedType.BnogObjects

}
