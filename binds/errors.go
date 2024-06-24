package binds

import "errors"

var NoFee error = errors.New("no fee for the given pair")
var EmptyAddress error = errors.New("no address was provided")
var NoEthClient error = errors.New("no ethclient was provided")
