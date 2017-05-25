package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"strconv"
)

type SimpleChaincode struct {
}
type PO_tier1 struct {
	Order_Id       string `json:"order_Id"`
	Order_Desc     string `json:"order_desc"`
	Order_Quantity string `json:"order_quantity"`
	Assigned_To_Id string `json:"assigned_to_id"`
	Created_By_Id  string `json:"created_by_id"`
	SubOrder_Id    string `json:"subOrder_Id"`
	Order_Status   string `json:"order_status"`
	Asset_ID       string `json:"asset_ID"`
}

type PO_OEM struct {
	Order_Id       string `json:"order_Id"`
	Order_Desc     string `json:"order_desc"`
	Order_Quantity string `json:"order_quantity"`
	Assigned_To_Id string `json:"assigned_to_id"`
	Created_By_Id  string `json:"created_by_id"`
	Order_Status   string `json:"order_status"`
	Asset_ID       string `json:"asset_ID"`
}

func main() {

	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)

	}
}

func (t *SimpleChaincode) convert(row shim.Row) PO_tier1 {
	var po = PO_tier1{}

	po.Order_Id = row.Columns[0].GetString_()
	po.Order_Desc = row.Columns[1].GetString_()
	po.Order_Quantity = row.Columns[2].GetString_()
	po.Assigned_To_Id = row.Columns[3].GetString_()
	po.Created_By_Id = row.Columns[4].GetString_()
	po.SubOrder_Id = row.Columns[5].GetString_()
	po.Order_Status = row.Columns[6].GetString_()
	po.Asset_ID = row.Columns[7].GetString_()

	return po
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	var err error
	err = stub.DelState("role_OEM")
	if err != nil {
		return nil, fmt.Errorf("remove operation failed. Error updating state: %s", err)
	}

	err = stub.DelState("role_tier_1")
	if err != nil {
		return nil, fmt.Errorf("remove operation failed. Error updating state: %s", err)
	}
	err = stub.DelState("role_first_tier_2")
	if err != nil {
		return nil, fmt.Errorf("remove operation failed. Error updating state: %s", err)
	}
	err = stub.DelState("role_second_tier_2")
	if err != nil {
		return nil, fmt.Errorf("remove operation failed. Error updating state: %s", err)
	}

	stub.PutState("role_OEM", []byte("OEM"))
	stub.PutState("role_tier_1", []byte("tier_1"))
	stub.PutState("role_first_tier_2", []byte("first_tier_2"))
	stub.PutState("role_second_tier_2", []byte("second_tier_2"))

	var columnsOrderTable []*shim.ColumnDefinition

	columnOne := shim.ColumnDefinition{Name: "Order_Id",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwo := shim.ColumnDefinition{Name: "Order_Desc",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnThree := shim.ColumnDefinition{Name: "Order_Quantity",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFour := shim.ColumnDefinition{Name: "Assigned_To_Id",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnFive := shim.ColumnDefinition{Name: "Created_By_Id",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnSix := shim.ColumnDefinition{Name: "SubOrder_Id",
		Type: shim.ColumnDefinition_STRING, Key: true}
	columnSeven := shim.ColumnDefinition{Name: "Order_Status",
		Type: shim.ColumnDefinition_STRING, Key: false}
	columnEight := shim.ColumnDefinition{Name: "Asset_ID",
		Type: shim.ColumnDefinition_STRING, Key: false}

	columnsOrderTable = append(columnsOrderTable, &columnOne)
	columnsOrderTable = append(columnsOrderTable, &columnTwo)
	columnsOrderTable = append(columnsOrderTable, &columnThree)
	columnsOrderTable = append(columnsOrderTable, &columnFour)
	columnsOrderTable = append(columnsOrderTable, &columnFive)
	columnsOrderTable = append(columnsOrderTable, &columnSix)
	columnsOrderTable = append(columnsOrderTable, &columnSeven)
	columnsOrderTable = append(columnsOrderTable, &columnEight)

	stub.CreateTable("PurchaseOrder", columnsOrderTable)

	orderId := "1000"
	stub.PutState("current_Order_Id", []byte(orderId))
	stub.PutState("current_SubOrder_Id", []byte(orderId))

	return nil, nil

}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "createOrder" {

		return createOrder(stub, args)
	}
	if function == "updateOrderStatus" {

		return updateOrderStatus(stub, args)
	}
	return nil, errors.New("Received unknown function invocation: " + function)

}
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "fetchAllOrders" {

		return fetchAllOrders(stub, args)
	}

	return nil, errors.New("Received unknown function invocation: " + function)
}

func createOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	//OrderId
	byteOrderId, err := stub.GetState("current_Order_Id")
	strOrderId := string(byteOrderId)
	intOrderId, _ := strconv.Atoi(strOrderId)

	currentId := intOrderId + 1
	strCurrentId := "PO" + strconv.Itoa(currentId)
	stub.PutState("current_Order_Id", []byte(strCurrentId))

	//Sub orderId
	byteSubOrderId, err := stub.GetState("current_SubOrder_Id")
	strSubOrderId := string(byteSubOrderId)
	intSubOrderId, _ := strconv.Atoi(strSubOrderId)

	currentSubId := intSubOrderId + 1
	strSubCurrentId := "PO" + strconv.Itoa(currentSubId)
	stub.PutState("current_SubOrder_Id", []byte(strSubCurrentId))

	col1Val := strCurrentId
	col2Val := args[0]
	col3Val := args[1]
	col4Val := args[2]
	col5Val := args[3]
	col6Val := strSubCurrentId
	col7Val := args[4]
	col8Val := args[5]

	var columns []*shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: col1Val}}
	col2 := shim.Column{Value: &shim.Column_String_{String_: col2Val}}
	col3 := shim.Column{Value: &shim.Column_String_{String_: col3Val}}
	col4 := shim.Column{Value: &shim.Column_String_{String_: col4Val}}
	col5 := shim.Column{Value: &shim.Column_String_{String_: col5Val}}
	col6 := shim.Column{Value: &shim.Column_String_{String_: col6Val}}
	col7 := shim.Column{Value: &shim.Column_String_{String_: col7Val}}
	col8 := shim.Column{Value: &shim.Column_String_{String_: col8Val}}

	columns = append(columns, &col1)
	columns = append(columns, &col2)
	columns = append(columns, &col3)
	columns = append(columns, &col4)
	columns = append(columns, &col5)
	columns = append(columns, &col6)
	columns = append(columns, &col7)
	columns = append(columns, &col8)

	row := shim.Row{Columns: columns}
	ok, err := stub.InsertRow("PurchaseOrder", row)

	if err != nil {
		return nil, fmt.Errorf("insertTableOne operation failed. %s", err)
		panic(err)

	}
	if !ok {
		return []byte("Row with given key" + args[0] + " already exists"), errors.New("insertTableOne operation failed. Row with given key already exists")
	}
	return nil, errors.New("Received unknown function invocation: ")
}

func fetchAllOrders(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var columns []shim.Column
	rowChannel, err := stub.GetRows("PurchaseOrder", columns)

	orderArray := []*PO_tier1{}

	for {
		select {

		case row, ok := <-rowChannel:

			if !ok {
				rowChannel = nil
			} else {
				po := new(PO_tier1)
				po.Order_Id = row.Columns[0].GetString_()
				po.Order_Desc = row.Columns[1].GetString_()
				po.Order_Quantity = row.Columns[2].GetString_()
				po.Assigned_To_Id = row.Columns[3].GetString_()
				po.Created_By_Id = row.Columns[4].GetString_()
				po.SubOrder_Id = row.Columns[5].GetString_()
				po.Order_Status = row.Columns[6].GetString_()
				po.Asset_ID = row.Columns[7].GetString_()

				orderArray = append(orderArray, po)
			}

		}
		if rowChannel == nil {
			break
		}
	}
	fmt.Println("Get Rows==================>" + orderArray)

	jsonRows, err := json.Marshal(orderArray)
	if err != nil {
		return nil, fmt.Errorf("getRowsTableFour operation failed. Error marshaling JSON: %s", err)
	}

	return jsonRows, nil

}

func updateOrderStatus(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}
