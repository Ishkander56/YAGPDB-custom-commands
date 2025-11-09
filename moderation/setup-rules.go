{{/*
	Sets up a "rules" channel.
	Deletes the previous rules list, sets up several simple embeds containing your server rules, and adds a button to use YAGPDB's built-in Ticket system.

	Trigger: Command (mention/cmd prefix) | "-setup"
	Channels: #rules, etc.
	Roles: Anybody with permission to post in your #rules channel.
	Misc: This assumes you have already set up the Ticket system in your server.
	Author: Ishkander <https://github.com/ishkander56>
*/}}

{{/* Customizable rules text. Every individual rule should be contained within the (cembed <text>) structure; add or remove new rules by following the existing format.*/}}
{{$rules := cslice
	(cembed
		"title" "Rule 0"
		"description" (joinStr "\n" "**The Golden Rule**" "By participating in this server, you agree to a broad social contract that can't be fully explained in a handful of rules posts." "" "Be nice." "Ask questions if you're unsure about certain behavior." "Avoid doing things that force us to add new rules." "Malicious rules-lawyering will not be tolerated." "\"Debate\" is not a good excuse for bad behavior.")
	)
	(cembed
		"title" "Rule 1"
		"description" (joinStr "\n" "**Act in good faith.**" "People across the internet come from a large range of backgrounds. They may not have the knowledge and experience that you do." "People deserve to be treated with a baseline level of decency. Do not mistreat others, especially not on the basis of identity or ability." "" "Ideologies that value or devalue people based on their race, religion, gender, orientation, or disability are expressly forbidden.")
	)
	(cembed
		"title" "Rule 2"
		"description" (joinStr "\n" "**Help others help you.**" "The best way to get help with an issue is to be forthcoming about information. The internet is full of enthusiasts who can answer questions on their own time, not coworkers who have to juggle tasks. State your issues quickly and clearly. Do not wait for people to commit to you before you go into proper detail. Tell people about the problem you're experiencing, not just the solution you're trying to use." "" "**Further reading**" "https://dontasktoask.com/" "https://nohello.net/en/" "https://xyproblem.info/" "https://stackoverflow.com/help/how-to-ask")
	)
}}

{{/* Actual processing begins here. */}}
{{deleteTrigger}}
{{execAdmin "clean 20"}}
{{range $i, $content := $rules}}
	{{sendMessage nil $content}}
{{end}}
{{execAdmin "tickets MenuCreate"}}
