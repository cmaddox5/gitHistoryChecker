# Pre Deploy Check

This little program will check each local project for commits made by team members in the last 2 weeks (duration customizable).

When ran on a "monday before deploy" it should give a good picture of which projects need to be tagged by your team members.

## Configuration

An example configuration is provided in the `config` directory. Rename the file to `config.json` and supply your team members names
and repository file locations you'd like to track. These must be absolute file system paths.

## Dependencies

This project requires `go-git`, you can find installation instructions on their github page. https://github.com/src-d/go-git