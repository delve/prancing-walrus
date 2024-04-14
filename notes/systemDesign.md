See also: terminology.md alongside this file

# Concepts

Walrus integrates two systems with distinct visions for the same concepts. Discord has users, which are individual humans and may have multiple roles. A human may have multiple Discord users, but Walrus explicitly ignores this complication. Walrus assumes a 1:1 relation between humans and Discord IDs. Any human with multiple Discord users will have to log in as the correct Discord user to access Walrus' features. This is, in large part, to allow offloading of many auth* functions to Discord simplifying Walrus' job.

Super Snail does not explicitly assume a 1:many relationship between humans and snail IDs, but the capability is readily available and so Walrus acccomodates this in order to fully serve the users. A user's snails may be in multiple clubs, and from Walrus' perspective that means the user is *also* in multiple clubs. This is common sense. A user may have elevated privileges (owner, officer, etc) in one club and not another. Walrus settles these authz concerns by looking at Discord role membership. This creates tight integration with the Snailverse Discord server management as roles must be defined correctly for Walrus features to work properly. In the future Walrus should learn commands allowing Discord management to create club setups from a simple command so that they don't have to remember specific implementation peculiarities.

# Clubs
## On Discord
A club has 2 Discord roles. For `Example Club`:
* `@ExampleClub`
* `@ExampleClub Officers`
Note the peculiarity of the spaces. At this time I don't know if there's a character limit in role names.

Members of the `@ExampleClub` role will have access to the open club channels and members of `@ExampleClub Officers` will have access to the open club channels and to club management channels.

## What Walrus cares about
Walrus doesn't care about the `@ExampleClub` role. Walrus calculate club 'membership' for a user to be all the clubs that any of the user's snails are a member of. This creates a subtle difference in 'membership' between Walrus and Discord, with Walrus being closer to the game's definition of club membership.

However Walrus does pay attention to `@ExampleClub Officers`. Members of this role can use the club management commands. This allows a club to have more Discord officers than the four that the game actually allows. In the future this may be leveraged for additional club based commands.

## Management Commands
Current club management commands allow officers to add or remove snails from their club *in Walrus*. Walrus does **not** update or read game state.

When a snail is inducted its player will be added to the club role automatically. When a snail is kicked its player will be removed from the club *and* officer roles *if and only if* the player has no other snails in the same club

Adding a player to a club officer role is currently a manual task for Discrod server management staff. In the future a "club president" role might be needed to ease managment of the `officer` role memberships.