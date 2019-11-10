package main

import (
	"flag"
	"fmt"
	"gofinancetoCSV/receiver"
	"gofinancetoCSV/sender"
	"log"
	"strings"
)

var FieldsOptions = map[string]struct{}{"open": struct{}{}, "low": struct{}{}, "high": struct{}{}, "close": struct{}{}, "timestamp": struct{}{}}

func main() {
	var (
		symbols     string
		startDay    string
		endDay      string
		fields      string
		outputFname string
		freq        string
	)
	flag.StringVar(&symbols, "symbols", "FB,TWTR", "Symbols to import to CSV")
	flag.StringVar(&startDay, "start", "2017-01-02T15:04:05Z", "Begin day")
	flag.StringVar(&endDay, "end", "2018-02-02T15:04:05Z", "end day")
	flag.StringVar(&fields, "fields", "open,low,high,close,timestamp", "fields to parse: open,close,low,high,adjclose,volume,timestamp")
	flag.StringVar(&freq, "frequency", "1d", "Data frequency: 1d,5d,1mo,3mo,6mo,1y,2y,5y,10y")
	flag.StringVar(&outputFname, "output", "outputfinance.csv", "output csv file name")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"Usage of financetocsv : \n\r"+
				"financetocsv -symbols \"<s1,s2>\" -start <start date> -end <end date> -fields \"<field names>\" -frequency <data frequecy> -output <csv file name>\n\r"+
				"Flags:\n\r")
		flag.PrintDefaults()
	}
	flag.Parse()
	normalizedFields := strings.Replace(fields, " ", "", -1)
	fieldsSlice := strings.Split(strings.ToLower(normalizedFields), ",")
	for _, f := range fieldsSlice {
		if _, ok := FieldsOptions[f]; !ok {
			log.Fatal("Invalid field ", f)
		}
	}
	csvrvcr, err := receiver.NewCSVRcv(outputFname, append([]string{"symbol"}, fieldsSlice...))
	if err != nil {
		log.Fatal(err)
	}
	normalizedSymbols := strings.Replace(symbols, " ", "", -1)
	config := sender.FinanceSenderConfig{normalizedSymbols, startDay, endDay, freq, fieldsSlice}
	financeSender := sender.NewFinanceSender(config)
	err = financeSender.Send(csvrvcr)
	if err != nil {
		log.Fatal(err)
	}
}
