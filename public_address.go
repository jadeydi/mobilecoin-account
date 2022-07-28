package api

import (
	"encoding/binary"
	"encoding/hex"
	"hash/crc32"

	"github.com/btcsuite/btcutil/base58"
	"github.com/jadeydi/mobilecoin-account/types"
	"google.golang.org/protobuf/proto"
)

type PublicAddress struct {
	ViewPublicKey   string `json:"view_public_key"`
	SpendPublicKey  string `json:"spend_public_key"`
	FogReportUrl    string `json:"fog_report_url"`
	FogReportId     string `json:"fog_report_id"`
	FogAuthoritySig string `json:"fog_authority_sig"`
}

func (addr *PublicAddress) B58Code() (string, error) {
	view, err := hex.DecodeString(addr.ViewPublicKey)
	if err != nil {
		return "", err
	}
	spend, err := hex.DecodeString(addr.SpendPublicKey)
	if err != nil {
		return "", err
	}
	sig, err := hex.DecodeString(addr.FogAuthoritySig)
	if err != nil {
		return "", err
	}
	address := &types.PublicAddress{
		ViewPublicKey:   &types.CompressedRistretto{Data: view},
		SpendPublicKey:  &types.CompressedRistretto{Data: spend},
		FogReportUrl:    addr.FogReportUrl,
		FogReportId:     addr.FogReportId,
		FogAuthoritySig: sig,
	}
	wrapper := &types.PrintableWrapper_PublicAddress{PublicAddress: address}
	data, err := proto.Marshal(&types.PrintableWrapper{Wrapper: wrapper})
	if err != nil {
		return "", err
	}

	bytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytes, crc32.ChecksumIEEE(data))
	bytes = append(bytes, data...)
	return base58.Encode(bytes), nil
}
