package DecentralizedABE

import "github.com/Nik-U/pbc"

type Cipher struct {
	C0         *pbc.Element   `field:"2"`
	C1s        []*pbc.Element `field:"2"`
	C2s        []*pbc.Element `field:"0"`
	C3s        []*pbc.Element `field:"0"`
	CipherText []byte
	Policy     string
}
