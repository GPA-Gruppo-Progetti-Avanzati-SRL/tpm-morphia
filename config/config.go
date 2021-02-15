package config

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system"
	"github.com/GPA-Gruppo-Progetti-Avanzati-SRL/tpm-morphia/system/resources"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"io/ioutil"
	"os"
	"path/filepath"
	"sigs.k8s.io/yaml"
	"strings"
)

const DefaultConfigFile = "tpm-morphia.yml"

type Config struct {
	flagSet *flag.FlagSet

	ConfigFile      string `json:"config-file"`
	TargetDirectory string `json:"out-dir"`
	Version         string `json:"tmpl-ver"`

	ResourceDirectory string
	FormatCode        bool `json:"format-code"`

	CollectionDefFile     string `json:"collection-def-file"`
	CollectionDefScanPath string `json:"collection-def-scan-path"`
}

var DefaultConfig = Config{
	flagSet:    nil,
	ConfigFile: DefaultConfigFile,
	Version:    "v1",
	FormatCode: true,
}

type configBuilder struct {
	configFile string
}

type ConfigBuilder interface {
	Build(ctx context.Context) (*Config, error)
	With(cfgFileName string) ConfigBuilder
}

func NewBuilder() ConfigBuilder {
	bld := &configBuilder{configFile: DefaultConfigFile}
	return bld
}

func (bld *configBuilder) With(fileName string) ConfigBuilder {
	bld.configFile = fileName
	return bld
}

func (bld *configBuilder) Build(ctx context.Context) (*Config, error) {
	return newConfig(ctx, bld.configFile)
}

func (cfg *Config) String() string {
	return fmt.Sprintf("%#v\n", cfg)
}

