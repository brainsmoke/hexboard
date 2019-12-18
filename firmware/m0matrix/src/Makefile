
GCC=arm-none-eabi-gcc
CFLAGS=-Wall -g -std=c99 -mlittle-endian -mcpu=cortex-m0 -march=armv6-m -mthumb -ffunction-sections -fdata-sections -I platform -I. -Os
LDFLAGS=-Wl,--gc-sections -Lplatform -T../$(LINKER_SCRIPT)

LINKER_SCRIPT=platform/stm32f030.ld
STARTUP_SRC=platform/startup_stm32f0xx.s

TARGETS=main.elf main.bin
TESTS=test/test.elf

MAIN_FLAGS=
TEST_FLAGS=


MAIN_SRCS=main.c bitbang.S util.c $(STARTUP_SRC)
TEST_SRCS=test/test.c bitbang.S util.c $(STARTUP_SRC)
HEADERS=bitbang.h util.h

.PHONY: all clean test

all:$(TARGETS)

clean:
	-rm $(TARGETS) $(TESTS) test/*.pyc

test: $(TESTS)
	arm-none-eabi-objdump -d test/test.elf | python test/test_bitbang.py

%.bin: %.elf
	arm-none-eabi-objcopy -O binary $< $@

main.elf: $(MAIN_SRCS) $(LINKER_SCRIPT) $(HEADERS)
	$(GCC) $(CFLAGS) $(LDFLAGS) $(MAIN_FLAGS) $(MAIN_SRCS) -o $@

test/test.elf: $(TEST_SRCS) $(LINKER_SCRIPT) $(HEADERS)
	$(GCC) $(CFLAGS) $(LDFLAGS) $(TEST_FLAGS) $(TEST_SRCS) -o $@

