# Lowest-Common-Ancestor
In graph theory and computer science, the lowest common ancestor (LCA) of two nodes *v* and *w* in a 
tree or directed acyclic graph (DAG) *T* is the lowest (i.e. deepest) node that has both *v* and *w* as 
descendants, where we define each node to be a descendant of itself (so if *v* has a direct connection 
from *w*, *w* is the lowest common ancestor). 

### Prerequisites
If you want to run these tests locally, you need the following installed:

- Go 1.x
- Testify - github.com/stretchr/testify/assert
- Python 3.8
- Coverage.py

### Running Golang Tests Locally
Clone the repository to your local machine.  
Open the terminal/command prompt, navigate to the cloned repository and run the following commands:
```
cd /golang/LCA
go test -cover -coverprofile=c.out
go tool cover -html=c.out -o coverage.html 
```
This should create a file called coverage.html that you can open in your browser. This html file will show code coverage
for the tests.

### Running the Python Tests Locally
Clone the repository to your local machine.  
Open the terminal/command prompt, navigate to the cloned repository and run the following commands:
```
cd python
coverage run -m pytest
```
This should create a folder called /htmlcov, which contains files called index.html and LCA_py.html. These html files will show 
code coverage for the tests. 
