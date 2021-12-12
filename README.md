<p align="center">
  <img src="screenshots/hero.jpg" alt="Fynodoro hero"/>
</p>

# Fynodoro, the Pomodoro Widget

![build](https://github.com/tomsquest/fynodoro/actions/workflows/build.yaml/badge.svg)

<img align="right" src="screenshots/app.png" alt="Fynodoro app screenshot">

Fynodoro is a tiny and cute Pomodoro **Widget**.

## Features

- :rocket: it counts from 25:00 to 0! And then from 5:00 to 0! :zap:

## Install

#### Downloads binaries

See the [Releases](https://github.com/tomsquest/fynodoro/releases) section for downloads

#### Deb

First, add the apt repo:

```shell
sudo echo "deb [trusted=yes] https://apt.fury.io/tomsquest/ /" > /etc/apt/sources.list.d/fury.fynodoro.list
sudo apt update
```

Then, install Fynodoro:

```shell
sudo apt install fynodoro 
```

#### Rpm

To enable, add the following file `/etc/yum.repos.d/fury.fynodoro.repo`:

```
[fury]
name=Gemfury Private Repo
baseurl=https://yum.fury.io/tomsquest/
enabled=1
gpgcheck=0
```

## TODO

- [ ] Long breaks
- [ ] Pico/Nano/Normal UI
- [ ] Tons of options
- [ ] Release Deb/Rpm/Snap :smile:

## Development

Run:

```shell
go run .
```

## Credits

- Icons made by [Freepik](https://www.freepik.com) from [Flaticon](https://www.flaticon.com)
- Screenshot pimped with [PrettySnap](https://prettysnap.app)
