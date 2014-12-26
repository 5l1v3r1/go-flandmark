.PHONY: all clean

all: flandmark/libflandmark/libflandmark_static.a flandmark_binding.o

flandmark/libflandmark/libflandmark_static.a:
	git clone https://github.com/uricamic/flandmark.git && cd flandmark && git checkout a0981a3b09cc5534255dc1dcdae2179097231bdd && cd -
	cd flandmark && cmake . && make && cd -

flandmark_binding.o:
	$(CXX) $(CXXFLAGS) -Iflandmark/libflandmark -I/usr/local/include/opencv -I/usr/include/opencv -Wall -c flandmark_binding.cpp -o flandmark_binding.o

clean:
	rm -rf flandmark
