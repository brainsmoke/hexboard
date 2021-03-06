EESchema Schematic File Version 4
LIBS:hexdump-cache
EELAYER 26 0
EELAYER END
$Descr A4 11693 8268
encoding utf-8
Sheet 1 1
Title ""
Date ""
Rev ""
Comp ""
Comment1 ""
Comment2 ""
Comment3 ""
Comment4 ""
$EndDescr
$Comp
L Connector:Conn_01x09_Female J1
U 1 1 5D40479F
P 2700 4200
F 0 "J1" H 2727 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 2727 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 2700 4200 50  0001 C CNN
F 3 "~" H 2700 4200 50  0001 C CNN
	1    2700 4200
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR04
U 1 1 5D40488F
P 2400 4700
F 0 "#PWR04" H 2400 4450 50  0001 C CNN
F 1 "GND" H 2405 4527 50  0000 C CNN
F 2 "" H 2400 4700 50  0001 C CNN
F 3 "" H 2400 4700 50  0001 C CNN
	1    2400 4700
	1    0    0    -1  
$EndComp
Wire Wire Line
	2500 3800 2300 3800
Wire Wire Line
	2400 4700 2400 4600
Wire Wire Line
	2400 4600 2500 4600
Wire Wire Line
	2400 4600 2400 4300
Wire Wire Line
	2400 4300 2500 4300
Connection ~ 2400 4600
Wire Wire Line
	2400 4300 2400 3900
Wire Wire Line
	2400 3900 2500 3900
Connection ~ 2400 4300
Wire Wire Line
	2500 4000 1950 4000
Wire Wire Line
	1950 4000 1950 3600
Wire Wire Line
	1950 4000 1950 4100
Wire Wire Line
	1950 4100 2500 4100
Connection ~ 1950 4000
Wire Wire Line
	1950 4100 1950 4200
Wire Wire Line
	1950 4200 2500 4200
Connection ~ 1950 4100
$Comp
L power:VCC #PWR01
U 1 1 5D4755C2
P 1950 3600
F 0 "#PWR01" H 1950 3450 50  0001 C CNN
F 1 "VCC" H 1967 3773 50  0000 C CNN
F 2 "" H 1950 3600 50  0001 C CNN
F 3 "" H 1950 3600 50  0001 C CNN
	1    1950 3600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J3
U 1 1 5D680335
P 3950 4200
F 0 "J3" H 3977 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 3977 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 3950 4200 50  0001 C CNN
F 3 "~" H 3950 4200 50  0001 C CNN
	1    3950 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J5
U 1 1 5D680379
P 5200 4200
F 0 "J5" H 5227 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 5227 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 5200 4200 50  0001 C CNN
F 3 "~" H 5200 4200 50  0001 C CNN
	1    5200 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J7
U 1 1 5D6803E5
P 6200 4200
F 0 "J7" H 6227 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 6227 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 6200 4200 50  0001 C CNN
F 3 "~" H 6200 4200 50  0001 C CNN
	1    6200 4200
	1    0    0    -1  
$EndComp
Wire Wire Line
	2500 4100 3750 4100
Connection ~ 2500 4100
Wire Wire Line
	3750 4100 5000 4100
Connection ~ 3750 4100
Wire Wire Line
	5000 4100 6000 4100
Connection ~ 5000 4100
Wire Wire Line
	2500 4000 3750 4000
Connection ~ 2500 4000
Wire Wire Line
	2500 4200 3750 4200
Connection ~ 2500 4200
Wire Wire Line
	2500 4300 3750 4300
Connection ~ 2500 4300
Wire Wire Line
	2500 4600 3750 4600
Connection ~ 2500 4600
Wire Wire Line
	2500 3900 3750 3900
Connection ~ 2500 3900
Wire Wire Line
	3750 3900 5000 3900
Connection ~ 3750 3900
Wire Wire Line
	5000 4000 3750 4000
Connection ~ 3750 4000
Wire Wire Line
	3750 4200 5000 4200
Connection ~ 3750 4200
Wire Wire Line
	5000 4300 3750 4300
Connection ~ 3750 4300
Wire Wire Line
	3750 4600 5000 4600
Connection ~ 3750 4600
Wire Wire Line
	5000 3900 6000 3900
