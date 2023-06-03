package make

// https://leetcode.cn/problems/reverse-bits/?envType=study-plan-v2&envId=top-interview-150

// v1 : 4ms
func reverseBits(n uint32) uint32 {
	table := [][]uint32 {
		{16, 0xFFFF0000, 0x0000FFFF, },
		{8,  0xFF00FF00, 0x00FF00FF, },
		{4,  0xF0F0F0F0, 0x0F0F0F0F, },
		{2,  0xCCCCCCCC, 0x33333333, },
		{1,  0xAAAAAAAA, 0x55555555, },
	}

	for i := 0; i < len(table); i++ {
		n = (n & table[i][1] >> table[i][0]) | (n & table[i][2] << table[i][0])
	}

	return n
}