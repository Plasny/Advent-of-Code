# My solutions to Advent of Code

## Running solutions

Because my solutions are written in different languages I created a `flake.nix`
with development environments for each language used. 

If solution is in `Python` you can use the following commands to run it:

```sh
cd 2023/day01
nix flake develop .#py
run
```

Or if it is in `Rust`:

```sh 
cd 2023/day02
nix flake develop .#rs
run
```

To install `nix` on your system you should follow the steps on 
[nixos.org](https://nixos.org/download)

