CONST_2_16_1 = 2 ** 16 + 1


DATA_2_16_2_64 = []
for i in range(16, 64):
    DATA_2_16_2_64.append(2 ** i - 1)
    DATA_2_16_2_64.append(2 ** i + 1)
