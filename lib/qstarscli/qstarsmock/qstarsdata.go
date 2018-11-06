package qstarsmock

var mockdata = map[string]string{
	"accounts.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "pubKey": "cosmosaccpub1zcjduepqvsxn3hhahlqqlz4l76ecftaqjrad5ayra5vxzjdnc4w5kqn0tnsq6c5x5k",
    "privKey": "PWMRFuCeJdXIzLb89BlcaxK1Wn6v0YxLLZDjDeFs8kJkDTje/b/AD4q/9rOEr6CQ+tp0g+0YYUmzxV1LAm9c4A==",
    "addr": "address1knzpxm6fp2s07enzu2zdfnm05wledt5nkt8zhn",
    "mnemonic": "holiday runway crush boring include stuff axis setup meadow gown feel annual climb hybrid enrich addict bread among expect pond engine enemy later vicious",
    "type": "local"
  }
}`,
	"accounts_address.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": 0,
    "message": "decoding bech32 failed: invalid bech32 string length 7"
  }
}`,
	"accounts_address_send.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": 0,
    "message": "UnmarshalJSON cannot decode empty bytes"
  }
}`,
	"kv.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": 0,
    "message": "UnmarshalJSON cannot decode empty bytes"
  }
}`,
	"kv_key.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": 0,
    "message": "ABCIQuery: Error unmarshalling rpc response: invalid character 'p' after top-level value"
  }
}`,
	"node_version.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "error": {
    "code": 0,
    "message": "no RPC client defined"
  }
}`,
	"version.json": `
{
  "jsonrpc": "2.0",
  "id": "",
  "result": {
    "version": "0.24.2"
  }
}`,
}
