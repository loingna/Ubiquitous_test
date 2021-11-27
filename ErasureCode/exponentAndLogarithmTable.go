// Copyright 2021 ubiquitous Storage, Hangzhou Dianzi University
// See LICENSE for copying information

package ErasureCode

// This file is used to generate the exponent table and
// log table on GF(2^8),
// which is moduled by the prime polynomial: P(x) = x^8 + x^4 + x^3 + x^2 + 1
// 本文件用于计算生成GF(2^8)域上的指数表和对数表，质多项式为：P(x) = x^8 + x^4 + x^3 + x^2 + 1

type ExponentAndLogTable struct {
	ExponentTable  [256]int
	LogTable [256]int
}

func (this *ExponentAndLogTable) GenerateExponentAndLogTable() error {
	n := 1
	for i := 0; i < 256; i++ {
		this.ExponentTable[i] = n
		this.LogTable[n] = i

		n *= 2

		if n >= 256 {
			n = n ^ 0x11d
		}
	}
	this.LogTable[1] = 0 // otherwise this.LogTable[1]=255, but it should be 0
	 return nil
}
