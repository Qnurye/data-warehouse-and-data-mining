package fp

// Mine 从给定的头表和最小支持度挖掘频繁模式
func Mine(headerTable map[string]*Node, minSupport int) map[string]int {
	fps := make(map[string]int) // 频繁模式集

	for item, node := range headerTable {

		supt := 0            // 支持度
		var conPs [][]string // 条件模式基

		// 找到条件模式基
		for curr := node; curr != nil; curr = curr.next {
			supt += curr.count

			var path []string

			// 从当前节点向上遍历到根节点，构建路径
			for parent := curr.parent; parent != nil && parent.item != ""; parent = parent.parent {
				path = append([]string{parent.item}, path...)
			}

			// 根据节点计数，将路径加入条件模式基
			for i := 0; i < curr.count; i++ {
				conPs = append(conPs, path)
			}
		}

		// 如果支持度大于等于最小支持度
		if supt >= minSupport {

			fps[join([]string{item})] = supt      // 记录当前项的频繁模式
			_, ht := BuildTree(conPs, minSupport) // 构建条件 FP 树
			subPs := Mine(ht, minSupport)         // 递归挖掘条件 FP 树中的模式
			for subP, subSupt := range subPs {
				comP := append([]string{item}, split(subP)...) // 合并当前项和子模式
				fps[join(comP)] = subSupt                      // 记录合并后的频繁模式
			}

		}

	}
	return fps // 返回所有挖掘到的频繁模式
}
