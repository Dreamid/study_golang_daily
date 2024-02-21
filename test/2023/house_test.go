package test

import (
	"strings"
	"testing"
)

/*
*
total_lx = 91808
total_bj = 545600
total_all = 637408

total_lx = 93398
total_bj = 547800
total_all = 641198
------------------------

	93398
	345400
	438798

--------updatetime-------

	house_test.go:47: total_lx = 86270
	house_test.go:48: total_bj = 323400
	house_test.go:49: total_all = 409670
*/
func TestHouse(t *testing.T) {
	//剩余月份
	months := 146
	//基础金额
	baseAmount := 1210
	//递减差值
	reduce := 8
	//所有金额
	var totalAmount int
	var totalBJ int
	var totalAll int
	var BJ = 2200
	for i := 0; i < months; i++ {
		totalBJ += BJ
		totalAll += BJ
		temp := baseAmount - (i * reduce)
		if temp > 0 {
			//当月利息
			tempAmount := baseAmount - temp
			//利息总额
			totalAmount += tempAmount
			//利息纳入总额
			totalAll += tempAmount
		}

	}
	totalBJ += 2200
	totalAmount += 1590
	totalAll += 2200 + 1590
	t.Logf("total_lx = %v", totalAmount) //91808
	t.Logf("total_bj = %v", totalBJ)     //
	t.Logf("total_all = %v", totalAll)   //
}

/*
*

		total_lx = 231360
		total_bj = 547800
	 	total_all = 779160

------------------update

	total_lx = 231360
	total_bj = 530200
	total_all = 761560
*/
func TestHouseWithOut20(t *testing.T) {
	//剩余月份
	months := 236
	//基础金额
	baseAmount := 1925
	//递减差值
	reduce := 8
	//所有金额
	var totalAmount int
	var totalBJ int
	var total int
	var BJ = 2200
	for i := 0; i < months; i++ {
		//本金总额增加
		totalBJ += BJ
		//本金先纳入总额
		total += BJ
		temp := baseAmount - (i * reduce)
		if temp > 0 {
			//当月利息
			tempAmount := baseAmount - temp
			//利息总额
			totalAmount += tempAmount
			//利息纳入总额
			total += tempAmount
		}
	}
	t.Logf("total_lx = %v", totalAmount) //231360
	t.Logf("total_bj = %v", totalBJ)     //
	t.Logf("total_all = %v", total)      //
}

/*
*
total_lx = 91808
total_bj = 545600
total_all = 637408

total_lx = 93398
total_bj = 547800
total_all = 641198
------------------------

	93398
	345400
	438798

--------updatetime-------

	house_test.go:47: total_lx = 86270
	house_test.go:48: total_bj = 323400
	house_test.go:49: total_all = 409670
*/
func TestHowManyWith30(t *testing.T) {
	//剩余月份
	months := 146
	//基础金额
	baseAmount := 1210
	//递减差值
	reduce := 8
	//所有金额
	var totalAmount int
	var totalBJ int
	var totalAll int
	var BJ = 2200
	for i := 0; i < months; i++ {
		totalBJ += BJ
		totalAll += BJ
		temp := baseAmount - (i * reduce)
		if temp > 0 {
			//当月利息
			tempAmount := baseAmount - temp
			//利息总额
			totalAmount += tempAmount
			//利息纳入总额
			totalAll += tempAmount
		}
		if totalBJ >= 100000 {
			t.Logf("当月利息:%v", baseAmount-temp)
			t.Logf("当月还款总额:%v", BJ+baseAmount-temp)
			return
		}

	}
	totalBJ += 2200
	totalAmount += 1590
	totalAll += 2200 + 1590
	t.Logf("total_lx = %v", totalAmount) //91808
	t.Logf("total_bj = %v", totalBJ)     //
	t.Logf("total_all = %v", totalAll)   //
}

func TestSplitString(t *testing.T) {
	str := "http://192.168.50.204:5051/console/admin/remotebrowser/synchronization\r\nhttp://192.168.50.204:5051/console/admin/remotebrowser/synchronization\nhttp://192.168.50.204:5051/console/admin/remotebrowser/synchronization\nhttp://192.168.50.204:5051/console/admin/remotebrowser/synchronization\n\nhttp://192.168.50.204:5051/console/admin/remotebrowser/synchronization\n"
	result := SplitStringByLineBreaks(str)
	for _, v := range result {
		t.Logf("%v\n", v)
	}
}

// 根据换行符切割字符串
func SplitStringByLineBreaks(s string) []string {
	if len(s) == 0 {
		return []string{}
	}
	result := []string{}
	if strings.Contains(s, "\r\n") {
		result = strings.Split(s, "\r\n")
	} else {
		result = strings.Split(s, "\n")
	}
	return result
}
