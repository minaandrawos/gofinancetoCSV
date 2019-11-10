package receiver

import "gofinancetoCSV/dataconverter"

type Receiver interface {
	Receive(dataconverter.Params) error
	Stop() error
}
