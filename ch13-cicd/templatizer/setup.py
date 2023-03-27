
import setuptools

try:
    import multiprocessing
except ImportError:
    pass

setuptools.setup(
    setup_requires=['pbr>=1.8'],
    package_data={'templatizer': ['*.j2', 'templates/*.j2']},
    pbr=True)
