module github.com/ODIM-Project/ODIM/svc-systems

go 1.13

require (
	github.com/Joker/hpp v1.0.0 // indirect
	github.com/ODIM-Project/ODIM/lib-dmtf v0.0.0-20200727133207-df3dfb728bd1
	github.com/ODIM-Project/ODIM/lib-persistence-manager v0.0.0-20201021053518-88345d4a9988
	github.com/ODIM-Project/ODIM/lib-rest-client v0.0.0-20201103092246-3e5a86618649
	github.com/ODIM-Project/ODIM/lib-utilities v0.0.0-20201014070346-4417177cddff
	github.com/prometheus/common v0.6.0
	github.com/stretchr/testify v1.5.1
	gopkg.in/go-playground/validator.v9 v9.30.0
)

replace github.com/ODIM-Project/ODIM/lib-utilities => ../lib-utilities
