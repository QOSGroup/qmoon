// Copyright 2018 The QOS Authors

package tx

// QuerySequenceIn query sequence in
func (c Client) SequenctIn(chainID string) (int64, error) {
	return c.c.QuerySequenceIn(chainID)
}

// QuerySequenceOut query sequence out
func (c Client) SequenctOut(chainID string) (int64, error) {
	return c.c.QuerySequenceOut(chainID)
}
