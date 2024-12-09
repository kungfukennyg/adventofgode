## \-\-- Day 9: All in a Single Night \-\--

Every year, Santa manages to deliver all of his presents in a single
night.

This year, however, he has some [new
locations]{title="Bonus points if you recognize all of the locations."}
to visit; his elves have provided him the distances between every pair
of locations. He can start and end at any two (different) locations he
wants, but he must visit each location exactly once. What is the
*shortest distance* he can travel to achieve this?

For example, given the following distances:

    London to Dublin = 464
    London to Belfast = 518
    Dublin to Belfast = 141

The possible routes are therefore:

    Dublin -> London -> Belfast = 982
    London -> Dublin -> Belfast = 605
    London -> Belfast -> Dublin = 659
    Dublin -> Belfast -> London = 659
    Belfast -> Dublin -> London = 605
    Belfast -> London -> Dublin = 982

The shortest of these is `London -> Dublin -> Belfast = 605`, and so the
answer is `605` in this example.

What is the distance of the shortest route?

To begin, [get your puzzle input](9/input){target="_blank"}.

Answer:

You can also [\[Share[on
[Bluesky](https://bsky.app/intent/compose?text=%22All+in+a+Single+Night%22+%2D+Day+9+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F9){target="_blank"}
[Twitter](https://twitter.com/intent/tweet?text=%22All+in+a+Single+Night%22+%2D+Day+9+%2D+Advent+of+Code+2015&url=https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F9&related=ericwastl&hashtags=AdventOfCode){target="_blank"}
[Mastodon](javascript:void(0);){onclick="var ms; try{ms=localStorage.getItem('mastodon.server')}finally{} if(typeof ms!=='string')ms=''; ms=prompt('Mastodon Server?',ms); if(typeof ms==='string' && ms.length){this.href='https://'+ms+'/share?text=%22All+in+a+Single+Night%22+%2D+Day+9+%2D+Advent+of+Code+2015+%23AdventOfCode+https%3A%2F%2Fadventofcode%2Ecom%2F2015%2Fday%2F9';try{localStorage.setItem('mastodon.server',ms);}finally{}}else{return false;}"
target="_blank"}]{.share-content}\]]{.share} this puzzle.
