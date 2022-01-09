/*
Copyright 2020 IBM All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/gateway"
)

const (
	mspID         = "Org1MSP"
	appUser       = "alice"
	peerEndpoint  = "localhost:7051"
	gatewayPeer   = "peer0.org1.example.com"
	channelName   = "mychannel"
	chaincodeName = "medicinecontract"
)

func main() {
	log.Println("============ Application starts ============")

	err := os.Setenv("DISCOVERY_AS_LOCALHOST", "true")
	if err != nil {
		log.Fatalf("\nError setting DISCOVERY_AS_LOCALHOST environemnt variable: %v", err)
	}

	wallet, err := gateway.NewFileSystemWallet("wallet")
	if err != nil {
		log.Fatalf("\nFailed to create wallet: %v", err)
	}

	if !wallet.Exists(appUser) {
		err = populateWallet(wallet)
		if err != nil {
			log.Fatalf("\nFailed to populate wallet contents: %v", err)
		}
	} else {
		log.Println("============ Sucessfully populated wallet ============")
	}

	ccpPath := filepath.Join(
		"..",
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"connection-org1.yaml",
	)

	gw, err := gateway.Connect(
		gateway.WithConfig(config.FromFile(filepath.Clean(ccpPath))),
		gateway.WithIdentity(wallet, appUser),
	)
	if err != nil {
		log.Fatalf("\nFailed to connect to gateway: %v", err)
	}
	defer gw.Close()

	network, err := gw.GetNetwork(channelName)
	if err != nil {
		log.Fatalf("\nFailed to get network: %v", err)
	}

	contract := network.GetContract(chaincodeName)

	log.Println("--> Submit Transaction: InitLedger, function creates the initial set of medical supply on the ledger")
	result, initerr := contract.SubmitTransaction("InitLedger")
	if initerr != nil {
		log.Fatalf("Failed to Submit transaction: %v", initerr)
	}
	log.Println("Transaction succesfully submitted: \n", string(result))

	log.Println("--> Submit Transaction: Issue, function sends issue for medicine.")
	result, issueerr := contract.SubmitTransaction("Issue", "Aspirin", "00012", "Pain management", "2022.05.09", "$10")
	if issueerr != nil {
		log.Fatalf("\nFailed to Submit transaction: %v", issueerr)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: Request, function sends request for medicine.")
	result, requesterr := contract.SubmitTransaction("Request", "Aspirin", "00001", "Alice")
	if requesterr != nil {
		log.Fatalf("\nFailed to Submit transaction: %v", requesterr)
	}
	log.Println(string(result))

	log.Println("--> Submit Transaction: CheckHistory, function shows history.")
	result, historyerr := contract.SubmitTransaction("CheckHistory", "Alice")
	if historyerr != nil {
		log.Fatalf("\nFailed to Submit transaction: %v", historyerr)
	}
	if len(result) > 0 {
		log.Println(string(result))
	} else {
		log.Println("Ledger has no transaction history.")
	}
	log.Println("--> Submit Transaction: Approve, function that approves medicine and changes its state.")
	result, approveerr := contract.SubmitTransaction("Approve", "Aspirin", "00001")
	if approveerr != nil {
		log.Fatalf("\nFailed to Submit transaction: %v", approveerr)
	}
	log.Println(string(result))

	log.Println("\n============ Application ends ============")
}

func populateWallet(wallet *gateway.Wallet) error {
	credPath := filepath.Join(
		"..",
		"..",
		"..",
		"test-network",
		"organizations",
		"peerOrganizations",
		"org1.example.com",
		"users",
		"User1@org1.example.com",
		"msp",
	)

	certPath := filepath.Join(credPath, "signcerts", "cert.pem")
	// read the certificate pem
	cert, err := ioutil.ReadFile(filepath.Clean(certPath))
	if err != nil {
		return err
	}

	keyDir := filepath.Join(credPath, "keystore")
	// there's a single file in this dir containing the private key
	files, err := ioutil.ReadDir(keyDir)
	if err != nil {
		return err
	}
	if len(files) != 1 {
		return fmt.Errorf("keystore folder should have contain one file")
	}
	keyPath := filepath.Join(keyDir, files[0].Name())
	key, err := ioutil.ReadFile(filepath.Clean(keyPath))
	if err != nil {
		return err
	}

	identity := gateway.NewX509Identity(mspID, string(cert), string(key))

	return wallet.Put(appUser, identity)
}
