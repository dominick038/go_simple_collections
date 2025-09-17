# Go Collections: The Wild, Wacky World of Dominick’s Data Structures

Welcome to the most unhinged Go collections repository this side of the multiverse! Why did I build this? BECAUSE I CAN, BABY! No other reason, no grand plan, just pure, chaotic, code-slinging glee. This is my playground for crafting data structures in Go.

## Table of Contents
- [Why, Tho?](#why-tho)
- [What’s in the Box?](#whats-in-the-box)
- [How to Steal This Chaos](#how-to-steal-this-chaos)
- [How to Wield the Madness](#how-to-wield-the-madness)
- [The Collections (aka My Brainchildren)](#the-collections-aka-my-brainchildren)
- [Join the Chaos](#join-the-chaos)
- [License (The Boring Bit)](#license-the-boring-bit)

## Why, Tho?

Why climb a mountain? Why eat pizza for breakfast? Why write your own collections when Go’s got slices and maps? BECAUSE IT’S FUN, THAT’S WHY! This repo is my love letter to Go, a fever dream of stacks, queues, and hashmaps that I built for the sheer thrill of it. No corporate agendas, no deadlines—just me, my keyboard, and a questionable amount of coffee.

## What’s in the Box?

- **Hand-Crafted Data Structures**: Lists, stacks, queues, trees, and maybe a graph or two if I’m feeling extra spicy.
- **Go Generics Shenanigans**: Type-safe madness using Go 1.18+ generics, because who needs `interface{}` nightmares?
- **Concurrency Craziness**: Some collections are thread-safe (IN THE FUTURE HOPEFULLY), because Go routines are my spirit animal.
- **Zero Dependencies**: Just pure, unadulterated Go. No baggage, no drama.
- ~~**Tests That Slap**: Every collection is battle-tested with unit tests so you can trust it~~

## How to Steal This Chaos

Wanna bring this unhinged energy to your Go project? It’s easier than convincing a cat to chase a laser:

```bash
go get github.com/dominick038/go_simple_collections
```

Make sure you’re rocking Go 1.18 or later, because generics are the glitter glue holding this party together.

## How to Wield the Madness

Here’s a sneak peek at the chaos you’re about to unleash. Want a stack that can handle your wildest dreams? Check this out:

```go
package main

import (
	"fmt"
	"github.com/yourusername/go-collections/stack"
)

func main() {
	// Create a stack for all your unhinged ideas
	ideaStack := stack.NewArrayStack[string]()
	ideaStack.Push("Build a robot that makes tacos")
	ideaStack.Push("Invent time travel")
	ideaStack.Push("Teach my goldfish to code")

	// Pop the chaos
	for !ideaStack.IsEmpty() {
		idea, _ := ideaStack.Pop()
		fmt.Println("Unhinged idea:", idea)
	}
}
```

Run it, and watch the madness unfold like a glitter bomb at a coding convention.

## The Collections (aka My Brainchildren)

Here’s the lineup of my chaotic creations (more to come as I spiral further into this rabbit hole):

- **Stack**: Push it, pop it, love it. Perfect for when you need to undo your life choices.
- **Queue**: First in, first out, like a line at a taco truck but with better vibes.

## Join the Chaos

Got an idea for a new collection? Wanna make my stack even stackier? Fork this repo, tinker, and send me a pull request! I’m all about collaborative chaos. Found a bug? File an issue, and I’ll wrestle it into submission with my bare hands.

## License (The Boring Bit)

This project is licensed under the MIT License, because I’m chaotic but not *that* chaotic. Do whatever you want with it just don’t blame me if your code starts singing show tunes at 3 a.m.
