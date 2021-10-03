package unicorn
// For Unicorn Engine. AUTO-GENERATED FILE, DO NOT EDIT [riscv_const.go]
const (

// RISCV registers

	RISCV_REG_INVALID = 0

// General purpose registers
	RISCV_REG_X0 = 1
	RISCV_REG_X1 = 2
	RISCV_REG_X2 = 3
	RISCV_REG_X3 = 4
	RISCV_REG_X4 = 5
	RISCV_REG_X5 = 6
	RISCV_REG_X6 = 7
	RISCV_REG_X7 = 8
	RISCV_REG_X8 = 9
	RISCV_REG_X9 = 10
	RISCV_REG_X10 = 11
	RISCV_REG_X11 = 12
	RISCV_REG_X12 = 13
	RISCV_REG_X13 = 14
	RISCV_REG_X14 = 15
	RISCV_REG_X15 = 16
	RISCV_REG_X16 = 17
	RISCV_REG_X17 = 18
	RISCV_REG_X18 = 19
	RISCV_REG_X19 = 20
	RISCV_REG_X20 = 21
	RISCV_REG_X21 = 22
	RISCV_REG_X22 = 23
	RISCV_REG_X23 = 24
	RISCV_REG_X24 = 25
	RISCV_REG_X25 = 26
	RISCV_REG_X26 = 27
	RISCV_REG_X27 = 28
	RISCV_REG_X28 = 29
	RISCV_REG_X29 = 30
	RISCV_REG_X30 = 31
	RISCV_REG_X31 = 32

// Floating-point registers
	RISCV_REG_F0 = 33
	RISCV_REG_F1 = 34
	RISCV_REG_F2 = 35
	RISCV_REG_F3 = 36
	RISCV_REG_F4 = 37
	RISCV_REG_F5 = 38
	RISCV_REG_F6 = 39
	RISCV_REG_F7 = 40
	RISCV_REG_F8 = 41
	RISCV_REG_F9 = 42
	RISCV_REG_F10 = 43
	RISCV_REG_F11 = 44
	RISCV_REG_F12 = 45
	RISCV_REG_F13 = 46
	RISCV_REG_F14 = 47
	RISCV_REG_F15 = 48
	RISCV_REG_F16 = 49
	RISCV_REG_F17 = 50
	RISCV_REG_F18 = 51
	RISCV_REG_F19 = 52
	RISCV_REG_F20 = 53
	RISCV_REG_F21 = 54
	RISCV_REG_F22 = 55
	RISCV_REG_F23 = 56
	RISCV_REG_F24 = 57
	RISCV_REG_F25 = 58
	RISCV_REG_F26 = 59
	RISCV_REG_F27 = 60
	RISCV_REG_F28 = 61
	RISCV_REG_F29 = 62
	RISCV_REG_F30 = 63
	RISCV_REG_F31 = 64
	RISCV_REG_PC = 65
	RISCV_REG_ENDING = 66

// Alias registers
	RISCV_REG_ZERO = 1
	RISCV_REG_RA = 2
	RISCV_REG_SP = 3
	RISCV_REG_GP = 4
	RISCV_REG_TP = 5
	RISCV_REG_T0 = 6
	RISCV_REG_T1 = 7
	RISCV_REG_T2 = 8
	RISCV_REG_S0 = 9
	RISCV_REG_FP = 9
	RISCV_REG_S1 = 10
	RISCV_REG_A0 = 11
	RISCV_REG_A1 = 12
	RISCV_REG_A2 = 13
	RISCV_REG_A3 = 14
	RISCV_REG_A4 = 15
	RISCV_REG_A5 = 16
	RISCV_REG_A6 = 17
	RISCV_REG_A7 = 18
	RISCV_REG_S2 = 19
	RISCV_REG_S3 = 20
	RISCV_REG_S4 = 21
	RISCV_REG_S5 = 22
	RISCV_REG_S6 = 23
	RISCV_REG_S7 = 24
	RISCV_REG_S8 = 25
	RISCV_REG_S9 = 26
	RISCV_REG_S10 = 27
	RISCV_REG_S11 = 28
	RISCV_REG_T3 = 29
	RISCV_REG_T4 = 30
	RISCV_REG_T5 = 31
	RISCV_REG_T6 = 32
	RISCV_REG_FT0 = 33
	RISCV_REG_FT1 = 34
	RISCV_REG_FT2 = 35
	RISCV_REG_FT3 = 36
	RISCV_REG_FT4 = 37
	RISCV_REG_FT5 = 38
	RISCV_REG_FT6 = 39
	RISCV_REG_FT7 = 40
	RISCV_REG_FS0 = 41
	RISCV_REG_FS1 = 42
	RISCV_REG_FA0 = 43
	RISCV_REG_FA1 = 44
	RISCV_REG_FA2 = 45
	RISCV_REG_FA3 = 46
	RISCV_REG_FA4 = 47
	RISCV_REG_FA5 = 48
	RISCV_REG_FA6 = 49
	RISCV_REG_FA7 = 50
	RISCV_REG_FS2 = 51
	RISCV_REG_FS3 = 52
	RISCV_REG_FS4 = 53
	RISCV_REG_FS5 = 54
	RISCV_REG_FS6 = 55
	RISCV_REG_FS7 = 56
	RISCV_REG_FS8 = 57
	RISCV_REG_FS9 = 58
	RISCV_REG_FS10 = 59
	RISCV_REG_FS11 = 60
	RISCV_REG_FT8 = 61
	RISCV_REG_FT9 = 62
	RISCV_REG_FT10 = 63
	RISCV_REG_FT11 = 64
)