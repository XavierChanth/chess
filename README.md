# chess

My weekend project - A (soon to be) chess engine implemented from scratch

## Goals

- [ ] Two representations for state
  - [ ] Algebraic notation (e.g. e4 e5 ...)
  - [ ] "Board state" - Represent a momentary state in a game
    - Why? makes it such that transposed game states are agnostic to this data
      representation
- [ ] Move validators
- [ ] Win / draw condition checkers
- [ ] Want to be able to convert between representations easily
- [ ] TUI to be able to play PvP
  - [ ] Algebraic notation in a dialogue
  - [ ] moving around the board with hjkl/arrows
  - [ ] mouse (maybe)
- [ ] Do ML with a NN on board state to build a chess bot
- [ ] PvBot - vs NN
- [ ] Train bot to play itself
- [ ] Online support
