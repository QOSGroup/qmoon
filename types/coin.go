// Copyright 2018 The QOS Authors

package types

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Coin struct {
	Denom  string `json:"denom"`
	Amount int64  `json:"amount"`
}

type Coins []Coin

func (c Coins) String() string {
	var res []string
	for _, v := range c {
		res = append(res, fmt.Sprintf("%v%v", v.Amount, v.Denom))
	}

	return strings.Join(res, ",")
}

//----------------------------------------
// Parsing

var (
	// Denominations can be 3 ~ 16 characters long.
	reDnm  = `[[:alpha:]][[:alnum:]]{2,15}`
	reAmt  = `[[:digit:]]+`
	reSpc  = `[[:space:]]*`
	reCoin = regexp.MustCompile(fmt.Sprintf(`^(%s)%s(%s)$`, reAmt, reSpc, reDnm))
)

// ParseCoin parses a cli input for one coin type, returning errors if invalid.
// This returns an error on an empty string as well.
func ParseCoin(coinStr string) (coin Coin, err error) {
	coinStr = strings.TrimSpace(coinStr)

	matches := reCoin.FindStringSubmatch(coinStr)
	if matches == nil {
		err = fmt.Errorf("invalid coin expression: %s", coinStr)
		return
	}
	denomStr, amountStr := matches[2], matches[1]

	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		return
	}

	return Coin{denomStr, int64(amount)}, nil
}

func ParseCoins(s string) (coins Coins, err error) {
	ss := strings.Split(s, ",")
	for _, v := range ss {
		c, err := ParseCoin(v)
		if err != nil {
			return nil, err
		}

		coins = append(coins, c)
	}

	return coins, nil
}
