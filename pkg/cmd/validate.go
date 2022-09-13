package cmd

import (
	"errors"

	"github.com/Li-AnLin/boilr/pkg/util/exit"
	"github.com/Li-AnLin/boilr/pkg/util/validate"
	cli "github.com/spf13/cobra"
)

// ErrTemplateInvalid indicates that the template is invalid.
var ErrTemplateInvalid = errors.New("validate: given template is invalid")

// Validate contains the cli-command for validating templates.
var Validate = &cli.Command{
	Use:   "validate <template-path>",
	Short: "Validate a local project template",
	Run: func(_ *cli.Command, args []string) {
		MustValidateArgs(args, []validate.Argument{
			{"template-path", validate.UnixPath},
		})

		templatePath := args[0]

		MustValidateTemplate(templatePath)

		exit.OK("Template is valid")
	},
}
