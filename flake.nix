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

        run-c = pkgs.writeShellScriptBin "run" ''
          echo "Running task 1:"
          gcc task1.c -O3 -o task1.out
          ./task1.out input.in
          rm task1.out
          echo "Running task 2:"
          gcc task2.c -O3 -o task2.out
          ./task2.out input.in
          rm task2.out
        '';

        run-cs = pkgs.writeShellScriptBin "run" ''
          echo "Running task 1:"
          csc task1.cs > /dev/null
          mono task1.exe input.in
          rm task1.exe
          echo "Running task 2:"
          csc task2.cs > /dev/null
          mono task2.exe input.in
          rm task2.exe
        '';

        run-go = pkgs.writeShellScriptBin "run" ''
          echo "Running both tasks:"
          go run *.go input.in
        '';

        run-js = pkgs.writeShellScriptBin "run" ''
          echo "Running task 1:"
          bun run task1.js input.in
          echo "Running task 2:"
          bun run task2.js input.in
        '';
      in {
        devShells = {
          c = pkgs.mkShell {
            buildInputs = with pkgs; [
              gcc
              run-c
            ];

            shellHook = ''
                echo "hello from c"
            '';
          };

          cs = pkgs.mkShell {
            buildInputs = with pkgs; [
              mono
              # dotnet-sdk_8 # for lsp
              run-cs
            ];

            shellHook = ''
                echo "hello from c#"
            '';
          };

          go = pkgs.mkShell {
            buildInputs = with pkgs; [
              go
              run-go
            ];

            shellHook = ''
                echo "hello from go"
            '';
          };

          js = pkgs.mkShell {
            buildInputs = with pkgs; [
              bun
              run-js
            ];

            shellHook = ''
                echo "hello from js"
            '';
          };

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
