# m3uspiff
[![codebeat badge](https://codebeat.co/badges/0e4adbee-bafe-4e6c-b72c-a7ae645e4372)](https://codebeat.co/projects/github-com-ibrokemypie-m3uspiff_go-master)
[![License: GPL v2](https://img.shields.io/badge/License-GPL%20v2-blue.svg)](https://www.gnu.org/licenses/old-licenses/gpl-2.0.en.html)

An M3U to XSPF playlist converter.

## Usage
``./m3uspiff [playlist.m3u]``

Converts M3U format playlists into XSPF (XML Shareable Playlist Format) playlists, using FFMPEG's ffprobe to augment the new playlist with each included song's metadata when available.

## Requirements

* ffmpeg
