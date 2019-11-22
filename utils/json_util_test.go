package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson(t *testing.T) {
	const jstr = `{"qosval111":"adad",check_tx":{"sender":"qosval19hrl38w5lm6sklw2hzrzrjtsxudpy8hy3dh7h5","gasUsed":"7060","events":[]},"deliver_tx":{"receiver":"qosacc19hrl38w5lm6sklw2hzrzrjtsxudpy8hy3dh7h5","gasUsed":"9000","events":[{"type":"vote-proposal","attributes":[{"key":"cHJvcG9zYWwtaWQ=","value":"MQ=="},{"key":"dm90ZXI=","value":"cW9zYWNjMTlocmwzOHc1bG02c2tsdzJoenJ6cmp0c3h1ZHB5OGh5M2RoN2g1"}]},{"type":"message","attributes":[{"key":"bW9kdWxl","value":"Z292"},{"key":"YWN0aW9u","value":"dm90ZS1wcm9wb3NhbA=="},{"key":"Z2FzLnBheWVy","value":"cW9zYWNjMTlocmwzOHc1bG02c2tsdzJoenJ6cmp0c3h1ZHB5OGh5M2RoN2g1"}]}]},"hash":"4A1852C65AEC17EFE7D7E6531C09BF606D84640A0F9CC81C4D2A671886F232FE","height":"50095"}`
	addrs := FindAddressInJson(jstr)
	fmt.Println(addrs)
	assert.NotNil(t, addrs)
}
