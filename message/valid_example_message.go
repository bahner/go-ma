package message

// This requires a bit of foo to create these values. So we'll just hardcode them here.
// So we only have one single source of truth.

// It's really meant for testing, but we might want to use it in other places too.
const (
	// This is the packed message. You can generate it using cmd/pack_valid_message.go
	Message_test_packed_valid_message = "zs5VGbMa1L2aB1CMGpuHTgFiKMse1ZcfdZsKbL6gNNmpPt2u8NEPiHAqD9UAUHMcQDsNaWJuKnKGS8KQ4zXH4nyrWmLEcbF8nquUh2kFHLD3pJyTkRYBu5QNUnF6CnhxH9F8KqHWLqVn1snmM9R9SkrnBftAfFhkMfTsHgwfTJukHDwodDAQeaVaFZ6MZDpva2rwrdBQMBkhgaoXhnSHHjA5sBTsEbocMVYAoqTv9eQak3KoQ4ZoxZE59bmhUjiMbkCmPCQaZsLNKKPuYyTUB6mZL4kQgBp3w1QoKHV63sgKcGuuLtcmJtKuBvhaefzq7FuSahysGKf4okzbx6J3rRRhZWNEj9HttjDbQVus7k5atVD7fZLHTKsag8qt4ukdo43JyzTKdVPbnXLmYjGqbS4xucjUeGo4KCt1KWe1BDAk3vbtQriA6BVj3Ety33U7UPnxFzesXhGqkwALcVS1cbCYEbYnCbe5tRgmtS2JhmBPqA7zdh1jAFYfxcmyRMw7yqN2p8tgJPZbHnsv"
)

func ValidExampleMessage() *Message {

	msg := &Message{
		ID:          "CT6EklGVDpQpaYrth_O80",
		Type:        MESSAGE_MIME_TYPE,
		From:        "did:space:k51qzi5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd",
		To:          "did:space:k51qzi5uqu5dk3pkcowsu2jqmnby0ry551xud502v000dzftwf4bj68384j84l",
		CreatedTime: "2023-01-01T01:01:01Z",
		ExpiresTime: "2023-01-02T01:01:01Z",
		MimeType:    "text/plain",
		Body:        "Share and Enjoy!",
		Version:     MESSAGE_VERSION,
		Signature:   "",
	}

	return msg

}
