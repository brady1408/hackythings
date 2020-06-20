# Setps to shrink volume

## shutdown wsl distro

```command
wsl --list --verbose
```

If the instance shows are running then shut it down.

```command
wsl --terminiate <distro>
```

I've only had luck shutting down the entire wsl vm as I have a lot of distros installed.

```command
wsl --shutdown
```

After confirming that the instance you want to shrink is stopped then you will need to find the path of the .vhdx. The location for these by default will be.

> %LOCALAPPDATA%\Packages\<distro>\LocalState\ext4.vhdx

the `<distro>` will change depending on which linux you distro you have installed for my ubuntu install the `<distro>` folder began with `CanonicalGroupLimited.Ubuntu` if you are having trouble finding it you can use dir to search the directory.

```command
dir %LOCALAPPDATA%\Packages\*.vhdx /s
```

Now that you have your path you can use `DiskPart` to compact the vdisk.

Start diskpart

```command
diskpart
```

Next select our vdisk using the path provided from your search above.

```command
select vdisk file = "C:\Users\<User>\AppData\Local\Packages\<distro>\LocalState\ext4.vhdx"
```

If the last command was successful you can then compact you disk using the following command.

```command
compact vdisk
```

That should complete all the steps needed to compact your wsl vdisk.
