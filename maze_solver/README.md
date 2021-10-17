# Maze Solver

A tool designed in GoLang which takes a Map as parameter as a JSON file and extracts the *Shortest Path to Exit*.
The approach used at the core consists of loading data onto a Tree like structure and then running BFS to compute the Shortest path to the exit.




## Input Configuration

To configure the tool to run on a custom Map, specify the Absolute Path to your map file in the `configs/base.json` file under the field `map_file`.

**NOTE:** To test the tool out-of-the-box there are already 3 Maps available in `data/maps` that represent different scenarios: `multiple_exits.json`, `single_exit.json` and `no_exit.json` 





## Running the Tool

To run the tool without IDE support run following one-liner commands in the command prompt respective of OS.

*Windows*

 `bash scripts/run.sh`

 To use bash on windows: [Use bash on Windows](https://www.thewindowsclub.com/how-to-run-sh-or-shell-script-file-in-windows-10)

*Linux*

 `sh scripts/run.sh`

*MacOS* 

 `sh scripts/run.sh`




## Output

The result is displayed in the form a set of moves the Player has to make to reach the Exit fastest. For example:

`[right forward exit]`



## Unit Tests

The code-base also comes integrated with Unit tests.
To run Unit Tests run the following one-liner commands in the command prompt respective of OS.

*Windows*

 `bash scripts/test.sh`

 To use bash on windows: [Use bash on Windows](https://www.thewindowsclub.com/how-to-run-sh-or-shell-script-file-in-windows-10)

*Linux*

 `sh scripts/test.sh`

*MacOS* 

 `sh scripts/test.sh`


