from datetime import datetime, timedelta
from dataclasses import dataclass
from collections.abc import Callable
import matplotlib.pyplot as plt

from benchmark_data import CONST_2_16_1, DATA_2_16_2_64


@dataclass
class BenchmarkReport:
    start: datetime
    end: datetime
    delta: timedelta
    title: str
    time_array: list
    number_array: list


def run_benchmark_1_65637_for_algo(algo: Callable, title: str) -> BenchmarkReport:
    print(f"\nTesting all numbers from 1 to 2^16+1(it is {CONST_2_16_1}), {title}")
    time_array = []
    number_array = []
    start = datetime.now()
    for i in range(2, CONST_2_16_1):
        number_array.append(i)
        algo(i)
        time_array.append((datetime.now() - start).total_seconds())
    end = datetime.now()
    delta = end - start

    return BenchmarkReport(start=start, end=end, delta=delta, title=title, time_array=time_array, number_array=number_array)


def run_benchmark_2_16_2_64_for_algo(algo: Callable, title: str) -> BenchmarkReport:
    print(f"\nTesting all numbers from 2^16-1, 2^16+1 to 2^64-1, 2^64+1, {title}")
    test_data = DATA_2_16_2_64
    time_array = []
    start = datetime.now()
    for number in test_data:
        algo(number)
        time_array.append((datetime.now() - start).total_seconds())
    end = datetime.now()
    delta = end - start
    return BenchmarkReport(start=start, end=end, delta=delta, title=title, time_array=time_array, number_array=test_data)


def print_report(report: BenchmarkReport):
    print("----------------------------------")
    start = report.start
    end = report.end
    delta = report.delta
    delta_ms = delta.total_seconds() * 1000
    print(f"Started at: {start}")
    print(f"Finished at: {end}")
    print(f"Test has taken {delta_ms} ms")
    print("----------------------------------")
