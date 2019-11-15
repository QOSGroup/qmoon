package models

import (
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type Tx struct {
	Id          int64  `xorm:"pk autoincr BIGINT"`
	ChainId     string `xorm:"-"`
	Height      int64  `xorm:"index(txs_height_idx) unique(txs_height_index_idx) BIGINT"`
	TxType      string `xorm:"TEXT"`
	Index       int64  `xorm:"unique(txs_height_index_idx) BIGINT"`
	Hash        string
	Maxgas      int64 `xorm:"BIGINT"`
	GasWanted   int64
	Fee         string
	GasUsed     int64
	QcpFrom     string    `xorm:"TEXT"`
	QcpTo       string    `xorm:"TEXT"`
	QcpSequence int64     `xorm:"BIGINT"`
	QcpTxindex  int64     `xorm:"BIGINT"`
	QcpIsresult bool      `xorm:"BOOL"`
	TxStatus    int       `xorm:"default 0 INTEGER"`
	OriginTx    string    `xorm:"TEXT"`
	JsonTx      string    `xorm:"TEXT"`
	Log         string    `xorm:"TEXT"`
	Time        time.Time `xorm:"-"`
	TimeUnix    int64
	ITxs        []*ITx `xorm:"-"`
}

type ITx struct {
	Id       int64  `xorm:"pk autoincr BIGINT"`
	Hash     string `xorm:"TEXT"`
	Seq      int64  `xorm:"BIGINT"`
	TxType   string `xorm:"TEXT"`
	OriginTx string `xorm:"TEXT"`
	JsonTx   string `xorm:"TEXT"`
}

type ITxAddress struct {
	Id      int64  `xorm:"pk autoincr BIGINT"`
	ItxHash string `xorm:"TEXT"`
	TxHash  string `xorm:"TEXT"`
	Address string `xorm:"TEXT"`
}

func (t *Tx) BeforeInsert() {
	t.TimeUnix = t.Time.Unix()
}

func (t *Tx) BeforeUpdate() {
	t.TimeUnix = t.Time.Unix()
}

func (t *Tx) AfterSet(colName string, _ xorm.Cell) {
	switch colName {
	case "time_unix":
		t.Time = time.Unix(t.TimeUnix, 0).Local()
	}
}

func (t *Tx) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(t)
	if err != nil {
		return err
	}

	for _, itx := range t.ITxs {
		_, err = x.Insert(itx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *Tx) InsertOrUpdate(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = TxByHash(chainID, t.Hash)
	if err != nil {
		_, err = x.Insert(t)
		if err != nil {
			return err
		}

		for _, itx := range t.ITxs {
			_, err = x.Insert(itx)
			if err != nil {
				return err
			}
		}
		return nil
	}
	_, err = x.Update(t)
	if err != nil {
		return err
	}

	for _, itx := range t.ITxs {
		_, err = x.Update(itx)
		if err != nil {
			return err
		}
	}

	return nil
}

type TxOption struct {
	TxType        string
	MinHeight     int64
	MaxHeight     int64
	Address       string
	Offset, Limit int
}

func Txs(chainID string, opt *TxOption) ([]*Tx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	var txs = make([]*Tx, 0)

	sess := x.NewSession()
	if opt.MinHeight != 0 && opt.MaxHeight != 0 {
		sess = sess.Where("height >= ?", opt.MinHeight).Where("height <= ?", opt.MaxHeight)
	}

	if opt.Address != "" {
		var itxs = make([]*ITx, 0)
		sess = sess.Distinct("hash").Where("json_tx like ?", "%"+opt.Address+"%")
		sess.Find(&itxs)

		if len(itxs) > 0 {
			var hashS = ""
			for _, iTx := range itxs {
				hashS += ", '" + iTx.Hash + "'"
			}
			hashS = hashS[2 : len(hashS)-1]
			sess = sess.Where("hash like ?", "("+hashS+")")
		}
	}

	if opt.TxType != "" {
		sess = sess.Where("tx_type like ?", "%"+opt.TxType+"%")
	}

	if opt.Limit > 0 {
		sess = sess.Limit(opt.Limit, opt.Offset)
	}

	return txs, sess.OrderBy("height desc").Find(&txs)
}

func TxByHeightIndex(chainID string, height, index int64) (*Tx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	tx := &Tx{Height: height, Index: index}
	has, err := x.Get(tx)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "Tx"}
	}

	tx.ITxs, err = ITxByHash(chainID, tx.Hash)

	return tx, nil
}

func TxByHash(chainID string, hash string) (*Tx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}

	tx := &Tx{Hash: hash}
	has, err := x.Get(tx)
	if err != nil {
		return nil, err
	}

	if !has {
		return nil, errors.NotExist{Obj: "Tx"}
	}

	tx.ITxs, err = ITxByHash(chainID, hash)

	return tx, nil
}

func ITxByHash(chainID string, hash string) ([]*ITx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	itxs := make([]*ITx, 0)
	err = x.Where("hash = ?", hash).Find(&itxs)

	if err != nil {
		return nil, errors.NotExist{Obj: "ITx"}
	}

	return itxs, nil
}

func ITxByAddress(chainID string, address string) ([]*ITx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	itxadds := make([]*ITxAddress, 0)
	err = x.Where("address = ?", address).Find(&itxadds)
	if err != nil {
		return nil, errors.NotExist{Obj: "Address " + address}
	}

	if len(itxadds) > 0 {
		hashString := ""
		for _, itxadd := range itxadds {
			hashString += ", '" + itxadd.ItxHash + "'"
		}
		hashString = hashString[2 : len(hashString)-1]

		itxs := make([]*ITx, 0)
		err = x.Where(" hash in (" + hashString + ")").Find(&itxs)
		if err != nil {
			return nil, errors.NotExist{Obj: "ITx " + hashString}
		}
		return itxs, nil
	}
	return nil, errors.NotExist{Obj: "Address " + address}

}
