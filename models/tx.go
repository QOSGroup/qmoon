package models

import (
	"github.com/QOSGroup/qmoon/types"
	"github.com/QOSGroup/qmoon/utils"
	"strconv"
	"time"

	"github.com/QOSGroup/qmoon/models/errors"
	"github.com/go-xorm/xorm"
)

type Tx struct {
	Id          int64  `xorm:"pk autoincr BIGINT"`
	ChainId     string `xorm:"-"`
	Height      int64  `xorm:"index(txs_height_idx) BIGINT"`
	TxType      string `xorm:"TEXT"`
	Index       int64  `xorm:"BIGINT"`
	Hash        string `xorm:"index(hash_idx) TEXT"`
	Maxgas      int64 `xorm:"BIGINT"`
	GasWanted   int64
	Fee         string `xorm:"TEXT"`
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
	Hash     string `xorm:"unique(hash_seq_idx) TEXT"`
	Seq      int64  `xorm:"unique(hash_seq_idx) BIGINT"`
	TxType   string `xorm:"TEXT"`
	OriginTx string `xorm:"TEXT"`
	JsonTx   string `xorm:"TEXT"`
}

type ITxAddress struct {
	Id      int64  `xorm:"pk autoincr BIGINT"`
	TxHash  string `xorm:"index(hash_seq_idx) TEXT"`
	ITxSeq  int64 `xorm:"index(hash_seq_idx) BIGINT"`
	Address string `xorm:"index(address_idx) TEXT"`
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

// please note that tx status code in db is the opesite to it in the chain
func UpdateTxStatusByHash (chainID string, status int, hash string, gas_used int64, gas_wanted int64) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}
	if status == 0 || status == 1 {
		status = 1 - status
	}
	_, err = x.Exec("update tx set tx_status= " + strconv.FormatInt(int64(status), 10) + ", gas_used=" + strconv.FormatInt(gas_used, 10) + ", gas_wanted=" + strconv.FormatInt(gas_wanted, 10) + " where hash = ?", hash)
	return err
}

