## \-\-- Day 12: JSAbacusFramework.io \-\--

Santa\'s Accounting-Elves need help balancing the books after a recent
order. Unfortunately, their accounting software uses a peculiar storage
format. That\'s where you come in.

They have a [JSON](http://json.org/) document which contains a variety
of things: arrays (`[1,2,3]`), objects (`{"a":1, "b":2}`), numbers, and
strings. Your first job is to simply find all of the *numbers*
throughout the document and add them together.

For example:

-   `[1,2,3]` and `{"a":2,"b":4}` both have a sum of `6`.
-   `[[[3]]]` and `{"a":{"b":4},"c":-1}` both have a sum of `3`.
-   `{"a":[-1,1]}` and `[-1,{"a":1}]` both have a sum of `0`.
-   `[]` and `{}` both have a sum of `0`.

You will not
[encounter]{title="Nor are you likely to be eaten by a grue... during *this* puzzle, anyway."}
any strings containing numbers.

What is the *sum of all numbers* in the document?

To begin, [get your puzzle input](12/input){target="_blank"}.

Answer:

You can also [\[Share[on
[Bluesky](https://bsky.app/intent/compose?text=%22JSAbacusFramework%2Eio%22+%2D+Day+12+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F12){target="_blank"}
[Twitter](https://twitter.com/intent/tweet?text=%22JSAbacusFramework%2Eio%22+%2D+Day+12+%2D+Advent+of+Code+2015&url=https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F12&related=ericwastl&hashtags=AdventOfCode){target="_blank"}
[Mastodon](javascript:void(0);){onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=%22JSAbacusFramework%2Eio%22+%2D+Day+12+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F12';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}"
target="_blank"}]{.share-content}\]]{.share} this puzzle.
