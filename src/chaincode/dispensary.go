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

// Dispensary example simple Chaincode implementation
type dispensary struct {
}

// NewCertHandler creates a new reference to CertHandler
func NewDispensary() *dispensary {
    return &dispensary{}
}

func (t *dispensary) placedOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
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
  /*  product:= args[1]
    quantity:= args[2]*/

  /*  var data = {"dispensaryDetails":id,"product":product,"quantity":quantity}
    orderAsJson,err = json.Marshal(data)*/
    stub.PutState(id,[]byte("Dispensary Place Order"))
    return nil,nil
}

func (t *dispensary) receivedShippmentOrder(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	 if len(args) != 1 {
    error := Error{"Incorrect number of arguments. Expecting 1"}
    errorMarshal, _ := json.Marshal(error)
    stub.SetEvent("receivedShippmentOrder", errorMarshal)
    return nil, errors.New("Incorrect number of arguments. Expecting 1")
  }
	id:= args[0]
    stub.PutState(id,[]byte("Received Shippment from Grower"))
    return nil,nil
}