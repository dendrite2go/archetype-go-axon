{ buildGoModule
, nix-gitignore
}:

buildGoModule {
  pname = "archetype-go-axon";
  version = "0.0.1";
  src = nix-gitignore.gitignoreSource [] ./.;
  goPackagePath = "github.com/dendrite2go/archetype-go-axon";
  goDeps = ./deps.nix;
  modSha256 = "0x36885x4ccv1m3ff7dcllgdvpx71a5zkfbcjznxl1rpbsa43lw2";
}
