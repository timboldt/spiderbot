struct Servo {
    min_micros: u16,
    max_micros: u16,
    zero_degrees: i16,
    reversed: bool,
}

impl Servo {
    fn degrees_to_micros(&self, degrees: f32) -> u16 {
        let micros_per_deg = 100f32 / 9f32;
        let micros = if self.reversed {
            (-micros_per_deg * degrees).round() as i16
        } else {
            (micros_per_deg * degrees).round() as i16
        };
        core::cmp::max(
            core::cmp::min(
                core::cmp::max(micros + self.zero_degrees, 0) as u16,
                self.max_micros,
            ),
            self.min_micros,
        )
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_degrees_to_micros() {
        let s = Servo {
            min_micros: 800,
            max_micros: 2200,
            zero_degrees: 500,
            reversed: false,
        };
        assert_eq!(s.degrees_to_micros(-10f32), 800);
        assert_eq!(s.degrees_to_micros(0f32), 800);
        assert_eq!(s.degrees_to_micros(10f32), 800);
        assert_eq!(s.degrees_to_micros(45f32), 1000);
        assert_eq!(s.degrees_to_micros(90f32), 1500);
        assert_eq!(s.degrees_to_micros(200f32), 2200);
    }

    #[test]
    fn test_reverse_degrees_to_micros() {
        let s = Servo {
            min_micros: 200,
            max_micros: 2200,
            zero_degrees: 1700,
            reversed: true,
        };
        assert_eq!(s.degrees_to_micros(-90f32), 2200);
        assert_eq!(s.degrees_to_micros(-10f32), 1811);
        assert_eq!(s.degrees_to_micros(0f32), 1700);
        assert_eq!(s.degrees_to_micros(10f32), 1589);
        assert_eq!(s.degrees_to_micros(45f32), 1200);
        assert_eq!(s.degrees_to_micros(90f32), 700);
        assert_eq!(s.degrees_to_micros(200f32), 200);
    }
}
