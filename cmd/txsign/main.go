package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"

	"github.com/theQRL/go-zond/common"
	"github.com/theQRL/go-zond/core/types"
	"github.com/theQRL/go-zond/crypto/pqcrypto/wallet"
)

func parseUint64(name, s string) (uint64, error) {
	v, err := strconv.ParseUint(s, 0, 64)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", name, err)
	}
	return v, nil
}

func parseBig(name, s string) (*big.Int, error) {
	if s == "" {
		return big.NewInt(0), nil
	}
	v, ok := new(big.Int).SetString(s, 0)
	if !ok {
		return nil, fmt.Errorf("%s: invalid integer %q", name, s)
	}
	return v, nil
}

func parseTo(to string) (*common.Address, error) {
	if to == "" {
		return nil, nil
	}
	if strings.HasPrefix(to, "Q") {
		addr, err := common.NewAddressFromString(to)
		if err != nil {
			return nil, err
		}
		return &addr, nil
	}
	if strings.HasPrefix(to, "0x") {
		b := common.FromHex(to)
		if len(b) != common.AddressLength {
			return nil, fmt.Errorf("to: expected 20 bytes, got %d", len(b))
		}
		addr := common.BytesToAddress(b)
		return &addr, nil
	}
	return nil, fmt.Errorf("to: expected Q-address or 0x-hex")
}

func main() {
	var (
		seed       = flag.String("seed", "", "wallet seed hex (required)")
		chainIDStr = flag.String("chainid", "1", "chain id")
		nonceStr   = flag.String("nonce", "0", "tx nonce")
		gasStr     = flag.String("gas", "21000", "gas limit")
		maxFeeStr  = flag.String("maxfee", "0", "maxFeePerGas")
		maxTipStr  = flag.String("maxtip", "0", "maxPriorityFeePerGas")
		toStr      = flag.String("to", "", "recipient (Q... or 0x...); empty for contract creation")
		valueStr   = flag.String("value", "0", "value")
		dataStr    = flag.String("data", "0x", "hex data payload")
		printFrom  = flag.Bool("print-from", false, "print derived sender address to stderr")
	)
	flag.Parse()

	if *seed == "" {
		fmt.Fprintln(os.Stderr, "-seed is required")
		os.Exit(2)
	}

	chainID, err := parseBig("chainid", *chainIDStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	nonce, err := parseUint64("nonce", *nonceStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	gas, err := parseUint64("gas", *gasStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	maxFee, err := parseBig("maxfee", *maxFeeStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	maxTip, err := parseBig("maxtip", *maxTipStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	to, err := parseTo(*toStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	value, err := parseBig("value", *valueStr)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	data := common.FromHex(*dataStr)

	w, err := wallet.RestoreFromSeedHex(*seed)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	txdata := &types.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasTipCap: maxTip,
		GasFeeCap: maxFee,
		Gas:       gas,
		To:        to,
		Value:     value,
		Data:      data,
	}

	signer := types.NewShanghaiSigner(chainID)
	signed, err := types.SignNewTx(w, signer, txdata)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	if *printFrom {
		from, err := signer.Sender(signed)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(2)
		}
		fmt.Fprintln(os.Stderr, "from:", from.Hex())
	}

	blob, err := signed.MarshalBinary()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
	fmt.Printf("0x%s\n", common.Bytes2Hex(blob))
}
