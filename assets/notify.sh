#!/bin/bash
#
# This script is used to notify the user when the pomodoro is done.
#

SOUND_FILE="/usr/share/fynodoro/notification.mp3"

if ! [ -r "$SOUND_FILE" ]; then
    echo "Missing sound file: $SOUND_FILE"
    exit 1
fi

if command -v ffplay &> /dev/null; then
    ffplay "$SOUND_FILE" -nodisp -nostats -hide_banner -autoexit &> /dev/null
elif command -v mplayer &> /dev/null; then
    mplayer "$SOUND_FILE" -really-quiet -noconsolecontrols -nolirc -nojoystick -nomouseinput -nogui &> /dev/null
elif command -v mpv &> /dev/null; then
    mpv "$SOUND_FILE" --really-quiet --no-audio-display &> /dev/null
elif command -v play &> /dev/null; then
    play "$SOUND_FILE" &> /dev/null;
else
    echo "No player found. Please install ffplay, mplayer, mpv, or sox."
    exit 1
fi
