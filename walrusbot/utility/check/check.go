package check

import "walrusbot/utility/log"

// TODO: This should be a LITTLE more clever than just FATAL-ing all the time.
func Err(err error, optional_msg ...string) {
	msg := "unhandled error"
	if optional_msg != nil {
		msg = optional_msg[0]
	}
	if err != nil {
		log.Fatalw(msg, "err", err)
	}
}
