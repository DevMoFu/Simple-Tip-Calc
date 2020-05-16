# Simple-Tip-Calc-In-Go

This is a simple tip calculator app.

***Prerequisites:***
- Have [Go](https://golang.org/doc/install) installed
- Have [git](https://git-scm.com/book/en/v2/Getting-Started-Installing-Git) install

***Instructions:***
1. git clone https://github.com/DevMoFu/Simple-Tip-Calc-In-Go.git
2. go build SimpleTipCalcInGo.go
3. `SimpleTipCalcInGo -Price <price of meal> -Tax <state tax> -Tip <percent to tip> -People <number of people>`

***Args***
```
Usage of ./SimpleTipCalcInGo:
  -People uint
        Number of people (default 3)
  -Price float
        Price of goods (default 100)
  -Tax float
        State food tax (default 8.5)
  -Tip float
        Percentage to Tip (default 15)
```

***Sample Output***
```
[]# ./SimpleTipCalcInGo -Price 150.85 -Tax 6.5 -Tip 20 -People 4

Total Tip: $30.17
Total Cost post tip and tax: $190.83
Total cost per person: $47.71

```