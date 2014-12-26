FILES=flandmark_detector.h flandmark_detector.cpp liblbp.h liblbp.cpp msvc-compat.h

.PHONY: all clean clean-sources

all: $(FILES)

flandmark:
	git clone https://github.com/uricamic/flandmark.git && cd flandmark && git checkout a0981a3b09cc5534255dc1dcdae2179097231bdd && cd -

%.h: flandmark
	cp flandmark/libflandmark/$@ .

%.cpp: flandmark
	cp flandmark/libflandmark/$@ .

clean: clean-sources
	rm -rf flandmark

clean-sources:
	rm -f $(FILES)
