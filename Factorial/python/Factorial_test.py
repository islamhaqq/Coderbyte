from .Factorial import Factorial
import unittest

class TestFactorial(unittest.TestCase):
    def test_one(self):
        self.assertEqual(Factorial(1), 1)

    def test_five(self):
        self.assertEqual(Factorial(5), 120)

    def test_fifteen(self):
        self.assertEqual(Factorial(15), 1307674368000)

if __name__ == '__main__':
    unittest.main()
