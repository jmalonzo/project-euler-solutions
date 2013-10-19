#!/usr/bin/python

import sys


if __name__ == '__main__':
    limit = 10
    if len(sys.argv) > 1:
        limit = sys.argv[1]

    print "Multiples of 3 and 5 below %s is %d" % (limit, sum(filter(lambda x: not x % 3 or not x % 5, range(int(limit)))))
