{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-25.05";
  };

  outputs = { self, nixpkgs }:
    let
      allSystems = [
        "x86_64-linux" # 64-bit Intel/AMD Linux
        "aarch64-linux" # 64-bit ARM Linux
        "x86_64-darwin" # 64-bit Intel macOS
        "aarch64-darwin" # 64-bit ARM macOS
      ];

      forAllSystems = f: nixpkgs.lib.genAttrs allSystems (system: f {
        system = system;
        pkgs = import nixpkgs {
          inherit system;
        };
      });

      # Build app.
      app = { name, pkgs, system }: pkgs.buildGoModule {
        name = name;
        pname = name;
        src = ./.;
        go = pkgs.go;
        vendorHash = "sha256-tARJMLmO7+T1mwtJQSDZgV4KQlgGCpNYySP+Ik0jP44=";
        env = {
          CGO_ENABLED = 0;
        };
        flags = [
          "-trimpath"
        ];
        ldflags = [
          "-s"
          "-w"
          "-extldflags -static"
        ];
      };

      # Development tools used.
      devTools = { system, pkgs }: [
        pkgs.git
        pkgs.go
        pkgs.gopls
      ];

      name = "version";
    in
    {
      # `nix build` builds the app.
      packages = forAllSystems ({ system, pkgs }: {
        default = app { name = name; pkgs = pkgs; system = system; };
      });
      # `nix develop` provides a shell containing required tools.
      devShells = forAllSystems ({ system, pkgs }: {
        default = pkgs.mkShell {
          buildInputs = (devTools { system = system; pkgs = pkgs; });
        };
      });
    };
}
