#                 GIT COMMANDS

## GIT BASICS

1. **git init** <directory>
    
    -->Creates a empty Git repositry in specified directory
    -->Run with no arguments to initialize the current directory as the Git repository.
    
    `ex:git init (creates a empty repository)`

2. **git clone** <repo>
    
    -->Clones repo present at the <repo> into your local machine
    -->Original repo can be located at the local file system or the remote machine via HTTP or SSH .
    
    `ex:git clone https://github.com/TutorialEdge/go-grpc-services-course.git`
    
3. **git config user.name** <name>
    
    -->Define author name to be used for all the commits in the current repo.
    
    
4. **git add** <directory>
    
    
    -->Stages all the changes in <directory> for the next commit. 
    -->replace <directory> with <file> for a specific change.
    
    `ex: git add . (stages all the files inside the current directory)`
    
5. **git status**  
    
    -->List which files are staged,unstaged and untracked.
    
    ex: `git status `
    
6. **git log** 
    
    -->Display the entire commit history using the default format
    
    ex: `git log`

7. **git diff**
    
    -->Shows unstaged changes  between your index and working directory
    
   command:  `git diff`
    
    
    
## UNDOING CHANGES 
    
1. **git revert** 
    
    -->Creates a new repo that undoes all the changes  made in the <commit>,then apply it to the current branch
    
2. **git reset** <file>
    
    -->Remove <file> from the staging area,but leaves the working directory unchanged.This unstages the a file without overwriting the any changes.
    
3 . **git clean -n** 
    
    -->shows which files would be removed from the working directory.
    -->Inorder to remove to replace -n flag with -f
    
    Command: `git clean -n`  (shows which file would be removed)
    Command : `git clean -f` (removes the file)
    
    
## REWRITING GIT HISTORY
    
1. **git commit --amend**
    
    --> Replace the last commit with the staged changes and last commit combined. Use with nothing staged to edit the last commit’s message.
    
    
2. **git rebase** <base>
    
    -->Rebase the current branch onto <base>. <base> can be a commit ID,
branch name, a tag, or a relative reference to HEAD.
    
3. **git reflog**
    
    -->Show a log of changes to local repository's
    HEAD.

## GIT BRANCHES
    
1. **git branch**
    
    -->List all the branches present in the repo.
    
    Command : `git branch`
    
2. **git branch** <newbranch> 
    
    -->Adds the new branch to the repo.
    
    Command : `git branch <newbranch>`
    
3. **git checkout -b** <newbranch> 
    
    -->Creates a new branch and checkout to the new branch
    
    Command : `git checkout -b Gitcommands`
4. **git merge <branch> 
    
    -->Merges the <branch> into the current branch
    
## REMOTE REPOSITORIES
    
    
1. **git remote add** <name> <url>
    
    -->Creates a new Connection to remote repo.After adding you can use <name> as a shortcut for <url> in the other commands.

2. **git fetch <remote> <branch>**
    
    -->Fetches a specific branch from the repo.Leave off the <branch> to fetch all the remote refs.

    
3. **git pull <remote>**
    
    -->Fetches the specified remote copy of the current branch and immediately merges into the local copy.
    

4.  **git push**<remote> branch
    
    -->Push the branch to <remote>, along with necessary commits and
objects. Creates named branch in the remote repo if it doesn’t exist

     

## ADDITIONAL OPTIONS
    
### Git Log
    
1. **git log** -limit
    
    -->Limit number of commits by <limit>
    Command : `git log -1`
    
2. **git log --oneline**
    
    -->Condense each commit to one line.
    
    Command: `git log --oneline`
    
3. **git log -p**
    
    -->Display the full diff of each commit
    Command:`git log -p`

4. **git log --stat**
    
    -->Include which files were altered and the relative number of lines were added or deleted.
    
    Command : `git log --stat`
    
5. **git log author="author_name"**
    
    -->Searchs for a commits by a particular author
    
    command:`git log --author="author_name"`
    
6. **git log --grep="Commit_msg"**
    
    -->Searches for a commit that matches the Commit_msg
    
    command:`git log --grep="commit_msg"`
    

7. **git log <since> <then>**
    
    -->Shows the commits that occur between <since> and <then> 

    -->Args can be Commit_id ,Branch_name,head or any kind of revision reference.
    
    ex : `git log  580f9f29f4b862f5912e6895b1838b8e5ebf1c85  161b462ae92a6b328d7113c3d0f23311ef113b55
`

## GIT CONFIG
    
1. **git config --global user.name** <name> 
    
    -->Define author name to be for all the commits by the current user
    
2. **git config --global user.email <email> 
    
    -->Defines author email to be for all the commits by the current user
    

## GIT DIFF
    
1. **git diff HEAD**
    
    -->Shows the difference between working directory and last commit.
    
2.  **git diff --cached**
    
    -->Shows the difference between staged chanages and the last commit.
    

## GIT RESET
    
1. **git reset** 
    
    -->Replace the last commit with the staged changes and last commit

    
2. **git reset --hard**
    
    -->Reset Staging area and working directory to match most recent commit and overwrites all the changes in the working directory.
    
