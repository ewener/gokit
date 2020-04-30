/******************************************************
# DESC       : 数值转换与计算
# MAINTAINER : yamei
# EMAIL      : daixw@ecpark.cn
# DATE       : 2019/12/5
******************************************************/
package backend

import "encoding/binary"

var (
	asPress   = 101.325
	byteOrder = binary.BigEndian
)

// 处理大气压mba和AD值关系。输入是ad_value（0~1023），输出是mba值
func Ad2mbar(ad uint16) float64 {
	// 将ad值转化变成电压v_tmp_value，电压压力差是0-4.5v
	vTmpValue := float64(ad)*4/1024 + 0.5
	// 将电压v_tmp_value转化为标准kpa值。kpa值范围是0kpa~1200kpa
	kpaTmpValue := (vTmpValue/5 - 0.1) * 1000 / 0.75
	// 将kpa_tmp_value值转化变成mba值
	mbaValue := kpaTmpValue / asPress
	return mbaValue
}

func Mbar2ad(mbar float64) uint16 {
	kpaTmpValue := mbar * asPress
	vTmpValue := (kpaTmpValue*0.75/1000 + 0.1) * 5
	ad := (vTmpValue - 0.5) * 1024 / 4
	return uint16(ad)
}

func AD2WaterLevel(ad uint16) float64 {
	// 将ad值转化变成电压v_tmp_value，电压压力差是0-5v
	vTmpvalue := float64(ad) * 5 / 1023
	// 将电压v_tmp_value转化为标准kpa值。kpa值范围是0kpa~20kpa
	kpaTmpValue := vTmpvalue / 5 * 20
	// 将kpa_tmp_value值转化变成水位高度
	wlValue := kpaTmpValue / 9.8 * 100
	return wlValue
}

// Int16ToBytes is
func Int16ToBytes(i uint16) []byte {
	var buf = make([]byte, 2)
	byteOrder.PutUint16(buf, i)
	return buf
}

// BytesToInt16 is
func BytesToInt16(buf []byte) uint16 {
	return byteOrder.Uint16(buf)
}

// Int32ToBytes is
func Int32ToBytes(i uint16) []byte {
	var buf = make([]byte, 4)
	byteOrder.PutUint32(buf, uint32(i))
	return buf
}

// BytesToInt32 is
func BytesToInt32(buf []byte) uint32 {
	return byteOrder.Uint32(buf)
}
