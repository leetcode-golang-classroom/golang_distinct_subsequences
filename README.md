# golang_distinct_subsequence

Given two strings `s` and `t`, return *the number of distinct subsequences of `s` which equals `t`*.

A string's **subsequence** is a new string formed from the original string by deleting some (can be none) of the characters without disturbing the remaining characters' relative positions. (i.e., `"ACE"` is a subsequence of `"ABCDE"` while `"AEC"` is not).

The test cases are generated so that the answer fits on a 32-bit signed integer.

## Examples

**Example 1:**

```
Input: s = "rabbbit", t = "rabbit"
Output: 3
Explanation:
As shown below, there are 3 ways you can generate "rabbit" from S.
```

**Example 2:**

```
Input: s = "babgbag", t = "bag"
Output: 5
Explanation:
As shown below, there are 5 ways you can generate "bag" from S.

```

**Constraints:**

- `1 <= s.length, t.length <= 1000`
- `s` and `t` consist of English letters.

## 解析

題目給定兩個字串 s, t

要求寫一個演算法算出所有由 s 產生子字字元序列中可以組成 t 字串的個數

假設 p 是 s 的子字元序列代表

對每個 p中的字元 存在 s 字串中

且順序與原本在 s 字串中的一樣

要從 s 中找出能組成 t 的子字元序列

讓我們觀察以下範例

![](https://i.imgur.com/bSwNQNM.png)

用 t 每個字元做比較基礎

每當有發現 s 中有相符的字元

對 s 來說有兩種選擇 可以選擇當下的字元 或是換下一個字元來否配看看

以上面這個原則 我們可以畫出以下決策樹

![](https://i.imgur.com/6JHkpTa.png)

因為 (m,n) 只有在 m = len(s), 或是 n = len(t) 才會停止

所以數高會是 max(len(s), len(t)) = maxLen 

而每個點都會有最多 2 個選擇 所以 會是 $2^{maxLen}$

使用 cache 能夠降低時間複雜度到 O(len(s)*len(t))

因為 (m,n) 最多有 len(s)*len(t) 種可能

空間複雜度也是 len(s)*len(t), 因為要把可能的值暫存下來

然而，如果透過 DFS 遞迴呼叫還是需要 O(maxLen) 的 call stack space

要降低 call stack 空間複雜度 就需要使用 tabulation Dynamic Programming 的方式

 定義 dp[i][j] 代表 s 從 0 到 i 子字元序列能夠組成 字串 t 從 0 到 j 的方法數

![](https://i.imgur.com/clpiJkr.png)

從之前的比較方式 , 可以發現定義的 dp 陣列會有以下關係

dp[i][j] = dp[i-1][j-1] + dp[i-1][j]   if s[i-1] == t[j-]

dp[i][j] = dp[i-1][j]                       otherwise

初始化 對所有 dp[i][0] = 1 因為當 t 是空字串時 , s 可以使用空字串來組成因此方法數是 1

一樣地因為需要算所有的 dp[i][j]

dp[len(s)][len(t)] 即為所求

所以時間複雜度 是 O(len(s)*len(t)) 

空間複雜度也是 O(len(s)*len(t))

然而因為不需要使用像 DFS遞迴  不會需要額外耗費 call stack 的空間

## 程式碼
```go
package sol

func numDistinct(s string, t string) int {
	sLen, tLen := len(s), len(t)
	if sLen < tLen {
		return 0
	}
	// dp[i][j]: number of s[:i] subsequence of t[:j]
	dp := make([][]int, sLen+1)
	for row := range dp {
		dp[row] = make([]int, tLen+1)
		dp[row][0] = 1
	}
	for sEnd := 1; sEnd <= sLen; sEnd++ {
		for tEnd := 1; tEnd <= tLen; tEnd++ {
			if s[sEnd-1] == t[tEnd-1] {
				dp[sEnd][tEnd] = dp[sEnd-1][tEnd-1] + dp[sEnd-1][tEnd]
			} else {
				dp[sEnd][tEnd] = dp[sEnd-1][tEnd]
			}
		}
	}
	return dp[sLen][tLen]
}
```
## 困難點

1. 要看出子字元序列的組成遞迴關係不是很直觀

## Solve Point

- [x]  建立一個整數矩陣 dp 大小是 len(s) by len(t), 用來存放每個中間運算
- [x]  初始化所有 dp[i][0] = 1
- [x]  逐步推算每個 dp[i][j]
- [x]  回傳 dp[len(s)][len(t)]