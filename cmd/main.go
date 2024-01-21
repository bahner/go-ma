package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/entity"
	"github.com/bahner/go-ma/internal"
	"github.com/bahner/go-ma/key"
	keyset "github.com/bahner/go-ma/key/set"
	log "github.com/sirupsen/logrus"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	// bahnerKeyset := "z4zt5ELTEWA2DthCxJ4FNiwdeXwtQdJRKLyBmeDcNYMpzPW4K66Dn1NZR2atpAdMUyUkxebKSodEozCJPo9jqw4xJpXtCusvLZS3QtJN3nsa74fPXpygjh37GqFMdhv7ba7BwJE9eZKutTy5UABcrj3t7UFpiKNh9pn6Xhov7tpebSTF3z8C2VNRUeyLfFvMmy1Dq2TfdqBC3WNWXcLuSmirut8gdkSQMwYaXgKkY6dExgFfRXghsqL8V2uETPcmiezQT9EitqRL1yNMKCkxUBGVzCxT1zsBhLNzFgw8PrX8xRBC9RxKnm56UZyPnJ7EEBiQQHNMqzHCobVQajLNx5Cm8Q3w84PRyRqsoY2S15cKofMZqFDk21KpWiD5WzPSPRxcE9ELddqMbjg4H1NhJofr2yn6hRfmik4MgENYb3GqhXb5UFG94wfZAxuefacMrisTPrUsfkdYFBVGpBFpEZR3oZT31J61nyRbPXQd7pnruTyZnKb71Umes7ZY1y4oZH2PTGiAK1ghmUzFv2yxg4sppVrLqY56AajRiT3yFKVfHk1R6MrdG9mgbbfk37UwqTmbmzpAdRbQRgWpYghEoN6Ko5fMKMxg6cf7bTKgndEhsj2zirZBTokMEgSD7hoqirJpAvirE98ERRcLJ1ZL2pGLGCeSWPJotfMEnLKd8jEJ7x9RmruCyrVpi87CFUmSUn5Nv8MfH5sxZLacGKXAGEGwZYJyAYM3tNfDf3oEfQ14nqXLN32TkkU5bf6Ma8PwdP9LGGxMTeBjcExzVXU3q981gGn38gZ3h4DtRm9mhPLpxUv3EbepFzxAxswtVkqVf22UUkgYhwzGmjMKGTMcdAB5S8DAVj5VyVpLv6XQ54LiABF1cvYjKNMiwytxVw7jj2MUgKwwuiFAGLKrUaQCLHGrYzBQSgf2ZjATQCbPChDCPJP7Hxc3pkt7KQ6QAwgCooLZHiKPEUgD5mSHrxDNxCvNv23SKvEnnudnimyLa7NYpSVfFo4ff6pv9puqxyG4qVEzhiku4pfaHCcnnnUR98Ni5BKEp1XdVtERt6ALy5RqULhxSg2oVeo4QDzGdezYhjNCmaFoY1bZDhyun"
	// jobKeyset := "z3hQcSxxrHPZSLq8okEmDfahHtVGFTd2d9eADfyPxoGBBDdhg4fWvaKksPkyFJwSVfRorPZz3ZPt5gvzKBeUVetjgoG2Hdv2FgeW19aQqQUoDsxZhqyAU6QmNw8f422GAwFJxKchQubLYSJiQxw6gyEMkdEwp16ZYhZS8bT82jufpWDmZWtWDXgUD9h9jwubec3ZNPk3H9Pvw4NQ7VRMZz4mLq6xAYoVkKWGDEzPfgv6pHTNQSbW8MW3uRhmGX9rRiWo3X1FA2kWYN4a3UnMXMQSYQXVFAve16aGWF7WYRnKcAXhFnA5JkvCciapyn44dgQz6GxoMBGHnHxsYuA6NoLDg3pcCdXNMzzsqAvX7g3FKks7YTSUNX4dExjTjZBbxsV7Wiuu9rB2B78fDhRd25cWMUFRhakQyLxmSfRiXK4qJJCrWR8to9DwJXbstimUK4j3BbWaAzywqKw8L2AwuYnMcsSvrywSMiveXQnmUhhzcgnLT9WwDgsLbj5sRQw1keUMAJbRooHQuVM3JamrePDc5wD48Bcg6DSLEvLK2rNanbKi7gtA23eJ6MSzNpE1qHAffLzfZ9spj2ec35BoCsq417CopqTAdVrvm4nfAR9XJEwSToM1UYNV9t7HtCfidxCnspJ6f1R332JmQE1mWrwXggk4ARhPG9CZoWvyPmNrTwVkdi8eHhsK226msXt4P9XjbNvTvh2cnghNYLdLYHRKWdoMxdsUadB1WtBXzUyq2QDG76vzUmKaips8L6p2SwiMKxjZAy5GAWSp6tbxBEqd2dKJjZgE6YKMrYvBpQUn3dB7FkYVXaFCL46dmCV23WAPb2j1uU6wAAQyTByP2JmXpKge4ARKmmfzo63jWE8mSbLE3Ehpzq16AeRSyDmi3P8scPWwrdunbob1G2P3HyqYCDVNnvqF71fxhEUeumQtefqVaKxR27KVsUwzjpxG9KPvnR4aT4zWzzDAXvexL8vPtcGYBP5LT5WYsKmbe3C7QFxZn5C65CzVxKRUMC6HzYXNpbW5efLURJHqtsy8wLozS8GPzCtMcTDaBMw81jP2wy1vTbtxPQmRxS6w47KPYctcNzJyxUpB3"
	log.SetLevel(log.DebugLevel)

	os.Setenv("IPFS_API_MULTIADDR", "/ip4/tcp/127.0.0.1/45005")

	bahner, err := keyset.GetOrCreate("bahner")
	if err != nil {
		fmt.Printf("error creating new keyset: %v\n", err)
	}

	bahnerPacked, err := bahner.Pack()
	if err != nil {
		fmt.Printf("error packing keyset: %v\n", err)
	}
	fmt.Println(bahnerPacked)

	foo, err := keyset.Unpack(bahnerPacked)
	if err != nil {
		fmt.Printf("error unpacking keyset: %v\n", err)
	}

	fmt.Println(foo.IPFSKey.Fragment)

}

