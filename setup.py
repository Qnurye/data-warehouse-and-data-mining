from setuptools import setup, find_packages

setup(
    name='data_warehouse_and_mining',
    version='0.1',
    packages=find_packages(where='src'),
    package_dir={'': 'src'}
)
