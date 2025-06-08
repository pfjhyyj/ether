use async_trait::async_trait;
use casbin::{Adapter, Filter, Model, Result};

use super::model::Rule;

pub struct MemoryAdapter {
    rules: Vec<Rule>,
    is_filtered: bool,
}

impl MemoryAdapter {
    pub fn new(rules: Vec<Rule>) -> Self {
        MemoryAdapter {
            rules,
            is_filtered: false,
        }
    }

    fn normalize_policy(rule: &Rule) -> Option<Vec<String>> {
        let mut policy = vec![
            &rule.v0, &rule.v1, &rule.v2, &rule.v3, &rule.v4, &rule.v5,
        ];

        loop {
            match policy.last() {
                Some(last) if last.is_empty() => {
                    policy.pop();
                }
                _ => break,
            }
        }

        if policy.is_empty() {
            None
        } else {
            Some(policy.iter().map(|&x| x.to_owned()).collect())
        }
    }
}

#[async_trait]
impl Adapter for MemoryAdapter {
    async fn load_policy(&mut self, m: &mut dyn Model) -> Result<()> {
        for rule in &self.rules {
            let Some(sec) = rule.ptype.chars().next().map(|x| x.to_string()) else {
                continue;
            };
            let Some(t1) = m.get_mut_model().get_mut(&sec) else {
                continue;
            };
            let Some(t2) = t1.get_mut(&rule.ptype) else {
                continue;
            };
            let Some(policy) = Self::normalize_policy(&rule) else {
                continue;
            };
            t2.get_mut_policy().insert(policy);
        }
        Ok(())
    }

    async fn load_filtered_policy<'a>(
        &mut self,
        m: &mut dyn Model,
        f: Filter<'a>,
    ) -> Result<()> {
        let filter_rules = self
            .rules
            .iter()
            .filter(|rule| {
                f.p.iter().any(|p| rule.ptype == *p)
                    || f.g.iter().any(|g| rule.ptype == *g)
            })
            .cloned()
            .collect::<Vec<_>>();

        for rule in filter_rules {
            let Some(sec) = rule.ptype.chars().next().map(|x| x.to_string()) else {
                continue;
            };
            let Some(t1) = m.get_mut_model().get_mut(&sec) else {
                continue;
            };
            let Some(t2) = t1.get_mut(&rule.ptype) else {
                continue;
            };
            let Some(policy) = Self::normalize_policy(&rule) else {
                continue;
            };
            t2.get_mut_policy().insert(policy);
        }
        self.is_filtered = true;
        Ok(())
    }

    async fn save_policy(&mut self, _m: &mut dyn Model) -> Result<()> {
        // Here we do not need to save anything as we are using an in-memory adapter
        Ok(())
    }

    async fn clear_policy(&mut self) -> Result<()> {
        self.rules.clear();
        self.is_filtered = false;
        Ok(())
    }

    fn is_filtered(&self) -> bool {
        self.is_filtered
    }

    async fn add_policy(
        &mut self,
        _sec: &str,
        ptype: &str,
        rule: Vec<String>,
    ) -> Result<bool> {
        // Add policy to the database
        let rule = Rule {
            ptype: ptype.to_string(),
            v0: rule.get(0).cloned().unwrap_or_default(),
            v1: rule.get(1).cloned().unwrap_or_default(),
            v2: rule.get(2).cloned().unwrap_or_default(),
            v3: rule.get(3).cloned().unwrap_or_default(),
            v4: rule.get(4).cloned().unwrap_or_default(),
            v5: rule.get(5).cloned().unwrap_or_default(),
        };
        self.rules.push(rule);
        Ok(true)
    }

    async fn add_policies(
        &mut self,
        _sec: &str,
        ptype: &str,
        rules: Vec<Vec<String>>,
    ) -> Result<bool> {
        // Add policies to the database
        for rule in rules {
            let rule = Rule {
                ptype: ptype.to_string(),
                v0: rule.get(0).cloned().unwrap_or_default(),
                v1: rule.get(1).cloned().unwrap_or_default(),
                v2: rule.get(2).cloned().unwrap_or_default(),
                v3: rule.get(3).cloned().unwrap_or_default(),
                v4: rule.get(4).cloned().unwrap_or_default(),
                v5: rule.get(5).cloned().unwrap_or_default(),
            };
            self.rules.push(rule);
        }
        Ok(true)
    }

    async fn remove_policy(
        &mut self,
        _sec: &str,
        ptype: &str,
        rule: Vec<String>,
    ) -> Result<bool> {
        // Remove policy from the database
        let rule = Rule {
            ptype: ptype.to_string(),
            v0: rule.get(0).cloned().unwrap_or_default(),
            v1: rule.get(1).cloned().unwrap_or_default(),
            v2: rule.get(2).cloned().unwrap_or_default(),
            v3: rule.get(3).cloned().unwrap_or_default(),
            v4: rule.get(4).cloned().unwrap_or_default(),
            v5: rule.get(5).cloned().unwrap_or_default(),
        };
        self.rules.retain(|r| r != &rule);
        Ok(true)
    }

    async fn remove_policies(
        &mut self,
        _sec: &str,
        ptype: &str,
        rules: Vec<Vec<String>>,
    ) -> Result<bool> {
        // Remove policies from the database
        for rule in rules {
            let rule = Rule {
                ptype: ptype.to_string(),
                v0: rule.get(0).cloned().unwrap_or_default(),
                v1: rule.get(1).cloned().unwrap_or_default(),
                v2: rule.get(2).cloned().unwrap_or_default(),
                v3: rule.get(3).cloned().unwrap_or_default(),
                v4: rule.get(4).cloned().unwrap_or_default(),
                v5: rule.get(5).cloned().unwrap_or_default(),
            };
            self.rules.retain(|r| r != &rule);
        }
        Ok(true)
    }

    async fn remove_filtered_policy(
        &mut self,
        _sec: &str,
        ptype: &str,
        field_index: usize,
        field_values: Vec<String>,
    ) -> Result<bool> {
        // Remove filtered policy from the database
        self.rules.retain(|rule| {
            if rule.ptype != ptype {
                return true;
            }
            let fields = [
                &rule.v0, &rule.v1, &rule.v2, &rule.v3, &rule.v4, &rule.v5,
            ];
            !field_values.iter().enumerate().any(|(i, value)| {
                i == field_index && (fields[i] == value || value.is_empty())
            })
        });
        self.is_filtered = false;
        Ok(true)
    }
}