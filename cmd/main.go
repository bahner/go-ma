package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma/entity"
	"github.com/bahner/go-ma/message"
	"github.com/bahner/go-ma/message/envelope"
	log "github.com/sirupsen/logrus"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	bahnerKeyset := "z3FZbnDdLWN9fJvc6Cgce2D9P3nTdnLJzu7AudrZN1DFfjzXVvRUfURiHAhfGW7Ee46G93cUe3zF7yaT8hmJbuaRVdWZhgGtHchuUQs25MjhgcfDBqURwiABbVf6ETTn6PbcZyTUkMitvzQQbDuJrbVN8vkWkqp7EVtiJMEqcx9MRwHRpjsHBTQh6J1UqfsS1vAmiwZHVM8EpEfu5caW251tAZRCNxeGACjxNm7riB5WKbsDffPsLETD479MnRorphG4jJF5U5rejn4Tv7739bNG8FaZeraqFTsnjs4tNRLe3v55GiEaU46SuGu4gvdJuEHedyLggq5gtz8kqZJGutrT3H6RRWonrUJVUimd4Sf96P4yBTYgDQ64EoZMjJjdqkgsKPk1qnxG9gpNPzaaGTevh7W8TeA2izdKwhcynJmrsToRmE6vBGhFXruVgibA3brHdx93YJTNQG6Q3hdCpTrPuFEQfqE88MJNsYro2vuAmLxVHuuaLRKgKHpTjiBw4ug9kQ1aMr5UR82TL46GpJ3ie8KNiuXy1iAj7Am4asbiGL3UG98pJhGAFiy9vdAVqWASQsA555i1uGN112aqcJS6qe7dQmYBsXXoJRhW5jyTsWFjjNVqo1xhGF9pv2MWmHQ3JJF7oTcdUEgfZpjchd8XkMs9TXqyXvd2Yimfm1hywgPgtXJVZiujn9EtZGbiCwQkyJKbzVWyAR51jYJ6GZSF1LNjvKuMQaNAvxxuYidhmrkjHJJxaadahnZ9jr1dfjh3hJV7RUmG8mAxhYAhKYMH1s3XtW346ajW125cWToWNhssXhKnBttJfA1rKokWgTpxKNYquXx7U1THYiLeTmA3EFxxF4cT7wQ1ZrVgLKUbTr4JR35hKMobn6WV6FJ5uLQJ5FHr132pe4UFs81TmXX1K2uG12u1USKByLuu11EBLQmKuj4k3HRePZ268oktFJRTiRtbas7yeeAXJA2p8VBo49Xghr7YdKLdjdXPc5dFihDYRde3MsMijDiRs9C3pcTvt2chDWryRxv3nqiJV2hBTDQALPs7h6Qs6UFNSDeJkcp"
	jobKeyset := "z2X4qaUJR6muxsSVZZ5x4jk55PBGdocTLtuGGEQpybHzv1qGwWrdFzPXkg3ZJgHiYXqLjWU7Qdhbea44jC6jxM8TaZXpTzet64wLd2jegUutS4mmAAMGp2uP6nsP8EEZ83CZ46BZyLHdp1UgpXA8cy1Uf5U1Hy7VGSAMZM7iHbnv1dZq93k1ZnThkJ9uNfib9ruR5WDr1maYoZPX4kyVYAcazkdCgmsPm7Yi7pXQpqPhnXRKeDj99C1o2j7P5Jt7Qsv6oS9WYHUQ78PT1PU6eZHTmiNaAUqXmRfQYkat75SeLmaaRTQsmjV5mwxPrHbChQt5bAsk4tnk7mMwv3NG196hzSJehYGkuZnWodAbbHUMKG86AcCC3TQHWAu8eag6UkiAVZzUN5e88uRj8Ro55E7kANT8ZbUneeQzGcX9ewsTVuTxxyLAL5ytHdFWKFGowEr27tqG7v8cXZXfKaA4t8TaNVH1DbBf9TDD9gVf6NLip8pTcPnvn7fuP41iMdzdFwqewCzT18iL2H8oh9GJMa3iRwANNjHMKWSYd3LiMBXVx9GTwRdCEKFTQNXWotTYUmrthbNWZbRnPxMQhEVGBP6uz8TY4AGxhohQUBVNaCFCpVjzm7nGpPD5n3C9f7orNWUnvCtr4ucsMCA9BBjzSsqos4AmPpG7a2h5EhLgsTyj4oXw5vz8PbGyXSWGWgn5qCzYY6oHzhkk56fvfdUTdgmoDkdwdgqPtZYJbuDSfvfLSgVVFTrx93oyqarn4Guy4ecbAZM42DAyatQqKrSzERd6WRayAbVkNTMUyqv3dDePMswgsHLE8HUc8WfdAEM7P7K8HHAQC6n61sGcN8KreeWmPj9sBNMK4xbfYYfiC1DdFVXSL13fUjucS7cqdce7Bo2yuMPxwaS5E1j67tgNQztERfk2MwnHZDrj4RJciDnWtxxG1ZE8Gtn2DFZquF4X5NzgoFccDZSthgKNEmBfiqjMpteaWDkorpGhC5dQFx7QFUWkxHUZoHaLSU7trUiHqsUHyT4u6uVNY4XsiBeY17VEKT31wB7kHUngDdYqfen"
	log.SetLevel(log.DebugLevel)

	os.Setenv("IPFS_API_SOCKET", "localhost:45005")

	// shell := internal.GetShell()

	// Create a new person, object - an entity
	// id, _ := nanoid.New()
	bahner, err := entity.NewFromPackedKeyset(bahnerKeyset)
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}
	job, err := entity.NewFromPackedKeyset(jobKeyset)
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}

	msgBody := "Share and enjoy!"
	msgMimeType := "text/plain"

	myMessage, err := message.New(
		bahner.DID.String(),
		job.DID.String(),
		msgBody,
		msgMimeType)
	if err != nil {
		fmt.Printf("Error creating new message: %v\n", err)
	}

	fmt.Println(myMessage)

	msgEnvelope, err := envelope.Seal(myMessage)
	if err != nil {
		fmt.Printf("Error creating new envelope: %v\n", err)
	}
	messageJSON, err := msgEnvelope.MarshalToJSON()
	if err != nil {
		fmt.Printf("Error marshalling message to JSON: %v\n", err)
	}

	fmt.Println(string(messageJSON))

	letter, err := msgEnvelope.Open(job.Keyset.EncryptionKey.PrivKey)
	if err != nil {
		fmt.Printf("Error unsealing envelope: %v\n", err)
	}

	fmt.Println(letter)

}

