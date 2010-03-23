(ns project-euler-solutions-test
  (:require [project-euler-solutions :as pes])
  (:use clojure.test))

(deftest test-sum-of-multiples-of-three-or-five
  #^{:doc "Test solution for Project Euler Problem 1"}
  (is (= (pes/sum-of-multiples-of-three-or-five 10) 23))
  (is (= (pes/sum-of-multiples-of-three-or-five 1000) 233168)))

(deftest test-fibo-sum-of-even-terms
  #^{:doc "Test solution for Project Euler Problem 2"}
  (is (= (pes/fibo-sum-of-even-terms 4000000) 4613732)))
