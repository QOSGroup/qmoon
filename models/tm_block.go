package models

type TmBlock struct {
	Id     int64  `xorm:"pk autoincr BIGINT"`
	Height int64  `xorm:"unique(tm_blocks_height_idx) BIGINT"`
	Data   string `xorm:"TEXT"`
}

func (tb *TmBlock) Insert(chainID string) error {
	x, err := GetNodeEngine(chainID)
	if err != nil {
		return err
	}

	_, err = x.Insert(tb)
	if err != nil {
		return err
	}

	return nil
}
