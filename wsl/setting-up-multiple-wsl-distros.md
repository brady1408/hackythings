# Setting up multiple WSL distros

I have really been liking developing in wsl it is a great way to get a linux feel on windows and with wsl 2 it feels pretty complete. One thing I did notice however was that if you setup too many helpers in wsl (npm, rvm, etc...) these seem to slow the loading of directories considerably, and the more you seem to load that need to watch folders the slower this gets. To this end I have done a lot of research into setting up multiple instances of wsl to use as dev environments. I can then have a distro for Golang, Elixir, Ruby, etc... all seperate and clean keeping these performance issues to a minimum. Here is how we go about it. I should mention that I use Ubuntu pretty exclusively so if you need a different distro these instructions may change slightly.

## Getting a distro

## Duplicate an existing distro

If you have a distro in wsl that is setup and that you want to use a your base for future distros you can easily do that by exporting this distro.

From the command prompt you can use this command to export your distro to a compressed file

> `wsl --export <distro name> <export file name>`

## Getting a fresh image

If you need to setup a fresh instance of Ubuntu for WSL you can download the image from here https://cloud-images.ubuntu.com/focal/current/ look for the wsl image that matches your environment, for most it will be `focal-server-cloudimg-amd64-wsl.rootfs.tar.gz`

 - import the cloud image into wsl

    > `wsl --import <distro name> <install directory> <cloud image path>`

 - start a wsl instance from this distro

    > `wsl -d <distro name>`

- from with in wsl you will want to create a new user

    > `sudo useradd username`

- add user to sudoers

    > `usermod -aG sudo username`

- if you are running docker using the new windows docker desktop, you can add docker to this wsl instance by doing the following.

  - right click the docker icon in your system tray
  - select settings
  - select the `Resources` side menu group
  - select `WSL INTEGRATION` sub list
  - flip the switch to on from any WSL2 distros you want to have access to docker
  - click `Apply & Restart`

- add the user to the docker group so you don't have to sudo every docker command

    > `sudo usermod -aG docker username`

- update the default user that is used when running this wsl instance you could just start the instance with the -u flag in wsl but if you are using vs code to develop from this instance then you will notice that it starts up with the root user instead of the one you just created. To solve this follow these steps.

  - open the new file for edit

    > `sudo nano /etc/wsl.conf`

  - update it with the user you want wsl to use

        [user]
        default=username

  - now exit wsl

    > `exit`

  - shutdown the wsl instance you are updating the default user for

    > `wsl --shutdown <distro name>`

  - now enter your wsl instance as normal

    > `wsl -d <distro name>`

At this point you should be setup with your new wsl instance and ready to develop. the vs code remote - WSL extension is fantastic. https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-wsl
