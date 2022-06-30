CC := gcc
CFLAGS := -Wall -Wextra -std=c99 -pthread
CGO := go build
STATIC := -buildmode=c-archive
SHARED := -buildmode=c-shared
LIBS := build/strutil.a build/strutil.so

.PHONY: all
all: $(LIBS)

build/strutil.a: strutil.go
	$(CGO) $(STATIC) -o build/strutil.a $<

build/strutil.so: strutil.go
	$(CGO) $(SHARED) -o build/strutil.so $<

.PHONY: clean
clean:
	find build -type f \( -name '*.h' -o -name '*.so' -o -name '*.a' \) -delete
