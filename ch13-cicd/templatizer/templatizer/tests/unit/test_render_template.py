import unittest

from templatizer.templatizer import render_template


class TestRenderTemplate(unittest.TestCase):
    """Just a sample unit test to show that things are working
    """

    def runTest(self):
        with self.assertRaises(IOError):
            render_template("foo")