func createSubjectFromPackedKeyset(keyset string) (*entity.Entity, error) {
	// Create a new person, object - an entity
	// id, _ := nanoid.New()

	subject, err := entity.NewFromPackedKeyset(keyset)
	if err != nil {
		return nil, fmt.Errorf("error creating new identity in ma: %v", err)
	}
	log.Debugf("Created new entity: %s", subject.DID.String())
	DIDDoc, err := doc.New(subject.DID.String(), subject.DID.String())
	if err != nil {
		return nil, fmt.Errorf("error creating new identity in ma: %v", err)
	}

	encVM, err := doc.NewVerificationMethod(
		subject.DID.Identifier,
		subject.DID.String(),
		key.KEY_AGREEMENT_KEY_TYPE,
		internal.GetDIDFragment(subject.Keyset.EncryptionKey.DID),
		subject.Keyset.EncryptionKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("error creating new verification method: %v", err)
	}
	err = DIDDoc.AddVerificationMethod(encVM)
	if err != nil {
		return nil, fmt.Errorf("error adding new verification method: %v", err)
	}
	DIDDoc.KeyAgreement = encVM.ID
	log.Debugf("Added keyAgreement verification method: %s", DIDDoc.KeyAgreement)

	signvm, err := doc.NewVerificationMethod(
		subject.DID.Identifier,
		subject.DID.String(),
		key.ASSERTION_METHOD_KEY_TYPE,
		internal.GetDIDFragment(subject.Keyset.SigningKey.DID),
		subject.Keyset.SigningKey.PublicKeyMultibase)
	if err != nil {
		return nil, fmt.Errorf("error creating new verification method: %v", err)
	}
	err = DIDDoc.AddVerificationMethod(signvm)
	if err != nil {
		return nil, fmt.Errorf("error adding new verification method: %v", err)
	}
	DIDDoc.AssertionMethod = signvm.ID
	log.Debugf("Created new assertion method: %s", DIDDoc.AssertionMethod)

	err = DIDDoc.Sign(subject.Keyset.SigningKey, signvm)
	if err != nil {
		return nil, fmt.Errorf("error signing new identity in ma: %v", err)
	}

	_, err = DIDDoc.Publish(false)
	if err != nil {
		return nil, fmt.Errorf("error publishing new identity in ma: %v", err)
	}

	return subject, nil
}