Connection ~ 5000 3900
Wire Wire Line
	6000 4000 5000 4000
Connection ~ 5000 4000
Wire Wire Line
	5000 4200 6000 4200
Connection ~ 5000 4200
Wire Wire Line
	6000 4300 5000 4300
Connection ~ 5000 4300
Wire Wire Line
	5000 4600 6000 4600
Connection ~ 5000 4600
Wire Wire Line
	3750 3800 3550 3800
Wire Wire Line
	5000 3800 4800 3800
Wire Wire Line
	6000 3800 5800 3800
Text Label 2300 3800 2    50   ~ 0
uart_rx0
$Comp
L Connector:Conn_01x09_Female J2
U 1 1 5D6AE1EB
P 2700 5600
F 0 "J2" H 2727 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 2727 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 2700 5600 50  0001 C CNN
F 3 "~" H 2700 5600 50  0001 C CNN
	1    2700 5600
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR05
U 1 1 5D6AE1F2
P 2400 6100
F 0 "#PWR05" H 2400 5850 50  0001 C CNN
F 1 "GND" H 2405 5927 50  0000 C CNN
F 2 "" H 2400 6100 50  0001 C CNN
F 3 "" H 2400 6100 50  0001 C CNN
	1    2400 6100
	1    0    0    -1  
$EndComp
Wire Wire Line
	2500 5200 2300 5200
Wire Wire Line
	2400 6100 2400 6000
Wire Wire Line
	2400 6000 2500 6000
Wire Wire Line
	2400 6000 2400 5700
Wire Wire Line
	2400 5700 2500 5700
Connection ~ 2400 6000
Wire Wire Line
	2400 5700 2400 5300
Wire Wire Line
	2400 5300 2500 5300
Connection ~ 2400 5700
Wire Wire Line
	2500 5400 1950 5400
Wire Wire Line
	1950 5400 1950 5000
Wire Wire Line
	1950 5400 1950 5500
Wire Wire Line
	1950 5500 2500 5500
Connection ~ 1950 5400
Wire Wire Line
	1950 5500 1950 5600
Wire Wire Line
	1950 5600 2500 5600
Connection ~ 1950 5500
$Comp
L power:VCC #PWR02
U 1 1 5D6AE20D
P 1950 5000
F 0 "#PWR02" H 1950 4850 50  0001 C CNN
F 1 "VCC" H 1967 5173 50  0000 C CNN
F 2 "" H 1950 5000 50  0001 C CNN
F 3 "" H 1950 5000 50  0001 C CNN
	1    1950 5000
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J4
U 1 1 5D6AE213
P 3950 5600
F 0 "J4" H 3977 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 3977 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 3950 5600 50  0001 C CNN
F 3 "~" H 3950 5600 50  0001 C CNN
	1    3950 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J6
U 1 1 5D6AE219
P 5200 5600
F 0 "J6" H 5227 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 5227 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 5200 5600 50  0001 C CNN
F 3 "~" H 5200 5600 50  0001 C CNN
	1    5200 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J8
U 1 1 5D6AE21F
P 6200 5600
F 0 "J8" H 6227 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 6227 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 6200 5600 50  0001 C CNN
F 3 "~" H 6200 5600 50  0001 C CNN
	1    6200 5600
	1    0    0    -1  
$EndComp
Wire Wire Line
	2500 5500 3750 5500
Connection ~ 2500 5500
Wire Wire Line
	3750 5500 5000 5500
Connection ~ 3750 5500
Wire Wire Line
	5000 5500 6000 5500
Connection ~ 5000 5500
Wire Wire Line
	2500 5400 3750 5400
Connection ~ 2500 5400
Wire Wire Line
	2500 5600 3750 5600
Connection ~ 2500 5600
Wire Wire Line
	2500 5700 3750 5700
Connection ~ 2500 5700
Wire Wire Line
	2500 6000 3750 6000
Connection ~ 2500 6000
Wire Wire Line
	2500 5300 3750 5300
Connection ~ 2500 5300
Wire Wire Line
	3750 5300 5000 5300
Connection ~ 3750 5300
Wire Wire Line
	5000 5400 3750 5400
Connection ~ 3750 5400
Wire Wire Line
	3750 5600 5000 5600
