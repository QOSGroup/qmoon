package txs

type TxStd struct {
	ITxs      []ITx       `json:"itx"`      //ITx接口，将被具体Tx结构实例化
	Signature []Signature `json:"sigature"` //签名数组
	ChainID   string      `json:"chainid"`  //ChainID: 执行ITx.exec方法的链ID
	MaxGas    string      `json:"maxgas"`   //Gas消耗的最大值
}

// 签名结构体
type Signature struct {
	Pubkey    string `json:"pubkey"`    //可选
	Signature string `json:"signature"` //签名内容
	Nonce     int64  `json:"nonce"`     //nonce的值
}

//type Pubkey struct {
//	Yype  string `json:"type"`  //可选
//	Value string `json:"value"` //签名内容
//}

type ITx struct {
	Sender    string   `json:"sender,omitempty"`
	Senders   []Amount `json:"senders,omitempty"`
	Receivers []Amount `json:"receivers,omitempty"`
}

type Amount struct {
	Addr string `json:"addr"`
	Qos  string `json:"qos"`
	Qscs string `json:"qscs"`
}
