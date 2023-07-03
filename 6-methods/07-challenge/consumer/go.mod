module example.com/consumer

go 1.20

require example.com/store v0.0.0

replace example.com/store v0.0.0 => ../store
