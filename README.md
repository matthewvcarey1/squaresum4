# squaresum4
Taking my daughter's homework far too seriously

Find squares that are sums of two other squares by brute force without using square root math or floating points.

If invoked without a parameter it assumes that you want to look at the range 1 to 100.

./squaresum <positive_number>

It lists the square numbers found.

## To build

go build -o squaresum.exe ./cmd/squaresum4

## To test

go test ./...

## Example
    $ ./squaresum 30
    9 (3 * 3) + 16 (4 * 4) = 25 (5 * 5)
    25 (5 * 5) + 144 (12 * 12) = 169 (13 * 13)
    36 (6 * 6) + 64 (8 * 8) = 100 (10 * 10)
    49 (7 * 7) + 576 (24 * 24) = 625 (25 * 25)
    64 (8 * 8) + 225 (15 * 15) = 289 (17 * 17)
    81 (9 * 9) + 144 (12 * 12) = 225 (15 * 15)
    100 (10 * 10) + 576 (24 * 24) = 676 (26 * 26)
    144 (12 * 12) + 256 (16 * 16) = 400 (20 * 20)
    225 (15 * 15) + 400 (20 * 20) = 625 (25 * 25)
    256 (16 * 16) + 900 (30 * 30) = 1156 (34 * 34)
    324 (18 * 18) + 576 (24 * 24) = 900 (30 * 30)
    400 (20 * 20) + 441 (21 * 21) = 841 (29 * 29)
    441 (21 * 21) + 784 (28 * 28) = 1225 (35 * 35)
