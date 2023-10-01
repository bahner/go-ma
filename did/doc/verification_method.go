package doc

import (
	"github.com/bahner/go-ma/did/vm"
)

func (d *Document) AddVerificationMethod(method vm.VerificationMethod) error {

	d.VerificationMethod = append(d.VerificationMethod, method)

	return nil
}
