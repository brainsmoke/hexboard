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
$Comp
L Connector:Conn_01x09_Female J9
U 1 1 5D6805C4
P 7250 4200
F 0 "J9" H 7277 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 7277 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 7250 4200 50  0001 C CNN
F 3 "~" H 7250 4200 50  0001 C CNN
	1    7250 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J11
U 1 1 5D6805CA
P 8500 4200
F 0 "J11" H 8527 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 8527 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 8500 4200 50  0001 C CNN
F 3 "~" H 8500 4200 50  0001 C CNN
	1    8500 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J13
U 1 1 5D6805D0
P 9750 4200
F 0 "J13" H 9777 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 9777 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 9750 4200 50  0001 C CNN
F 3 "~" H 9750 4200 50  0001 C CNN
	1    9750 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J15
U 1 1 5D6805D6
P 10750 4200
F 0 "J15" H 10777 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 10777 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 10750 4200 50  0001 C CNN
F 3 "~" H 10750 4200 50  0001 C CNN
	1    10750 4200
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
	6000 4100 7050 4100
Connection ~ 6000 4100
Wire Wire Line
	7050 4100 8300 4100
Connection ~ 7050 4100
Wire Wire Line
	8300 4100 9550 4100
Connection ~ 8300 4100
Wire Wire Line
	9550 4100 10550 4100
Connection ~ 9550 4100
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
	6000 4600 7050 4600
Connection ~ 6000 4600
Wire Wire Line
	6000 4300 7050 4300
Connection ~ 6000 4300
Wire Wire Line
	6000 4200 7050 4200
Connection ~ 6000 4200
Wire Wire Line
	6000 4000 7050 4000
Connection ~ 6000 4000
Wire Wire Line
	6000 3900 7050 3900
Connection ~ 6000 3900
Wire Wire Line
	7050 3900 8300 3900
Connection ~ 7050 3900
Wire Wire Line
	8300 4000 7050 4000
Connection ~ 7050 4000
Wire Wire Line
	7050 4200 8300 4200
Connection ~ 7050 4200
Wire Wire Line
	8300 4300 7050 4300
Connection ~ 7050 4300
Wire Wire Line
	7050 4600 8300 4600
Connection ~ 7050 4600
Wire Wire Line
	8300 3900 9550 3900
Connection ~ 8300 3900
Wire Wire Line
	9550 4000 8300 4000
Connection ~ 8300 4000
Wire Wire Line
	8300 4200 9550 4200
Connection ~ 8300 4200
Wire Wire Line
	9550 4300 8300 4300
Connection ~ 8300 4300
Wire Wire Line
	8300 4600 9550 4600
Connection ~ 8300 4600
Wire Wire Line
	9550 3900 10550 3900
Connection ~ 9550 3900
Wire Wire Line
	10550 4000 9550 4000
Connection ~ 9550 4000
Wire Wire Line
	9550 4200 10550 4200
Connection ~ 9550 4200
Wire Wire Line
	10550 4300 9550 4300
Connection ~ 9550 4300
Wire Wire Line
	9550 4600 10550 4600
Connection ~ 9550 4600
Wire Wire Line
	3750 3800 3550 3800
Wire Wire Line
	5000 3800 4800 3800
Wire Wire Line
	6000 3800 5800 3800
Text Label 2300 3800 2    50   ~ 0
uart_rx2
Wire Wire Line
	7050 3800 6850 3800
Wire Wire Line
	8300 3800 8100 3800
Wire Wire Line
	9550 3800 9350 3800
Wire Wire Line
	10550 3800 10350 3800
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
Text Label 2300 5200 2    50   ~ 0
uart_rx2
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
$Comp
L Connector:Conn_01x09_Female J10
U 1 1 5D6AE225
P 7250 5600
F 0 "J10" H 7277 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 7277 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 7250 5600 50  0001 C CNN
F 3 "~" H 7250 5600 50  0001 C CNN
	1    7250 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J12
U 1 1 5D6AE22B
P 8500 5600
F 0 "J12" H 8527 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 8527 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 8500 5600 50  0001 C CNN
F 3 "~" H 8500 5600 50  0001 C CNN
	1    8500 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J14
