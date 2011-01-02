CC = gcc
CFLAGS = -I../bash-4.1 -I. -I../bash-4.1/include -I../bash-4.1/lib \
	-DPROGRAM='"bash"' -DCONF_HOSTTYPE='"x86_64"' \
	-DCONF_OSTYPE='"linux-gnu"' \
	-DCONF_MACHTYPE='"x86_64-unknown-linux-gnu"' \
	-DCONF_VENDOR='"unknown"' -DLOCALEDIR='"/usr/local/share/locale"' \
	-DPACKAGE='"bash"' -DSHELL -DHAVE_CONFIG_H

all: bash-min

main.o: main.c
y.tab.o: y.tab.c
xmalloc.o: xmalloc.c

bash-min:	main.o y.tab.o xmalloc.o
	$(CC) -o bash-min \
	main.o \
	y.tab.o \
	xmalloc.o \

clean:
	rm -rf *.o bash-min
