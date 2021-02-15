<p>
<img src="images/git-logo.jpg">
<h1 align="center">Git</h1>
</p>

# Table of Content

  - [Some Problems](#some-problems)
  - [What is VCS](#what-is-vcs)
  - [Advantages](#advantages)
  - [VCS Types](#vcs-types)
    - [CVCS](#cvcs)
    - [DVCS](#dvcs)
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



## What is VCS
In software engineering, version control is a class of systems responsible for managing changes to computer programs, documents, large web sites, or other collections of information. Version control is a component of software configuration management.

## Advantages of VCS
The answer ot the questions below are the advantages of VCS:

- Have you ever tried to return to previous versions in your projects?
- Have you ever lost your code?
- Have you shared your code with others?
- Have you felt the need of reviewing the history of the project?

## VCS Types
### CVCS
<p align="center">
<img src="images/cvcs.png" width="480" height="300">
<br>
<b>Centrialiazed version control system</b>
</p>

### DVCS

<p align="center">
<img src="images/dcvs.png" height="300" width="480">
<br>
<b>Centrialiazed version control system</b>
</p>

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
This will show you the history of commits in reversal order. Some fields such as name, email, and the commit IDs will be shown. By using ```git log --oneline```, you can see history of commits in line. For instance, the command ```git log -2``` will show you the last 2 commits.

### Restting
#### part1
```
git checkout --<file name>
```
It might happen that you want to return to your previous state. In this state, you haven;t entered stage area.
In case you want to extend it to all files, use ```.``` instead of file name.
It acts like ```ctrl + z```.

#### part2
```
git reset HEAD <filename>
```
You have entered stage area, but you want to remove your changes there.

**Note**: Although you will quit from stage area, but you still can use command number --.

<p align="center"> <img src="images/reset1.png"</p>

#### part3
In this state, you want to return to a specific commit. It is called reset commit. For this, you have to know the ID of that commit.
```
git reset --hard <commit id>
```
For getting commit IDs, you should refer to command number --.

<p align="center"> <img src="images/reset3.png"</p>


### Remote
```
git remote add origin <repository link>
```
Use this command if you want to remote to your project in github or gitlab.

### Push
```
git push origin master
```
This will upload your project to your github/gitlab repository.

### Pull
```
git pull origin master
```
This will fetch the latest updates made by your collaborators or in other ways.

### Clone
```
git clone <repository link>
```
This will download the repository from the link given.

## .gitignore
There are some files or directories that there is no need to be tracked by git.
Files such as ```.class``` are unncessary files which may cause some problems after being fetched.

```
touch .gitignore
```
This will create a .gitignore file for you. In this file, you write files or directories that you don't want to be tracked by git.

<p align="center"><img src="images/gitignore.png"</p>
<p align="center"><img src="images/gitign.png"</p>
<p align="center">gitignore file content</p>

## Branches

The default branch is master. You can create different branches in your project to use them.

<p align="center"><img src="images/branch1.png"></p>


## Branch usage
Suppose you're woking on the project and have developed different versions of that. You have to keep each version seperate from others. Beacuse we have learnt git by now, we keep different versions in different branches. Now we can keep the first version of the project and start to develop the next versions.

<p align="center"><img src="images/branch.png"></p>


## commands

### Creating Branch

```
git branch <name>
```
The command above will create a branch named **name**.


### Showing all Branches

```
git branch -a
```
This will show you the list of all branches.


### Changing between branches
```
git checkout <name>
```
By using this command, you will switch from the current branch to the **name** branch.


### Deleting Branch

```
git branch -d <name>
```
This will delete the **name** branch.

## Merge
Suppose you and your friend are developing an application. You do the back-end and your friend do the front-end. At last, These two should be merged.

### command
```
git merge <name>
```
Suppose you want to merge branch **name** with **master**. Use this command in master branch.
## Conflicts

Some conflicts may occur when you remote to your server and you won't be able to push. So first you must fetch tha latest update, merge the conflicts and then push your commit.

