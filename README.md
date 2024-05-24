# fraction_book

This came from a tweet or a tumblr post or reddit idk... somewhere on the internet. 
The idea is that if you assign each letter, punctuation, and number its own value (i.e. 001, 002, 003), you could turn an entire book into a decimal. Becuase this decimal will always be a rational number, it can be converted into a fraction, and therefore you will have a simple fraction representing an entire book.

~~Considering how much less space it takes to store a fraction than the entire text, as long as you have the "key" to what is essentially a cryptographic cipher, you could easily turn it back into text by converting from the fraction back into a decimal and parsing that decimal using the key.~~

I thought the above was going to be true. But it turns out these fractions, when exact, are not simplifiable. So basically it's a numerator that's roughly 3x as long as the input text (if you use the key where each character is represented by a 3 digit number) over a denominator that's basically the same length. Plus, you need to store the "/" symbol so people know it's a fraction. In short, the only thing this would ever be useful for would be hiding data I guess, but it's not exactly cryptographically secure. 

I hope someone comes along and looks at this and thinks "actually Mike, you're dumb, this could totally work if you didn't just (1) see the first few fractions that the code created, (2) ask ChatGPT to simplify them, only for it to tell you they were already simplified, and then (3) give up because you thought it was impossible" and submits a beautiful PR that solves all the problems it has. I don't think that's super likely because (a) who is gonna stumble on to this, and (b) if they do, who is gonna care enough to solve it. (It could have potentially revolutionized the way we store large text files if it actually created a fraction smaller than the input text but I digress)

Please feel free to submit any suggestions/solutions/criticisms/etc as well, or to just chastise me for taking the easy way out and asking ChatGPT for help

## Example output
Take a peek at `main.go` for an explanation, although I guess the printlns kinda explain it too...
```
======= Decimal String: ==========
0.020058059069000059069000051000070055069070000065056000070058055000055063055068057055064053075000051062055068070000069075069070055063
======= Text From Decimal String: ==========
This is a test of the emergency alert system
======= Fraction String: ==========
20058059069000059069000051000070055069070000065056000070058055000055063055068057055064053075000051062055068070000069075069070055063/1000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
======== Result Text: ===========
This is a test of the emergency alert system
```

## Ideas worth exploring at a later date?
- Since we know the denominator will pretty much always be 10^n, we could technically just represent the fraction in scientific notation, i.e. numeratorE10^n or whatever... been a while since I did that and again, it wouldn't make it efficient enough space-wise to be a practical storage solution for text

# todo
- optimize I guess? I did not write any of this code with optimization in mind. I wrote it on a plane to germany, and then finished it while deliriously waiting until a reasonable time to fall asleep so I could wake up at 5am to go to the gym with my brother the next day. I imagine for very long texts this will be INCREDIBLY slow, both from text -> fraction and vice versa ~~(UPDATE: now that I learned it doesn't currently work anyway, this note doesn't really matter)~~
(UPDATE 2: plus, I added some ChatGPT-generated code at the end, and that may not be very optimized either tbh)
