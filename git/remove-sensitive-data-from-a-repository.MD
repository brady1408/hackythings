## [Removing sensitive data from a repository](https://help.github.com/en/github/authenticating-to-github/removing-sensitive-data-from-a-repository)

**If you commit sensitive data, such as a password or SSH key into a Git repository, you can remove it from the history. To entirely remove unwanted files from a repository's history you can use either the `git filter-branch` command or the BFG Repo-Cleaner open source tool.**

The `git filter-branch` command and the BFG Repo-Cleaner rewrite your repository's history, which changes the SHAs for existing commits that you alter and any dependent commits. Changed commit SHAs may affect open pull requests in your repository. We recommend merging or closing all open pull requests before removing files from your repository.

You can remove the file from the latest commit with `git rm`. For information on removing a file that was added with the latest commit, see "Removing files from a repository's history."

1. clone the repo
2. cd YOUR-REPOSITORY
3. Run the following command, replacing `PATH-TO-YOUR-FILE-WITH-SENSITIVE-DATA` with the **path to the file you want to remove, not just it's filename.** These arguments will: 

    - Force Git to process, but not check out, the entire history of every branch and tag
    - Remove the specified file, as well as any empty commits generated as a result
    - **Overwrite your existing tags**

		```bash
		$ git filter-branch --force --index-filter \
		"git rm --cached --ignore-unmatch PATH-TO-YOUR-FILE-WITH-SENSITIVE-DATA" \
		--prune-empty --tag-name-filter cat -- --all
		> Rewrite 48dc599c80e20527ed902928085e7861e6b3cbe6 (266/266)
		> Ref 'refs/heads/master' was rewritten
		```

	>**Note:** If the file with sensitive data used to exist at any other paths (because it was moved or renamed), you must run this command on those paths, as well.

4. Add your file with sensitive data to `.gitignore` to ensure that you don't accidentally commit it again.

	```bash
	$ echo "YOUR-FILE-WITH-SENSITIVE-DATA" >> .gitignore
	$ git add .gitignore
	$ git commit -m "Add YOUR-FILE-WITH-SENSITIVE-DATA to .gitignore"
	> [master 051452f] Add YOUR-FILE-WITH-SENSITIVE-DATA to .gitignore
	>  1 files changed, 1 insertions(+), 0 deletions(-)
	```

5. Double-check that you've removed everything you wanted to from your repository's history, and that all of your branches are checked out.
6. Once you're happy with the state of your repository, force-push your local changes to overwrite your GitHub repository, as well as all the branches you've pushed up:

	```bash
	$ git push origin --force --all
	> Counting objects: 1074, done.
	> Delta compression using 2 threads.
	> Compressing objects: 100% (677/677), done.
	> Writing objects: 100% (1058/1058), 148.85 KiB, done.
	> Total 1058 (delta 590), reused 602 (delta 378)
	> To https://github.com/YOUR-USERNAME/YOUR-REPOSITORY.git
	>  + 48dc599...051452f master -> master (forced update)
	```

7. In order to remove the sensitive file from [your tagged releases](https://help.github.com/en/articles/about-releases), you'll also need to force-push against your Git tags:

	```bash
	$ git push origin --force --tags
	> Counting objects: 321, done.
	> Delta compression using up to 8 threads.
	> Compressing objects: 100% (166/166), done.
	> Writing objects: 100% (321/321), 331.74 KiB | 0 bytes/s, done.
	> Total 321 (delta 124), reused 269 (delta 108)
	> To https://github.com/YOUR-USERNAME/YOUR-REPOSITORY.git
	>  + 48dc599...051452f master -> master (forced update)
	```

8. Contact [GitHub Support](https://support.github.com/contact) or [GitHub Premium Support](https://premium.githubsupport.com/) asking them to remove cached views and references to the sensitive data in pull requests on GitHub.
9. Tell your collaborators to [rebase](https://git-scm.com/book/en/Git-Branching-Rebasing), not merge, any branches they created off of your old (tainted) repository history. One merge commit could reintroduce some or all of the tainted history that you just went to the trouble of purging.
10. After some time has passed and you're confident that git filter-branch had no unintended side effects, you can force all objects in your local repository to be dereferenced and garbage collected with the following commands (using Git 1.8.5 or newer):

	```bash
	$ git for-each-ref --format="delete %(refname)" refs/original | git update-ref --stdin
	$ git reflog expire --expire=now --all
	$ git gc --prune=now
	> Counting objects: 2437, done.
	> Delta compression using up to 4 threads.
	> Compressing objects: 100% (1378/1378), done.
	> Writing objects: 100% (2437/2437), done.
	> Total 2437 (delta 1461), reused 1802 (delta 1048)
	```

	>Note: You can also achieve this by pushing your filtered history to a new or empty repository and then making a fresh clone from GitHub.

### Avoiding accidental commits in the future
---

There are a few simple tricks to avoid committing things you don't want committed:

- Use a visual program like [GitHub Desktop](https://desktop.github.com/) or [gitk](https://git-scm.com/docs/gitk) to commit changes. Visual programs generally make it easier to see exactly which files will be added, deleted, and modified with each commit.
- Avoid the catch-all commands `git add .` and `git commit -a` on the command line—use `git add filename` and `git rm filename` to individually stage files, instead.
- Use `git add --interactive` to individually review and stage changes within each file.
- Use `git diff --cached` to review the changes that you have staged for commit. This is the exact diff that `git commit` will produce as long as you don't use the `-a` flag.

### [Further reading](https://help.github.com/en/github/authenticating-to-github/removing-sensitive-data-from-a-repository#further-reading)

- [git filter-branch main page](https://git-scm.com/docs/git-filter-branch)
- [Pro Git: Git Tools - Rewriting History](https://git-scm.com/book/en/Git-Tools-Rewriting-History)