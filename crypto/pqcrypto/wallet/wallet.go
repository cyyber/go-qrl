package wallet

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	cryptomldsa87 "github.com/theQRL/go-qrllib/crypto/ml_dsa_87"
	walletcommon "github.com/theQRL/go-qrllib/wallet/common"
	"github.com/theQRL/go-qrllib/wallet/common/descriptor"
	"github.com/theQRL/go-qrllib/wallet/common/wallettype"
	walletmldsa87 "github.com/theQRL/go-qrllib/wallet/ml_dsa_87"
	"github.com/theQRL/go-zond/common"
)

const (
	SeedSizeBytes = walletcommon.ExtendedSeedSize

	ML_DSA_87 wallettype.WalletType = wallettype.ML_DSA_87
)

var ErrBadWalletType = errors.New("unsupported wallet type")

// NOTE(rgeraldes24): work-in-progress
type Wallet interface {
	GetExtendedSeed() walletcommon.ExtendedSeed
	GetAddress() [common.AddressLength]uint8
	GetDescriptor() walletmldsa87.Descriptor
	GetPK() walletmldsa87.PK
	Sign([]uint8) ([cryptomldsa87.CryptoBytes]uint8, error)
}

func Generate(t wallettype.WalletType) (Wallet, error) {
	var (
		w   Wallet
		err error
	)
	switch t {
	case ML_DSA_87:
		w, err = walletmldsa87.NewWallet()
	default:
		return nil, ErrBadWalletType
	}
	if err != nil {
		return nil, err
	}

	return w, nil
}

func RestoreFromSeedBytes(seed []byte) (Wallet, error) {
	ext, err := walletcommon.NewExtendedSeedFromBytes(seed)
	if err != nil {
		return nil, err
	}
	return restoreWalletFromExtendedSeed(ext)
}

func RestoreFromSeedHex(seed string) (Wallet, error) {
	// NOTE(rgeraldes24): NewExtendedSeedFromHexString does not support 0x prefix
	seedBytes := common.FromHex(seed)
	return RestoreFromSeedBytes(seedBytes)
}

func restoreWalletFromExtendedSeed(ext walletcommon.ExtendedSeed) (Wallet, error) {
	var (
		wallet Wallet
		err    error
	)
	desc := descriptor.New(ext.GetDescriptorBytes())
	switch desc.Type() {
	case byte(wallettype.ML_DSA_87):
		wallet, err = walletmldsa87.NewWalletFromSeed(ext.GetSeed())
		if err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("unsupported wallet type in descriptor: %v", desc.Type())
	}

	return wallet, nil
}

func RestoreFromFile(file string) (Wallet, error) {
	seed, err := ReadSeedFromFile(file)
	if err != nil {
		return nil, err
	}

	w, err := RestoreFromSeedHex(seed)
	if err != nil {
		return nil, err
	}

	return w, nil
}

func ReadSeedFromFile(file string) (string, error) {
	fd, err := os.Open(file)
	if err != nil {
		return "", err
	}
	defer fd.Close()

	r := bufio.NewReader(fd)
	buf := make([]byte, SeedSizeBytes*2)
	n, err := common.ReadASCII(buf, r)
	if err != nil {
		return "", err
	} else if n != len(buf) {
		return "", fmt.Errorf("seed too short, want %d hex characters", SeedSizeBytes*2)
	}
	if err := common.CheckKeyFileEnd(r); err != nil {
		return "", err
	}

	return string(buf), nil
}
