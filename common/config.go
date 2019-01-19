package common

import (
	"errors"
	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/cfgutil"
	"github.com/elastic/go-ucfg/yaml"
)

// Config object to store hierarchical configurations into.
// See https://godoc.org/github.com/elastic/go-ucfg#Config
type Config ucfg.Config

// ConfigNamespace storing at most one configuration section by name and sub-section.
type ConfigNamespace struct {
	name   string
	config *Config
}

var configOpts = []ucfg.Option{
	ucfg.PathSep("."),
	ucfg.ResolveEnv,
	ucfg.VarExp,
}

func NewConfig() *Config {
	return fromConfig(ucfg.New())
}

// NewConfigFrom creates a new Config object from the given input.
// From can be any kind of structured data (struct, map, array, slice).
//
// If from is a string, the contents is treated like raw YAML input. The string
// will be parsed and a structure config object is build from the parsed
// result.
func NewConfigFrom(from interface{}) (*Config, error) {
	if str, ok := from.(string); ok {
		c, err := yaml.NewConfig([]byte(str), configOpts...)
		return fromConfig(c), err
	}

	c, err := ucfg.NewFrom(from, configOpts...)
	return fromConfig(c), err
}

// MustNewConfigFrom creates a new Config object from the given input.
// From can be any kind of structured data (struct, map, array, slice).
//
// If from is a string, the contents is treated like raw YAML input. The string
// will be parsed and a structure config object is build from the parsed
// result.
//
// MustNewConfigFrom panics if an error occurs.
func MustNewConfigFrom(from interface{}) *Config {
	cfg, err := NewConfigFrom(from)
	if err != nil {
		panic(err)
	}
	return cfg
}

func MergeConfigs(cfgs ...*Config) (*Config, error) {
	config := NewConfig()
	for _, c := range cfgs {
		if err := config.Merge(c); err != nil {
			return nil, err
		}
	}
	return config, nil
}

func NewConfigWithYAML(in []byte, source string) (*Config, error) {
	opts := append(
		[]ucfg.Option{
			ucfg.MetaData(ucfg.Meta{Source: source}),
		},
		configOpts...,
	)
	c, err := yaml.NewConfig(in, opts...)
	return fromConfig(c), err
}

// OverwriteConfigOpts allow to change the globally set config option
func OverwriteConfigOpts(options []ucfg.Option) {
	configOpts = options
}

func LoadFile(path string) (*Config, error) {

	c, err := yaml.NewConfigWithFile(path, configOpts...)
	if err != nil {
		return nil, err
	}

	cfg := fromConfig(c)
	return cfg, err
}

func LoadFiles(paths ...string) (*Config, error) {
	merger := cfgutil.NewCollector(nil, configOpts...)
	for _, path := range paths {
		cfg, err := LoadFile(path)
		if err := merger.Add(cfg.access(), err); err != nil {
			return nil, err
		}
	}
	return fromConfig(merger.Config()), nil
}

func (c *Config) Merge(from interface{}) error {
	return c.access().Merge(from, configOpts...)
}

func (c *Config) Unpack(to interface{}) error {
	return c.access().Unpack(to, configOpts...)
}

func (c *Config) Path() string {
	return c.access().Path(".")
}

func (c *Config) PathOf(field string) string {
	return c.access().PathOf(field, ".")
}

func (c *Config) HasField(name string) bool {
	return c.access().HasField(name)
}

func (c *Config) CountField(name string) (int, error) {
	return c.access().CountField(name)
}

func (c *Config) Bool(name string, idx int) (bool, error) {
	return c.access().Bool(name, idx, configOpts...)
}

func (c *Config) String(name string, idx int) (string, error) {
	return c.access().String(name, idx, configOpts...)
}

func (c *Config) Int(name string, idx int) (int64, error) {
	return c.access().Int(name, idx, configOpts...)
}

func (c *Config) Float(name string, idx int) (float64, error) {
	return c.access().Float(name, idx, configOpts...)
}

func (c *Config) Child(name string, idx int) (*Config, error) {
	sub, err := c.access().Child(name, idx, configOpts...)
	return fromConfig(sub), err
}

func (c *Config) SetBool(name string, idx int, value bool) error {
	return c.access().SetBool(name, idx, value, configOpts...)
}

func (c *Config) SetInt(name string, idx int, value int64) error {
	return c.access().SetInt(name, idx, value, configOpts...)
}

func (c *Config) SetFloat(name string, idx int, value float64) error {
	return c.access().SetFloat(name, idx, value, configOpts...)
}

func (c *Config) SetString(name string, idx int, value string) error {
	return c.access().SetString(name, idx, value, configOpts...)
}

func (c *Config) SetChild(name string, idx int, value *Config) error {
	return c.access().SetChild(name, idx, value.access(), configOpts...)
}

func (c *Config) IsDict() bool {
	return c.access().IsDict()
}

func (c *Config) IsArray() bool {
	return c.access().IsArray()
}

func fromConfig(in *ucfg.Config) *Config {
	return (*Config)(in)
}

func (c *Config) access() *ucfg.Config {
	return (*ucfg.Config)(c)
}

func (c *Config) GetFields() []string {
	return c.access().GetFields()
}

// Unpack unpacks a configuration with at most one sub object. An sub object is
// ignored if it is disabled by setting `enabled: false`. If the configuration
// passed contains multiple active sub objects, Unpack will return an error.
func (ns *ConfigNamespace) Unpack(cfg *Config) error {
	fields := cfg.GetFields()
	if len(fields) == 0 {
		return nil
	}

	var (
		err   error
		found bool
	)

	for _, name := range fields {
		var sub *Config

		sub, err = cfg.Child(name, -1)
		if err != nil {
			// element is no configuration object -> continue so a namespace
			// Config unpacked as a namespace can have other configuration
			// values as well
			continue
		}

		if ns.name != "" {
			return errors.New("more than one namespace configured")
		}

		ns.name = name
		ns.config = sub
		found = true
	}

	if !found {
		return err
	}
	return nil
}

// Name returns the configuration sections it's name if a section has been set.
func (ns *ConfigNamespace) Name() string {
	return ns.name
}

// Config return the sub-configuration section if a section has been set.
func (ns *ConfigNamespace) Config() *Config {
	return ns.config
}

// IsSet returns true if a sub-configuration section has been set.
func (ns *ConfigNamespace) IsSet() bool {
	return ns.config != nil
}