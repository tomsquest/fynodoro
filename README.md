<p align="center">
  <img src="screenshots/hero.jpg" alt="Fynodoro hero"/>
</p>

# **Fynodoro** is a tiny and cute Pomodoro **Widget**.

![GitHub release)](https://img.shields.io/github/v/release/tomsquest/fynodoro?style=flat-square)
![GitHub status](https://img.shields.io/github/workflow/status/tomsquest/fynodoro/build/master?style=flat-square)

## Screenshots

### Timer

<p align="center">
    <img src="screenshots/app.pimped.png" alt="Fynodoro app screenshot">
</p>

### Settings

<p align="center">
    <img src="screenshots/settings,pimped.png" alt="Fynodoro settings screenshot">
</p>

## Features

- ⏲️ Pomodoro Timer ✨
- 🗒️ Configurable: work duration, short break/long break, work rounds 🖊️
- 🏆 Small download size
- 💼 Releases for Ubuntu/Debian, Fedora/Redhat, and as linux binaries

## Configuration

The Pomodoro technique defaults to 4 work rounds of 25 minutes, with a 5 minutes pause ("short break") in-between and a final 15 minutes pause (the "long" break), for a total of 2 hours (4x25m Work + 3x5m Short breaks + 1x15m Long break).

You can **configure**:

- the duration in minutes of the Work period (default: `25` minutes)
- the duration in minutes of the Short breaks (default: `5` minutes)
- the duration in minutes of the Long breaks (default: `15` minutes)
- the number of Work rounds before a long break (default: `4` rounds)

You can **disable** Long breaks by setting the duration of Long breaks to `0` or the number of work rounds to `0`. This will make the timer do a Work period, then a Short break, and so-on, and never do Long break.

You can **disable** Short breaks by setting the duration of Short breaks to `0`. This will make the timer do a Work period, then a Long break, and so-on, and never do Short break.

Tips: you can **disable** both Short and Long breaks by setting them to `0`. The timer will then act as a ticker, notifying you after each Work period.

## Changelog

See the [Releases](https://github.com/tomsquest/fynodoro/releases) section.

## Install

### Install Ubuntu/Debian (.deb)

Latest: [![Latest version of 'fynodoro' @ Cloudsmith](https://api-prd.cloudsmith.io/v1/badges/version/tomsquest/fynodoro/deb/fynodoro/latest/a=amd64;d=any-distro%252Fany-version;t=binary/?render=true&show_latest=true)](https://cloudsmith.io/~tomsquest/repos/fynodoro/packages/detail/deb/fynodoro/latest/a=amd64;d=any-distro%252Fany-version;t=binary/#install)

Add the repository and install Fynodoro:

```shell
curl -1sLf 'https://dl.cloudsmith.io/public/tomsquest/fynodoro/setup.deb.sh' | sudo -E bash
sudo apt install fynodoro 
```

_[Complete instructions at CloudSmith.io](https://cloudsmith.io/~tomsquest/repos/fynodoro/packages/detail/deb/fynodoro/latest/a=amd64;d=any-distro%252Fany-version;t=binary/#install)_

### Install Fedora/Redhat (.rpm)

Latest: [![Latest version of 'fynodoro' @ Cloudsmith](https://api-prd.cloudsmith.io/v1/badges/version/tomsquest/fynodoro/rpm/fynodoro/latest/a=x86_64;d=any-distro%252Fany-version;t=binary/?render=true&show_latest=true)](https://cloudsmith.io/~tomsquest/repos/fynodoro/packages/detail/rpm/fynodoro/latest/a=x86_64;d=any-distro%252Fany-version;t=binary/#install)

Add the repository and install Fynodoro:

```shell
curl -1sLf 'https://dl.cloudsmith.io/public/tomsquest/fynodoro/setup.rpm.sh' | sudo -E bash
# Choose between:
sudo dnf install fynodoro
sudo yum install fynodoro
sudo microdnf install fynodoro
sudo zypper install fynodoro
```

_[Complete instructions at CloudSmith.io](https://cloudsmith.io/~tomsquest/repos/fynodoro/packages/detail/rpm/fynodoro/latest/a=x86_64;d=any-distro%252Fany-version;t=binary/#install)_

### Downloads Linux binaries

Latest: [![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/tomsquest/fynodoro?style=flat-square)](https://github.com/tomsquest/fynodoro/releases)

See the [Releases](https://github.com/tomsquest/fynodoro/releases) section for downloads.

## Releasing the project

Push a new tag from a clean master:

```shell
git checkout master && git pull
if ! [ -z "$(git status --untracked-files=no --porcelain)" ]; then 
  echo "Warning: there are some local changes"
fi
git tag v1.3.0 && git push --tags
# Publish Release draft: https://github.com/tomsquest/fynodoro/releases
```

## TODO

- [ ] Pico/Nano/Normal UI
- [ ] Tons of options (run script on pomodoro end, notification sound, ...)
- [ ] Release Windows, macOS, Android, IOS versions

## Credits

- Icon made by [Freepik](https://www.freepik.com) from [Flaticon](https://www.flaticon.com/free-icon/tomato_877814)
- Screenshot pimped with [PrettySnap](https://prettysnap.app)
