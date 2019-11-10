package sender

import (
	"gofinancetoCSV/dataconverter"
	"gofinancetoCSV/receiver"
	"strings"
	"time"

	"github.com/piquette/finance-go/chart"
	"github.com/piquette/finance-go/datetime"
)

type FinanceSenderConfig struct {
	Symbols  string
	StartDay string
	EndDay   string
	Freq     string
	Fields   []string
}

type FinanceSender struct {
	FinanceSenderConfig
}

func NewFinanceSender(fsc FinanceSenderConfig) *FinanceSender {
	return &FinanceSender{
		fsc,
	}
}

func (fs *FinanceSender) Send(recv receiver.Receiver) error {
	fieldsMap := make(map[string]struct{})
	for _, f := range fs.Fields {
		fieldsMap[strings.ToLower(f)] = struct{}{}
	}
	t1, err := time.Parse(time.RFC3339, fs.StartDay)
	if err != nil {
		return err
	}
	t2, err := time.Parse(time.RFC3339, fs.EndDay)
	if err != nil {
		return err
	}
	symbols := strings.Split(fs.Symbols, ",")
	for _, symbol := range symbols {
		params := &chart.Params{
			Symbol:   symbol,
			Start:    datetime.New(&t1),
			End:      datetime.New(&t2),
			Interval: datetime.Interval(fs.Freq),
		}
		iter := chart.Get(params)
		for iter.Next() {
			p, err := dataconverter.FromGFStructToParam(fieldsMap, *iter.Bar())
			if err != nil {
				return err
			}
			if _, ok := p["symbol"]; !ok {
				p["symbol"] = symbol
			}
			err = recv.Receive(p)
			if err != nil {
				return err
			}
		}
	}
	return recv.Stop()
}