U 1 1 5D6AE231
P 9750 5600
F 0 "J14" H 9777 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 9777 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 9750 5600 50  0001 C CNN
F 3 "~" H 9750 5600 50  0001 C CNN
	1    9750 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J16
U 1 1 5D6AE237
P 10750 5600
F 0 "J16" H 10777 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 10777 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 10750 5600 50  0001 C CNN
F 3 "~" H 10750 5600 50  0001 C CNN
	1    10750 5600
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
	6000 5500 7050 5500
Connection ~ 6000 5500
Wire Wire Line
	7050 5500 8300 5500
Connection ~ 7050 5500
Wire Wire Line
	8300 5500 9550 5500
Connection ~ 8300 5500
Wire Wire Line
	9550 5500 10550 5500
Connection ~ 9550 5500
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
	6000 6000 7050 6000
Connection ~ 6000 6000
Wire Wire Line
	6000 5700 7050 5700
Connection ~ 6000 5700
Wire Wire Line
	6000 5600 7050 5600
Connection ~ 6000 5600
Wire Wire Line
	6000 5400 7050 5400
Connection ~ 6000 5400
Wire Wire Line
	6000 5300 7050 5300
Connection ~ 6000 5300
Wire Wire Line
	7050 5300 8300 5300
Connection ~ 7050 5300
Wire Wire Line
	8300 5400 7050 5400
Connection ~ 7050 5400
Wire Wire Line
	7050 5600 8300 5600
Connection ~ 7050 5600
Wire Wire Line
	8300 5700 7050 5700
Connection ~ 7050 5700
Wire Wire Line
	7050 6000 8300 6000
Connection ~ 7050 6000
Wire Wire Line
	8300 5300 9550 5300
Connection ~ 8300 5300
Wire Wire Line
	9550 5400 8300 5400
Connection ~ 8300 5400
Wire Wire Line
	8300 5600 9550 5600
Connection ~ 8300 5600
Wire Wire Line
	9550 5700 8300 5700
Connection ~ 8300 5700
Wire Wire Line
	8300 6000 9550 6000
Connection ~ 8300 6000
Wire Wire Line
	9550 5300 10550 5300
Connection ~ 9550 5300
Wire Wire Line
	10550 5400 9550 5400
Connection ~ 9550 5400
Wire Wire Line
	9550 5600 10550 5600
Connection ~ 9550 5600
Wire Wire Line
	10550 5700 9550 5700
Connection ~ 9550 5700
Wire Wire Line
	9550 6000 10550 6000
Connection ~ 9550 6000
Wire Wire Line
	3750 5200 3550 5200
Wire Wire Line
	5000 5200 4800 5200
Wire Wire Line
	6000 5200 5800 5200
$Comp
L Connector:Conn_01x04_Female J25
U 1 1 5D6BB480
P 17700 5000
F 0 "J25" H 17727 4976 50  0000 L CNN
F 1 "Conn_01x04_Female" H 17727 4885 50  0000 L CNN
F 2 "footprints:header-1x4" H 17700 5000 50  0001 C CNN
F 3 "~" H 17700 5000 50  0001 C CNN
	1    17700 5000
	1    0    0    -1  
$EndComp
Wire Wire Line
	18250 4900 18150 4900
Wire Wire Line
	18150 4900 18150 5000
Wire Wire Line
	18150 5200 18250 5200
Wire Wire Line
	18250 5000 18150 5000
Connection ~ 18150 5000
Wire Wire Line
	18150 5000 18150 5100
Wire Wire Line
	18250 5100 18150 5100
Connection ~ 18150 5100
Wire Wire Line
	18150 5100 18150 5200
Wire Wire Line
	18150 5200 18150 5500
Connection ~ 18150 5200
$Comp
L Connector:Conn_01x04_Female J26
U 1 1 5D6D7309
P 18450 5000
F 0 "J26" H 18477 4976 50  0000 L CNN
F 1 "Conn_01x04_Female" H 18477 4885 50  0000 L CNN
F 2 "footprints:header-1x4" H 18450 5000 50  0001 C CNN
F 3 "~" H 18450 5000 50  0001 C CNN
	1    18450 5000
	1    0    0    -1  
$EndComp
Wire Wire Line
	18850 4900 18750 4900
Wire Wire Line
	18750 4900 18750 5000
