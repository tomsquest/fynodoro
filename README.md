<p align="center">
  <img src="screenshots/hero.jpg" alt="Fynodoro hero"/>
</p>

# Fynodoro, the Pomodoro Widget

![GitHub release)](https://img.shields.io/github/v/release/tomsquest/fynodoro?style=flat-square)
![GitHub status](https://img.shields.io/github/workflow/status/tomsquest/fynodoro/build/master?style=flat-square)

**Fynodoro** is a tiny and cute Pomodoro **Widget**.

## Application

<p align="center">
    <img src="screenshots/app.pimped.png" alt="Fynodoro app screenshot">
</p>

## Settings

<p align="center">
    <img src="screenshots/settings,pimped.png" alt="Fynodoro settings screenshot">
</p>

## Features

- ⏲️ Timer with 25 minutes of work, then short break of 5 min. Do it four times before a long break of 15 minutes. Total: 2 hours. ✨
- 🚀 Configurable Work rounds, Short breaks, Long breaks ⚡
- 🏆 Small download size
- 💼 Releases for linux: Ubuntu/Debian, Fedora/Redhat, and linux binaries

## Changelog

See the [Releases](https://github.com/tomsquest/fynodoro/releases) section.

## Install

#### Downloads binaries

Latest: [![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/tomsquest/fynodoro?style=flat-square)](https://github.com/tomsquest/fynodoro/releases)

See the [Releases](https://github.com/tomsquest/fynodoro/releases) section for downloads.

#### Install Ubuntu/Debian (.deb)

Latest: [![Latest version of 'fynodoro' @ Cloudsmith](https://api-prd.cloudsmith.io/v1/badges/version/tomsquest/fynodoro/deb/fynodoro/latest/a=amd64;d=any-distro%252Fany-version;t=binary/?render=true&show_latest=true)](https://cloudsmith.io/~tomsquest/repos/fynodoro/packages/detail/deb/fynodoro/latest/a=amd64;d=any-distro%252Fany-version;t=binary/#install)

Add the repository and install Fynodoro:

```shell
curl -1sLf 'https://dl.cloudsmith.io/public/tomsquest/fynodoro/setup.deb.sh' | sudo -E bash
sudo apt install fynodoro 
```

_[Complete instructions at CloudSmith.io](https://cloudsmith.io/~tomsquest/repos/fynodoro/packages/detail/deb/fynodoro/latest/a=amd64;d=any-distro%252Fany-version;t=binary/#install)_

#### Install Fedora/Redhat (.rpm)

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
