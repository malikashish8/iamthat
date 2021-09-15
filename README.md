# IAmThat

> Easy switch between differrent profiles for differrent tools such as ssh and git

## How does it work
* Details of all files to be switched for any profile are written in user folder in `~/.config/iamthat/profile.yaml`.
* Location of config files for various profiles list in `profile.yaml` is relative `~/.config/iamthat` folder.
* Decision of which config is currently used is based on which `~/.ssh/config` file is used. Therefore, ensure that each profile has a unique file.

## Sample `profile.yaml`

```yaml
profile:
  - name: "personal"
    git-user-config: "config/git/personal"
    ssh-user-config: "config/ssh/personal"
  - name: "office"
    git-user-config: "config/git/office"
    ssh-user-config: "config/ssh/office"
```

## Supported Configs

| In profile.yaml | Location on System |
| --------------- | ------------------ |
| git-user-config | $HOME/.gitconfig   |
| ssh-user-config | $HOME/.ssh/config  |

Not all configs for all profiles need to be present in the config folder. If a config for target profile is missing the behaviour is to delete the current file to reset settings to default.

## ToDo

* Use `config-def.yaml` as the file which holds the config mappings rather than hard coding them
* Allow using more that 2 configs
* Persist current config file in iamthat config so that updated config does not confuse it.
