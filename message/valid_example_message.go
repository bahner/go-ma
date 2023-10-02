package message

import "github.com/bahner/go-ma"

// This requires a bit of foo to create these values. So we'll just hardcode them here.
// So we only have one single source of truth.

// It's really meant for testing, but we might want to use it in other places too.
const (
	// This is the packed message. You can generate it using cmd/pack_valid_message.go
	Message_test_packed_valid_message = "zPnTcfHPut9RzxmUSYHE2SnTnpq1TMLL9LZMBQusnjfD5EHk5SRv7RHiewJc1kVJ6SXTFigWXQNMeovu9gDKYZ8SfQ1xn51cLyZvVtJhHoUk4YVebezEBwiW5xFpv6LVAVk7LxS6pxrRxhRGgciqCiDswe6PfnP4jr3R8jVc5T9Ad5WimrCq84NAC8ZEHk4Bts3ZtQMZhswZnXcVVHDfsyoLL4qobL6VT68YnWk1yhP1sTjHHdj3vTz79TxtU3T6WherGufZ7R9ELVTj25xSV9y3QfhXz875kp9zhW4mgpuJxCzjDuEddHX2ZAw3k4QjmTJAUBAW3Bv4WxLxqBdnWMbULLQhridfVJTdiTgYeic1mLEa3UwwFDCL6Cu99qrzWVaJNx8GmYXt42ronzuGko7Mm9pi8JrfC9EXyWNWemuxuXqQueVdg3JX4drBLAAk8WrcBo1UXdXKif2gNwpenmoGx8MMM1Dz9CCuuwfyrJ7ttwLvURawqtdC4WBmNpxGSGnjnDyN"
)

func ValidExampleMessage() *Message {

	msg := &Message{
		ID:           "CT6EklGVDpQpaYrth_O80",
		MimeType:     ma.MESSAGE_MIME_TYPE,
		From:         "did:ma:k51qzi5uqu5djy7ca9encml5bqicdz47khiww4dvcvso4iqg3z7xy0amwnwcwd",
		To:           "did:ma:k51qzi5uqu5dk3pkcowsu2jqmnby0ry551xud502v000dzftwf4bj68384j84l",
		CreatedTime:  "2023-01-01T01:01:01Z",
		ExpiresTime:  "2023-01-02T01:01:01Z",
		BodyMimeType: "text/plain",
		Body:         "Share and Enjoy!",
		Version:      ma.VERSION,
		Signature:    "",
	}

	return msg

}
