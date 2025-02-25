import pytest
from src.sum import custom_sum

def test_custom_sum_empty_list():
    assert custom_sum([]) == 0

def test_custom_sum_single_number():
    assert custom_sum([5]) == 5

def test_custom_sum_multiple_numbers():
    assert custom_sum([1, 2, 3, 4, 5]) == 15

def test_custom_sum_negative_numbers():
    assert custom_sum([-1, -2, -3]) == -6

def test_custom_sum_mixed_numbers():
    assert custom_sum([-5, 0, 5]) == 0

def test_custom_sum_float_numbers():
    assert custom_sum([1.5, 2.5, 3.5]) == 7.5

def test_custom_sum_large_numbers():
    assert custom_sum([1000000, 2000000, 3000000]) == 6000000

def test_custom_sum_zero():
    assert custom_sum([0, 0, 0]) == 0

def test_custom_sum_single_zero():
    assert custom_sum([0]) == 0

def test_custom_sum_decimal_precision():
    assert abs(custom_sum([0.1, 0.2, 0.3]) - 0.6) < 1e-10

def test_custom_sum_mixed_types():
    assert custom_sum([1, 2.5, -3.7, 4]) == 3.8
