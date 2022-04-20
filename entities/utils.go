package entities

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/jackc/pgtype"
)

type VegaPublicKey string

func (pk *VegaPublicKey) Bytes() ([]byte, error) {
	strPK := pk.String()

	bytes, err := hex.DecodeString(strPK)
	if err != nil {
		return nil, fmt.Errorf("decoding '%v': %w", string(pk.String()), ErrInvalidID)
	}
	return bytes, nil
}

func (pk *VegaPublicKey) Error() error {
	_, err := pk.Bytes()
	return err
}

func (pk *VegaPublicKey) String() string {
	return string(*pk)
}

func (pk VegaPublicKey) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
	bytes, err := pk.Bytes()
	if err != nil {
		return buf, err
	}
	return append(buf, bytes...), nil
}

func (pk *VegaPublicKey) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	strPK := hex.EncodeToString(src)

	*pk = VegaPublicKey(strPK)
	return nil
}

type TendermintPublicKey string

func (pk *TendermintPublicKey) Bytes() ([]byte, error) {
	strPK := pk.String()

	bytes, err := base64.StdEncoding.DecodeString(strPK)
	if err != nil {
		return nil, fmt.Errorf("decoding '%v': %w", string(pk.String()), ErrInvalidID)
	}
	return bytes, nil
}

func (pk *TendermintPublicKey) Error() error {
	_, err := pk.Bytes()
	return err
}

func (pk *TendermintPublicKey) String() string {
	return string(*pk)
}

func (pk TendermintPublicKey) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
	bytes, err := pk.Bytes()
	if err != nil {
		return buf, err
	}
	return append(buf, bytes...), nil
}

func (pk *TendermintPublicKey) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	strPK := base64.StdEncoding.EncodeToString(src)

	*pk = TendermintPublicKey(strPK)
	return nil
}

type EthereumAddress string

func (addr *EthereumAddress) Bytes() ([]byte, error) {
	strAddr := addr.String()

	if strings.HasPrefix(strAddr, "0x") == false {
		return nil, fmt.Errorf("invalid '%v': %w", string(addr.String()), ErrInvalidID)
	}

	bytes, err := hex.DecodeString(strAddr[2:])
	if err != nil {
		return nil, fmt.Errorf("decoding '%v': %w", string(addr.String()), ErrInvalidID)
	}
	return bytes, nil
}

func (addr *EthereumAddress) Error() error {
	_, err := addr.Bytes()
	return err
}

func (addr *EthereumAddress) String() string {
	return string(*addr)
}

func (addr EthereumAddress) EncodeBinary(ci *pgtype.ConnInfo, buf []byte) ([]byte, error) {
	bytes, err := addr.Bytes()
	if err != nil {
		return buf, err
	}
	return append(buf, bytes...), nil
}

func (addr *EthereumAddress) DecodeBinary(ci *pgtype.ConnInfo, src []byte) error {
	strAddr := "0x" + hex.EncodeToString(src)

	*addr = EthereumAddress(strAddr)
	return nil
}
