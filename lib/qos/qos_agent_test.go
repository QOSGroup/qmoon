package qos

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"testing"
)

func TestQueryProposals(t *testing.T) {
	proposals, err := NewQosCli("").QueryProposals("39.97.234.227:26657")
	//log.Printf("res:%+v, err:%+v", proposals, err.Error())
	//bytes, err := json.Marshal(proposals)
	//log.Printf("res:%+v, err:%+v", string(bytes), err.Error())

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
