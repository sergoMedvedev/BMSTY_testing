from random import randrange
from math import log2


def is_prime(n: int) -> bool:
    """just brut force"""
    i = 2
    stop = n - 1
    while i <= stop:
        if not n % i:
            return False
        i += 1
    return True


def is_prime_optimized(n: int) -> bool:
    """6k+1 optimization"""
    if n <= 3:
        return n >= 1
    if not n % 2 or not n % 3:
        return False
    i = 5
    """use sqrt(n) instead of n in loop"""
    stop = int(n ** 0.5)
    while i <= stop:
        if not n % i or not n % (i + 2):
            return False
        i += 6
    return True


def is_prime_miller(n: int, rounds: int | None = None) -> bool:
    """6k+1 optimization"""
    if n <= 3:
        return n >= 1
    if not n % 2 or not n % 3:
        return False
    """present n as 2^s * t + 1"""
    buff = n - 1
    s, t = 0, 0
    while t == 0:
        if not buff % 2:
            buff = buff / 2
            s += 1
        else:
            t = int(buff)
    """prepare rounds, log2(n) recommended and used as default"""
    if rounds is None:
        rounds = int(log2(n)) + 1
    x, y = 0, 0
    for i in range(0, rounds):
        a = randrange(2, n - 2)
        x = power_mod(x=a, pow_=t, mod_=n)
        for j in range(0, max(s, 1)):
            y = power_mod(x=x, pow_=2, mod_=n)
            if y == 1 and x != 1 and x != n - 1:
                return False
            x = y
        if y != 1:
            return False
    return True


def power_mod(x: int, pow_: int, mod_: int):
    res = 1
    pow_res = 0
    while pow_res < pow_:
        pow_res_1 = 2
        res1 = x
        while pow_res + pow_res_1 <= pow_:
            res1 = (res1 * res1) % mod_
            pow_res_1 *= 2
        pow_res_1 //= 2
        res = (res * res1) % mod_
        pow_res += pow_res_1
    return res % mod_