Wire Wire Line
	18750 5200 18850 5200
Wire Wire Line
	18850 5000 18750 5000
Connection ~ 18750 5000
Wire Wire Line
	18750 5000 18750 5100
Wire Wire Line
	18850 5100 18750 5100
Connection ~ 18750 5100
Wire Wire Line
	18750 5100 18750 5200
Wire Wire Line
	18750 5200 18750 5500
Connection ~ 18750 5200
$Comp
L Connector:Conn_01x04_Female J27
U 1 1 5D6D73DE
P 19050 5000
F 0 "J27" H 19077 4976 50  0000 L CNN
F 1 "Conn_01x04_Female" H 19077 4885 50  0000 L CNN
F 2 "footprints:header-1x4" H 19050 5000 50  0001 C CNN
F 3 "~" H 19050 5000 50  0001 C CNN
	1    19050 5000
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x04_Female J28
U 1 1 5D6DEEA9
P 19850 5000
F 0 "J28" H 19877 4976 50  0000 L CNN
F 1 "Conn_01x04_Female" H 19877 4885 50  0000 L CNN
F 2 "footprints:header-1x4" H 19850 5000 50  0001 C CNN
F 3 "~" H 19850 5000 50  0001 C CNN
	1    19850 5000
	1    0    0    -1  
$EndComp
Text Label 17250 4900 2    50   ~ 0
uart_rx0
Text Label 17250 5000 2    50   ~ 0
uart_rx1
Text Label 17250 5100 2    50   ~ 0
uart_rx2
Text Label 17250 5200 2    50   ~ 0
uart_rx3
Wire Wire Line
	17250 4900 17500 4900
Wire Wire Line
	17500 5000 17250 5000
Wire Wire Line
	17250 5100 17500 5100
Wire Wire Line
	17500 5200 17250 5200
Text Label 19400 4900 2    50   ~ 0
uart_rx4
Text Label 19400 5000 2    50   ~ 0
uart_rx5
Text Label 19400 5100 2    50   ~ 0
uart_rx6
Text Label 19400 5200 2    50   ~ 0
uart_rx7
Wire Wire Line
	19400 4900 19650 4900
Wire Wire Line
	19650 5000 19400 5000
Wire Wire Line
	19400 5100 19650 5100
Wire Wire Line
	19650 5200 19400 5200
$Comp
L power:GND #PWR017
U 1 1 5D75DFFA
P 18150 5500
F 0 "#PWR017" H 18150 5250 50  0001 C CNN
F 1 "GND" H 18155 5327 50  0000 C CNN
F 2 "" H 18150 5500 50  0001 C CNN
F 3 "" H 18150 5500 50  0001 C CNN
	1    18150 5500
	1    0    0    -1  
$EndComp
$Comp
L power:GND #PWR018
U 1 1 5D75E083
P 18750 5500
F 0 "#PWR018" H 18750 5250 50  0001 C CNN
F 1 "GND" H 18755 5327 50  0000 C CNN
F 2 "" H 18750 5500 50  0001 C CNN
F 3 "" H 18750 5500 50  0001 C CNN
	1    18750 5500
	1    0    0    -1  
$EndComp
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
	10300 5900 10550 5900
$Comp
L power:GND #PWR012
U 1 1 5D664754
P 10300 6100
F 0 "#PWR012" H 10300 5850 50  0001 C CNN
F 1 "GND" H 10305 5927 50  0000 C CNN
F 2 "" H 10300 6100 50  0001 C CNN
F 3 "" H 10300 6100 50  0001 C CNN
	1    10300 6100
	1    0    0    -1  
$EndComp
Wire Wire Line
	10300 5900 10300 6100
Wire Wire Line
	6800 4500 7050 4500
$Comp
L power:GND #PWR09
U 1 1 5D66E5BC
P 6800 4700
F 0 "#PWR09" H 6800 4450 50  0001 C CNN
F 1 "GND" H 6805 4527 50  0000 C CNN
F 2 "" H 6800 4700 50  0001 C CNN
F 3 "" H 6800 4700 50  0001 C CNN
	1    6800 4700
	1    0    0    -1  
$EndComp
Wire Wire Line
	6800 4500 6800 4700
Wire Wire Line
	9300 4500 9550 4500
