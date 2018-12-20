#Idea

Read in a list of entities and a taxonomy.

Use max entropy to assign probability, while remembering the choice as log
(in a separate table).



tOffer user a choice of Others, so that the choice is mutually exclusive and collectively exhaustive.
Number of choices is a param you can tweak then.

When user didn't choose any of that, deduct "1" from the options shown.
Don't deduct others, or deduct an even number. Either way.

Learning rate is also a param, set that to 1.0.
Have a log of that so system can be replayed.


Now, derive probability.

* If system starts with zero state.
* If system starts with probability.


Let's say system starts with probability.
We need a way to exploit/explore.
(and graduate, that is a cron job)

Given a set of probabilities, which one to show to the user?

The pricinple of explore exploit. Bayesian Bandit. yeah.




Each time user makes a choice,

MVP:
Give a list of entities, assuming all equal probability. Start outputting the choices and record them.
In mem: maintain a probability distribution (type).

def spread(prob, category).
// this is to calculate probability.

def entropy(prob)
// this is sampling.
def sample(prob) categories.
// this is Bayesian update.
def modify(prob, category)
User update one, we log the choices and user pick. And then do a frequency update on the count. This gives us the new probability.
How to incorporate that with prior when you have prior?
Two tablets a day.
In mem:
Heap of probability pop the max entropy to curate.
Sampling algorithm can be plug in. No need to fix on one Right now
When you spread to sub categories the entropy goes up. Which is understandable.
You can maintain a heap (value sorted) entropy of entities so you always have something to show to user.
Probably need to mutex protect the log. Which is fine I can do it via an event queue. Go routine on that protected by mutex.
You can reproduce that anyway. But sure you can save a snapshot of the entropy.
Data structure wise:
Need to have a way to store entity and their labels probability distribution.
No need a way to store entity and their entropy, derived anyway.
LevelDB does it just fine. Log based anyway.
Keep adding the entry to the end.
Get entry to probability done tonight. (Optimization) Treat this as a snapshot. Every N seconds fire up a snapshot so you can load it. Add entry to snapshot.
Then get entropy in memory.
Then get sampling done. The


cognitive
