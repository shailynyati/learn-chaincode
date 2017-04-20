package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"log"
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

var userAsbytes []byte
var UserRegistrationInput UserRegistrationDetails

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
	var user = UserRegistrationDetails{
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
	//var output string
	//var UserRegistrationBytes []byte
	//var err string
	UserRegistrationBytes, err := json.Marshal(user)
	err = stub.PutState(user.Ffid, UserRegistrationBytes)

	if err != nil {
		fmt.Println("Could not save UserRegistration to ledger", err)
		return nil, err
	}

	fmt.Println("Successfully saved User Registration")
	return nil, nil
}

//func AddDeletePoints(ffId string, operator string, points int)(string){
//
//	var output string
//	var totalPoints int = getPoints(ffId)
//	if(operator=="Add"){
//		totalPoints = totalPoints+points
//		output = "success"
//	}
//	else if((totalPoints == 0) && (points>totalPoints)){
//		output = "failure"
//	}
//	 else
//	 {
//		totalPoints = totalPoints-points
//		output="success"
//	 }
//	return output
//}

//func getPoints(string ffid) int {
//	var user UserRegistrationDetails = getUser(ffId)
//	return strconv.Atoi(user.TotalPoints)
//}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
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
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "RegisterUser" {
		return t.RegisterUser(stub, args)
	}
	//	fmt.Println("invoke did not find func: " + function)
	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// Get User - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}
	ffId := args[0] //keys to read from chaincode

	userAsbytes, err = stub.GetState(ffId)
	//getUser(ffId)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}
	return userAsbytes, nil
}

//func getUser(ffId string) (UserRegistrationDetails){
//	var user UserRegistrationDetails
//	var err = json.Unmarshal(userAsbytes,&user)
//	if(user.Ffid==UserRegistrationInput.Ffid)
//		return user
//
//
//}
