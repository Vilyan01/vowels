
# Vowels

Vowels is a AI engine made in Golang. It is based on simple principles of cause and effect. 

## Core Structure

The idea behind vowels is that every action has an effect on an entity. This is broken down into several stages:

1. Entities - The entity is the actor which has certain interests. These interests will affect what actions the entity commits.
2. Interests - A certain object or motivation for an entity.
3. Actions - Something an entity will do to satisfy interests. Actions that best fit an entities interests will be given priority over other actions.
4. Outcomes - Actions will satisfy interests, as these interests are satisfied, the outcome will affect the level of the interest.
5. Updates - As interests are satisfied, the weight of that interest will go down and the weight of unsatisfied interests will go up.

Lets say for example we have a race of orcs (*entities*). This race is driven by gold, territorial dominance, and self preservation (*interests*). There are many things an orc can do. It can sleep, roam, attack players, hunt, and socialize (*actions*). All of these actions will satisfy the interests of the orc in some way. However, lets say roaming will satisfy the interests the best. The orcs will form parties and conduct patrols throughout its territories. They encounter a caravan and come to a decision point. Their desire for gold and territorial dominance outweighs their desire for self preservation and they decide to attack the caravan. There are two *outcomes* for this scenario. They either win and are enriched by gold, or lose the fight. Each of these *outcomes* will *update* the weight of their interests in a different manner.

## Usage

1. Open your **Terminal** application.
2. Type `go get github.com/Vilyan01/vowels` to download the library.