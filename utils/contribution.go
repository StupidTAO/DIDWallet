package utils

import (
	"fmt"
	"math"
)

const e = 2.7182818284
const pi = 3.141592653589793

func GetContribution(baseAmount uint32, middleAmount int32, basePriority float32,  middlePriority float32, outDgree int) (contribute float64, err error) {

	//计算捐赠中值指标
	r := 2 - 0.1 * float64(middleAmount)
	r1 := 1 + math.Pow(e, r)
	middleAmountIndex := 10 * (float64(middleAmount) / r1) + float64(baseAmount)

	//计算优先级中值指标
	r = 2 - 0.1 * float64(middlePriority)
	r1 = 1 + math.Pow(e, r)
	middlePriorityIndex := (float64(middlePriority) / r1) + float64(basePriority)

	//计算出度指标
	atan1 := math.Atan(float64(outDgree))
	r = pi/3 -atan1
	function := -2 * math.Pow(r, 2)
	dgreeIndex := float64(outDgree) * math.Pow(e, function)
	fmt.Println(middleAmountIndex, middlePriorityIndex, dgreeIndex)
	fmt.Println("outDgreeIndex", dgreeIndex)

	index := dgreeIndex + 2*middleAmountIndex + middlePriorityIndex
	fmt.Println("index: ", index)

	return index, nil
}
