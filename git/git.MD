# Git things

## Setup Multiple Git Accounts on one machine

**Setting up RSA and configuring Git**
Instructions to set up an RSA key and configuring Git with it can be found on most git sites.
github.com is here [Adding a new SSH key to your GitHub account](https://help.github.com/en/articles/adding-a-new-ssh-key-to-your-github-account)

## Setting up your machine to use multiple git accounts
First you need to setup your ~/.ssh/config using the following example. This example has multiple accounts one for the user brady and one for the user ydarb. 

**~/.ssh/config**

	Host github.com-brady
		HostName github.com
		User git
		IdentityFile ~/.ssh/brady_github.com_rsa

	Host github.com-brady
		HostName github.com
		User git
		IdentifyFile ~/.ssh/brady_github.com_rsa

  

## usage
**clone repo**

    git@github.com-brady:brady/project_a.git

**cd into project and set the git user**

	cd project_a
	git config user.name "brady"
	git config user.email "brady@example.com"

**clone repo from second account**

    git@github.com-ydarb:ydarb/project_b.git

**cd into project and set the git user**

	cd project_b
	git config user.name "ydarb"
	git config user.email "ydarb@example.com"

**use git as normal**

