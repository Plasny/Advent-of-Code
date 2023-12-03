{
  description = "Advent of Code in multiple languages";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs/nixos-23.05";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }: 
    flake-utils.lib.eachDefaultSystem (system:
      let 
        pkgs = nixpkgs.legacyPackages.${system};

        cargoBuildInputs = with pkgs; lib.optionals stdenv.isDarwin [
          darwin.apple_sdk.frameworks.CoreServices
        ];

        run-python = pkgs.writeShellScriptBin "run" ''
          echo "Running task 1:"
          python3 task1.py input.in
          echo "Running task 2:"
          python3 task2.py input.in
        '';

        run-rust = pkgs.writeShellScriptBin "run" ''
          echo "Running task 1:"
          cargo run --bin task1 input.in
          echo "Running task 2:"
          cargo run --bin task2 input.in
        '';
      in {
        devShells = {
          py = pkgs.mkShell {
            buildInputs = with pkgs; [
              python3
              run-python
            ];

            shellHook = ''
                echo "hello from python"
            '';
          };

          rs = pkgs.mkShell {
            RUST_SRC_PATH = "${pkgs.rust.packages.stable.rustPlatform.rustLibSrc}";

            buildInputs = with pkgs; [
              cargo
              rustc
              rust-analyzer
              rustfmt
              clippy
              run-rust
            ] ++ cargoBuildInputs;

            shellHook = ''
                echo "hello from rust"
            '';
          };
        };
      }
    );
}