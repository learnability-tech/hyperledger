// SPDX-License-Identifier: Apache-2.0

/*
  Sample Chaincode based on Demonstrated Scenario

 This code is based on code written by the Hyperledger Fabric community.
  Original code can be found here: https://github.com/hyperledger/fabric-samples/blob/release/chaincode/fabcar/fabcar.go
 */

package main

/* Imports  
* 4 utility libraries for handling bytes, reading and writing JSON, 
formatting, and string manipulation  
* 2 specific Hyperledger Fabric specific libraries for Smart Contracts  
*/ 
import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	// "time"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}

/* Define Tuna structure, with 4 properties.  
Structure tags are used by encoding/json library
*/

type Person struct {
        PersonId string `json:"personId"`
        Name string `json:"name"`
        Gender  string `json:"gender"`
}

type Policy struct {
	PolicyNum string `json:"policyNum"`
	Plan      string `json:"plan"`
	CreateTime string `json:"createTime"`
	InsuredAmt int `json:"insuredAmt"`
	PersonId string `json:"personId"`
}

type Claim struct {
	ClaimId string `json:"claimId"`
	ClaimAmt int `json:"claimAmt"`
	ClaimedAmt int `json:"claimedAmt"`
	PersonId string `json:"personId"`
	PolicyNum string `json:"policyNum"`
	CreateTime string `json:"createTime"`
	Status string `json:"status"`
	Remarks string `json:"remarks"`
}

/*
 * The Init method *
 called when the Smart Contract "tuna-chaincode" is instantiated by the network
 * Best practice is to have any Ledger initialization in separate function 
 -- see initLedger()
 */
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

/*
 * The Invoke method *
 called when an application requests to run the Smart Contract "tuna-chaincode"
 The app also specifies the specific smart contract function to call with args
 */
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger
	if function == "createPerson" {
		return s.createPerson(APIstub, args)
	}else if function == "createPolicy" {
		return s.createPolicy(APIstub, args)
	}else if function == "createClaim" {
		return s.createClaim(APIstub, args)
	}else if function == "queryPerson" {
		return s.queryPerson(APIstub, args)
	}else if function == "queryPolicy" {
		return s.queryPolicy(APIstub, args)
	}else if function == "queryClaim" {
		return s.queryClaim(APIstub, args)
	}else if function == "confirmClaim" {
		return s.confirmClaim(APIstub, args)
	}else if function == "rejectClaim" {
		return s.rejectClaim(APIstub, args)
	}else if function == "queryByPersonId" {
		return s.queryByPersonId(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

/*
confirm claim
*/
func (s *SmartContract) confirmClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    claimAsBytes, _ := APIstub.GetState(args[0])
	if claimAsBytes == nil {
		return shim.Error("Could not locate Claim ")
	}
	claim := Claim{}

	json.Unmarshal(claimAsBytes, &claim)
	j, err := strconv.Atoi(args[1])
	if err != nil {
          return shim.Error(fmt.Sprintf("ClaimedAmt is not numeric: %s", args[0]))
    }
    if(strings.Compare(claim.Status, "open")!=0){
    	return shim.Error(fmt.Sprintf("Claim already processed %s", args[0]))
    }
	claim.ClaimedAmt = j
	claim.Status="claimed"
	claim.Remarks=args[2]
	claimAsBytes, _ = json.Marshal(claim)
	APIstub.PutState(args[0], claimAsBytes)
	
	return shim.Success(nil)	
}

/*
confirm claim
*/
func (s *SmartContract) rejectClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    claimAsBytes, _ := APIstub.GetState(args[0])
	if claimAsBytes == nil {
		return shim.Error("Could not locate Claim ")
	}
	claim := Claim{}

	json.Unmarshal(claimAsBytes, &claim)
	if(strings.Compare(claim.Status, "open")!=0){
    	return shim.Error(fmt.Sprintf("Claim already processed %s", args[0]))
    }
	claim.Status="rejected"
	claim.Remarks=args[1]
	claimAsBytes, _ = json.Marshal(claim)
	APIstub.PutState(args[0], claimAsBytes)
	
	return shim.Success(nil)	
}


/*
Query preson
*/
func (t *SmartContract) queryPerson(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    tunaAsBytes, _ := APIstub.GetState(args[0])
	if tunaAsBytes == nil {
		return shim.Error("Could not find people")
	}
	return shim.Success(tunaAsBytes)	
}

/*
Query policy
*/
func (t *SmartContract) queryPolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    tunaAsBytes, _ := APIstub.GetState(args[0])
	if tunaAsBytes == nil {
		return shim.Error("Could not locate policy")
	}
	return shim.Success(tunaAsBytes)	
}

/*
Query Claim
*/
func (t *SmartContract) queryClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
    tunaAsBytes, _ := APIstub.GetState(args[0])
	if tunaAsBytes == nil {
		return shim.Error("Could not locate claim")
	}
	return shim.Success(tunaAsBytes)	
}

/*
Create preson
*/
func (s *SmartContract) createPerson(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
     person := Person{PersonId: args[0], Name: args[1], Gender: args[2]}
     personAsBytes, _ :=json.Marshal(person)
     err := APIstub.PutState(args[0], personAsBytes)
     if err != nil {
                return shim.Error(fmt.Sprintf("Failed to record tuna catch: %s", args[0]))
     }

    return shim.Success(nil)		
}

/*
Create policy
*/
func (s *SmartContract) createPolicy(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	 personAsBytes, _ := APIstub.GetState(args[4])
	if personAsBytes == nil {
		return shim.Error("Could not find person ")
	}
	 j, err := strconv.Atoi(args[3])
     policy := Policy{PolicyNum: args[0], Plan: args[1], CreateTime: args[2], InsuredAmt: j,PersonId: args[4]}
     policyAsBytes, _ :=json.Marshal(policy)
     APIstub.PutState(args[0], policyAsBytes)
     if err != nil {
                return shim.Error(fmt.Sprintf("Failed to record policy catch: %s", args[0]))
     }

    return shim.Success(nil)		
}

/*
Create policy
*/
func (s *SmartContract) createClaim(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	policyAsBytes, _ := APIstub.GetState(args[4])
	if policyAsBytes == nil {
		return shim.Error("Could not find policy ")
	}
	personAsBytes, _ := APIstub.GetState(args[3])
	if personAsBytes == nil {
		return shim.Error("Could not find person ")
	}

	i, _ := strconv.Atoi(args[1])
	 // j, err := strconv.Atoi(args[3])
    claim := Claim{ClaimId: args[0], ClaimAmt: i, CreateTime: args[2], ClaimedAmt: 0,
				PersonId: args[3],PolicyNum: args[4], Status: "open", Remarks: " "}
     claimAsBytes, _ :=json.Marshal(claim)
     APIstub.PutState(args[0], claimAsBytes)
     

    return shim.Success(nil)		
}

func (t *SmartContract) queryByPersonId(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {
	holder := args[0]

	queryString := fmt.Sprintf("{\"selector\":{\"personId\":\"%s\"}}", holder)
	queryResults, err := APIstub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	defer queryResults.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for queryResults.HasNext() {
		queryResponse, err := queryResults.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// Add comma before array members,suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- queryAllTuna:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}




/*
 * main function *
calls the Start function 
The main function starts the chaincode in the container during instantiation.
 */
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
