package token

import "github.com/o1egl/paseto"

//PasetoMaker is a PASETO token maker
type PasetoMaker struct {
	paseto *paseto.V2
	symmetricKey []key
}