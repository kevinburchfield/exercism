// This stub file contains items that aren't used yet; feel free to remove this module attribute
// to enable stricter warnings.
#![allow(unused)]

use std::collections::HashMap;

pub fn can_construct_note(magazine: &[&str], note: &[&str]) -> bool {
    let mut hm: HashMap<&str, u32> = HashMap::new();
    for word in magazine {
        let count = hm.get(word);
        match count {
            Some(c) => hm.insert(word, c + 1),
            _ => hm.insert(word, 1)
        };
    }
    for word in note {
        let count = hm.get(word);
        match count {
            Some(&c) => {
                println!("{:?}", c);
                if c == 0 {
                    return false;
                } else {
                    hm.insert(word, c - 1);
                }
            },
            None => return false
        }
    }
    true
}
