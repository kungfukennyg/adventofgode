## \-\-- Day 4: The Ideal Stocking Stuffer \-\--

Santa needs help [mining](https://en.wikipedia.org/wiki/Bitcoin#Mining)
some [AdventCoins]{title="Hey, mined your own business!"} (very similar
to [bitcoins](https://en.wikipedia.org/wiki/Bitcoin)) to use as gifts
for all the economically forward-thinking little girls and boys.

To do this, he needs to find [MD5](https://en.wikipedia.org/wiki/MD5)
hashes which, in
[hexadecimal](https://en.wikipedia.org/wiki/Hexadecimal), start with at
least *five zeroes*. The input to the MD5 hash is some secret key (your
puzzle input, given below) followed by a number in decimal. To mine
AdventCoins, you must find Santa the lowest positive number (no leading
zeroes: `1`, `2`, `3`, \...) that produces such a hash.

For example:

-   If your secret key is `abcdef`, the answer is `609043`, because the
    MD5 hash of `abcdef609043` starts with five zeroes
    (`000001dbbfa...`), and it is the lowest such number to do so.
-   If your secret key is `pqrstuv`, the lowest number it combines with
    to make an MD5 hash starting with five zeroes is `1048970`; that is,
    the MD5 hash of `pqrstuv1048970` looks like `000006136ef...`.

Your puzzle input is `ckczppom`{.puzzle-input}.

Answer:

You can also [\[Share[on
[Bluesky](https://bsky.app/intent/compose?text=%22The+Ideal+Stocking+Stuffer%22+%2D+Day+4+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F4){target="_blank"}
[Twitter](https://twitter.com/intent/tweet?text=%22The+Ideal+Stocking+Stuffer%22+%2D+Day+4+%2D+Advent+of+Code+2015&url=https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F4&related=ericwastl&hashtags=AdventOfCode){target="_blank"}
[Mastodon](javascript:void(0);){onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=%22The+Ideal+Stocking+Stuffer%22+%2D+Day+4+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F4';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}"
target="_blank"}]{.share-content}\]]{.share} this puzzle.