// func createSubjectFromPackedKeyset(keyset string) (*entity.Entity, error) {
// 	// Create a new person, object - an entity
// 	// id, _ := nanoid.New()

// 	subject, err := entity.NewFromPackedKeyset(keyset)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating new identity in ma: %v", err)
// 	}
// 	log.Debugf("Created new entity: %s", subject.DID.String())
// 	DIDDoc, err := doc.New(subject.DID.String(), subject.DID.String())
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating new identity in ma: %v", err)
// 	}

// 	encVM, err := doc.NewVerificationMethod(
// 		subject.DID.Identifier,
// 		subject.DID.String(),
// 		ma.KEY_AGREEMENT_KEY_TYPE,
// 		subject.Keyset.EncryptionKey.PublicKeyMultibase)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating new verification method: %v", err)
// 	}
// 	err = DIDDoc.AddVerificationMethod(encVM)
// 	if err != nil {
// 		return nil, fmt.Errorf("error adding new verification method: %v", err)
// 	}
// 	DIDDoc.KeyAgreement = encVM.ID
// 	log.Debugf("Added keyAgreement verification method: %s", DIDDoc.KeyAgreement)

// 	signvm, err := doc.NewVerificationMethod(
// 		subject.DID.Identifier,
// 		subject.DID.String(),
// 		ma.VERIFICATION_KEY_TYPE,
// 		subject.Keyset.SigningKey.PublicKeyMultibase)
// 	if err != nil {
// 		return nil, fmt.Errorf("error creating new verification method: %v", err)
// 	}
// 	err = DIDDoc.AddVerificationMethod(signvm)
// 	if err != nil {
// 		return nil, fmt.Errorf("error adding new verification method: %v", err)
// 	}
// 	DIDDoc.AssertionMethod = signvm.ID
// 	log.Debugf("Created new assertion method: %s", DIDDoc.AssertionMethod)

// 	err = DIDDoc.Sign(subject.Keyset.SigningKey, signvm)
// 	if err != nil {
// 		return nil, fmt.Errorf("error signing new identity in ma: %v", err)
// 	}

// 	_, err = DIDDoc.Publish()
// 	if err != nil {
// 		return nil, fmt.Errorf("error publishing new identity in ma: %v", err)
// 	}

// 	return subject, nil
// }
