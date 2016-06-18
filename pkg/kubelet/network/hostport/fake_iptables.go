/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package hostport

import (
	"bytes"
	"fmt"
	"net"
	"strings"

	utiliptables "k8s.io/kubernetes/pkg/util/iptables"
)

type fakeChain struct {
	name  utiliptables.Chain
	rules []string
}

type fakeTable struct {
	name   utiliptables.Table
	chains map[string]*fakeChain
}

type fakeIptables struct {
	tables map[string]*fakeTable
}

func NewFakeIptables() *fakeIptables {
	return &fakeIptables{
		tables: make(map[string]*fakeTable, 0),
	}
}

func (f *fakeIptables) GetVersion() (string, error) {
	return "1.4.21", nil
}

func (f *fakeIptables) getTable(tableName utiliptables.Table) (*fakeTable, error) {
	table, ok := f.tables[string(tableName)]
	if !ok {
		return nil, fmt.Errorf("Table %s does not exist", tableName)
	}
	return table, nil
}

func (f *fakeIptables) getChain(tableName utiliptables.Table, chainName utiliptables.Chain) (*fakeTable, *fakeChain, error) {
	table, err := f.getTable(tableName)
	if err != nil {
		return nil, nil, err
	}

	chain, ok := table.chains[string(chainName)]
	if !ok {
		return table, nil, fmt.Errorf("Chain %s/%s does not exist", tableName, chainName)
	}

	return table, chain, nil
}

func (f *fakeIptables) ensureChain(tableName utiliptables.Table, chainName utiliptables.Chain) (bool, *fakeChain) {
	table, chain, err := f.getChain(tableName, chainName)
	if err != nil {
		// either table or table+chain don't exist yet
		if table == nil {
			table = &fakeTable{
				name:   tableName,
				chains: make(map[string]*fakeChain),
			}
			f.tables[string(tableName)] = table
		}
		chain := &fakeChain{
			name:  chainName,
			rules: make([]string, 0),
		}
		table.chains[string(chainName)] = chain
		return false, chain
	}
	return true, chain
}

func (f *fakeIptables) EnsureChain(tableName utiliptables.Table, chainName utiliptables.Chain) (bool, error) {
	existed, _ := f.ensureChain(tableName, chainName)
	return existed, nil
}

func (f *fakeIptables) FlushChain(tableName utiliptables.Table, chainName utiliptables.Chain) error {
	_, chain, err := f.getChain(tableName, chainName)
	if err != nil {
		return err
	}
	chain.rules = make([]string, 0)
	return nil
}

func (f *fakeIptables) DeleteChain(tableName utiliptables.Table, chainName utiliptables.Chain) error {
	table, _, err := f.getChain(tableName, chainName)
	if err != nil {
		return err
	}
	delete(table.chains, string(chainName))
	return nil
}

// Returns index of rule in array; < 0 if rule is not found
func findRule(chain *fakeChain, rule string) int {
	for i, candidate := range chain.rules {
		if rule == candidate {
			return i
		}
	}
	return -1
}

func (f *fakeIptables) ensureRule(position utiliptables.RulePosition, tableName utiliptables.Table, chainName utiliptables.Chain, rule string) (bool, error) {
	_, chain, err := f.getChain(tableName, chainName)
	if err != nil {
		_, chain = f.ensureChain(tableName, chainName)
	}

	rule, err = normalizeRule(rule)
	if err != nil {
		return false, err
	}
	ruleIdx := findRule(chain, rule)
	if ruleIdx >= 0 {
		return true, nil
	}

	if position == utiliptables.Prepend {
		chain.rules = append([]string{rule}, chain.rules...)
	} else if position == utiliptables.Append {
		chain.rules = append(chain.rules, rule)
	} else {
		return false, fmt.Errorf("Unknown position argument %q", position)
	}

	return false, nil
}

func normalizeRule(rule string) (string, error) {
	normalized := ""
	remaining := strings.TrimSpace(rule)
	for {
		var end int

		if strings.HasPrefix(remaining, "--to-destination=") {
			remaining = strings.Replace(remaining, "=", " ", 1)
		}

		if remaining[0] == '"' {
			end = strings.Index(remaining[1:], "\"")
			if end < 0 {
				return "", fmt.Errorf("Invalid rule syntax: mismatched quotes")
			}
			end += 2
		} else {
			end = strings.Index(remaining, " ")
			if end < 0 {
				end = len(remaining)
			}
		}
		arg := remaining[:end]

		// Normalize un-prefixed IP addresses like iptables does
		if net.ParseIP(arg) != nil {
			arg = arg + "/32"
		}

		if len(normalized) > 0 {
			normalized += " "
		}
		normalized += strings.TrimSpace(arg)
		if len(remaining) == end {
			break
		}
		remaining = remaining[end+1:]
	}
	return normalized, nil
}

