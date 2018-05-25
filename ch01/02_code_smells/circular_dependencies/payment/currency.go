// +build bad

package payment

import (
	"errors"

	"github.com/corsc/Hands-On-Dependency-Injection-in-Go/ch01/02_code_smells/circular_dependencies/config"
)

// Currency is custom type for currency
type Currency string

// Processor processes payments
type Processor struct {
	Config *config.Config
}

// Pay makes a payment in the default currency
func (p *Processor) Pay(amount float64) error {
	// TODO: implement me
	return errors.New("not implemented yet")
}
