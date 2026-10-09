{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-26.05";
  };

  outputs =
    { self, nixpkgs }:
    let
      allSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        "x86_64-darwin" # 64-bit Intel macOS
        "aarch64-darwin" # 64-bit ARM macOS
      ];

      forAllSystems = f: nixpkgs.lib.genAttrs allSystems (system: f (import nixpkgs { inherit system; }));
    in
    {
      # `nix build` builds the app.
      packages = forAllSystems (pkgs: rec {
        version = pkgs.buildGoModule {
          name = "version";
          src = ./.;
          vendorHash = "sha256-tARJMLmO7+T1mwtJQSDZgV4KQlgGCpNYySP+Ik0jP44=";
          env.CGO_ENABLED = 0;
          flags = [ "-trimpath" ];
          ldflags = [
            "-s"
            "-w"
            "-extldflags -static"
          ];
        };
        default = version;
      });
      # `nix develop` provides a shell containing required tools.
      devShells = forAllSystems (pkgs: {
        default = pkgs.mkShell {
          buildInputs = [
            pkgs.git
            pkgs.go
            pkgs.gopls
          ];
        };
      });
    };
}
