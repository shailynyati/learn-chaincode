package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"encoding/json"
)

type SimpleChaincode struct {
}

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
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting User registration: %s", err)
	}
}

func (t *SimpleChaincode) RegisterUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	fmt.Println("Entering UserRegistration")

	if len(args) < 2 {
		fmt.Println("Invalid number of args")
		return nil, errors.New("Expected at least two arguments for User registration")
	}

	var ffId = args[0]
	var UserRegistrationInput = args[1]
	var UserRegistrationInput UserRegistrationDetails
	//var output string
	UserRegistrationBytes,err := json.Marshal(&UserRegistrationInput)
	err = stub.PutState(ffId, UserRegistrationBytes))

	if err != nil {
		//output = "failure"
		//stub.PutState(output, nil)

		fmt.Println("Could not save UserRegistration to ledger", err)
		return nil, err
	}

	//output = "success"
	fmt.Println("Successfully saved User Registration")
	return nil, nil
}

func AddDeletePoints(ffId string, operator string, points int)(string){
	
	var output string
	user UserRegistrationDetails := getUser(ffId)
	totalPoints := user.
	if(operator=="Add"){
		totalPoints = totalPoints+points
		output = "success"
	}
		else if(totalPoints == 0 or points>totalPoints){
			return output = "failure"
			else{
				totalPoints = totalPoints-points
				output="success"	
			}
	}
		
		return output	
}


func getPoints(string ffid)(int){
	
}






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
	//
	return nil, errors.New("Received unknown function invocation: " + function)
	//
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "getUser" { //read a variable
		return t.getUser(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// Get User - query function to read key/value pair
func (t *SimpleChaincode) getUser(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var  jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}
	ffId := args[0] //keys to read from chaincode
	
	userAsbytes, err := stub.GetState(ffId)
	gettUserAsJson(userAsbytes)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + ffId + "\"}"
		return nil, errors.New(jsonResp)
	}
	return userAsbytes, nil
}


func  getUserAsJson([]byte userAsbytes) (string){
	
	var user UserRegistrationsDetails
	userAsJson := json.Unmarshal(userAsbytes,&user)
	
	return userAsJson
	
	
}
