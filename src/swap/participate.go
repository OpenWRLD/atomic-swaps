// Package swap provides various key functions to execute an atomic swap
package swap

import (
	"atomic-swaps/src/api"
	"bytes"
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/DiviProject/diviutil"
)

// Participate : Participates in an atomic swap
// Participating in an atomic swap requires a valid address and amount
// You also need to provide the secret from the original atomic swap contract
// You also need to be connected to an RPC client with a valid address and amount
func Participate(address string, amount float64, secretHash []byte, currency string, autopublish bool) (api.ParticipateResponse, error) {
	fmt.Println(address)
	fmt.Println(currency)

	network := RetrieveNetwork(currency)

	decodedAddress, err := diviutil.DecodeAddress(address, network)
	if err != nil {
		return api.ParticipateResponse{"", "", "", "", "", "", "", "", struct{}{}, nil, 51200}, err
	}

	if !decodedAddress.IsForNet(network) {
		return api.ParticipateResponse{"", "", "", "", "", "", "", "", struct{}{}, nil, 51200}, fmt.Errorf("This address is not for this currency")
	}

	p2pkh, ok := decodedAddress.(*diviutil.AddressPubKeyHash)
	if !ok {
		return api.ParticipateResponse{"", "", "", "", "", "", "", "", struct{}{}, nil, 51200}, fmt.Errorf("participant address is not P2PKH")
	}

	cryptoAmount, err := diviutil.NewAmount(amount)
	if err != nil {
		return api.ParticipateResponse{"", "", "", "", "", "", "", "", struct{}{}, nil, 51200}, err
	}

	if len(secretHash) != sha256.Size {
		return api.ParticipateResponse{"", "", "", "", "", "", "", "", struct{}{}, nil, 51200}, fmt.Errorf("secret hash has wrong size")
	}

	client := GetRPCClient(currency)

	defer func() {
		client.Shutdown()
		client.WaitForShutdown()
	}()

	locktime := time.Now().Add(24 * time.Hour).Unix()

	b, err := BuildContract(
		client,
		&ContractArgs{
			them:       p2pkh,
			amount:     cryptoAmount,
			locktime:   locktime,
			secretHash: secretHash,
		},
		currency,
	)
	if err != nil {
		return api.ParticipateResponse{"", "", "", "", "", "", "", "", struct{}{}, nil, 51200}, err
	}

	var contractBuf bytes.Buffer
	contractBuf.Grow(b.contractTx.SerializeSize())
	b.contractTx.Serialize(&contractBuf)

	refundTxHash := b.refundTx.TxHash()
	contractFeePerKb := CalcFeePerKb(b.contractFee, b.contractTx.SerializeSize())
	refundFeePerKb := CalcFeePerKb(b.refundFee, b.refundTx.SerializeSize())

	var refundBuf bytes.Buffer
	refundBuf.Grow(b.refundTx.SerializeSize())
	b.refundTx.Serialize(&refundBuf)

	if autopublish == false {
		fmt.Printf("Contract fee: %v (%0.8f BTC/kB)\n", b.contractFee, contractFeePerKb)
		fmt.Printf("Refund fee:   %v (%0.8f BTC/kB)\n\n", b.refundFee, refundFeePerKb)
		fmt.Printf("Contract (%v):\n", b.contractP2SH)
		fmt.Printf("%x\n\n", b.contract)

		fmt.Printf("Contract transaction (%v):\n", b.contractTxHash)
		fmt.Printf("%x\n\n", contractBuf.Bytes())

		fmt.Printf("Refund transaction (%v):\n", &refundTxHash)
		fmt.Printf("%x\n\n", refundBuf.Bytes())
	}

	PromptPublishTx(client, b.contractTx, "contract", autopublish)

	return api.ParticipateResponse{
		b.contractFee.String(),
		b.refundFee.String(),
		b.contractP2SH.String(),
		fmt.Sprintf("%x", b.contract),
		b.contractTxHash.String(),
		fmt.Sprintf("%x", contractBuf.Bytes()),
		refundTxHash.String(),
		fmt.Sprintf("%x", refundBuf.Bytes()),
		struct{}{},
		nil,
		51200,
	}, nil
}
