package check

import "walrusbot/utility/log"

func Err(err error, optional_msg ...string) {
	msg := "unhandled error"
	if optional_msg != nil {
		msg = optional_msg[0]
	}
	if err != nil {
		log.Fatalw(msg, "err", err)
	}
}
