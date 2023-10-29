/*
Copyright © 2021-2023 Macaroni OS Linux
See AUTHORS and LICENSE for the license details and contributors.
*/
package browser

import (
	"fmt"
)

func (p *BrowserPackage) HasOptions() bool {
	return len(p.EnabledOptions) > 0
}

func (p *BrowserPackage) ClearOptions() {
	p.EnabledOptions = []*BrowserOpt{}
}

func (p *BrowserPackage) GetAllOptions() []string {
	ans := []string{}

	for i, _ := range p.EnabledOptions {
		ans = append(ans, p.EnabledOptions[i].Option...)
	}

	return ans
}

func (p *BrowserPackage) SetEnabledOptions(opts []*BrowserOpt) {
	p.EnabledOptions = []*BrowserOpt{}

	for _, o := range opts {
		o.Minimize()
		p.EnabledOptions = append(p.EnabledOptions, o)
	}
}

func (p *BrowserPackage) Merge(m *BrowserPackage) error {
	// Creating a map of the first option enabled
	mOpts := make(map[string]*BrowserOpt, 0)

	if p.Package != m.Package {
		return fmt.Errorf("Trying to merge package %s with %s",
			p.Package, m.Package)
	}

	for i, _ := range p.EnabledOptions {
		mOpts[p.EnabledOptions[i].Option[0]] = p.EnabledOptions[i]
	}

	// Check if add new options
	for i, _ := range m.EnabledOptions {
		if _, present := mOpts[m.EnabledOptions[i].Option[0]]; !present {
			p.EnabledOptions = append(
				p.EnabledOptions,
				m.EnabledOptions[i],
			)
		}
	}

	return nil
}
