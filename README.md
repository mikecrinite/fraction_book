# decimal_book

This came from a tweet or a tumblr post or reddit idk something. 
The idea is that if you assign each letter, punctuation, and number its own value (i.e. 001, 002, 003), you could turn an entire book into a decimal. Becuase this decimal will always be a rational number, it can be converted into a fraction, and therefore you will have a simple fraction representing an entire book.

Considering how much less space it takes to store a fraction than the entire text, as long as you have the "key" to what is essentially a cryptographic cipher, you could easily turn it back into text by converting from the fraction back into a decimal and parsing that decimal using the key.

I called it decimal_book but I really meant fraction_book. I'm not changing it, you get the idea

## Problems
Unfortunately, I forgot that decimal precision would be an issue. As far as I can tell, there is no way to have infinite precision in golang in any built-in classes. This means that we're always going to be limited in the length of the book, and sadly that limit is very low when you consider that each character is 3 decimal places long.

It's possible that this is also an issue with creating the big.Rat that represents the book, but for the inputs I've given the program so far, it seems it hasn't had any issues with that. As a caveat of the above, it is impossible to test this once it becomes a fraction, so while I can't test that it encoded properly, I can verify that it encoded differently than last time, and I did a couple pretty cursory tests of that. It's definitely possible I just haven't reached the point where it breaks though. 

## Solutions?
- Long division? Maybe we can write our own division class that can turn the big.Rat into a String quotient rather than any built-in numeric value. While this would likely require a ton of memory, it would completely bypass the precision issue. It would also take a while, and I'd have to remember how to do decimal long division. It's been a while since I divided any numbers without a calculator.
- Find some API that can do infinite precision? I'm not even sure if "infinite precision" is what I should be looking for. It's possible that "arbitrary precision" is precise enough. All I know is that specifying precision will probably be really difficult, and there may be built-in maximum precisions anyway in golang. Part of why I was writing this was to learn golang so I can't exactly claim to be an expert.

# Notes
Please feel free to submit any suggestions/solutions/criticisms/etc
I would love to get this working. I'd prefer to write the code myself but I guess feel free to just write a PR that makes it work, too. It'd just be sick for it to work in the first place


# todo
- optimize: I did not write any of this code with optimization in mind. I wrote it on a plane to germany, and then finished it while deliriously waiting until a reasonable time to fall asleep so I could wake up at 5am to go to the gym with my brother the next day. I imagine for very long texts this will be INCREDIBLY slow, both from text -> fraction and vice versa (UPDATE: now that I learned it doesn't currently work anyway, this note doesn't really matter)
