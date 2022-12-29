class Solution:
  def numWaterBottles(self, numBottles: int, numExchange: int) -> int:
    # First drink
    result = numBottles
    emptyBottles = numBottles

    # While empty bottles >= numExchange
    while emptyBottles >= numExchange:
      numBottles = emptyBottles // numExchange
      emptyBottles = emptyBottles % numExchange

      result += numBottles
      emptyBottles += numBottles

    return result
