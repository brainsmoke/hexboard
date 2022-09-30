#ifndef BITBANG_H
#define BITBANG_H

/* platform constants */

#define GPIO_ODR_OFFSET (0x14)
#define GPIO_BSRR_OFFSET (0x18)

/* GPIO, BSRR */

#define CLEAR(X) (X<<16)
#define SET(X)   (X)

/* GPIO A */
#define PIN_DATA_SOUTH         (0)
#define PIN_DATA_NORTH         (4)
#define PIN_CLK_SOUTH          (1)
#define PIN_CLK_NORTH          (2)
#define PIN_LATCH              (7)
#define PIN_NOT_OUTPUT_ENABLE  (6)
#define PIN_ENABLE_HIGH        (5)
#define PIN_UART_RX            (3)
#define PIN_SELECT_A          (10)
#define PIN_SELECT_B           (9)

#define BIT_DATA_NORTH         (1<<PIN_DATA_NORTH)
#define BIT_CLK_NORTH          (1<<PIN_CLK_NORTH)
#define BIT_DATA_SOUTH         (1<<PIN_DATA_SOUTH)
#define BIT_CLK_SOUTH          (1<<PIN_CLK_SOUTH)
#define BIT_LATCH              (1<<PIN_LATCH)
#define BIT_NOT_OUTPUT_ENABLE  (1<<PIN_NOT_OUTPUT_ENABLE)
#define BIT_ENABLE_HIGH        (1<<PIN_ENABLE_HIGH)
#define BIT_SELECT_A           (1<<PIN_SELECT_A)
#define BIT_SELECT_B           (1<<PIN_SELECT_B)

#define MASK_CLK (BIT_CLK_NORTH|BIT_CLK_SOUTH)
#define MASK_DATA (BIT_DATA_NORTH|BIT_DATA_SOUTH)

/* GPIOB */

#define PIN_ENABLE_LOW  (1)

#define BIT_ENABLE_LOW  (1<<PIN_ENABLE_LOW)

/* precomputation */

#define X4(bits) (bits*0x01010101)

/* static constants (code changes needed to change this) */

#define N_ROWS               (4) /* code expects power of two */
#define N_CHANNELS           (2)
#define N_CHIPS_PER_CHANNEL  (4)
#define N_PINS_PER_CHIP      (16)

/* derived constants */

#define N_BITS_PER_CHANNEL   (N_CHIPS_PER_CHANNEL*N_PINS_PER_CHIP)
#define N_BITS_PER_ROW       (N_BITS_PER_CHANNEL*N_CHANNELS)
#define N_LEDS               (N_BITS_PER_CHANNEL*N_CHANNELS*N_ROWS)

#ifndef __ASSEMBLER__

#include <stdint.h>

/* sends 64x (max 8) bits parallel at ~1/6th the clockspeed to GPIO[0-7]
 * gpio is the GPIO base address
 * 
 * probably 340 cycles excluding call
 */
void bitbang64_clk_stm32(uint8_t *buffer, volatile uint16_t *gpio);
void bitbang64_clk_no_enable_stm32(uint8_t *buffer, volatile uint16_t *gpio);


/* equivalent to:
 *
 * uint8_t *p = precomp_buf;
 * for (i=0; i<128; i+=8)
 *     for (j=0; j<4; j++)
 *         *p++ = bit_set(frame[i+j], bit) * PIN_DATA_NORTH |
 *                bit_set(frame[i+j+4], bit) * PIN_DATA_SOUTH |
 *                nThByteLittleEndian(otherpins4x, j);
 *
 */
void precomp64(uint8_t *precomp_buf, uint8_t *frame, uint32_t bit, uint32_t otherpins4x);

void write_wait_write(volatile uint32_t *addr, uint32_t pre_data, uint32_t post_data, uint32_t cycles);

#endif

#endif // BITBANG_H
