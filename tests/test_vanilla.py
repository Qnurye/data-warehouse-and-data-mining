import unittest
from Apriori.vanilla import calculate_support, apriori

class TestApriori(unittest.TestCase):

    def test_calculate_support_single_itemset(self):
        dataset = [{'a', 'b'}, {'b', 'c'}, {'a', 'c'}]
        itemset = frozenset(['a'])
        self.assertAlmostEqual(calculate_support(itemset, dataset), 2/3)

    def test_calculate_support_empty_dataset(self):
        dataset = []
        itemset = frozenset(['a'])
        self.assertAlmostEqual(calculate_support(itemset, dataset), 0.0)

    def test_apriori_single_itemset(self):
        dataset = [{'a', 'b'}, {'b', 'c'}, {'a', 'c'}]
        min_support = 0.5
        expected = {frozenset(['a']): 2/3, frozenset(['b']): 2/3, frozenset(['c']): 2/3}
        self.assertEqual(apriori(dataset, min_support), expected)

    def test_apriori_no_frequent_itemsets(self):
        dataset = [{'a', 'b'}, {'b', 'c'}, {'a', 'c'}]
        min_support = 1.0
        expected = {}
        self.assertEqual(apriori(dataset, min_support), expected)

    def test_apriori_multiple_itemsets(self):
        dataset = [{'a', 'b'}, {'b', 'c'}, {'a', 'c'}, {'a', 'b', 'c'}]
        min_support = 0.5
        expected = {
            frozenset(['a']): 3/4,
            frozenset(['b']): 3/4,
            frozenset(['c']): 3/4,
            frozenset(['a', 'b']): 2/4,
            frozenset(['a', 'c']): 2/4,
            frozenset(['b', 'c']): 2/4
        }
        self.assertEqual(apriori(dataset, min_support), expected)

if __name__ == '__main__':
    unittest.main()
