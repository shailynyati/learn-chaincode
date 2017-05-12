package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}
type PurchaseOrder struct {
	poID           string `json:"poID"`
	DealerID       string `json:"dealerID"`
	poAmount       string `json:"poamount"`
	poCreationDate string `json:"pocreationdate"`
	DateOfDelivery string `json:"dateOfdelivery"`
	//lineItems []LineItems `json:"lineItems"`
}

type Dealer struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
	Zip     string `json:"zip"`
}

type LineItems struct {
	itemId         string `json:id"`
	itemType       string `json:"type"`
	itemDesc       string `json:"desc"`
	itemQuantity   string `json:"quantity"`
	itemPrice      string `json:"unitPrice"`
	itemTotalPrice string `json:"totalPrice"`
}

//
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error in creating Purchase Order", err)
	}
}
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error
	var pobytes []byte
	//	var invoicebytes []byte

	if err != nil {
		return nil, errors.New("Error creating Purchase Order")
	}
	err = stub.PutState("po_Id", pobytes)
	return nil, nil

}

// Invoke is your entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "createDealer" {

		return createDealer(stub, args)
	}
	if function == "createPO" {

		return createPO(stub, args)
	}
	if function == "createInvoice" {

		return createInvoice(stub, args)
	}

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "getPO" {
		return t.getPO(stub, args)
	}
	if function == "getDealer" {
		return t.getDealer(stub, args)
	}
	if function == "getInvoice" {
		return t.getInvoice(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

func createDealer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	dealer := Dealer{}

	dealer.Id = args[0]
	dealer.Name = args[1]
	dealer.City = args[2]
	dealer.Country = args[3]
	dealer.Zip = args[4]
	dealerBytes, err := json.Marshal(dealer)

	if err != nil {
		return nil, errors.New("Problem while saving Purchase Order in BlockChain Network")
	}
	err = stub.PutState("Dealer-"+args[0], dealerBytes)
	return nil, nil

}
func createPO(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	purchaseOrder := PurchaseOrder{}

	purchaseOrder.poID = args[0]
	purchaseOrder.DealerID = args[1]
	purchaseOrder.poAmount = args[2]
	purchaseOrder.poCreationDate = args[3]
	purchaseOrder.DateOfDelivery = args[4]
	//purchaseOrder. = args[5]
	purchaseOrderBytes, err := json.Marshal(purchaseOrder)

	if err != nil {
		return nil, errors.New("Problem while saving Purchase Order in BlockChain Network")
	}
	err = stub.PutState("PO-"+args[0], purchaseOrderBytes)
	return nil, nil

}
func createInvoice(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//TODO
	return nil, nil
}

// Get PO - query function to read key/value pair

func (t *SimpleChaincode) getPO(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	if len(args) < 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}
	poID := args[0] //keys to read from chaincode
	fmt.Print(poID + " this is is the key ")
	poAsbytes, err := stub.GetState(poID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + poID + "\"}"
		return nil, errors.New(jsonResp)
	}
	return poAsbytes, nil
}

func (t *SimpleChaincode) getInvoice(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	//TODO
	return nil, nil
}

func (t *SimpleChaincode) getDealer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var jsonResp string
	var err error

	if len(args) < 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}
	dealerID := args[0] //keys to read from chaincode
	fmt.Print(dealerID + " this is is the key ")
	dealerAsbytes, err := stub.GetState(dealerID)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + dealerID + "\"}"
		return nil, errors.New(jsonResp)
	}
	return dealerAsbytes, nil

	return nil, nil
}
