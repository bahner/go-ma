package key

import (
	"github.com/multiformats/go-multibase"
)

// Im not a big fan of the enormmous RSA keys,
// byt they are battle tested AND allow for
// encryption directly with the public key.

// This way we can encrypt messages directly, and
// avoid publication of the public key. This drastically
// improves the resilience of the system.

// Long generation times be damned.
// Do not multiply objects beyond necessity.
//   - That Occam guy

// What we could do is multicodec the key type.
// Then we can change it without anyone noticing.

// What could possibly go wrong?
const PRIVATE_KEY_BITS = 2048
const PRIVATE_KEY_ENCODING = multibase.Base58BTC
