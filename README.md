<p align="center">
  <img src="screenshots/hero.jpg" alt="Fynodoro hero"/>
</p>

# Fynodoro, the Pomodoro Widget

![build](https://github.com/tomsquest/fynodoro/actions/workflows/build.yml/badge.svg)

<img align="right" src="screenshots/app.png" alt="Fynodoro app screenshot">

Fynodoro is a tiny and cute Pomodoro **Widget**.

## Features

- :rocket: it counts from 25:00 to 0! And then from 5:00 to 0! :zap:

## TODO

- [ ] Release v1.0.0 :smile:
- [ ] Long break
- [ ] Tons of options

## Development

Run:

```shell
go generate && go run .
```

Package the app with Fyne (requires: `go get fyne.io/fyne/v2/cmd/fyne`):

```shell
fyne package -os linux -icon assets/Icon.png
```

## Credits

- Icons made by [Freepik](https://www.freepik.com) from [Flaticon](https://www.flaticon.com)
- Screenshot pimped with [PrettySnap](https://prettysnap.app)
