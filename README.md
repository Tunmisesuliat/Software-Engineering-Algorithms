# Software-Engineering-Algorithms
Algorithms to some software engineering challenges

Technologies used:
● Javascript
● Python
● Golang

Concepts used:
Arrays, Array methods, slices, control structures(if, for, range)

Questions:

Question 1:
An array A consisting of N integers is given. A triplet (P, Q, R) is triangular if 0 ≤
P < Q < R < N and:
● A[P] + A[Q] > A[R],
● A[Q] + A[R] > A[P],
● A[R] + A[P] > A[Q].
For example, consider array A such that:
A[0] = 10 A[1] = 2 A[2] = 5
A[3] = 1 A[4] = 8 A[5] = 20
Triplet (0, 2, 4) is triangular.
that, given an array A consisting of N integers, returns 1 if there exists a
triangular triplet for this array and returns 0 otherwise.

Question 2:
A string S consisting of N characters is considered to be properly nested if any
of the following conditions is true:

● S is empty;
● S has the form "(U)" or "[U]" or "{U}" where U is a
properly nested string;
● S has the form "VW" where V and W are properly nested
strings.

For example, the string "{[()()]}" is properly nested but "([)()]" is not.
Write a function:
class Solution { public int solution(String S); }
that, given a string S consisting of N characters, returns 1 if S is properly
nested and 0 otherwise.
For example, given S = "{[()()]}", the function should return 1 and given S =
"([)()]", the function should return 0, as explained above.
Write an efficient algorithm for the following assumptions:
● N is an integer within the range [0..200,000];
● string S consists only of the following characters: "(", "{",
"[", "]", "}" and/or ")".

Question 3
An array A consisting of N integers is given. The dominator of array A is the
value that occurs in more than half of the elements of A.
For example, consider array A such that
A[0] = 3 A[1] = 4 A[2] = 3
A[3] = 2 A[4] = 3 A[5] = -1
A[6] = 3 A[7] = 3
The dominator of A is 3 because it occurs in 5 out of 8 elements of A (namely
in those with indices 0, 2, 4, 6 and 7) and 5 is more than a half of 8.
Write a function
class Solution { public int solution(int[] A); }
that, given an array A consisting of N integers, returns index of any element of array A in
which the dominator of A occurs. The function should return −1 if array A does not have a
dominator.
For example, given array A such that
A[0] = 3 A[1] = 4 A[2] = 3
A[3] = 2 A[4] = 3 A[5] = -1
A[6] = 3 A[7] = 3
the function may return 0, 2, 4, 6 or 7, as explained above.
Write an efficient algorithm for the following assumptions:
● N is an integer within the range [0..100,000];
● each element of array A is an integer within the range
[−2,147,483,648..2,147,483,647].