$Comp
L power:GND #PWR011
U 1 1 5D6786CF
P 9300 4700
F 0 "#PWR011" H 9300 4450 50  0001 C CNN
F 1 "GND" H 9305 4527 50  0000 C CNN
F 2 "" H 9300 4700 50  0001 C CNN
F 3 "" H 9300 4700 50  0001 C CNN
	1    9300 4700
	1    0    0    -1  
$EndComp
Wire Wire Line
	9300 4500 9300 4700
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
Wire Wire Line
	8050 5900 8300 5900
$Comp
L power:GND #PWR010
U 1 1 5D697A55
P 8050 6100
F 0 "#PWR010" H 8050 5850 50  0001 C CNN
F 1 "GND" H 8055 5927 50  0000 C CNN
F 2 "" H 8050 6100 50  0001 C CNN
F 3 "" H 8050 6100 50  0001 C CNN
	1    8050 6100
	1    0    0    -1  
$EndComp
Wire Wire Line
	8050 5900 8050 6100
$Comp
L Connector:Conn_01x09_Female J17
U 1 1 5D67067E
P 11800 4200
F 0 "J17" H 11827 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 11827 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 11800 4200 50  0001 C CNN
F 3 "~" H 11800 4200 50  0001 C CNN
	1    11800 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J19
U 1 1 5D670684
P 13050 4200
F 0 "J19" H 13077 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 13077 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 13050 4200 50  0001 C CNN
F 3 "~" H 13050 4200 50  0001 C CNN
	1    13050 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J21
U 1 1 5D67068A
P 14300 4200
F 0 "J21" H 14327 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 14327 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 14300 4200 50  0001 C CNN
F 3 "~" H 14300 4200 50  0001 C CNN
	1    14300 4200
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J23
U 1 1 5D670690
P 15300 4200
F 0 "J23" H 15327 4226 50  0000 L CNN
F 1 "Conn_01x07_Female" H 15327 4135 50  0000 L CNN
F 2 "footprints:header-1x9" H 15300 4200 50  0001 C CNN
F 3 "~" H 15300 4200 50  0001 C CNN
	1    15300 4200
	1    0    0    -1  
$EndComp
Wire Wire Line
	10550 4100 11600 4100
Wire Wire Line
	11600 4100 12850 4100
Connection ~ 11600 4100
Wire Wire Line
	12850 4100 14100 4100
Connection ~ 12850 4100
Wire Wire Line
	14100 4100 15100 4100
Connection ~ 14100 4100
Wire Wire Line
	10550 4600 11600 4600
Wire Wire Line
	10550 4300 11600 4300
Wire Wire Line
	10550 4200 11600 4200
Wire Wire Line
	10550 4000 11600 4000
Wire Wire Line
	10550 3900 11600 3900
Wire Wire Line
	11600 3900 12850 3900
Connection ~ 11600 3900
Wire Wire Line
	12850 4000 11600 4000
Connection ~ 11600 4000
Wire Wire Line
	11600 4200 12850 4200
Connection ~ 11600 4200
Wire Wire Line
	12850 4300 11600 4300
Connection ~ 11600 4300
Wire Wire Line
	11600 4600 12850 4600
Connection ~ 11600 4600
Wire Wire Line
	12850 3900 14100 3900
Connection ~ 12850 3900
Wire Wire Line
	14100 4000 12850 4000
Connection ~ 12850 4000
Wire Wire Line
	12850 4200 14100 4200
Connection ~ 12850 4200
Wire Wire Line
	14100 4300 12850 4300
Connection ~ 12850 4300
Wire Wire Line
	12850 4600 14100 4600
Connection ~ 12850 4600
Wire Wire Line
	14100 3900 15100 3900
Connection ~ 14100 3900
Wire Wire Line
	15100 4000 14100 4000
Connection ~ 14100 4000
Wire Wire Line
	14100 4200 15100 4200
Connection ~ 14100 4200
Wire Wire Line
	15100 4300 14100 4300
Connection ~ 14100 4300
Wire Wire Line
	14100 4600 15100 4600
Connection ~ 14100 4600
Wire Wire Line
	11600 3800 11400 3800
Wire Wire Line
	12850 3800 12650 3800
Text Label 9350 3800 2    50   ~ 0
uart_rx5
Wire Wire Line
	14100 3800 13900 3800
