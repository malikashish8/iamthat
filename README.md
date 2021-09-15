# IAmThat

> Easy switch between differrent profiles for differrent tools such as ssh and git

## How does it work
* Details of all files to be switched for any profile are written in user folder in `~/.config/iamthat/profile.yaml`.
* Location of config files for various profiles list in `profile.yaml` is relative `~/.config/iamthat` folder.
* A state file is created to hold the name of the profile currently in use
* If only two profiles are configured switch to the other profile happens.
* If more that two profiles are configured then user is prompted to select the profile.

## Sample `profile.yaml`

```yaml
profile:
  config-template:
    - name: "git-user-config"
      to-path: "~/.gitconfig"
    - name: "ssh-user-config"
      to-path: "~/.ssh/config"
  config:
    - name: "personal"
      config-file:
        - type: "git-user-config"
          from-path: "config/git/personal"
        - type: "ssh-user-config"
          from-path: "config/ssh/personal"
    - name: "office"
      config-file:
        - type: "git-user-config"
          from-path: "config/git/office"
        - type: "ssh-user-config"
          from-path: "config/ssh/office"
```

If a from-file is missing the behaviour is to delete the current to-file to reset settings to default.

## ToDo

* Improve readability of `profile.yaml` by choosing better names or restructuring it. 
