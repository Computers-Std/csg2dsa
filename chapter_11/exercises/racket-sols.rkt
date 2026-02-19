#lang htdp/isl+

;; Ex 2
(define (get-evens lon)
  (cond
    [(empty? lon) '()]
    [else (if (even? (first lon))
              (cons (first lon) (get-evens (rest lon)))
              (get-evens (rest lon)))]))

;; Ex 3
(define (nthnum n)
  (cond
    [(= 1 n) 1]
    [else (+ n (nthnum (- n 1)))]))

;; Ex 4
(define (indexofx loc idx)
  (cond
    [(string=? (first loc) "x") idx]
    [else (indexofx (rest loc) (+ 1 idx))]))
;; Usage: (indexofx (explode "abcdefghijklmnopqrstuvwxyz") 0)

;; Ex 5
(define (uniqPaths rows cols)
  (cond
    [(or (= rows 1) (= cols 1)) 1]
    [else (+ (uniqPaths (- rows 1) cols)
             (uniqPaths rows (- cols 1)))]))
