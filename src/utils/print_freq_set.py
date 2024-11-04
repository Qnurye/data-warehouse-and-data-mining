def print_freq_set(frequent_itemsets: dict[frozenset, float]) -> None:
    """
    输出频繁项集及其支持度
    :param frequent_itemsets: 频繁项集及其支持度
    :return: None
    """
    for itemset, support in frequent_itemsets.items():
        items = sorted(list(itemset))
        print(f"{'{' + ', '.join(items) + '}'}: {support:.2f}")
