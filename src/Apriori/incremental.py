from Apriori import apriori


def incremental_update(
        DB: list[set], db: list[set],
        FP: dict[frozenset, float],
        s: float
) -> dict[frozenset, float]:
    """
    增量更新频繁项集
    :param DB: 原数据集
    :param db: 新增数据集
    :param FP: 原频繁项集
    :param s: 最小支持度
    :return: dict 更新后的频繁项集
    """

    fp = apriori(db, s)  # 新增数据集的频繁项集
    total_size = len(DB) + len(db)

    # c1 为 FP 和 fp 的键的交集
    c1 = set(FP.keys()) & set(fp.keys())
    # c2 = FP - c1
    c2 = set(FP.keys()) - c1
    # c3 = fp - c1
    c3 = set(fp.keys()) - c1

    for k in c1:
        FP[k] = (FP[k] * len(DB) + fp[k] * len(db)) / total_size

    count = 0
    for t in db:
        for k in c2:
            if k.issubset(t):
                count += 1
    FP.update({k: (FP[k] * len(DB) + count) / total_size for k in c2})

    count = 0
    for t in db:
        for k in c3:
            if k.issubset(t):
                count += 1
    FP.update({k: (fp[k] * len(db)) / total_size for k in c3})

    return FP
