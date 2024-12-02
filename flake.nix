{
  description = "Advent of Code";

  inputs = {
    nixos.url = "nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = {
    nixos,
    flake-utils,
    ...
  }:
    flake-utils.lib.eachDefaultSystem (system: let
      pkgs = import nixos {inherit system;};
    in {
      devShell = pkgs.mkShell {buildInputs = with pkgs; [go golangci-lint];};
    });
}
