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

#### Install Debian/Ubuntu (.deb)

First, add the apt repo:

```shell
curl -1sLf \
  'https://dl.cloudsmith.io/public/tomsquest/fynodoro/setup.deb.sh' \
  | sudo -E bash
```

Then, install Fynodoro:

```shell
sudo apt install fynodoro 
```

_Fynodoro uses Cloudsmith to host deb/rpm. [See the complete instructions](https://cloudsmith.io/~tomsquest/repos/fynodoro/setup/#formats-deb)_

#### Install Fedora/Redhat (.rpm)

Add the repository:

```
curl -1sLf \
  'https://dl.cloudsmith.io/public/tomsquest/fynodoro/setup.rpm.sh' \
  | sudo -E bash
```

_Fynodoro uses Cloudsmith. [See the complete instructions](https://cloudsmith.io/~tomsquest/repos/fynodoro/setup/#repository-setup-yum)_

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
