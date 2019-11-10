package sender

import "gofinancetoCSV/receiver"

type Sender interface {
	Send(receiver.Receiver) error
}
