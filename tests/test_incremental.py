import unittest

from Apriori import incremental_update, apriori


class TestIncrementalUpdate(unittest.TestCase):
    def setUp(self):
        self.D = [
            {'A', 'B', 'C'},
            {'A', 'C'},
            {'B', 'C'},
            {'A', 'B'},
        ]

        self.D_new = [
            {'A', 'B', 'C'},
            {'A', 'B'},
            {'C'},
        ]

        self.min_support = 0.5

        self.freq_set = apriori(self.D, self.min_support)

        self.f_new = apriori(self.D_new, self.min_support)

    def test_incremental_update(self):
        # 调用函数进行增量更新
        updated_freq_set = incremental_update(
            self.D, self.D_new, self.freq_set, self.min_support
        )

        # 期望结果
        expected_result = apriori(self.D + self.D_new, self.min_support)

        # 检查结果
        for itemset in expected_result:
            self.assertAlmostEqual(updated_freq_set[itemset], expected_result[itemset])


if __name__ == "__main__":
    unittest.main()
