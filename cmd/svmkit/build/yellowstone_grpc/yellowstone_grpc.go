package yellowstone_grpc

import (
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/abklabs/svmkit/pkg/runner"
	"github.com/abklabs/svmkit/pkg/runner/deb"
	"github.com/abklabs/svmkit/pkg/runner/deployer"

	"github.com/abklabs/svmkit/cmd/svmkit/build/yellowstone_grpc/assets"
)

type Build struct {
	runner.RunnerCommand
	BuildDir    string
	KeepPayload bool
	PackageName string
	Maintainer  string
	NoBuild     bool
	AgavePath   string
}

func (cmd *Build) Env() *runner.EnvBuilder {
	env := runner.NewEnvBuilder()

	env.Set("BUILD_DIR", cmd.BuildDir)
	env.Set("PACKAGE_NAME", cmd.PackageName)
	env.Set("MAINTAINER", cmd.Maintainer)
	env.SetBool("NO_BUILD", cmd.NoBuild)
	env.Set("AGAVE_PATH", cmd.AgavePath)

	return env
}

func (cmd *Build) Check() error {
	cmd.SetConfigDefaults()

	pkgGrp := deb.Package{}.MakePackageGroup()

	pkgGrp.Add(deb.Package{
		Name:          "protobuf-compiler",
	})

	if err := cmd.UpdatePackageGroup(pkgGrp); err != nil {
		return err
	}

	return nil
}

func (cmd *Build) AddToPayload(p *runner.Payload) error {
	script, err := assets.FS.Open(assets.BuildScript)

	if err != nil {
		return err
	}

	p.AddReader("steps.sh", script)

	if err := cmd.RunnerCommand.AddToPayload(p); err != nil {
		return err
	}

	return nil
}

var YellowstoneGRPCCmd = &cobra.Command{
	Use:   "yellowstone-grpc",
	Short: "Build a Yellowstone-GRPC Debian package",
	RunE: func(cmd *cobra.Command, args []string) error {
		cwd, err := os.Getwd()

		if err != nil {
			return err
		}

		flags := cmd.Flags()

		keepPayload, err := flags.GetBool("keep-payload")

		if err != nil {
			return err
		}

		maintainer, err := flags.GetString("maintainer")

		if err != nil {
			return err
		}

		packageName, err := flags.GetString("package-name")

		if err != nil {
			return err
		}

		noBuild, err := flags.GetBool("no-build")

		if err != nil {
			return err
		}

		agavePath, err := flags.GetString("agave-path")

		if err != nil {
			return err
		}

		runnerCommand := &Build{
			BuildDir:    cwd,
			KeepPayload: keepPayload,
			PackageName: packageName,
			Maintainer:  maintainer,
			NoBuild:     noBuild,
			AgavePath:   agavePath,
		}

		if err = runnerCommand.Check(); err != nil {
			return err
		}

		outputDir, err := os.MkdirTemp("", "build-*")

		if err != nil {
			return err
		}

		p := &runner.Payload{
			RootPath:    outputDir,
			DefaultMode: 0640,
		}

		if err := runner.PrepareCommandPayload(p, runnerCommand); err != nil {
			return err
		}

		d := &deployer.Local{
			Payload:     p,
			KeepPayload: keepPayload,
		}

		log.Printf("writing to '%s'...", outputDir)

		if err := d.Deploy(); err != nil {
			return err
		}

		handler := &deployer.LoggerHandler{
			LogCallback: func(s string) {
				log.Print(s)
			},
		}

		return d.Run([]string{"./run.sh"}, handler)
	},
}

func init() {
	flags := YellowstoneGRPCCmd.Flags()

	flags.String("maintainer", "Engineering <engineering@abklabs.com>", "name and email of the maintainer of the package")
	flags.String("package-name", "yellowstone-grpc", "name of the package")
	flags.Bool("keep-payload", false, "don't remove build scripts after completion")
	flags.Bool("no-build", false, "configure the repository, but do not build")
	flags.String("agave-path", "", "path to local agave repo (required)")
	_ = YellowstoneGRPCCmd.MarkFlagRequired("agave-path")
}
