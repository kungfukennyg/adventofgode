## \-\-- Day 10: Elves Look, Elves Say \-\--

Today, the Elves are playing a game called
[look-and-say](https://en.wikipedia.org/wiki/Look-and-say_sequence).
They take turns making sequences by reading aloud the previous sequence
and using that reading as the next sequence. For example, `211` is read
as \"one two, two ones\", which becomes `1221` (`1` `2`, `2` `1`s).

Look-and-say sequences are generated iteratively, using the previous
value as input for the next step. For each step, take the previous
value, and replace each run of digits (like `111`) with the number of
digits (`3`) followed by the digit itself (`1`).

For example:

-   `1` becomes `11` (`1` copy of digit `1`).
-   `11` becomes `21` (`2` copies of digit `1`).
-   `21` becomes `1211` (one `2` followed by one `1`).
-   `1211` becomes `111221` (one `1`, one `2`, and two `1`s).
-   `111221` becomes `312211` (three `1`s, two `2`s, and one `1`).

Starting with the digits in your puzzle input, apply this process 40
times. What is *the length of the result*?

Your puzzle input is `1321131112`{.puzzle-input}.

Answer:

You can also [\[Share[on
[Bluesky](https://bsky.app/intent/compose?text=%22Elves+Look%2C+Elves+Say%22+%2D+Day+10+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F10){target="_blank"}
[Twitter](https://twitter.com/intent/tweet?text=%22Elves+Look%2C+Elves+Say%22+%2D+Day+10+%2D+Advent+of+Code+2015&url=https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F10&related=ericwastl&hashtags=AdventOfCode){target="_blank"}
[Mastodon](javascript:void(0);){onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=%22Elves+Look%2C+Elves+Say%22+%2D+Day+10+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F10';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}"
target="_blank"}]{.share-content}\]]{.share} this puzzle.
