1. `go run gen_prod.go 1024 > prod_data.go` - 1024 arg is number of megabytes
2. `go build -tags prod` - takes 1 min for me
3. Pick a key out of the file
4. `time ./example m355pOs9BEVCdzWceLQ747x9MY1KasZF` - replace with some key from the file

I get ~150ms reliably on my laptop for 1G, ~20ms for 100M, just crudely using `time` as above.
