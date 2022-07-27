package api

import (
	"encoding/hex"
	"log"
	"testing"

	"github.com/jadeydi/mobilecoin-account/block"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/proto"
)

func TestAccount(t *testing.T) {
	assert := assert.New(t)
	account, err := DecodeAccount("d9V5WDNZxa7fNRw24JwaqDpCrRnKjpsGcN2CpLJNJZQjt7Vhwsmm2w2CoY4g2u7vy5HFxL5S8uGUUogApWteXdgrd3GnowaFUWU1HefMdGK7fdhWEcznR9SacnddL2KA3NmEzgDRoqqwHzTv8cPXo5udtRAy4Q4xYsPTmXkcZTH212SNXudQwA6KfwUqS3aKvJFMLcr1iUmfupMikwVYcfboJ6i3gejGhua5BVX1GRhL2BRWMHhnRCThqicQAy")
	assert.Nil(err)
	address, err := account.B58Code()
	assert.Nil(err)
	log.Println(address)
	assert.Equal(address, "d9V5WDNZxa7fNRw24JwaqDpCrRnKjpsGcN2CpLJNJZQjt7Vhwsmm2w2CoY4g2u7vy5HFxL5S8uGUUogApWteXdgrd3GnowaFUWU1HefMdGK7fdhWEcznR9SacnddL2KA3NmEzgDRoqqwHzTv8cPXo5udtRAy4Q4xYsPTmXkcZTH212SNXudQwA6KfwUqS3aKvJFMLcr1iUmfupMikwVYcfboJ6i3gejGhua5BVX1GRhL2BRWMHhnRCThqicQAy")

	addr := "G57w8Br44AYd6aEKfagTyLFvt4tTLhDdzGsX6PbYwfumwpjc1htSpWfoey2FLYNKMJA28q8YyqYb83dh66A7BTVA4XNZzXsNNUDv1nTmaw"
	account, err = DecodeAccount(addr)
	address, err = account.B58Code()
	assert.Nil(err)
	log.Println(address)
	assert.Equal("G57w8Br44AYd6aEKfagTyLFvt4tTLhDdzGsX6PbYwfumwpjc1htSpWfoey2FLYNKMJA28q8YyqYb83dh66A7BTVA4XNZzXsNNUDv1nTmaw", address)

	account, err = DecodeAccount("3DsoPZDzxAtCCvwuNKvmSQcqHaa6w99U2dxjroC8DXZTvMM2LJdUHDXXDVQEf3vqew5Xzc3nw5jKEUMUuqvNREfV5GVCyiEvusZUFJeeRt3ykVEfAnp6suXZe15R9vvNKMBf2sTR2Bk6PiRrwyNNFFchY1kYkdCsH4dAsGJmFpRB4kGkhXANxFbK7r2K1FTPGpPxav31XXkCaXB3KcTVEEWptGNnEYrDKJdvJvwuMdkzv8")
	assert.Nil(err)
	log.Printf("1:::%#v", account)

	account, err = DecodeAccount("CpJVvph9KVdpngREbnazkU1Z5XvmYNeDzQxhf9ft1agRZYCBk9gdfY489ibBAHPEtbFTmcrK4tdXa47bbsgffq21FRErBWAHje83S67XvcybFA5CVkwZghjm8GCWEZGyFqdaZiHB4B54A1GBnmeHGJra97oz7vQUGPyQNYZnrS3CY2Et6SRoAweDwemuukGKmTcBG1kaJKWRjRVVktzuBYbXPvSzx9edbk9UirbKRoZHyu")
	assert.Nil(err)
	log.Printf("2:: %#v", account)

	var pa block.PublicAddress
	buf, _ := hex.DecodeString("0a220a201c548b856c6b0400c2a99b2e918f267cca1bdee0367dc01258bd2ff9d8d2167412220a201e34eb6fc825788b7532943f1bef76554a8e2624e65190ec65572a29cd2e811c")
	err = proto.Unmarshal(buf, &pa)
	if err != nil {
		log.Println(err)
	}
	log.Printf(hex.EncodeToString(pa.ViewPublicKey.GetData()))
	log.Printf(hex.EncodeToString(pa.SpendPublicKey.GetData()))
}
