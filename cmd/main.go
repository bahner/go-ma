package main

import (
	"fmt"
	"os"

	"github.com/bahner/go-ma"
	"github.com/bahner/go-ma/did/doc"
	"github.com/bahner/go-ma/entity"
	"github.com/bahner/go-ma/message"
	"github.com/bahner/go-ma/message/envelope"
	log "github.com/sirupsen/logrus"
)

// const subEthaMessage = "Share and enjoy!"

func main() {

	bahnerKeyset := "z3FZbnDdLWN9fJvc6Cgce2D9P3nTdnLJzu7AudrZN1DFfjzXVvRUfURiHAhfGW7Ee46G93cUe3zF7yaT8hmJbuaRVdWZhgGtHchuUQs25MjhgcfDBqURwiABbVf6ETTn6PbcZyTUkMitvzQQbDuJrbVN8vkWkqp7EVtiJMEpSdifSvNJ3X6LTw7N1aWXf9dRaeBvDUtmSk6WxJRpSNiUWsW3xnZFwXdkq6hBNxeRJSA2GAd81m9SFBgy3A1qzZ9RGZvnUGgX9JUxc9UcNyBasU5tcXH7HCUS8aXXN33CpCsEi9aUwTMuyTq8fchUpAWQWsDHhgi54YWZAkfpSY7Lei3qc4cyoLm1mDc5yPrEEawFe8GA3XjCxCm6iAbbVhQvXdQzDP95VLa7Jze3rYNJJxYc6TCDY8bei6wXtu9wycSL57arnTE7zEwE4QT7rpWAqEz25ymtVhbDLABrq16byrWta9Y2MsR4BfVTfEHR7S32k7ZFvRM4f3DLp2wfwutUz7bxLRrTFzQs8Rkmvgys1VvzSBygUUmh3mKvaPikifiAq9aqBYgcTWC8UxosMgWm59NwwYoFhQ9GuCeySS6fvrsiiTxUoiypiAWi2gLdCfUGsjUyk7rK6ZsiKxQeqN65LJfQMUbYv1a2c65wUd3TYH2jXkHXbfTMpstTdTugFmf8m5JdL8HNZY7xVrVagZLbrAPTHW1uD3WgF2cdjGq9QwvHaa6TkW1DexWdC6AZ9YwCPCnRzijcg3UigdTm5ngNj2YEFerjYNkCKyhpJSpwjfJLFvDCMrUUVASiJEqQ5rWRL1jXVMWgFKygcKe6FQDMqeRnZbm95aWz2meLTuNtnU28KtRYfeLMQdKHBizpoweGGp7dSsqQsNXSehZ5PcV71p4Yhp2Qn7Rm4c9FU248ZrsGeVUxf1sChb6bQuLXbzdAirD9gMwGevT3wmS4zwgo7x5zYDezDfUFd9GsdgxxbuzvyUxAeQV6igL9MHDnQzY1XvogteZ1YGF7AzabV6j1fUZz96H54bhvozmv5yD4p3fDersgFcxYxsRTs7gucJwzkCq"
	jobKeyset := "z2X4qaUJR6muxsSVZZ5x4jk55PBGdocTLtuGGEQpybHzv1qGwWrdFzPXkg3ZJgHiYXqLjWU7Qdhbea44jC6jxM8TaZXpTzet64wLd2jegUutS4mmAAMGp2uP6nsP8EEZ83CZ46BZyLHdp1UgpXA8cy1Uf5U1Hy7VGSANSVfLmxZYpEXpxSeGKzn2M2zJ7q7tSzVA6kUqZaGdQJngwse97Py9Y2K6j35BWYeuK2CrsUG6Q1ckmJc23gZaZu87ri8anMYvDDXbs4jo1JWwrCZkqBXo2dmCy2utCmPhTnhpYwt3GeTNq9g96sk7LGWNJJFMR7mcwAviq3SvB7KLpcdfMPfLtkufRSbSJGj6NdvNNcSZ1mtPsW2vvAsbMx1GFiKzWnNcM8sv8VH1GTsf2A3zRoAxehCc1j1MkHSBjckpAWYDq1LbJ48ZRwJAsBpaMYiJJDQ4ur7pASzqTaH6mFRL9SAMiRQpWk2KtyzT6C5YcyGFQCTnHCNyjWL9EcHLZ5N1G4e34ZUNLEnGbfea8GWFb1hFe9fqEi9aXdHbQpVtXt2czE3iaNi3ZYGCxxLxx2RMkJFkFTHybJ1vaRjbcwW3PSLraETjupCSu4BVVxY5xtQJ6AiVxdgyNU8YLwPqzFQ6tLzNXjvQzwqvzeyky8N38V7Cz6KVpPsp1z5TB5MJXouPNq9QWGuWo5xE4JNNNtboJz2f4e6F46kptgzDfdaiXVwgNfMQBedHskxEZguR5F2WhrvhtZyXU7LwNSMwFe1aEDrRwoX7Z7CRsd4aeLE2KnuAAASmP3gBEAiP9kskaWeFYz5Mu4fNXYA8X2iRdpSy9nYDP53YdTnaMWUKaYZ6cLeJFy1qNuiNuTDGbwpbAyiEziwvBzSLJY7RqSDpwnR5EJVpxLxBK7BqBcAjUgP7TxR5Pk4k2w7JYofBZrHa69WLtFbKZBnk5SSnFRJ2T8D6gdFzYcGkuKpgdWAWJqjLtFc588hzMQhRb2fLaQSp9QF7CKG2vTJXip476DPQyCN12wYu9wwDNk4AvNzf9KxcDUsZhw2KLrQ8WQq7fYe4vFB"
	log.SetLevel(log.DebugLevel)

	os.Setenv("IPFS_API_SOCKET", "localhost:45005")

	// shell := internal.GetShell()

	// Create a new person, object - an entity
	// id, _ := nanoid.New()
	bahner, err := entity.NewFromPackedKeyset(bahnerKeyset)
	// bahner, err := createSubjectFromPackedKeyset(bahnerKeyset)
	if err != nil {
		fmt.Printf("Error creating new identity in ma: %v\n", err)
	}
	job, err := entity.NewFromPackedKeyset(jobKeyset)
	// job, err := createSubjectFromPackedKeyset(jobKeyset)
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
		ma.KEY_AGREEMENT_KEY_TYPE,
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
		ma.VERIFICATION_KEY_TYPE,
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

	_, err = DIDDoc.Publish()
	if err != nil {
		return nil, fmt.Errorf("error publishing new identity in ma: %v", err)
	}

	return subject, nil
}
