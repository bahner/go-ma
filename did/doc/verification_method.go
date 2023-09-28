package doc

import (
	"fmt"

	"github.com/bahner/go-ma/did/vm"
)

func (d *Document) VerificationMethodsOfType(keyType string) ([]vm.VerificationMethod, error) {

	var methods []vm.VerificationMethod

	// if internal.Contains(vm.ValidVerificationMethods(), keyType) == false {
	if !vm.IsValidVerificationMethod(keyType) {
		return methods, fmt.Errorf("doc: Verification method type %s not allowed", keyType)
	}

	for _, method := range d.VerificationMethod {
		if method.Type == keyType {
			methods = append(methods, method)
		}
	}

	return methods, nil
}

func (d *Document) AddVerificationMethod(method vm.VerificationMethod) error {

	if !vm.IsValidVerificationMethod(method.Type) {
		return fmt.Errorf("doc: Verification method type %s not allowed", method.Type)
	}

	d.VerificationMethod = append(d.VerificationMethod, method)

	return nil
}
