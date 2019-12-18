#ifndef UTIL_H
#define UTIL_H

#define F_CPU (48e6)
#define F_SYS_TICK_CLK (F_CPU/8)

#define O(pin) (1<<(2*pin))
#define ALT_FN(pin) (2<<(2*pin))
#define SWD (ALT_FN(13)|ALT_FN(14))

#define OSPEED_SWD (OSPEED_HIGH(13))

#define OSPEED_LOW(pin) (0)
#define OSPEED_MEDIUM(pin) (1<<(2*pin))
#define OSPEED_HIGH(pin) (3<<(2*pin))

void clock48mhz(void);
void enable_sys_tick(uint32_t ticks);
void usart1_rx_pa3_dma3_enable(volatile uint8_t *buf, uint32_t size, long baudrate_prescale);

#endif // UTIL_H
