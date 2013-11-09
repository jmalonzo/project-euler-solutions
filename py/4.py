#!/usr/bin/env python

import sys

def palindrome(s_num):
    length = len(s_num)
    for x in range(length / 2):
        if s_num[x] is not s_num[length - x - 1]:
            return False
    return s_num


if __name__ == '__main__':
    start = 999
    end = 99
    lp = 0
    for i in range(start, end, -1):
        for j in range(start, end, -1):
            num = i * j
            p = int(palindrome(str(num)))
            if p and p > lp:
                lp = p

    print >> sys.stdout, "The largest palindrome made from the product of two 3-digit numbers is: %s" % lp
