package qos

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"testing"
)

func TestTx(t *testing.T) {
	proposals, err := NewQosCli("").QueryTx("47.103.78.91:26657", "443259A455F99566C901CD9A10F9541D26F8EED70B7480BD4F2312EC637A2876")
	bytes, err := json.Marshal(proposals)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}

func TestQueryProposals(t *testing.T) {
	proposals, err := NewQosCli("").QueryProposals("39.97.234.227:26657")
	bytes, err := json.Marshal(proposals)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}

func TestQueryProposal(t *testing.T) {
	proposals, err := NewQosCli("").QueryProposal("39.97.234.227:26657", 1)
	bytes, err := json.Marshal(proposals)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}

func TestQueryVotes(t *testing.T) {
	proposals, err := NewQosCli("").QueryVotes("39.97.234.227:26657", 1)
	bytes, err := json.Marshal(proposals)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}

func TestQueryDeposits(t *testing.T) {
	proposals, err := NewQosCli("").QueryDeposits("39.97.234.227:26657", 1)
	bytes, err := json.Marshal(proposals)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}

func TestQueryTally(t *testing.T) {
	proposals, err := NewQosCli("").QueryTally("39.97.234.227:26657", 1)
	bytes, err := json.Marshal(proposals)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}

func TestQueryDelegationsWithValidator(t *testing.T) {
	proposals, err := NewQosCli("").QueryDelegationsWithValidator("47.103.79.28:26657", "qosval19hrl38w5lm6sklw2hzrzrjtsxudpy8hyfaea3e")
	bytes, err := json.Marshal(proposals)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}
