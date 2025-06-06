{
  description = "";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs =
    {
      self,
      flake-utils,
      nixpkgs,
      ...
    }:
    flake-utils.lib.eachDefaultSystem (
      system:
      let
        pkgs = import nixpkgs { inherit system; };
      in
      {
        packages.default = self.packages.${system}.reumes;
        packages.reumes = pkgs.buildGoModule {
          pname = "reumes";
          version = "0.1.0";

          src = ./.;

          vendorHash = "sha256-ZU6t3B6U1JZwbA205monoknlrowa+5yCAG5U4gW1jO0=";

          meta = {
            homepage = "https://github.com/4jamesccraven/reumes";
            license = pkgs.lib.licenses.gpl3;
            mainProgram = "reumes";
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            go-tools
            gotools

            typst
          ];
        };
      }
    );
}
