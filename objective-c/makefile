client: libnsspeech.dylib
	clang spclient.c libnsspeech.dylib -o client

libnsspeech.dylib:
	clang -dynamiclib splib.m -framework AppKit -o libnsspeech.dylib

clean:
	rm *.o
	rm libnsspeech.dylib
	rm client
