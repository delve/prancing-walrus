# What do I need from a SS API
I need an API that provide roughly this function:
* I can send it an arbitrary snail ID
* I get back public stats for that snail, particularly
* * Club membership
* * Sim power
* * Total power
* * Club offices (officer, mascot, etc)
* * Any HARD or AFFCT stats

TBH I don't remember exactly what information is available on other snails through the game interface, but I believe the first 3 are definitely available, and those are the only ones I'd *need*. Club rank would be super useful as well.

# What do I want to do with it
I want to allow Discord users in a SS club to coordinate better during Species War. Currently Prancing Walrus reads a spreadsheet maintained by humans and provides coordination on role and mining assignment, allowing a club to be more effective. The trouble is the manual sheet requires humans to do a lot of data gathering and entry.

With an API as described I could
* Let players tell me their snail ID
* Associate thier discord ID (which is unchangable, not their server) with the snail ID (which is also unchangeable AFAIK) in my database
* Call the API to get their info
* Cache that info in my database
* Refresh that info automatically every week, just before SW starts
* Update cached info on demand, with rate limiting
* Automatically calculate role assignments for clubs based on relatively simple algorithms (current users are using descending sim power in the spreadsheet)
* Automatically assign Discord roles based on Species War role, to allow for targeted pings (EG @Vanguards Need someone to finish off the boss so we can get to drilling)
* Automatically maintain Prancing Walrus' knowledge of club memberships (when a snail leaves a club they won't be calculated in role assigments, and vice versa)
* When a snail leaves a club they can be automatically removed from the club's Discord roles, blocking them from accessing club specific channels
* If club officer information is available it streamlines controlling access to certain Prancing Walrus commands such as refreshing the stats for snails in their club (EG to facilitate removal from roles as in the last point)

# What controls do I intend to have in place
* Rate limiting - acceptable rate TBD, but this will be a necessary control for both automated and on-demand refreshing
* Snail validation - I'll need some way to ensure a Discord user is giving an ID for a snail they own. My intial thoughts here are having them tweak their stats in a spefici way. EG Prancing Walrus says "Make your sim power end in '77' then tell me to check." This isn't attack-proof but should be a reasonably high bar.
