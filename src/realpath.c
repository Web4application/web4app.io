CC = gcc
CFLAGS = -Wall -O2
LOCAL_CFLAGS = 
DEFS = 
LOCAL_DEFS = 

CCFLAGS = $(DEFS) $(LOCAL_DEFS) $(LOCAL_CFLAGS) $(CFLAGS)

SHOBJ_CC = $(CC)
SHOBJ_CFLAGS = -fPIC $(CCFLAGS)
SHOBJ_LD = $(CC)
SHOBJ_LDFLAGS = -shared
SHOBJ_XLDFLAGS = 
SHOBJ_LIBS = 
SHOBJ_STATUS = 

.c.o:
	$(SHOBJ_CC) $(SHOBJ_CFLAGS) -c -o $@ $<

../libexec/nodenv-realpath.dylib: realpath.o
	$(SHOBJ_LD) $(SHOBJ_LDFLAGS) $(SHOBJ_XLDFLAGS) -o $@ realpath.o $(SHOBJ_LIBS)

clean:
	rm -f *.o ../libexec/*.dylib
