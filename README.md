# Hacker Laws CLI

A simple command line tool to give you a fancy command line interface to dive into laws, theories, principles and patterns listed in the repo [hacker-laws](https://github.com/dwmkerr/hacker-laws) by [Dave Kerr](https://github.com/dwmkerr).

## How To Install

### Basic

Follow the steps;

```bash
git clone git@github.com:umutphp/hacker-laws-cli.git

cd hacker-laws-cli

go run main.go
```

### Build as binary

Follow the steps;

```bash
git clone git@github.com:umutphp/hacker-laws-cli.git

cd hacker-laws-cli

sudo go build -o /usr/local/bin/hacker-laws-cli .

hacker-laws-cli list
```

### Download and use official binary

Visit the latest release page and download the binary correspondingly.

```bash
wget -O /usr/local/bin/hacker-laws-cli https://latest-binary-url

hacker-laws-cli
```

## How To Use

The build will create an executable with name *hacker-laws-cli*.

Execute it without argument to see the argument option;

```bash
$ ./hacker-laws-cli
Options for the command:
        help   To display argument list.
        list   To list the laws and principles.
      random   To display random law or principles.
```

Sample execution #1

```bash
$ ./hacker-laws-cli list
Laws
     -  Amdahl's Law
     -  The Broken Windows Theory
     -  Brooks' Law
     -  Conway's Law
     -  Cunningham's Law
     -  Dunbar's Number
     -  Gall's Law
     -  Goodhart's Law
     -  Hanlon's Razor
     -  Hofstadter's Law
     -  Hutber's Law
     -  The Hype Cycle & Amara's Law
     -  Hyrum's Law (The Law of Implicit Interfaces)
     -  Metcalfe's Law
     -  Moore's Law
     -  Murphy's Law / Sod's Law
     -  Occam's Razor
     -  Parkinson's Law
     -  Premature Optimization Effect
     -  Putt's Law
     -  Reed's Law
     -  The Law of Conservation of Complexity (Tesler's Law)
     -  The Law of Leaky Abstractions
     -  The Law of Triviality
     -  The Unix Philosophy
     -  The Spotify Model
     -  Wadler's Law
Principles
     -  Wheaton's Law
     -  The Dilbert Principle
     -  The Pareto Principle (The 80/20 Rule)
     -  The Peter Principle
     -  The Robustness Principle (Postel's Law)
     -  SOLID
     -  The Single Responsibility Principle
     -  The Open/Closed Principle
     -  The Liskov Substitution Principle
     -  The Interface Segregation Principle
     -  The Dependency Inversion Principle
     -  The DRY Principle
     -  The KISS principle
     -  YAGNI
```

Sample execution #2

```bash
$ ./hacker-laws-cli random

-----------------------------------------------------
The Peter Principle
-----------------------------------------------------

The Peter Principle on Wikipedia

> People in a hierarchy tend to rise to their "level of incompetence".
>
> _Laurence J. Peter_

A management concept developed by Laurence J. Peter, the Peter Principle observes that people who are good at their jobs are promoted, until they reach a level where they are no longer successful (their "level of incompetence". At this point, as they are more senior, they are less likely to be removed from the organisation (unless they perform spectacularly badly) and will continue to reside in a role which they have few intrinsic skills at, as their original skills which made them successful are not necessarily the skills required for their new jobs.

This is of particular interest to engineers - who initial start out in deeply technical roles, but often have a career path which leads to _managing_ other engineers - which requires a fundamentally different skills-set.

See Also:

- The Dilbert Principle
- Putt's Law

-----------------------------------------------------
         github.com/dwmkerr/hacker-laws by Dave Kerr
-----------------------------------------------------
```
