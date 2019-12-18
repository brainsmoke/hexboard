#include <stdlib.h>
#include "stm32f0xx.h"

#include "util.h"

void enable_sys_tick(uint32_t ticks)
{
	SysTick->LOAD = ticks - 1;
	SysTick->VAL = 0;
	SysTick->CTRL = SysTick_CTRL_ENABLE_Msk|SysTick_CTRL_TICKINT_Msk;
}

void clock48mhz(void)
{
	/* 1 wait state required for flash accesses at clockspeeds > 24MHz */
	/* enable prefetching */
	FLASH->ACR |= (1 * FLASH_ACR_LATENCY)  |  FLASH_ACR_PRFTBE;

	/* HSI is 8MHz, PLL has a 2x divider & config a 12x multplier -> 48MHz */
	RCC->CFGR = ( RCC->CFGR & ~RCC_CFGR_PLLMULL ) | RCC_CFGR_PLLMULL12;

	/* turn on PLL */
	RCC->CR |= RCC_CR_PLLON;
	while( !(RCC->CR & RCC_CR_PLLRDY) );

	/* switch to PLL for system clock */
	RCC->CFGR |= RCC_CFGR_SW_PLL;
	while ((RCC->CFGR & RCC_CFGR_SWS) != RCC_CFGR_SWS_PLL);
}

void usart1_rx_pa3_dma3_enable(volatile uint8_t *buf, uint32_t size, long baudrate_prescale)
{
	GPIOA->MODER |= ALT_FN(3); /* alternate function mode for PA3 */
	GPIOA->AFR[0] |= 1 << 12; /* mux PA3 to usart1_rx */
	RCC->AHBENR |= RCC_AHBENR_DMA1EN;
	RCC->APB2ENR |= RCC_APB2ENR_USART1EN;

	DMA1_Channel3->CPAR = (uint32_t)&USART1->RDR;
	DMA1_Channel3->CMAR = (uint32_t)buf;
	DMA1_Channel3->CNDTR = size;
	DMA1_Channel3->CCR = DMA_CCR_MINC | DMA_CCR_CIRC | (0*DMA_CCR_MSIZE_0) | (0*DMA_CCR_PSIZE_0);

	USART1->CR1 = 0;
	USART1->BRR = baudrate_prescale;
	USART1->CR3 = USART_CR3_DMAR;
	USART1->CR1 |= USART_CR1_RE | USART_CR1_UE;
	/* enable dma on usart1_rx */
	DMA1_Channel3->CCR |= DMA_CCR_EN;
}

