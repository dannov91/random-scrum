# Scrum Random Order

This short `Golang` program takes the name of the members of your daily scrum meeting and randomize their order, but organized in groups. It then takes that output and gets copied to your clipboard.

## Usage

The program takes a `.yaml` file as argument, where the groups and members of the scrum meeting are listed without any specific order:

```bash
go run main.go ./scrum-groups.yml
```

Expected Output:

```bash
Success! Copied to clipboard.
```

## Notes

As mentioned above, groups order are randomized as well as their members. In the `scrum-groups.yml` example file, groups order may vary is as follows:

* QA, Development, Design
* Development, QA, Design
* Design, Development, QA
* Etc...

This is done this way in order to keep the members of the same group together while being internally randomized in each group.
