package bnog_io

import (
	"encoding/gob"
	"fmt"
	"github.com/OntoLedgy/core_ontology/code/core/object_model/objects"
	uuid "github.com/satori/go.uuid"
	"os"
)

//TODO - add full structure of objects required.
type ObjectRegister []uuid.UUID

func WriteBnog(
	fullFileName string,
	bnogObjects *objects.BnogObjects) {

	objectRegister := ObjectRegister{}

	dataFile, err :=
		os.Create(
			fullFileName)

	if err != nil {

		fmt.Println(
			err)

		os.Exit(
			1)
	}

	// serialize the data
	dataEncoder :=
		gob.NewEncoder(
			dataFile)

	for _, value := range bnogObjects.RegistryKeyedOnUuid {

		objectRegister =
			append(objectRegister, *value.ObjectUuid.UUID)

	}

	dataEncoder.
		Encode(
			&objectRegister)

	dataFile.Close()
}

func ReadBnog(
	fullFileName string) {

	objectRegister := &ObjectRegister{}

	// open data file
	dataFile, err :=
		os.Open(
			fullFileName)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	dataDecoder :=
		gob.NewDecoder(
			dataFile)

	err =
		dataDecoder.Decode(
			objectRegister)

	if err != nil {

		fmt.Println(
			err)

		os.Exit(
			1)
	}

	dataFile.Close()

	fmt.Println("bnog detailes", objectRegister)
}
