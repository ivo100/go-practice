NPY_INC = $(shell python -c 'import numpy; print(numpy.get_include())')

test:
	PYTHONPATH=$(PWD) CGO_CFLAGS="-I $(NPY_INC)" CGO_LDFLAGS="-L /opt/homebrew/Cellar/python@3.11/3.11.9/Frameworks/Python.framework/Versions/3.11/lib" \
		go test -v

bench:
	PYTHONPATH=$(PWD) CGO_CFLAGS="-I $(NPY_INC)" \
	       go test -run NONE -bench .
