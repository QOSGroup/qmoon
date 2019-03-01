// Copyright 2018 The QOS Authors

package metric

//
//func TestPush(t *testing.T) {
//	var vars []types.Validator
//	vars = append(vars, types.Validator{
//		ChainID:     "qos_test",
//		Address:     "a",
//		VotingPower: time.Now().Unix() % 100 * 2,
//	})
//	vars = append(vars, types.Validator{
//		ChainID:     "qos_test",
//		Address:     "b",
//		VotingPower: time.Now().Unix() % 80 * 2,
//	})
//
//	ValidatorVotingPower(vars)
//}
//
//func TestQuery(t *testing.T) {
//	cli, err := api.NewClient(api.Config{Address: prometheusServer})
//	assert.Nil(t, err)
//
//	papi := v1.NewAPI(cli)
//
//	now := time.Now()
//	ds, err := papi.QueryRange(context.Background(), "validator_voting_power_percent_capricorn_2000", v1.Range{
//		Start: utils.NDaysAgo(now, 1),
//		End:   now,
//		Step:  time.Minute * 24,
//	})
//	assert.Nil(t, err)
//
//	if ds.Type() == model.ValMatrix {
//		m := ds.(model.Matrix)
//		v := m[0]
//		t.Logf("---m:%+v %v", v.Values[0].Timestamp, v.Values[0].Value.String())
//		t.Logf("---m:%+v %v", v.Values[1].Timestamp, v.Values[1].Value.String())
//		t.Logf("---m:%+v %v", v.Values[2].Timestamp, v.Values[2].Value.String())
//		t.Logf("---m:%+v %v", v.Values[3].Timestamp, v.Values[3].Value.String())
//
//	} else {
//		t.Logf("------unkown")
//	}
//}
//
//func TestQueryValidatorVotingPower(t *testing.T) {
//	ds, err := QueryValidatorVotingPower("capricorn-2000", "38C6373076DC69C5B4F370A0EF10AF13E8ACFAE0")
//	assert.Nil(t, err)
//	t.Logf("%v", ds)
//}
