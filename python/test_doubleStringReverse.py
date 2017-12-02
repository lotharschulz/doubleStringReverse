#!/usr/bin/env python3


i = "be thEre oR Be sqUare"
r = "ERAuQS Eb rO EReHT EB"


def double_str_reverse(string):
    if isinstance(string, str):
        return string[::-1].swapcase()


def double_str_reverse_alt(string):
    if isinstance(string, str):
        return ''.join(reversed(string)).swapcase()


def double_str_reverse_alt2(string):
    if isinstance(string, str):
        return "".join([letter.lower() if letter.isupper() else letter.upper() for letter in string[::-1]])


def test_double_str_reverse():
    assert double_str_reverse(i) == r
    assert double_str_reverse_alt(i) == r
    assert double_str_reverse_alt2(i) == r


if __name__ == "__main__":
    print("'" + i + "' double string reversed:      '" + double_str_reverse(i) + "'")
    print("'" + i + "' double string reversed alt:  '" + double_str_reverse_alt(i) + "'")
    print("'" + i + "' double string reversed alt2: '" + double_str_reverse_alt2(i) + "'")
