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
	"github.com/hyperledger/fabric/core/chaincode/shim"
        "errors"
"encoding/json"
)

// Customer example simple Chaincode implementation
type customer struct {
}



// NewCertHandler creates a new reference to CertHandler
func NewCustomer() *customer {
    return &customer{}
}

func (t *customer) purchase(stub shim.ChaincodeStubInterface,args []string) ([]byte, error) {
     if len(args) != 1 {
    error := Error{"Incorrect number of arguments. Expecting 1"}
    errorMarshal, _ := json.Marshal(error)
    stub.SetEvent("purchase", errorMarshal)
    return nil, errors.New("Incorrect number of arguments. Expecting 1")
  }
    stub.PutState(args[0],[]byte("Purchased the Product"))
    return nil,nil
}