func (t *Tx) InsertOrUpdate(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = TxByHash(chainID, t.Hash, 0, 0, 0, 0)
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

			addrs := utils.FindAddressInJson(itx.JsonTx)
			for _, addr := range addrs {
				addITx := ITxAddress{
					TxHash: itx.Hash,
					ITxSeq: itx.Seq,
					Address: addr,
				}
				_, err = x.Insert(addITx)
				if err != nil {
					return err
				}
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
	Height		int64
	MinId     int64
	MaxId     int64
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
	defer sess.Close()

	if opt.Height > 0 {
		sess = sess.Where("height = ?", opt.Height)
	}

	if opt.MinId > 0 {
		sess = sess.Where("id > ?", opt.MinId)
	}

	if opt.MaxId > 0 {
		sess = sess.Where("id < ?", opt.MaxId)
	}

	if opt.TxType != "" {
		sess = sess.Where("tx_type like ?", "%"+opt.TxType+"%")
	}

	if opt.Limit > 0 {
		sess = sess.Limit(opt.Limit, opt.Offset)
	} else {
		sess = sess.Limit(20, 0)
	}

	return txs, sess.OrderBy("id desc").Find(&txs)
}

func TxByHeightIndex(chainID string, height, index int64, minId, maxId, limit, offset int64) (*Tx, error) {
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

	tx.ITxs, err = ITxByHash(chainID, tx.Hash, minId, maxId, limit, offset)

	return tx, nil
}

func TxByHash(chainID string, hash string, minId, maxId, limit, offset int64) (*Tx, error) {
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

	tx.ITxs, err = ITxByHash(chainID, hash, minId, maxId, limit, offset)

	return tx, nil
}

func ITxByHash(chainID string, hash string, minId, maxId, limit, offset int64) ([]*ITx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	itxs := make([]*ITx, 0)
	sess := x.NewSession()
	defer sess.Close()

	if minId > 0 {
		sess = sess.Where("id > ?", minId)
	}
	if maxId > 0 {
		sess = sess.Where("id < ?", maxId)
	}
	sess = sess.Where("hash = ?", hash)

	if limit >0 {
		sess = sess.Limit(int(limit), int(offset))
	}

	return itxs, sess.OrderBy("id desc").Find(&itxs)
}

func ITxByAddress(chainID string, address string) ([]*ITx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	itxadds := make([]*ITxAddress, 0)
	result := make([]*ITx,0)
	err = x.Where("address like ?", address).Find(&itxadds)
	if err != nil {
		return nil, errors.NotExist{Obj: "Address " + address}
	}

	if len(itxadds) > 0 {
		hashString := ""
		for _, itxadd := range itxadds {
			hashString += ", '" + itxadd.TxHash + "'"
		}
		hashString = hashString[2 :]

		err = x.Where(" hash in (" + hashString + ")").Find(&result)
		if err != nil {
			return nil, errors.NotExist{Obj: "ITx " + hashString}
		}
		//return result, nil
	}
	return result, nil
}

func TxByAddress(chainID string, address string, minId int64, maxId int64, offset int, limit int) ([]*Tx, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	itxadds := make([]*ITxAddress, 0)
	result := make([]*Tx,0)

	sess := x.NewSession()
	defer sess.Close()

	err = x.Where("address like ?", address).Limit(1000).Find(&itxadds)
	if err != nil {
		//return nil, errors.NotExist{Obj: "Address " + address}
		return nil, err
	}

	if len(itxadds) > 0 {
		hashString := ""
		for _, itxadd := range itxadds {
			hashString += ", '" + itxadd.TxHash + "'"
		}
		hashString = hashString[2 :]
		sess = sess.Where(" hash in ( "+hashString+" )")
		if maxId > 0 {
			sess = sess.Where("id < ?", maxId)
		}
		if minId > 0 {
			sess = sess.Where("id > ?", minId)
		}
		if limit > 0 {
			sess = sess.Limit(limit, offset)
		} else {
			sess = sess.Limit(20, 0)
		}
		err = sess.Desc("id").Find(&result)
		if err != nil {
			return nil, errors.NotExist{Obj: "ITx " + hashString}
			//return nil, err
		}
	}
	return result, nil
}

func TxByAddressAndType(chainID string, address string, minHeight int64, maxHeight int64, offset int, limit int, txTypes... string) ([]*types.ValidatorOperations, error) {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return nil, err
	}
	itxadds := make([]*ITxAddress, 0)

	itxResult := make([]*ITx,0)
	txResult := make([]*Tx,0)

	result := make([]*types.ValidatorOperations, 0)
	err = x.Where("address like ?", address).Limit(1000).Find(&itxadds)
	if err != nil {
		// return nil, errors.NotExist{Obj: "Address " + address}
		return nil, err
	}

	sess := x.NewSession()
	defer sess.Close()

	if len(itxadds) > 0 {
		hashString := ""
		for _, itxadd := range itxadds {
			hashString += ", '" + itxadd.TxHash + "'"
		}


		hashString = hashString[2 :]
		sess = sess.Where("hash in ( "+hashString+" )")


		if maxHeight > 0 {
			sess = sess.Where("id <= ?", maxHeight)
		}
		if minHeight > 0 {
			sess = sess.Where("id >= ?", minHeight)
		}
		if limit <= 0 {
			limit = 0
			offset = 0
		}
		if limit > 0 {
			sess = sess.Limit(limit, offset)
		} else {
			sess = sess.Limit(20, 0)
		}

		if len(txTypes) != 0 {
			typeString := ""
			for _, t := range txTypes {
				typeString += ", '" + t + "'"
			}
			typeString = typeString[2 :]
			typeString = "tx_type in (" + typeString +")"
			err = sess.Where(typeString).Limit(limit, offset).Desc("id").Find(&itxResult)
		} else {
			err = sess.Limit(limit, offset).Desc("id").Find(&itxResult)
		}

		err = sess.Limit(limit, offset).Desc("id").Find(&txResult)


		//if maxHeight > minHeight && minHeight > 0 {
		//	err = x.Where(" hash in ( "+hashString+" )"+ typeString + " and height between ? and ?", minHeight, maxHeight).Limit(limit, offset).Desc("id").Find(&itxResult)
		//} else {
		//	err = x.Where(" hash in ( "+hashString+" )"+ typeString).Limit(limit, offset).Desc("id").Find(&itxResult)
		//}
		//if err != nil {
		//	//return nil, errors.NotExist{Obj: "ITx " + hashString}
		//	return nil, err
		//}
		//
		//if maxHeight > minHeight && minHeight > 0 {
		//	err = x.Where(" hash in ( "+hashString+" ) and height between ? and ?", minHeight, maxHeight).Limit(limit, offset).Desc("height").Find(&txResult)
		//} else {
		//	err = x.Where(" hash in ( "+hashString+" )").Limit(limit, offset).Desc("height").Find(&txResult)
		//}
		if err != nil {
			//return nil, errors.NotExist{Obj: "ITx " + hashString}
			return nil, err
		}

		itxsMap := make(map[string]*ITx)
		for _, itx :=range itxResult {
			itxsMap[itx.Hash] = itx
		}
		for _, tx := range txResult {
			if x, ok := itxsMap[tx.Hash]; ok {
				result = append(result,
					&types.ValidatorOperations{Height:tx.Height, Operation:x.TxType})
			}
		}
	}
	return result, nil
}


