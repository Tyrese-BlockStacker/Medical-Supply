## Installing and setup:
1. install [Hyperledger Fabric](https://hyperledger-fabric.readthedocs.io/en/latest/getting_started.html)  
2. Place this medical-supply repository in the fabric-samples repository  
3. Make sure to run ```go mod vendor``` in each folder with a go.mod file to get all the dependencies

_________________________
## Starting and stopping the network:
Go to ```cd medical-supply```  
Starting the network: ```source networkDeploy.sh```  
Stopping the network: ```source networkClean.sh```  

_________________________
## Deploying the chaincode (smart contract):

### For customers:
```
cd medical-supply/customers

source customers.sh
```

### For Regulators:
```
cd medical-supply/regulators

source regulators.sh
```

__________________________
## Application:
```
cd medical-supply/customers/application/
go run app.go
```

__________________________
## For Monitoring docker

To see the Fabric nodes running on the local machine: ```docker ps```
To view the network: ```docker network inspect fabric_test```

peer0.org1.example.com will be used for the Customers  
peer0.org2.example.com will be used for the Regulators


## For Monitoring as either customer or regulator
Go to their respective folder
``` cd medical-supply/customers ``` or ``` cd medical-supply/regulators ```

To show output from the Docker containers:
```./configuration/cli/monitordocker.sh fabric_test``` or alternatively if port number doesn't work: ```./monitordocker.sh fabric_test <port_number>```

## For Monitoring using Hyperledger Explorer

![alt](images/explorer.png?raw=true "Hyperledger Explorer")
1. Start test-network using networkDeploy.sh
2. Go to explorer folder: ```cd medical-supply/explorer```
3. Run: ```docker-compose up -d``` to start the Hyperledger Explorer 
4. Go to ```https://localhost:8080``` for the Hyperledger Explorer.   
For the login screen:    
username: exploreradmin   
password: exploreradminpw  
Note: These can be changed in the ```test-network.json``` file.
5. Run: ```docker-compose down -v``` to stop the Hyperledger Explorer

## For Running benchmark tests using Hyperledger Caliper
![alt](images/caliper.png?raw=true "Hyperledger Caliper")
1. Install npm and run ```npm install``` inside caliper folder
2. Start test-network using ```source networkDeploy.sh```.
3. run ```source customersSetup.sh``` and ```source regulatorsSetup.sh``` just like when deploying the chaincode.
4. run ```npx caliper bind --caliper-bind-sut fabric:2.2``` to bind hyperledger caliper to hyperledger fabric. Note: fabric version 2.3 did not work at the time of this project.
5. run ```npx caliper launch manager --caliper-fabric-gateway-enabled``` to start running the tests.


## possible errors:
problem: Got permission denied while trying to connect to the Docker daemon socket at unix:///var/run/docker.sock
run: sudo chmod 666 /var/run/docker.sock

problem: permission denied to /dev/tpm0 or /dev/tpmrm0
run: sudo chown <username> /dev/tpm0

replace - CORE_VM_ENDPOINT=unix:///host/var/run/docker.sock in 
  
package main

import (
	"fmt"

	"github.com/google/go-tpm/tpm2"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	pb "github.com/hyperledger/fabric-protos-go/peer"
)

// For samples check https://github.com/hyperledger-archives/education/blob/master/LFS171x/fabric-material/chaincode/sample-chaincode.go

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func (s *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

func (s *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Printf("The function was invoked\n")
	// init code
	rwc, err := tpm2.OpenTPM("/dev/tpmrm0")
	if err != nil {
		shim.Error(fmt.Sprintf("could not open tpm:, %v", err))
	}
	res, err := tpm2.GetRandom(rwc, 2)
	if err != nil {
		shim.Error(fmt.Sprintf("could not get random numbers: %v", err))
	}
	fmt.Printf("Yoo it works res is %v\n", res)
	return shim.Success(res)
}

//NOTE - parameters such as ccid and endpoint information are hard coded here for illustration. This can be passed in in a variety of standard ways
func main() {
	// rwc, err_ := tpm2.OpenTPM("/dev/tpmrm0")
	// if err_ != nil {
	// 	fmt.Print("oeps eerste")
	// 	// shim.Error(fmt.Errorf("could not open tpm:, %v", err_).Error())
	// }
	// res, err_ := tpm2.GetRandom(rwc, 2)
	// if err_ != nil {
	// 	fmt.Print("oeps tweede")
	// 	// shim.Error(fmt.Errorf("could not get random numbers: %v", err_).Error())
	// }
	// fmt.Printf("%v", res)
	//The ccid is assigned to the chaincode on install (using the “peer lifecycle chaincode install <package>” command) for instance
	ccid := "mycc:724bf2be51e5a9e98d79d15d482d4fb1666af022c7f1368de18dec355d839da8"

	server := &shim.ChaincodeServer{
		CCID:    ccid,
		Address: "localhost:7075",
		CC:      new(SimpleChaincode),
		TLSProps: shim.TLSProperties{
			Disabled: true,
		},
	}
	fmt.Printf("Starting chaincode %s at %s\n", server.CCID, server.Address)
	err := server.Start()
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
