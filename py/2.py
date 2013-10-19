#!/usr/bin/python


def fibo(terms, limit):
    length = len(terms)
    s = terms[length - 1] + terms[length - 2]
    if s > limit:
        return terms

    terms.append(s)
    return fibo(terms, limit)

if __name__ == '__main__':
    print "The sum of the even terms of a fibo sequence under 4M is %s " % sum(filter(lambda x: not x % 2, fibo([1, 2], 4000000)))
