module helloworld

require (
	github.com/ebeechsrc/microsoft-learn-go-first-steps-calculator v0.0.0
	rsc.io/quote v1.5.2
)

replace github.com/ebeechsrc/microsoft-learn-go-first-steps-calculator v0.0.0 => ../calculator

go 1.20