Wire Wire Line
	15100 3800 14900 3800
Wire Wire Line
	11350 4500 11600 4500
$Comp
L power:GND #PWR013
U 1 1 5D6706C9
P 11350 4700
F 0 "#PWR013" H 11350 4450 50  0001 C CNN
F 1 "GND" H 11355 4527 50  0000 C CNN
F 2 "" H 11350 4700 50  0001 C CNN
F 3 "" H 11350 4700 50  0001 C CNN
	1    11350 4700
	1    0    0    -1  
$EndComp
Wire Wire Line
	11350 4500 11350 4700
Wire Wire Line
	13850 4500 14100 4500
$Comp
L power:GND #PWR015
U 1 1 5D6706D1
P 13850 4700
F 0 "#PWR015" H 13850 4450 50  0001 C CNN
F 1 "GND" H 13855 4527 50  0000 C CNN
F 2 "" H 13850 4700 50  0001 C CNN
F 3 "" H 13850 4700 50  0001 C CNN
	1    13850 4700
	1    0    0    -1  
$EndComp
Wire Wire Line
	13850 4500 13850 4700
Connection ~ 10550 3900
Connection ~ 10550 4000
Connection ~ 10550 4100
Connection ~ 10550 4200
Connection ~ 10550 4300
Connection ~ 10550 4600
$Comp
L Connector:Conn_01x09_Female J18
U 1 1 5D67ED1E
P 11800 5600
F 0 "J18" H 11827 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 11827 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 11800 5600 50  0001 C CNN
F 3 "~" H 11800 5600 50  0001 C CNN
	1    11800 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J20
U 1 1 5D67ED24
P 13050 5600
F 0 "J20" H 13077 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 13077 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 13050 5600 50  0001 C CNN
F 3 "~" H 13050 5600 50  0001 C CNN
	1    13050 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J22
U 1 1 5D67ED2A
P 14300 5600
F 0 "J22" H 14327 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 14327 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 14300 5600 50  0001 C CNN
F 3 "~" H 14300 5600 50  0001 C CNN
	1    14300 5600
	1    0    0    -1  
$EndComp
$Comp
L Connector:Conn_01x09_Female J24
U 1 1 5D67ED30
P 15300 5600
F 0 "J24" H 15327 5626 50  0000 L CNN
F 1 "Conn_01x07_Female" H 15327 5535 50  0000 L CNN
F 2 "footprints:header-1x9" H 15300 5600 50  0001 C CNN
F 3 "~" H 15300 5600 50  0001 C CNN
	1    15300 5600
	1    0    0    -1  
$EndComp
Wire Wire Line
	10550 5500 11600 5500
Wire Wire Line
	11600 5500 12850 5500
Connection ~ 11600 5500
Wire Wire Line
	12850 5500 14100 5500
Connection ~ 12850 5500
Wire Wire Line
	14100 5500 15100 5500
Connection ~ 14100 5500
Wire Wire Line
	10550 6000 11600 6000
Wire Wire Line
	10550 5700 11600 5700
Wire Wire Line
	10550 5600 11600 5600
Wire Wire Line
	10550 5400 11600 5400
Wire Wire Line
	10550 5300 11600 5300
Wire Wire Line
	11600 5300 12850 5300
Connection ~ 11600 5300
Wire Wire Line
	12850 5400 11600 5400
Connection ~ 11600 5400
Wire Wire Line
	11600 5600 12850 5600
Connection ~ 11600 5600
Wire Wire Line
	12850 5700 11600 5700
Connection ~ 11600 5700
Wire Wire Line
	11600 6000 12850 6000
Connection ~ 11600 6000
Wire Wire Line
	12850 5300 14100 5300
Connection ~ 12850 5300
Wire Wire Line
	14100 5400 12850 5400
Connection ~ 12850 5400
Wire Wire Line
	12850 5600 14100 5600
Connection ~ 12850 5600
Wire Wire Line
	14100 5700 12850 5700
Connection ~ 12850 5700
Wire Wire Line
	12850 6000 14100 6000
Connection ~ 12850 6000
Wire Wire Line
	14100 5300 15100 5300
Connection ~ 14100 5300
Wire Wire Line
	15100 5400 14100 5400
