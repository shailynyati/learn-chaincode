package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type UserRegistrationsDetails struct {
	Ffid        string `json:"ffid"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	DOB         string `json:"DOB"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Zip         string `json:"zip"`
	CreatedBy   string `json:"createdby"`
	Title       string `json:"title"`
	Gender      string `json:"gender"`
	TotalPoints string `json:"totalPoints"`
}

func main() {
	err := shim.Start(new(UserRegistrationsDetails))
	if err != nil {
		fmt.Printf("Error starting User registration: %s", err)
	}
}

func (t *UserRegistrationsDetails) RegisterUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering UserRegistration")

	if len(args) < 2 {
		fmt.Println("Invalid number of args")
		return "error", errors.New("Expected at least two arguments for User registration")
	}

	var ffId = args[0]
	var UserRegistrationInput = args[1]
	var output string
	err := stub.PutState(ffId, []byte(UserRegistrationInput))
	if err != nil {
		output = "failure"
		stub.putState("output", err)

		fmt.Println("Could not save UserRegistration to ledger", err)
		return stub.getState(output), err
	}

	output = "success"
	fmt.Println("Successfully saved User Registration")
	return stub.getState(output), nil
}

// Init resets all the things
func (t *UserRegistrationsDetails) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("User-1", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke is your entry point to invoke a chaincode function
func (t *UserRegistrationsDetails) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "RegisterUser" {
		return t.RegisterUser(stub, args)
	}
	//	fmt.Println("invoke did not find func: " + function)
	//
	return nil, errors.New("Received unknown function invocation: " + function)
	//
}

// Query is our entry point for queries
func (t *UserRegistrationsDetails) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
//func (t *UserRegistrationsDetails) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
//	var key, value string
//	key = args[0] //rename for funsies
//	value = args[1]
//	if len(args) != 2 {
//		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
//	}
//
//	err = stub.PutState("sum", []byte(sum1)) //write the variable in chaincode state
//	//err = stub.PutState(key, []byte(sum))
//	if err != nil {
//		return nil, err
//	}
//	return nil, nil
//}

// read - query funct	ion to read key/value pair
func (t *UserRegistrationsDetails) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}
	key = args[0] //keys to read from chaincode
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}
	return valAsbytes, nil
}
