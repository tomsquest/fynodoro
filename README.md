<p align="center">
  <img src="https://github.com/tomsquest/fynodoro/blob/master/screenshots/app.png?raw=true" alt="Fynodoro screenshot"/>
</p>

# Fynodoro

![build](https://github.com/tomsquest/fynodoro/actions/workflows/checks.yml/badge.svg)

Fynodoro is the little widget for doing Pomodoro.

## Development

Run:

```shell
go generate && go run .
```

Package tne app with Fyne (requires: `go get fyne.io/fyne/v2/cmd/fyne`):

```shell
fyne package -os linux -icon assets/Icon.png
```

## Credits

- Icons made by [Freepik](https://www.freepik.com) from [Flaticon](https://www.flaticon.com)
- Screenshot pimped with [PrettySnap](https://prettysnap.app)
