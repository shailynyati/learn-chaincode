/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func SumProductDiff(i, j string) (string, string, string) {

	var a, b int
	//var sum, prod, diff int

	a, err1 := strconv.Atoi(i)

	if err1 != nil {
		// handle error
	}
	b, err2 := strconv.Atoi(j)
	if err2 != nil {
		// handle error
	}

	sum := a + b

	prod := a * b

	diff := a - b

	return strconv.Itoa(sum), strconv.Itoa(prod), strconv.Itoa(diff)
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "write" {
		return t.write(stub, args)
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
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	key = args[0] //rename for funsies
	value = args[1]
	//var sum string
	sum1, prod2, diff3 := SumProductDiff(key, value)
	fmt.Println("Sum:", sum1, "| Product:", prod2, "| Diff:", diff3)
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	//s := []string{value, " From Shaily"}
	//s1 := strings.Join(s, ",")

	err = stub.PutState("sum", []byte(sum1)) //write the variable in chaincode state
	//err = stub.PutState(key, []byte(sum))
	if err != nil {
		return nil, err
	}
	return nil, nil
}

//added function concatstring

// read - query funct	ion to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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
