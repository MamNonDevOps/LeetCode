import unittest
import main

class TestNumWaterBottles(unittest.TestCase):
  s = main.Solution()

  def test_numWaterBottles(self):
    self.assertEqual(self.s.numWaterBottles(9, 3), 13, "Should be 13")

if __name__ == '__main__':
  unittest.main()