Connection ~ 3750 5600
Wire Wire Line
	5000 5700 3750 5700
Connection ~ 3750 5700
Wire Wire Line
	3750 6000 5000 6000
Connection ~ 3750 6000
Wire Wire Line
	5000 5300 6000 5300
Connection ~ 5000 5300
Wire Wire Line
	6000 5400 5000 5400
Connection ~ 5000 5400
Wire Wire Line
	5000 5600 6000 5600
Connection ~ 5000 5600
Wire Wire Line
	6000 5700 5000 5700
Connection ~ 5000 5700
Wire Wire Line
	5000 6000 6000 6000
Connection ~ 5000 6000
Wire Wire Line
	3750 5200 3550 5200
Wire Wire Line
	5000 5200 4800 5200
Wire Wire Line
	6000 5200 5800 5200
Wire Wire Line
	2250 4500 2500 4500
$Comp
L power:GND #PWR03
U 1 1 5D650DF9
P 2250 4700
F 0 "#PWR03" H 2250 4450 50  0001 C CNN
F 1 "GND" H 2255 4527 50  0000 C CNN
F 2 "" H 2250 4700 50  0001 C CNN
F 3 "" H 2250 4700 50  0001 C CNN
	1    2250 4700
	1    0    0    -1  
$EndComp
Wire Wire Line
	2250 4500 2250 4700
Wire Wire Line
	4750 4500 5000 4500
$Comp
L power:GND #PWR07
U 1 1 5D65A927
P 4750 4700
F 0 "#PWR07" H 4750 4450 50  0001 C CNN
F 1 "GND" H 4755 4527 50  0000 C CNN
F 2 "" H 4750 4700 50  0001 C CNN
F 3 "" H 4750 4700 50  0001 C CNN
	1    4750 4700
	1    0    0    -1  
$EndComp
Wire Wire Line
	4750 4500 4750 4700
Wire Wire Line
	3500 5900 3750 5900
$Comp
L power:GND #PWR06
U 1 1 5D682A4B
P 3500 6100
F 0 "#PWR06" H 3500 5850 50  0001 C CNN
F 1 "GND" H 3505 5927 50  0000 C CNN
F 2 "" H 3500 6100 50  0001 C CNN
F 3 "" H 3500 6100 50  0001 C CNN
	1    3500 6100
	1    0    0    -1  
$EndComp
Wire Wire Line
	3500 5900 3500 6100
Wire Wire Line
	5750 5900 6000 5900
$Comp
L power:GND #PWR08
U 1 1 5D68D0D3
P 5750 6100
F 0 "#PWR08" H 5750 5850 50  0001 C CNN
F 1 "GND" H 5755 5927 50  0000 C CNN
F 2 "" H 5750 6100 50  0001 C CNN
F 3 "" H 5750 6100 50  0001 C CNN
	1    5750 6100
	1    0    0    -1  
$EndComp
Wire Wire Line
	5750 5900 5750 6100
Text Label 5800 3800 2    50   ~ 0
uart_rx1
Text Label 4800 3800 2    50   ~ 0
uart_rx1
Text Label 3550 3800 2    50   ~ 0
uart_rx0
Text Label 7900 4800 2    50   ~ 0
uart_rx0
Text Label 7900 4900 2    50   ~ 0
uart_rx1
Wire Wire Line
	8150 4800 7900 4800
$Comp
L Connector:Conn_01x02_Female J9
U 1 1 5D73B508
P 8350 4800
F 0 "J9" H 8377 4776 50  0000 L CNN
F 1 "Conn_01x02_Female" H 8377 4685 50  0000 L CNN
F 2 "footprints:header-1x2" H 8350 4800 50  0001 C CNN
F 3 "~" H 8350 4800 50  0001 C CNN
	1    8350 4800
	1    0    0    -1  
$EndComp
Wire Wire Line
	8150 4900 7900 4900
Text Label 2300 5200 2    50   ~ 0
uart_rx0
Text Label 5800 5200 2    50   ~ 0
uart_rx1
Text Label 4800 5200 2    50   ~ 0
uart_rx1
Text Label 3550 5200 2    50   ~ 0
uart_rx0
$EndSCHEMATC
