(ns project-euler-solutions)

(defn sum-of-multiples-of-three-or-five [x]
  "Problem 1: If we list all the natural numbers below 10 that are
  multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples
  is 23.

  Find the sum of all the multiples of 3 or 5 below 1000."
  (apply + (filter #(or (zero? (rem % 3)) (zero? (rem % 5))) (range 1 x))
))

(defn fibo-sum-of-even-terms [x]
  "Problem 2: Each new term in the Fibonacci sequence is generated by
  adding the previous two terms. By starting with 1 and 2, the first
  10 terms will be:

  1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...

  Find the sum of all the even-valued terms in the sequence which do
  not exceed four million."
  (apply + (filter #(even? %) (rest (loop [f '(1 0)]
                                      (if (> (first f) x)
                                        f
                                        (recur (conj f (+ (first f) (second f))))))))))