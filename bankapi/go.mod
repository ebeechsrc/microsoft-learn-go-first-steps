module example.com/bankapi

go 1.20

require (
	github.com/msft/bank v0.0.1
)

replace github.com/msft/bank => ../bankcore
