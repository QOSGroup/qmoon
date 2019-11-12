package lib

import (
	"github.com/gin-gonic/gin/json"
	"log"
	"testing"
)

func TestValidators(t *testing.T) {
	validators, err := NewCosmosCli("").Validators("47.98.253.9:36657")
	bytes, err := json.Marshal(validators)
	log.Printf("res:%+v, err:%+v", string(bytes), err)
}
