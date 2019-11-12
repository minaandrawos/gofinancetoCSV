# gofinancetoCSV
A command line tool that generates CSV reports of stock data, written in Go

# Build

```
go get github.com/minaandrawos/gofinancetoCSV/...
go build -o financetoCSV
```

# Usage

```
  -symbols string
        Symbols to import to CSV (default "FB,TWTR")
  -start string
        Begin day (default "2017-01-02T15:04:05Z")
  -end string
        end day (default "2018-02-02T15:04:05Z")
  -fields string
        fields to parse: open,close,low,high,adjclose,volume,timestamp (default "open,low,high,close,timestamp")
  -frequency string
        Data frequency: 1d,5d,1mo,3mo,6mo,1y,2y,5y,10y (default "1d")
  -output string
        output csv file name (default "outputfinance.csv")
```

# Usage example
  
```
financetoCSV -symbols "fb,amzn,twtr" -start 2017-03-12T00:00:00Z -end 2018-01-03T00:00:00Z -fields "open,close,timestamp" -frequency 1d -output outputfbwork.csv
```

