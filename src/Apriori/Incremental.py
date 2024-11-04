from Apriori.vanilla import calculate_support


def incremental_update(
        D: list[set], D_new: list[set],
        freq_set: dict[frozenset, float], f_new: dict[frozenset, float],
        min_support: float
) -> dict[frozenset, float]:
    """
    增量更新频繁项集
    :param D: 原数据集
    :param D_new: 新增数据集
    :param freq_set: 原频繁项集
    :param f_new: 新增频繁项集
    :param min_support: 最小支持度
    :return: dict 更新后的频繁项集
    """

    total_transactions = len(D) + len(D_new)

    for itemset in f_new:
        count_d = calculate_support(itemset, D) * len(D)
        count_d_new = calculate_support(itemset, D_new) * len(D_new)
        total_count = count_d + count_d_new
        support_total = total_count / total_transactions

        if support_total >= min_support:
            freq_set[itemset] = support_total

    for itemset in list(freq_set.keys()):
        count_d_new = calculate_support(itemset, D_new) * len(D_new)
        total_count = freq_set[itemset] * len(D) + count_d_new
        support_total = total_count / total_transactions

        if support_total >= min_support:
            freq_set[itemset] = support_total
        else:
            del freq_set[itemset]

    return freq_set
