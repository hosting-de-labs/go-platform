{
  description = "hosting.de netbox tools collection";

  inputs = {
    flake-parts.url = "github:hercules-ci/flake-parts";
    nixpkgs.url = "github:nixos/nixpkgs/nixos-24.05";
  };

  outputs = inputs:
    inputs.flake-parts.lib.mkFlake { inherit inputs; } {
      systems = [ "x86_64-linux" ];
      perSystem = { self', system, lib, config, pkgs, ... }: {
        devShells.default = pkgs.mkShell {
          name = "crispd development shell";
          nativeBuildInputs = with pkgs; [
            go_1_21
            gofumpt
            golangci-lint
          ];
        };
      };
    };
}
