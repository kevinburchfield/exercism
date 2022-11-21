// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

pub struct Player {
    pub health: u32,
    pub mana: Option<u32>,
    pub level: u32,
}

impl Player {
    pub fn revive(&self) -> Option<Player> {
        match self.level {
            l if l < 10 && l > 0 && self.health <= 0 => Some(Player {
                level: self.level,
                mana: None,
                health: 100
            }),
            l if l >= 10 && self.health <= 0 => Some(Player{
                level: self.level,
                mana: Some(100),
                health: 100
            }),
            _ => None
        }
    }

    pub fn cast_spell(&mut self, mana_cost: u32) -> u32 {
        match self.mana {
            Some(m) => {
                if mana_cost <= m {
                    self.mana = Some(m - mana_cost);
                    2 * mana_cost
                } else {
                    0
                }
            },
            _ => {
                if self.health > mana_cost {
                    self.health -= mana_cost
                } else {
                    self.health = 0;
                }
                0
            }
        }
    }
}
