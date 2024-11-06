def calculate_support(itemset: frozenset, dataset: list[set]) -> float:
    """
    计算候选项集的支持度
    :param itemset: 项集
    :param dataset: 数据集
    :return: float 支持度
    """

    if len(dataset) == 0:
        return 0.0
    count = sum(1 for transaction in dataset if itemset.issubset(transaction))
    return count / len(dataset)


def apriori(dataset: list[set], min_support: float) -> dict[frozenset, float]:
    """
    Apriori 算法
    :param dataset: 数据集
    :param min_support: 最小支持度
    :return: dict 频繁项集及其支持度
    """

    # k = 1
    itemsets = {frozenset([item]) for transaction in dataset for item in transaction}
    freq_set = {itemset: calculate_support(itemset, dataset) for itemset in itemsets if
                calculate_support(itemset, dataset) >= min_support}

    # 不断生成更大的频繁项集，直到无法生成
    k = 2
    while freq_set:
        current_itemsets = set(freq_set.keys())
        f_k = {}
        # 生成候选项集
        candidate_itemsets = {i.union(j) for i in current_itemsets for j in current_itemsets if len(i.union(j)) == k}

        # 减枝
        for itemset in candidate_itemsets:
            subsets = [itemset.difference({item}) for item in itemset]
            if any(subset not in current_itemsets for subset in subsets):
                candidate_itemsets.remove(itemset)

        # 计算候选项集的支持度并筛选出频繁项集
        for itemset in candidate_itemsets:
            support = calculate_support(itemset, dataset)
            if support >= min_support:
                f_k[itemset] = support

        if not f_k:
            break
        freq_set.update(f_k)
        k += 1

    return freq_set
