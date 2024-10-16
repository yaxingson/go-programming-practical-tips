from unittest import TestCase, main

class BaseTest(TestCase):
  def test_upper(self):
    self.assertFalse(18 == 56)


if __name__ == '__main__':
  main()
