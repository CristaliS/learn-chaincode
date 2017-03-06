package main

import (
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// CarFleetManagement simple Chaincode implementation
type CarFleetManagement struct {
}

func main() {
	err := shim.Start(new(CarFleetManagement))
	if err != nil {
		fmt.Printf("Error starting CarFleetManagement chaincode: %s", err)
	}
}

// Init resets the chaincode
func (t *CarFleetManagement) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2")
	}

	err1 := stub.PutState("rentInfo", []byte(args[0]))

	if err1 != nil {
		return nil, err1
	}

	err2 := stub.PutState("agency", []byte(args[1]))

	if err2 != nil {
		return nil, err2
	}

	return nil, nil
}

// Invoke is our entry point to invoke a chaincode function
func (t *CarFleetManagement) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.Write(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

func (t *CarFleetManagement) Write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	value = args[0]
	err = stub.PutState("rentInfo", []byte(value))
	if err != nil {
		return nil, err
	}
	value = args[1]
	err = stub.PutState("agency", []byte(value))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Query is our entry point for queries
func (t *CarFleetManagement) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	if function == "read" {
		return t.Read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

func (t *CarFleetManagement) Read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}
