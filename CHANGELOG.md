# Changelog

All notable changes to this project are documented in this file.

## [0.0.2] - 2026-11-11

### Added

- Open command now accepts a today option that opens the file for the current date in TextEdit, creating it if needed.
- Save command that archives the source file, resets it to a blank document, commits, and reopens it in TextEdit. With a title, the archive is saved in the folder for the current date under that title; otherwise it is saved as a timestamped copy.
- Commit command that stages and commits all changes in the textedit directory.
- Version flags that print the version.
- Usage message listing the supported commands, shown when no valid command is given.

### Changed

- Running textedit with no arguments now prints the list of supported commands.

## [0.0.1] - 2026-11-10

### Added

- Open command that opens a file in TextEdit.
