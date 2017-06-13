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
"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// Grower example simple Chaincode implementation
type Grower struct {
}

type Error struct{
  Err string
}


var ddispensary = NewDispensary()
var ccustomer = NewCustomer()


// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
    err := shim.Start(new(Grower))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}


// Init resets all the things
func (t *Grower) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    if len(args) != 1 {
        return nil, errors.New("Incorrect number of arguments. Expecting 1")
    }

    err := stub.PutState("hello_world", []byte(args[0]))
    if err != nil {
        return nil, err
    }

    return nil, nil
}


// Invoke is our entry point to invoke a chaincode function
func (t *Grower) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("********************************Invoke****************************************")

    //   Handle different functions
    if function == "receiveOrder" {
        // Receive Order from Dispensary
        return t.receiveOrder(stub, args)
    } else if function == "shipOrder" {
        // Transfer product to Dispensary
        return t.shipOrder(stub,args)
    }else if function == "dispensaryPlacedOrder" {
        // Order Placed from Dispensary
        return t.dispensaryPlacedOrder(stub, args)
    }else if function == "shipmentReceivedByDispensary" {
        // Grower shippment to Dispensary
        return t.shipmentReceivedByDispensary(stub, args)
    }else if function == "customerPurchase" {
        // Purchased by Customer
        return t.customerPurchase(stub,args)
    }
    return nil, errors.New("Received unknown function invocation")
}


func (t *Grower) receiveOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
 if len(args) != 3 {
    error := Error{"Incorrect number of arguments. Expecting 3"}
    errorMarshal, _ := json.Marshal(error)
    stub.SetEvent("receiveOrderError", errorMarshal)
    return nil, errors.New("Incorrect number of arguments. Expecting 3")
  }
  // ==== Input sanitation ====
  fmt.Println("- start Receiving Order")
  if len(args[0]) <= 0 {
    error := Error{"1st argument must be a non-empty string"}
    errorMarshal, _ := json.Marshal(error)
    stub.SetEvent("receiveOrderError", errorMarshal)
    return nil, errors.New("1st argument must be a non-empty string")
  }
  if len(args[1]) <= 0 {
    error := Error{"2nd argument must be a non-empty string"}
    errorMarshal, _ := json.Marshal(error)
    stub.SetEvent("receiveOrderError", errorMarshal)
    return nil, errors.New("2nd argument must be a non-empty string")
  }
  if len(args[2]) <= 0 {
    error := Error{"3rd argument must be a non-empty string"}
    errorMarshal, _ := json.Marshal(error)
    stub.SetEvent("receiveOrderError", errorMarshal)
    return nil, errors.New("3rd argument must be a non-empty string")
  }

id:= args[0]
    /*product:= args[1]
    quantity:= args[2]*/

    /*var data = {"dispensaryDetails":id,"product":product,"quantity":quantity}
    orderAsJson,err = json.Marshal(data)*/
    stub.PutState(id,[]byte("Grower Received Order"))
    return nil,nil
}


func (t *Grower) dispensaryPlacedOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
         ddispensary.placedOrder(stub, args)
        return nil,nil
}


func (t *Grower) shipmentReceivedByDispensary(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
      ddispensary.receivedShippmentOrder(stub,args)
        return nil,nil
}


func (t *Grower) customerPurchase(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
      ccustomer.purchase(stub,args)
        return nil,nil
}


func (t *Grower) getOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte,error) {
    fmt.Println("getCircle called")
    if len(args) !=1{
        return nil,errors.New("Incorrect number of arguments. Expecting 1")
    }
    Id := args[0]
    order, err := stub.GetState(Id)
    if err != nil {
        return nil,err
    }

    return order,nil
}

func (t *Grower) shipOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

if len(args) != 2 {
    error := Error{"Incorrect number of arguments. Expecting 2"}
    errorMarshal, _ := json.Marshal(error)
    stub.SetEvent("transferOrderError", errorMarshal)
    return nil, errors.New("Incorrect number of arguments. Expecting 2")
  }
    id:= args[0]
    orderAsJson,err := stub.GetState(id)
if err != nil {
        return nil,err
    }
    if len(orderAsJson)==0{
        fmt.Println("Order doesnt exists")
        error := Error{"Order doesnt exists"}
        errorMarshal, _ := json.Marshal(error)
        stub.SetEvent("transferOrderError", errorMarshal)
        return nil, errors.New("Order doesnt exists")
    }

    /*var order
    json.Unmarshal(orderAsJson,&order)*/
    stub.PutState(id,[]byte("Shipped By Grower"))
    return nil,nil
}

// Query is our entry point for queries
func (t *Grower) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("Query is running " + function)

    if function =="getOrder"{
            return t.getOrder(stub,args)
        }
    return nil,errors.New("Received unknown function query")
}