func (f *fakeIptables) EnsureRule(position utiliptables.RulePosition, tableName utiliptables.Table, chainName utiliptables.Chain, args ...string) (bool, error) {
	ruleArgs := make([]string, 0)
	for _, arg := range args {
		// quote args with internal spaces (like comments)
		if strings.Index(arg, " ") >= 0 {
			arg = fmt.Sprintf("\"%s\"", arg)
		}
		ruleArgs = append(ruleArgs, arg)
	}
	return f.ensureRule(position, tableName, chainName, strings.Join(ruleArgs, " "))
}

func (f *fakeIptables) DeleteRule(tableName utiliptables.Table, chainName utiliptables.Chain, args ...string) error {
	_, chain, err := f.getChain(tableName, chainName)
	if err == nil {
		rule := strings.Join(args, " ")
		ruleIdx := findRule(chain, rule)
		if ruleIdx < 0 {
			return nil
		}
		chain.rules = append(chain.rules[:ruleIdx], chain.rules[ruleIdx+1:]...)
	}
	return nil
}

func (f *fakeIptables) IsIpv6() bool {
	return false
}

func saveChain(chain *fakeChain, data *bytes.Buffer) {
	for _, rule := range chain.rules {
		data.WriteString(fmt.Sprintf("-A %s %s\n", chain.name, rule))
	}
}

func (f *fakeIptables) Save(tableName utiliptables.Table) ([]byte, error) {
	table, err := f.getTable(tableName)
	if err != nil {
		return nil, err
	}

	data := bytes.NewBuffer(nil)
	data.WriteString(fmt.Sprintf("*%s\n", table.name))

	rules := bytes.NewBuffer(nil)
	for _, chain := range table.chains {
		data.WriteString(fmt.Sprintf(":%s - [0:0]\n", string(chain.name)))
		saveChain(chain, rules)
	}
	data.Write(rules.Bytes())
	data.WriteString("COMMIT\n")
	return data.Bytes(), nil
}

func (f *fakeIptables) SaveAll() ([]byte, error) {
	data := bytes.NewBuffer(nil)
	for _, table := range f.tables {
		tableData, err := f.Save(table.name)
		if err != nil {
			return nil, err
		}
		if _, err = data.Write(tableData); err != nil {
			return nil, err
		}
	}
	return data.Bytes(), nil
}

func (f *fakeIptables) restore(restoreTableName utiliptables.Table, data []byte, flush utiliptables.FlushFlag) error {
	buf := bytes.NewBuffer(data)
	var tableName utiliptables.Table
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			break
		}
		if line[0] == '#' {
			continue
		}

		line = strings.TrimSuffix(line, "\n")
		if strings.HasPrefix(line, "*") {
			tableName = utiliptables.Table(line[1:])
		}
		if tableName != "" {
			if restoreTableName != "" && restoreTableName != tableName {
				continue
			}
			if strings.HasPrefix(line, ":") {
				chainName := utiliptables.Chain(strings.Split(line[1:], " ")[0])
				if flush == utiliptables.FlushTables {
					table, chain, _ := f.getChain(tableName, chainName)
					if chain != nil {
						delete(table.chains, string(chainName))
					}
				}
				_, _ = f.ensureChain(tableName, chainName)
			} else if strings.HasPrefix(line, "-A") {
				parts := strings.Split(line, " ")
				if len(parts) < 3 {
					return fmt.Errorf("Invalid iptables rule '%s'", line)
				}
				chainName := utiliptables.Chain(parts[1])
				rule := strings.TrimPrefix(line, fmt.Sprintf("-A %s ", chainName))
				_, err := f.ensureRule(utiliptables.Append, tableName, chainName, rule)
				if err != nil {
					return err
				}
			} else if strings.HasPrefix(line, "-X") {
				parts := strings.Split(line, " ")
				if len(parts) < 3 {
					return fmt.Errorf("Invalid iptables rule '%s'", line)
				}
				if err := f.DeleteChain(tableName, utiliptables.Chain(parts[1])); err != nil {
					return err
				}
			} else if line == "COMMIT" {
				if restoreTableName == tableName {
					return nil
				}
				tableName = ""
			}
		}
	}

	return nil
}

func (f *fakeIptables) Restore(tableName utiliptables.Table, data []byte, flush utiliptables.FlushFlag, counters utiliptables.RestoreCountersFlag) error {
	return f.restore(tableName, data, flush)
}

func (f *fakeIptables) RestoreAll(data []byte, flush utiliptables.FlushFlag, counters utiliptables.RestoreCountersFlag) error {
	return f.restore("", data, flush)
}

func (f *fakeIptables) AddReloadFunc(reloadFunc func()) {
}

func (f *fakeIptables) Destroy() {
}