func newConfig(_ context.Context, cfgFileName string) (*Config, error) {
	logger := system.GetLogger()

	cfg := DefaultConfig
	cfg.ConfigFile = cfgFileName

	_ = level.Debug(logger).Log(system.DefaultLogMessageField, "Embedded Configuration Loaded", "Config", cfg)

	if _, err := cfg.readConfigFromFile(logger, cfg.ConfigFile, false); err != nil {
		return &cfg, err
	}

	_ = level.Info(logger).Log("Message", "Initializing Flag Set")
	cfg.initializeFlagSet()

	currentConfigFile := cfg.ConfigFile

	_ = level.Info(logger).Log("Message", "Parsing Cmd Line Param")
	if err := cfg.flagSet.Parse(os.Args[1:]); err != nil {
		return &cfg, err
	}

	if len(cfg.flagSet.Args()) != 0 {
		_ = level.Warn(logger).Log(system.DefaultLogMessageField, "Invalid Command Line flag", "Flag", cfg.flagSet.Arg(0))
	}

	_ = level.Debug(logger).Log(system.DefaultLogMessageField, "Command Line Parsed", "Config", cfg)

	if cfg.ConfigFile != currentConfigFile {
		/*
		 * Il caricamento dell'ultimo file disponibile non modifica il flagSet. Eventuali flag per path 'dinamici' eventualmente inseriti sono
		 * censiti tra gli errori....
		 */
		if _, err := cfg.readConfigFromFile(logger, cfg.ConfigFile, true); err != nil {
			return &cfg, err
		}
	}

	_ = level.Info(logger).Log(system.DefaultLogMessageField, "Configuration Loaded", "Config", cfg)

	if err := cfg.checkValid(); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (cfg *Config) initializeFlagSet() {
	cfg.flagSet = flag.NewFlagSet("tpm_router", flag.ContinueOnError)

	cfg.flagSet.StringVar(&cfg.ConfigFile, "config-file", cfg.ConfigFile, "Path to the morphia configuration file.")

	/*
	 */
	cfg.flagSet.BoolVar(&cfg.FormatCode, "format-code", cfg.FormatCode, "Boolean: format code?")
	cfg.flagSet.StringVar(&cfg.TargetDirectory, "out-dir", cfg.TargetDirectory, "mount point of generated files.")
	cfg.flagSet.StringVar(&cfg.Version, "tmpl-ver", cfg.Version, "Version of templates (v1)")
	cfg.flagSet.StringVar(&cfg.CollectionDefFile, "collection-def-file", cfg.CollectionDefFile, "collection definition filename")
	cfg.flagSet.StringVar(&cfg.CollectionDefScanPath, "collection-def-scan-path", cfg.CollectionDefScanPath, "scan directory for collection definition")
}

func (cfg *Config) readConfigFromFile(logger log.Logger, aConfigFileName string, mustExists bool) (bool, error) {

	fileLoaded := false

	_ = level.Info(logger).Log(system.DefaultLogMessageField, "Loading Config File", "FileName", aConfigFileName)

	var configContent []byte
	var err error
	if _, err = os.Stat(aConfigFileName); err == nil {
		if configContent, err = ioutil.ReadFile(aConfigFileName); err != nil {
			return false, err
		}
	} else if os.IsNotExist(err) {

		if !resources.Has("/resources/" + aConfigFileName) {
			if mustExists {
				_ = level.Error(logger).Log(system.DefaultLogMessageField, "config file not found", "FileName", aConfigFileName, "mustExists", true)
				return false, err
			} else {
				_ = level.Warn(logger).Log(system.DefaultLogMessageField, "config file not found", "FileName", aConfigFileName, "mustExists", false)
				return false, nil
			}
		} else {
			configContent, _ = resources.Get("/resources/" + aConfigFileName)
			aConfigFileName = "/resources/" + aConfigFileName
		}

		// path/to/whatever does *not* exist
	} else {
		// Schrodinger: file may or may not exist. See err for details.
		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
	}

	if configContent != nil {
		if yerr := cfg.readConfigFromByteArray(configContent); yerr != nil {
			_ = level.Error(logger).Log(system.DefaultLogMessageField, "Config file YAML error", "FileName", aConfigFileName, "YAML Error", yerr.Error())
			return false, yerr
		} else {
			_ = level.Info(logger).Log(system.DefaultLogMessageField, "Configuration file loaded", "FileName", aConfigFileName)
			cfg.ConfigFile = aConfigFileName
			fileLoaded = true
		}
	}

	return fileLoaded, nil
}

func (cfg *Config) readConfigFromByteArray(configContent []byte) error {

	var tpmexporterCfg struct{ TPMExporter *Config }
	tpmexporterCfg.TPMExporter = cfg
	yerr := yaml.Unmarshal(configContent, &tpmexporterCfg)
	return yerr
}

func (cfg *Config) checkValid() error {

	if cfg.CollectionDefFile == "" && cfg.CollectionDefScanPath == "" {
		return errors.New("no definition files or scan directories specified in config")
	}

	if cfg.CollectionDefFile != "" {
		if f, err := os.Stat(cfg.CollectionDefFile); err != nil {
			return err
		} else {
			if f.IsDir() {
				return errors.New("collection def file is directory")
			}
		}
	}

	if cfg.CollectionDefScanPath != "" {
		if f, err := os.Stat(cfg.CollectionDefScanPath); err != nil {
			return err
		} else {
			if !f.IsDir() {
				return errors.New("collection def file is not directory")
			}
		}
	}

	if cfg.TargetDirectory == "" {
		return errors.New("missing out-dir config")
	} else if f, err := os.Stat(cfg.TargetDirectory); err != nil {
		return err
	} else {
		if !f.IsDir() {
			return errors.New("TargetDirectory def file is not directory")
		}
	}

	if cfg.Version != "v1" {
		return errors.New("just have v1 version of templates sorry")
	}
	return nil
}

func (cfg *Config) FindCollectionToProcess(logger log.Logger) ([]string, error) {

	if cfg.CollectionDefFile != "" {
		if _, err := os.Stat(cfg.CollectionDefFile); err == nil {
			return []string{cfg.CollectionDefFile}, nil
		} else {
			return nil, err
		}
	}

	if cfg.CollectionDefScanPath != "" {

		defs := make([]string, 0)
		err := filepath.Walk(cfg.CollectionDefScanPath, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if filterPath(info.Name(), info.IsDir()) {
				if !info.IsDir() {
					_ = level.Debug(logger).Log("msg", "visited file or dir", "name", path)
					defs = append(defs, path)
				}
				return nil
			} else {
				if info.IsDir() {
					// fmt.Printf("skipping dir: %+v \n", info.Name())
					_ = level.Debug(logger).Log("msg", "skipping dir", "name", info.Name())
					return filepath.SkipDir
				}
			}

			return nil
		})

		if err != nil {
			return nil, err
		} else if len(defs) > 0 {
			return defs, nil
		}
	}

	return nil, errors.New("no files specified in config")
}

func filterPath(n string, isDir bool) bool {

	if isDir {
		if strings.HasPrefix(n, ".") && n != "." && n != ".." {
			return false
		}

		return true
	} else {
		if strings.HasSuffix(n, "-tpmm.json") {
			return true
		}

		return false
	}
}
