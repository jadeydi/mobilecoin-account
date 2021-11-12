package api

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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
}