Connection ~ 14100 5400
Wire Wire Line
	14100 5600 15100 5600
Connection ~ 14100 5600
Wire Wire Line
	15100 5700 14100 5700
Connection ~ 14100 5700
Wire Wire Line
	14100 6000 15100 6000
Connection ~ 14100 6000
Wire Wire Line
	14850 5900 15100 5900
$Comp
L power:GND #PWR016
U 1 1 5D67ED69
P 14850 6100
F 0 "#PWR016" H 14850 5850 50  0001 C CNN
F 1 "GND" H 14855 5927 50  0000 C CNN
F 2 "" H 14850 6100 50  0001 C CNN
F 3 "" H 14850 6100 50  0001 C CNN
	1    14850 6100
	1    0    0    -1  
$EndComp
Wire Wire Line
	14850 5900 14850 6100
Wire Wire Line
	12600 5900 12850 5900
$Comp
L power:GND #PWR014
U 1 1 5D67ED71
P 12600 6100
F 0 "#PWR014" H 12600 5850 50  0001 C CNN
F 1 "GND" H 12605 5927 50  0000 C CNN
F 2 "" H 12600 6100 50  0001 C CNN
F 3 "" H 12600 6100 50  0001 C CNN
	1    12600 6100
	1    0    0    -1  
$EndComp
Wire Wire Line
	12600 5900 12600 6100
Connection ~ 10550 5300
Connection ~ 10550 5400
Connection ~ 10550 5500
Connection ~ 10550 5600
Connection ~ 10550 5700
Connection ~ 10550 6000
Text Label 6850 3800 2    50   ~ 0
uart_rx4
Text Label 5800 3800 2    50   ~ 0
uart_rx3
Text Label 4800 3800 2    50   ~ 0
uart_rx3
Text Label 3550 3800 2    50   ~ 0
uart_rx2
Text Label 8100 3800 2    50   ~ 0
uart_rx4
Text Label 10350 3800 2    50   ~ 0
uart_rx5
Text Label 11400 3800 2    50   ~ 0
uart_rx6
Text Label 12650 3800 2    50   ~ 0
uart_rx6
Text Label 13900 3800 2    50   ~ 0
uart_rx7
Text Label 14900 3800 2    50   ~ 0
uart_rx7
Wire Wire Line
	7050 5200 6850 5200
Wire Wire Line
	8300 5200 8100 5200
Wire Wire Line
	9550 5200 9350 5200
Wire Wire Line
	10550 5200 10350 5200
Wire Wire Line
	11600 5200 11400 5200
Wire Wire Line
	12850 5200 12650 5200
Wire Wire Line
	14100 5200 13900 5200
Wire Wire Line
	15100 5200 14900 5200
Text Label 9350 5200 2    50   ~ 0
uart_rx5
Text Label 6850 5200 2    50   ~ 0
uart_rx4
Text Label 5800 5200 2    50   ~ 0
uart_rx3
Text Label 4800 5200 2    50   ~ 0
uart_rx3
Text Label 3550 5200 2    50   ~ 0
uart_rx2
Text Label 8100 5200 2    50   ~ 0
uart_rx4
Text Label 10350 5200 2    50   ~ 0
uart_rx5
Text Label 11400 5200 2    50   ~ 0
uart_rx6
Text Label 12650 5200 2    50   ~ 0
uart_rx6
Text Label 13900 5200 2    50   ~ 0
uart_rx7
Text Label 14900 5200 2    50   ~ 0
uart_rx7
Text Label 21150 4950 2    50   ~ 0
uart_rx0
Text Label 21150 5050 2    50   ~ 0
uart_rx1
Wire Wire Line
	21400 4950 21150 4950
$Comp
L Connector:Conn_01x02_Female J29
U 1 1 5D73B508
P 21600 4950
F 0 "J29" H 21627 4926 50  0000 L CNN
F 1 "Conn_01x02_Female" H 21627 4835 50  0000 L CNN
F 2 "footprints:header-1x2" H 21600 4950 50  0001 C CNN
F 3 "~" H 21600 4950 50  0001 C CNN
	1    21600 4950
	1    0    0    -1  
$EndComp
Wire Wire Line
	21400 5050 21150 5050
$EndSCHEMATC
