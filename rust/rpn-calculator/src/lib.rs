#[derive(Debug)]
pub enum CalculatorInput {
    Add,
    Subtract,
    Multiply,
    Divide,
    Value(i32),
}

pub fn evaluate(inputs: &[CalculatorInput]) -> Option<i32> {
    let mut stack: Vec<i32> = vec![];

    for i in inputs {
        if let CalculatorInput::Value(v) = i {
            stack.push(*v);
        } else {
            let r = stack.pop()?;
            let l = stack.pop()?;
            match i {
                CalculatorInput::Add => stack.push(l + r),
                CalculatorInput::Multiply => stack.push(l * r),
                CalculatorInput::Subtract => stack.push(l - r),
                CalculatorInput::Divide => stack.push(l / r),
                _ => ()
            }
        }
    }
    if stack.len() == 1 {
        stack.pop()
    } else {
        None
    }
}
