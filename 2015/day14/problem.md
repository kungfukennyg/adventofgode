## \-\-- Day 14: Reindeer Olympics \-\--

This year is the Reindeer Olympics! Reindeer can fly at high speeds, but
must rest occasionally to recover their energy. Santa would like to know
which of his reindeer is fastest, and so he has them race.

Reindeer can only either be *flying* (always at their top speed) or
*resting* (not moving at all), and always spend whole seconds in either
state.

For example, suppose you have the following Reindeer:

-   Comet can fly *14 km/s for 10 seconds*, but then must rest for *127
    seconds*.
-   Dancer can fly *16 km/s for 11 seconds*, but then must rest for *162
    seconds*.

After one second, Comet has gone 14 km, while Dancer has gone 16 km.
After ten seconds, Comet has gone 140 km, while Dancer has gone 160 km.
On the eleventh second, Comet begins resting (staying at 140 km), and
Dancer continues on for a total distance of 176 km. On the 12th second,
both reindeer are resting. They continue to rest until the 138th second,
when Comet flies for another ten seconds. On the 174th second, Dancer
flies for another 11 seconds.

In this example, after the 1000th second, both reindeer are resting, and
Comet is in the lead at *`1120`* km (poor Dancer has only gotten `1056`
km by that point). So, in this situation, Comet would win (if the race
ended at 1000 seconds).

Given the descriptions of each reindeer (in your puzzle input), after
exactly `2503` seconds, *what distance has the winning reindeer
traveled*?

To begin, [get your puzzle input](14/input){target="_blank"}.

Answer:

You can also [\[Share[on
[Bluesky](https://bsky.app/intent/compose?text=%22Reindeer+Olympics%22+%2D+Day+14+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F14){target="_blank"}
[Twitter](https://twitter.com/intent/tweet?text=%22Reindeer+Olympics%22+%2D+Day+14+%2D+Advent+of+Code+2015&url=https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F14&related=ericwastl&hashtags=AdventOfCode){target="_blank"}
[Mastodon](javascript:void(0);){onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=%22Reindeer+Olympics%22+%2D+Day+14+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F14';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}"
target="_blank"}]{.share-content}\]]{.share} this puzzle.
