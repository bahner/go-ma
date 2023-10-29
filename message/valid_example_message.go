package message

import "github.com/bahner/go-ma"

// This requires a bit of foo to create these values. So we'll just hardcode them here.
// So we only have one single source of truth.

// It's really meant for testing, but we might want to use it in other places too.
const (
	// This is the packed message. You can generate it using cmd/pack_valid_message.go
	Message_test_packed_valid_message = "zPnTcfHPut9RzxmUSYHE2SnTnpq1TMLL9LZMBQusnjfCjp8SU7DnMrfNcqf9f5tTTD37L2L5ZJCQ51QrJEYMVghmA8UnXtJ1jwabYo3on9UwRtmo2fFWJ5aH3fVnpYVEDbieRgbb2UKnhxfcAiMmoEGVBeYGrAjgg5crywtSstKFuVLHVGAUxu5BhM3P9wDSgYnJ9CNrgv87SR7RFYSwaHMUjF4q7fFNet2jNHtFqVxcthwLU39dPVLayx4wAoCEXurhkYqoNzD3NoJ99Z7FvLoemiLAwUy3JmWvTVr9QwBvWu4pqSq2deKsLXJ7BhPVjAbLabnCS8xwaNh6FNaEsbxssDgJMhP45VkWN9BQ9qis8wUwuCtUdM9ieXmq5E7aJszmCA7DjXCxc1tqiCD6djr2H3LwCrqHxMwdMRvbSY1jFobxAy6vjrjDc9xLHjZmT2S8gKnYCHhw9Pc49arxgwJR5kwDsXrrMKJW27dcjRXdBCLVUU31GXTEFPE65Dfa2crv32XE"
)

func ValidExampleMessage() *Message {

	msg := &Message{
		ID:           "CT6EklGVDpQpaYrth_O80",
		MimeType:     ma.MESSAGE_MIME_TYPE,
		From:         "did:ma:k51qzi5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd",
		To:           "did:ma:k51qzi5uqu5dk3pkcowsu2jqmnby0ry551xud502v000dzftwf4bj68384j84l",
		Created:      "2023-01-01T01:01:01Z",
		Expires:      "2023-01-02T01:01:01Z",
		BodyMimeType: "text/plain",
		Body:         "Share and Enjoy!",
		Version:      ma.VERSION,
		Signature:    "",
	}

	return msg

}
