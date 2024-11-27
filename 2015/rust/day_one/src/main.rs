fn part_one(input: &str) -> isize {
    input.chars().fold(0, |val, ch| match ch {
        '(' => val + 1,
        ')' => val - 1,
        ch => unreachable!("Char is unexpected {ch}"),
    })
}

fn part_two(input: &str) -> usize {
    input
        .chars()
        .scan(0, |val, ch| {
            let add_value = match ch {
                '(' => 1,
                ')' => -1,
                ch => unreachable!("Char is unexpected {ch}"),
            };
            *val = *val + add_value;
            Some(val.clone())
        })
        .take_while(|x| *x >= 0)
        .count()
        + 1
}

fn main() {
    let input = include_str!("../input.txt").trim();

    let part_one_result = part_one(input);
    println!("Part 1 result: {part_one_result}");

    let part_two_result = part_two(input);
    println!("Part 1 result: {part_two_result}");
}

#[cfg(test)]
mod tests {
    use crate::{part_one, part_two};

    #[test]
    pub fn part_one_example() {
        let test_inputs = [
            ("(())", 0),
            ("()()", 0),
            ("(((", 3),
            ("(()(()(", 3),
            ("))(((((", 3),
            ("())", -1),
            ("))(", -1),
            (")))", -3),
            (")())())", -3),
        ];
        for (input, expected) in test_inputs.iter() {
            let result = part_one(*input);
            assert_eq!(result, *expected);
        }
    }

    #[test]
    pub fn part_two_example() {
        let test_inputs = [(")", 1), ("()())", 5)];
        for (input, expected) in test_inputs.iter() {
            let result = part_two(*input);
            assert_eq!(result, *expected, "input: {input}");
        }
    }
}
