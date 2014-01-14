#Right now the makefile only for linux
VER=0.1
CC=g++

FILES = Game.cpp pyramid.cpp

FLAGS := -Wall -O 

#OS = MSDOS
OS = LINUX

OBJECTS := $(FILES: .cpp=.o)

all: $(OBJECTS)
	$(CC) $(FLAGS) -D$(OS) -o pyramid $(OBJECTS)

clean:
	rm -f *.o pyramid *.tgz
D=pyramid
tar:
	tar -C .. -zcf pyramid-$(VER).tgz $D/Game.cpp $D/Game.h $D/pyramid.cpp $D/README $D/Makefile $D/wins.txt 
	

