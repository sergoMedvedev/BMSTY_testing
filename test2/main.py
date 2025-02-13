import sys

from algo import is_prime, is_prime_optimized, is_prime_miller
import benchmark as benchmark
import matplotlib.pyplot as plt

def main():
    for i in range(101, 1000):
        if is_prime_miller(i) != is_prime_optimized(i):
            print(f"{i}, miller: {is_prime_miller(i)}, simple optimized: {is_prime_optimized(i)}")
    print(sys.version)
    report = benchmark.run_benchmark_1_65637_for_algo(is_prime, "Brut-force algo")
    benchmark.print_report(report)
    plt.plot(report.number_array, report.time_array, label="Brut-force algo")


    report = benchmark.run_benchmark_1_65637_for_algo(is_prime_optimized,
                                                      "Brut-force algo with sqrt(n) and 6k+1 optimizations")
    benchmark.print_report(report)
    plt.plot(report.number_array, report.time_array, label="Brut-force algo with sqrt(n)")


    report = benchmark.run_benchmark_1_65637_for_algo(is_prime_miller,
                                                      "Miller test")
    benchmark.print_report(report)
    plt.plot(report.number_array, report.time_array, label="Miller test")
    plt.xlabel("number")
    plt.ylabel("time (s)")
    plt.legend()
    plt.grid()
    plt.show()


    report = benchmark.run_benchmark_2_16_2_64_for_algo(is_prime_optimized, "Brut-force algo")
    benchmark.print_report(report)
    plt.plot(report.number_array, report.time_array, label="Brut-force algo 2_16_2_64")

    report = benchmark.run_benchmark_2_16_2_64_for_algo(is_prime_optimized,
                                                        "Brut-force algo with sqrt(n) and 6k+1 optimizations")
    benchmark.print_report(report)
    plt.plot(report.number_array, report.time_array, label="Brut-force algo with sqrt(n) 2_16_2_64")


    report = benchmark.run_benchmark_2_16_2_64_for_algo(is_prime_miller,
                                                        "Miller test")
    benchmark.print_report(report)
    plt.plot(report.number_array, report.time_array, label="Miller test 2_16_2_64")

    plt.xlabel("number")
    plt.ylabel("time (s)")
    plt.legend()
    plt.grid()
    plt.show()




if __name__ == '__main__':
    main()