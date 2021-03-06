import argparse
import os

EMPTY_JUPYTER_NOTEBOOK = """{
 "cells": [
  {
   "cell_type": "code",
   "execution_count": null,
   "metadata": {},
   "outputs": [],
   "source": []
  }
 ],
 "metadata": {
  "kernelspec": {
   "display_name": "Python 3",
   "language": "python",
   "name": "python3"
  },
  "language_info": {
   "codemirror_mode": {
    "name": "ipython",
    "version": 3
   },
   "file_extension": ".py",
   "mimetype": "text/x-python",
   "name": "python",
   "nbconvert_exporter": "python",
   "pygments_lexer": "ipython3",
   "version": "3.6.5"
  }
 },
 "nbformat": 4,
 "nbformat_minor": 2
}
"""


def main():
    """Set up files I use for a day of Advent of Code."""
    parser = argparse.ArgumentParser("Bootstrap files for a day of Advent of Code.")
    parser.add_argument("year", type=str)
    parser.add_argument("day", type=int)
    args = parser.parse_args()

    # Make directory for year
    try:
        os.mkdir(args.year)
    except FileExistsError:
        pass

    # Make directory for day
    day_dir = os.path.join(args.year, f"day{args.day:02}")
    try:
        os.mkdir(day_dir)
    except FileExistsError:
        # We've already set up files for this day
        print(f"adventofcode/{day_dir} already set up :)")
        return

    # Create empty file to paste in example input
    with open(os.path.join(day_dir, f"day-{args.day}-input-example.txt"), "w"):
        pass

    # Create empty Jupyter notebook
    with open(os.path.join(day_dir, f"Day {args.day} - .ipynb"), "w") as file:
        file.write(EMPTY_JUPYTER_NOTEBOOK)

    print(f"Done setting up files for adventofcode/{day_dir} :)")


if __name__ == "__main__":
    main()
