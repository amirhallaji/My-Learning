<p>
<img src="images/git-logo.jpg">
<h1 align="center">Git</h1>
</p>

# Table of Content

  - [Some Problems](#some-problems)
  - [What is VCS](#what-is-vcs)
  - [Advantages](#advantages)
  - [VCS Types](#vcs-types)
  - [Git commands](#git-commands)
    - [config](#config)
    - [initialize](#initialize)
    - [adding to stage](#adding-to-stage)
    - [status](#status)
    - [commit](#commit)
    - [Deleting from stage area](#deleting-from-stage-area)
    - [Showing differences](#showing-differences)
    - [showing log](#showing-log)
    - [Resting](#resting)
      - [part1](#part1)
      - [part2](#part2)
      - [part3](#part3)
    - [Remote](#remote)
    - [Push](#push)
    - [Pull](#pull)
    - [Clone](#clone)
  - [.gitignore](#gitignore)
  - [Branches](#branches)
  - [Branch usage](#branch-usage)
  - [commands](#commands)
    - [Creating Branch](#creating-branch)
    - [Showing all Branches](#showing-all-branches)
    - [Changing between branches](#changing-between-branches)
    - [Deleting Branch](#deleting-branch)
  - [Merge](#merge)
    - [command](#command)
  - [Conflicts](#conflicts)


## Some Problems
...

## What is VCS
In software engineering, version control is a class of systems responsible for managing changes to computer programs, documents, large web sites, or other collections of information. Version control is a component of software configuration management.

## Advantages

## VCS Types

## Git commands
### config
```
git config --global user.name "your name"
git config --global user.email "your email"
```

This will save your name and email and will be shown in your commit history.
### initialize
```commandline
git init
```
This command initializes an empty repository. It's the first step to create a repository.

### adding to stage
```commandline
git add <filename>
```
- In case you want to add all of the files to the stage area, you should use one of the commands below.
```commandline
git add -A
git add .
```

### status
```
git status
```
By this command, we understand the conditions of the files. For instance, if we had used ```git add -A``` before, the files are shown in green color in terminal, otherwise they are shown in red.


### commit
```
git commit -m "commit message"
```
By using this command, all of the changes are saved in repository. After using this command, use ```git status``` and you'll see that the phrase below will be shown.

<p align="center"><u>nothing to commit, working tree clean</u></p>

There is one other command you can use. The command below will both add and commit, but you must pay attention that you can only use it when you have added your files to the stage area at least once.
```
git commit -a -m "commit message"
```

### Deleting from stage area

```
git rm --cached <filename>
```
This command will remove your file from stage area.

```
git rm --cached -r .
```
This will remove all files from stage area.

**Note**:If you use the second command, you won't be able to use command number , because the files are removed from stage area.

### Showing differences 
```
git diff
```
This will show the differences between **now** and the **last time adding to the stage area**. 


### showing log
```
git log
```
This will show you the history of commits in reversal order. Some fields such as name, email, and the commit history will be shown. By using ```git log --oneline```, you can see history of commits in line. For instance, the command ```git log -2``` will show you the last 2 commits.

### Resting

#### part1

#### part2

#### part3

### Remote

### Push

### Pull

### Clone

## .gitignore

## Branches

## Branch usage

## commands

### Creating Branch

### Showing all Branches

### Changing between branches

### Deleting Branch

## Merge

### command

## Conflicts

