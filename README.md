# batch-renamer

Batch-Renamer is a tool to facilitate - well - renaming files in a batch.
It provides an executable command that effects the current working directory in which it is run.
When issued for the first time inside a directory, it creates a list of all files (not recursing into subdirectories).
This list is then written to two separate files: `names` & `.names.old`.
Then, `names` can be modified in place to the newly desired names.
*Note that the order of the lines must not be changed or it will end in a desaster.* (This is because the order of files must be the same in `names` and `.names.old`.)
When the file is modified to liking, `batch-renamer` is executed again, this time moving the files according to the name files which are deleted after the task is completed.

So the entire workflow of `batch-renamer` looks like this:
1. `cd` into the directory of choice
2. execute `batch-renamer` once
3. change names in `names` file
4. execute `batch-renamer` a second time

Although this project has been tested on many occasions, it is provided as is without any guarantee on functionality or safety.