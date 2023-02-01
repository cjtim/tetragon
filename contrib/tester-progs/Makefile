
GCC ?= gcc

PROGS = sigkill-tester \
	capabilities-tester \
	namespace-tester \
	fork-tester \
	dup-tester \
	trigger-test-events \
	sigkill-unprivileged-user-ns-tester \
	nop \
	exit-leader \
	exit-tester \
	lseek-pipe

all: $(PROGS)

%: %.c
	$(GCC) -Wall $< -o $@

capabilities-tester: capabilities-tester.c
	$(GCC) -Wall $< -o $@ -lcap

exit-tester: exit-tester.c
	$(GCC) -Wall $< -o $@ -lpthread

sigkill-unprivileged-user-ns-tester: sigkill-unprivileged-user-ns-tester.c
	$(GCC) -Wall $< -o $@ -lcap

exit-leader: exit-leader.c
	$(GCC) -Wall $< -o $@ -lpthread

lseek-pipe: FORCE
	go build -o lseek-pipe ./go/lseek-pipe

.PHONY: clean
clean:
	rm -f $(PROGS)

FORCE:
