package DecentralizedABE

import "github.com/Nik-U/pbc"

type Authority interface {
	GetPK() *pbc.Element
	GetAPKMap() map[string]*APK
}
