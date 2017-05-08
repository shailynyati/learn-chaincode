package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"log"
	"strconv"
)

type SimpleChaincode struct {
}

type UserRegistrationDetails struct {
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
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting User registration: %s", err)
	}
}

func (t *SimpleChaincode) RegisterUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering UserRegistration")
	log.Print("Entering UserRegistration")
	if len(args) < 2 {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected at least two arguments for User registration")
	}
	user := UserRegistrationDetails{
		Ffid:        args[0],
		Firstname:   args[1],
		Lastname:    args[2],
		DOB:         args[3],
		Email:       args[4],
		Address:     args[5],
		Country:     args[6],
		City:        args[7],
		Zip:         args[8],
		CreatedBy:   args[9],
		Title:       args[10],
		Gender:      args[11],
		TotalPoints: args[12]}

	UserRegistrationBytes, err := json.Marshal(user)
	//err = stub.PutState(args[0], UserRegistrationBytes)
	err = stub.PutState(args[0], UserRegistrationBytes)

	if err != nil {
		fmt.Println("Could not save UserRegistration to ledger", err)
		return nil, err
	}

	fmt.Println("Successfully saved User Registration")
	return nil, nil
}

//args[0] = id [<number>]
//args[1] = operator [add, delete]
//args[2] = points [<number>]
func (t *SimpleChaincode) AddDeletePoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var totalPoints int
	var pointsToModifyInt int

	operator := args[1]
	pointsToModify := args[2]

	userAsbytes, _ := t.getUser(stub, args)

	user := UserRegistrationDetails{}
	err := json.Unmarshal(userAsbytes, &user)

	fmt.Println("***********************")

	if err != nil {
		return nil, err
	}

	totalPoints, _ = strconv.Atoi(user.TotalPoints)
	pointsToModifyInt, _ = strconv.Atoi(pointsToModify)

	fmt.Println(totalPoints)
	fmt.Println(pointsToModifyInt)

	if operator == "Add" {
		totalPoints = totalPoints + pointsToModifyInt
	}

	if operator == "Delete" {
		if totalPoints < pointsToModifyInt {
			return nil, errors.New("Points not sufficient to spend")
		}
		totalPoints = totalPoints - pointsToModifyInt
	}

	user.TotalPoints = strconv.Itoa(totalPoints)
	UserRegistrationBytes, _ := json.Marshal(user)
	err = stub.PutState(args[0], UserRegistrationBytes)

	if err != nil {
		return nil, err
	}

	return nil, err
}

func (t *SimpleChaincode) getPoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	//var jsonResp error
	if len(args) < 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	user, err := t.getUser(stub, args)
	//fmt.Println("getpoints " + user)

	if err != nil {
		return nil, err
	}

	u := UserRegistrationDetails{}
	jsonResp := json.Unmarshal(user, &u)
	points := []byte(u.TotalPoints)
	return points, nil
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err := stub.PutState("User", []byte(args[0]))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Invoke is your entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "RegisterUser" {
		return t.RegisterUser(stub, args)
	}
	if function == "AddDeletePoints" {
		return t.AddDeletePoints(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	//	if function == "read" { //read a variable
	//		return t.read(stub, args)
	//	}
	if function == "getUser" {
		return t.getUser(stub, args)
	}
	if function == "getPoints" {
		return t.getPoints(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// Get User - query function to read key/value pair

func (t *SimpleChaincode) getUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	if len(args) < 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}
	ffId := args[0] //keys to read from chaincode
	fmt.Print(ffId + " this is is the key ")
	userAsbytes, err := stub.GetState(ffId)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}
	return userAsbytes, nil
}
