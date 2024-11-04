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

    total_size = len(D) + len(D_new)

    # 更新频繁项集
    updated_freq_set = freq_set.copy()  # 复制原频繁项集

    for itemset in freq_set:
        updated_freq_set[itemset] = (freq_set[itemset] * len(D)) / total_size

    for itemset in f_new:
        if itemset in updated_freq_set:
            # 如果项集已经存在，累加支持度
            updated_freq_set[itemset] += (f_new[itemset] * len(D_new)) / total_size
        else:
            # 否则直接加入
            updated_freq_set[itemset] = (f_new[itemset] * len(D_new)) / total_size

    final_freq_set = {itemset: support for itemset, support in updated_freq_set.items() if support >= min_support}

    return final_freq_set
