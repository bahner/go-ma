package did

import "github.com/ipfs/boxo/ipns"

// Returns the ipns.IPNSName of the DID
// Not the error, as verification is done in the constructor
func (d *DID) IPNSName() ipns.Name {
	name, _ := ipns.NameFromString(d.IPNS)

	return name
}
