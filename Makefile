.PHONY: all clean

all: flandmark/libflandmark/libflandmark_static.a

flandmark/libflandmark/libflandmark_static.a:
	git clone https://github.com/uricamic/flandmark.git && cd flandmark && git checkout a0981a3b09cc5534255dc1dcdae2179097231bdd && cd -
	cd flandmark && cmake . && make && cd -

clean:
	rm -rf flandmark
