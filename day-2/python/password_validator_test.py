import unittest

from password_validator import valid, valid_strict


def load_input():
    with open("./input") as f:
        lines = f.read().splitlines()
    return lines


class PasswordValidatorTest(unittest.TestCase):
    def test_simple_valid(self):
        input = ["1-3 a: abcde", "1-3 b: cdefg", "2-9 c: cccccccc"]
        self.assertEqual(
            valid(input),
            2
        )

    def test_input_valid(self):
        input = load_input()
        self.assertEqual(
            valid(input),
            383
        )

    def test_simple_valid_strict(self):
        input = ["1-3 a: abcde", "1-3 b: cdefg", "2-9 c: cccccccc"]
        self.assertEqual(
            valid_strict(input),
            1
        )

    def test_input_valid_strict(self):
        input = load_input()
        self.assertEqual(
            valid_strict(input),
            272
        )


if __name__ == "__main__":
    unittest.main()
