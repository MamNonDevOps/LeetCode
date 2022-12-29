import unittest
import main

class TestNumWaterBottles(unittest.TestCase):
  s = main.Solution()

  def test_numWaterBottles(self):
    self.assertEqual(self.s.numWaterBottles(9, 3), 13, "Should be 13")
    self.assertEqual(self.s.numWaterBottles(15, 4), 19, "Should be 19")

if __name__ == '__main__':
  unittest.main()
