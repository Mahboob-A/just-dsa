# Addition without plus (+) or minus (-)
Given two integers `a` and `b`, return the sum of the two integers without using the operators `+` and `-`.

## Intuition:
We observe that the actual sum is formed by XORing the digits, while the carry is obtained by performing bitwise AND operations and left-shifting it by one. We continue this process until the carry is zero.

## Approach:
1. Initialize variables `sum` and `carry` to store the intermediate sum and carry during the addition process.
2. Iterate until there's no carry left (i.e., `b != 0`):
   - Calculate the sum without considering the carry by performing a bitwise XOR operation (`a ^ b`).
   - Compute the carry generated by adding the corresponding bits of `a` and `b` using a bitwise AND operation followed by a left shift by one (`(a & b) << 1`).
   - Update `a` to the computed sum and `b` to the calculated carry shifted by one for the next iteration.
3. Once the carry becomes zero, indicating that there are no more bits to carry, return the final sum stored in `a`.

## Time Complexity: O(log(max(a, b))) - The number of iterations is proportional to the number of bits in the larger of the two numbers.
## Space Complexity: O(1) - Constant space is used regardless of the input size.