package config

// Validate confirms the necessary values to support the desired Clair mode
// exist and sets default values.
func Validate(c *Config) ([]Warning, error) {
	m, err := ParseMode(c.Mode)
	if err != nil {
		return nil, err
	}
	return forEach(c, func(i interface{}) ([]Warning, error) {
		if v, ok := i.(validator); ok {
			return v.validate(m)
		}
		return nil, nil
	})
}

// Types that want complex defaults or to fail validation can implement the
// validator interface.
type validator interface {
	validate(Mode) ([]Warning, error)
}
