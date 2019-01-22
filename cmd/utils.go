package cmd

import (
	"github.com/jeffguorg/go-fuse-template/logger"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func OnCobraInitialized() {
	logrus.SetFormatter(&logger.TemplateFormatter{})
}

func init() {
	cobra.OnInitialize(OnCobraInitialized)
}
